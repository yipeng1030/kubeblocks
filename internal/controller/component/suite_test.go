/*
Copyright ApeCloud, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package component

import (
	"context"
	"go/build"
	"path/filepath"
	"testing"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	dbaasv1alpha1 "github.com/apecloud/kubeblocks/apis/dbaas/v1alpha1"
	//+kubebuilder:scaffold:imports

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	ctrl "sigs.k8s.io/controller-runtime"

	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"

	dataprotectionv1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	intctrlutil "github.com/apecloud/kubeblocks/internal/controllerutil"
	"github.com/apecloud/kubeblocks/internal/testutil"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment
var ctx context.Context
var cancel context.CancelFunc
var testCtx testutil.TestContext
var logger logr.Logger

func init() {
	viper.AutomaticEnv()
	// viper.Set("ENABLE_DEBUG_LOG", "true")
}

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	if viper.GetBool("ENABLE_DEBUG_LOG") {
		logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true), func(o *zap.Options) {
			o.TimeEncoder = zapcore.ISO8601TimeEncoder
		}))
	}

	ctx, cancel = context.WithCancel(context.TODO())
	logger = logf.FromContext(ctx).WithValues()
	logger.Info("logger start")

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "..", "..", "config", "crd", "bases"),
			// use dependent external CRDs.
			// resolved by ref: https://github.com/operator-framework/operator-sdk/issues/4434#issuecomment-786794418
			filepath.Join(build.Default.GOPATH, "pkg", "mod", "github.com", "kubernetes-csi/external-snapshotter/",
				"client/v6@v6.0.1", "config", "crd")},
		ErrorIfCRDPathMissing: true,
	}

	var err error
	// cfg is defined in this file globally.
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = dbaasv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = dataprotectionv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = snapshotv1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	//+kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	// run reconcile
	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:                scheme.Scheme,
		MetricsBindAddress:    "0",
		ClientDisableCacheFor: intctrlutil.GetUncachedObjects(),
	})
	Expect(err).ToNot(HaveOccurred())

	testCtx = testutil.NewDefaultTestContext(ctx, k8sClient, testEnv)

	go func() {
		defer GinkgoRecover()
		err = k8sManager.Start(ctx)
		Expect(err).ToNot(HaveOccurred(), "failed to run manager")
	}()
})

var _ = AfterSuite(func() {
	cancel()
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})
