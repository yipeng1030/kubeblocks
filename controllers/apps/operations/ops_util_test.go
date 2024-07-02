/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package operations

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	opsutil "github.com/apecloud/kubeblocks/controllers/apps/operations/util"
	"github.com/apecloud/kubeblocks/pkg/constant"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/generics"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
	testk8s "github.com/apecloud/kubeblocks/pkg/testutil/k8s"
)

var _ = Describe("OpsUtil functions", func() {

	var (
		randomStr             = testCtx.GetRandomStr()
		clusterDefinitionName = "cluster-definition-for-ops-" + randomStr
		clusterVersionName    = "clusterversion-for-ops-" + randomStr
		clusterName           = "cluster-for-ops-" + randomStr
	)

	cleanEnv := func() {
		// must wait till resources deleted and no longer existed before the testcases start,
		// otherwise if later it needs to create some new resource objects with the same name,
		// in race conditions, it will find the existence of old objects, resulting failure to
		// create the new objects.
		By("clean resources")

		// delete cluster(and all dependent sub-resources), clusterversion and clusterdef
		testapps.ClearClusterResources(&testCtx)

		// delete rest resources
		inNS := client.InNamespace(testCtx.DefaultNamespace)
		ml := client.HasLabels{testCtx.TestObjLabelKey}
		// namespaced
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.InstanceSetSignature, true, inNS, ml)
		testapps.ClearResources(&testCtx, generics.ConfigMapSignature, inNS, ml)
		testapps.ClearResources(&testCtx, generics.OpsRequestSignature, inNS, ml)
	}

	BeforeEach(cleanEnv)

	AfterEach(cleanEnv)

	Context("Test ops_util functions", func() {
		It("Test ops_util functions", func() {
			By("init operations resources ")
			opsRes, _, _ := initOperationsResources(clusterDefinitionName, clusterVersionName, clusterName)
			testapps.MockInstanceSetComponent(&testCtx, clusterName, consensusComp)

			By("Test the functions in ops_util.go")
			opsRes.OpsRequest = createHorizontalScaling(clusterName, appsv1alpha1.HorizontalScaling{
				ComponentOps: appsv1alpha1.ComponentOps{ComponentName: consensusComp},
				Replicas:     pointer.Int32(1),
			})
			Expect(patchValidateErrorCondition(ctx, k8sClient, opsRes, "validate error")).Should(Succeed())
			Expect(PatchOpsHandlerNotSupported(ctx, k8sClient, opsRes)).Should(Succeed())
			Expect(isOpsRequestFailedPhase(appsv1alpha1.OpsFailedPhase)).Should(BeTrue())
			Expect(PatchClusterNotFound(ctx, k8sClient, opsRes)).Should(Succeed())
		})

		It("Test opsRequest failed cases", func() {
			By("init operations resources ")
			opsRes, _, _ := initOperationsResources(clusterDefinitionName, clusterVersionName, clusterName)
			testapps.MockInstanceSetComponent(&testCtx, clusterName, consensusComp)
			pods := testapps.MockInstanceSetPods(&testCtx, nil, opsRes.Cluster, consensusComp)
			time.Sleep(time.Second)
			By("Test the functions in ops_util.go")
			ops := testapps.NewOpsRequestObj("restart-ops-"+randomStr, testCtx.DefaultNamespace,
				clusterName, appsv1alpha1.RestartType)
			ops.Spec.RestartList = []appsv1alpha1.ComponentOps{{ComponentName: consensusComp}}
			opsRes.OpsRequest = testapps.CreateOpsRequest(ctx, testCtx, ops)
			opsRes.OpsRequest.Status.Phase = appsv1alpha1.OpsRunningPhase
			opsRes.OpsRequest.Status.StartTimestamp = metav1.Now()

			By("mock component failed")
			clusterComp := opsRes.Cluster.Status.Components[consensusComp]
			clusterComp.Phase = appsv1alpha1.FailedClusterCompPhase
			opsRes.Cluster.Status.SetComponentStatus(consensusComp, clusterComp)

			By("expect for opsRequest is running")
			handleRestartProgress := func(reqCtx intctrlutil.RequestCtx,
				cli client.Client,
				opsRes *OpsResource,
				pgRes *progressResource,
				compStatus *appsv1alpha1.OpsRequestComponentStatus) (expectProgressCount int32, completedCount int32, err error) {
				return handleComponentStatusProgress(reqCtx, cli, opsRes, pgRes, compStatus,
					func(ops *appsv1alpha1.OpsRequest, pod *corev1.Pod, compOps ComponentOpsInterface, s string) bool {
						return !pod.CreationTimestamp.Before(&ops.Status.StartTimestamp)
					})
			}

			reqCtx := intctrlutil.RequestCtx{Ctx: ctx}
			compOpsHelper := newComponentOpsHelper(opsRes.OpsRequest.Spec.RestartList)

			opsPhase, _, err := compOpsHelper.reconcileActionWithComponentOps(reqCtx, k8sClient, opsRes,
				"test", handleRestartProgress)
			Expect(err).Should(BeNil())
			Expect(opsPhase).Should(Equal(appsv1alpha1.OpsRunningPhase))

			By("mock one pod recreates failed, expect for opsRequest is Failed")
			testk8s.MockPodIsTerminating(ctx, testCtx, pods[2])
			testk8s.RemovePodFinalizer(ctx, testCtx, pods[2])
			// recreate it
			pod := testapps.MockInstanceSetPod(&testCtx, nil, clusterName, consensusComp, pods[2].Name, "follower", "Readonly")
			// mock pod is failed
			testk8s.MockPodIsFailed(ctx, testCtx, pod)
			opsPhase, _, err = compOpsHelper.reconcileActionWithComponentOps(reqCtx, k8sClient, opsRes, "test", handleRestartProgress)
			Expect(err).Should(BeNil())
			Expect(opsPhase).Should(Equal(appsv1alpha1.OpsFailedPhase))
		})

		It("Test opsRequest with disable ha", func() {
			By("init operations resources ")
			opsRes, _, _ := initOperationsResources(clusterDefinitionName, clusterVersionName, clusterName)

			By("Test the functions in ops_util.go")
			ops := testapps.NewOpsRequestObj("restart-ops-"+randomStr, testCtx.DefaultNamespace,
				clusterName, appsv1alpha1.RestartType)
			ops.Spec.RestartList = []appsv1alpha1.ComponentOps{{ComponentName: consensusComp}}
			opsRes.OpsRequest = testapps.CreateOpsRequest(ctx, testCtx, ops)
			Expect(testapps.ChangeObjStatus(&testCtx, opsRes.OpsRequest, func() {
				opsRes.OpsRequest.Status.Phase = appsv1alpha1.OpsCreatingPhase
				opsRes.OpsRequest.Status.StartTimestamp = metav1.Time{Time: time.Now()}
			})).Should(Succeed())

			By("create ha configmap and do horizontalScaling with disable ha")
			haConfigName := "ha-config"
			haConfig := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      haConfigName,
					Namespace: testCtx.DefaultNamespace,
					Annotations: map[string]string{
						"enable": "true",
					},
				},
			}
			Expect(k8sClient.Create(ctx, haConfig)).Should(Succeed())
			opsRes.OpsRequest.Annotations = map[string]string{
				constant.DisableHAAnnotationKey: haConfigName,
			}

			By("mock instance set")
			its := testapps.MockInstanceSetComponent(&testCtx, clusterName, consensusComp)

			By("expect to disable ha")
			reqCtx := intctrlutil.RequestCtx{Ctx: testCtx.Ctx}
			_, err := GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(haConfig), func(g Gomega, cm *corev1.ConfigMap) {
				cm.Annotations["enable"] = "false"
			})).Should(Succeed())

			By("mock restart ops to succeed and expect to enable ha")
			opsRes.OpsRequest.Status.Phase = appsv1alpha1.OpsRunningPhase
			_ = testapps.MockInstanceSetPods(&testCtx, its, opsRes.Cluster, consensusComp)
			_, err = GetOpsManager().Reconcile(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(testapps.GetOpsRequestPhase(&testCtx, client.ObjectKeyFromObject(opsRes.OpsRequest))).Should(Equal(appsv1alpha1.OpsSucceedPhase))
			Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(haConfig), func(g Gomega, cm *corev1.ConfigMap) {
				cm.Annotations["enable"] = "true"
			})).Should(Succeed())
		})

		It("Test opsRequest Queue functions", func() {
			By("init operations resources ")
			reqCtx := intctrlutil.RequestCtx{Ctx: testCtx.Ctx}
			opsRes, _, _ := initOperationsResources(clusterDefinitionName, clusterVersionName, clusterName)

			runHscaleOps := func(expectPhase appsv1alpha1.OpsPhase) *appsv1alpha1.OpsRequest {
				ops := createHorizontalScaling(clusterName, appsv1alpha1.HorizontalScaling{
					ComponentOps: appsv1alpha1.ComponentOps{ComponentName: consensusComp},
					Replicas:     pointer.Int32(1),
				})
				opsRes.OpsRequest = ops
				_, err := GetOpsManager().Do(reqCtx, k8sClient, opsRes)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(opsRes.OpsRequest.Status.Phase).Should(Equal(expectPhase))
				return ops
			}

			By("run first h-scale ops, expect phase to Creating")
			ops1 := runHscaleOps(appsv1alpha1.OpsCreatingPhase)

			By("run next h-scale ops, expect phase to Pending")
			ops2 := runHscaleOps(appsv1alpha1.OpsPendingPhase)

			By("check opsRequest annotation in cluster")
			cluster := &appsv1alpha1.Cluster{}
			Expect(k8sClient.Get(ctx, client.ObjectKeyFromObject(opsRes.Cluster), cluster)).Should(Succeed())
			opsSlice, _ := opsutil.GetOpsRequestSliceFromCluster(cluster)
			Expect(len(opsSlice)).Should(Equal(2))
			Expect(opsSlice[0].InQueue).Should(BeFalse())
			Expect(opsSlice[1].InQueue).Should(BeTrue())

			By("test enqueueOpsRequestToClusterAnnotation function with Reentry")
			opsBehaviour := opsManager.OpsMap[ops2.Spec.Type]
			_, _ = enqueueOpsRequestToClusterAnnotation(ctx, k8sClient, opsRes, opsBehaviour)
			Expect(k8sClient.Get(ctx, client.ObjectKeyFromObject(opsRes.Cluster), cluster)).Should(Succeed())
			opsSlice, _ = opsutil.GetOpsRequestSliceFromCluster(cluster)
			Expect(len(opsSlice)).Should(Equal(2))

			By("test DequeueOpsRequestInClusterAnnotation function when first opsRequest is Failed")
			// mock ops1 is Failed
			ops1.Status.Phase = appsv1alpha1.OpsFailedPhase
			opsRes.OpsRequest = ops1
			Expect(DequeueOpsRequestInClusterAnnotation(ctx, k8sClient, opsRes)).Should(Succeed())
			testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(ops2), func(g Gomega, ops *appsv1alpha1.OpsRequest) {
				// expect ops2 is Cancelled
				g.Expect(ops.Status.Phase).Should(Equal(appsv1alpha1.OpsCancelledPhase))
			})

			testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(cluster), func(g Gomega, cluster *appsv1alpha1.Cluster) {
				opsSlice, _ = opsutil.GetOpsRequestSliceFromCluster(cluster)
				// expect cluster's opsRequest queue is empty
				g.Expect(opsSlice).Should(BeEmpty())
			})
		})

		It("Test opsRequest dependency", func() {
			By("init operations resources ")
			reqCtx := intctrlutil.RequestCtx{Ctx: testCtx.Ctx}
			opsRes, _, _ := initOperationsResources(clusterDefinitionName, clusterVersionName, clusterName)

			By("create a first horizontal opsRequest")
			ops1 := createHorizontalScaling(clusterName, appsv1alpha1.HorizontalScaling{
				ComponentOps: appsv1alpha1.ComponentOps{ComponentName: consensusComp},
				ScaleIn: &appsv1alpha1.ScaleIn{
					ReplicaChanger: appsv1alpha1.ReplicaChanger{ReplicaChanges: pointer.Int32(1)},
				},
			})
			opsRes.OpsRequest = ops1
			_, err := GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(opsRes.OpsRequest.Status.Phase).Should(Equal(appsv1alpha1.OpsCreatingPhase))

			By("create another horizontal opsRequest with force flag and dependent the first opsRequest")
			ops2 := createHorizontalScaling(clusterName, appsv1alpha1.HorizontalScaling{
				ComponentOps: appsv1alpha1.ComponentOps{ComponentName: consensusComp},
				ScaleOut: &appsv1alpha1.ScaleOut{
					ReplicaChanger: appsv1alpha1.ReplicaChanger{ReplicaChanges: pointer.Int32(1)},
				},
			})
			ops2.Annotations = map[string]string{constant.OpsDependentOnSuccessfulOpsAnnoKey: ops1.Name}
			ops2.Spec.Force = true
			opsRes.OpsRequest = ops2
			_, err = GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(opsRes.OpsRequest.Status.Phase).Should(Equal(appsv1alpha1.OpsPendingPhase))
			// expect the dependent ops has been annotated
			Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(ops1), func(g Gomega, ops *appsv1alpha1.OpsRequest) {
				g.Expect(ops.Annotations[constant.RelatedOpsAnnotationKey]).Should(Equal(ops2.Name))
			})).Should(Succeed())

			By("expect for the ops is Creating when dependent ops is succeed")
			Expect(testapps.ChangeObjStatus(&testCtx, ops1, func() {
				ops1.Status.Phase = appsv1alpha1.OpsSucceedPhase
			})).Should(Succeed())

			_, err = GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(opsRes.OpsRequest.Status.Phase).Should(Equal(appsv1alpha1.OpsCreatingPhase))

			By("expect for the ops is Cancelled when dependent ops is Failed")
			Expect(testapps.ChangeObjStatus(&testCtx, ops1, func() {
				ops1.Status.Phase = appsv1alpha1.OpsFailedPhase
			})).Should(Succeed())

			ops2.Annotations = map[string]string{constant.OpsDependentOnSuccessfulOpsAnnoKey: ops1.Name}
			ops2.Status.Phase = appsv1alpha1.OpsPendingPhase
			_, err = GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(testapps.GetOpsRequestPhase(&testCtx, client.ObjectKeyFromObject(ops2))).Should(Equal(appsv1alpha1.OpsCancelledPhase))
		})
	})
})
