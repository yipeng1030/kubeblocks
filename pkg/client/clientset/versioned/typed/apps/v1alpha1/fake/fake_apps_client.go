/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/typed/apps/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAppsV1alpha1) BackupPolicyTemplates() v1alpha1.BackupPolicyTemplateInterface {
	return &FakeBackupPolicyTemplates{c}
}

func (c *FakeAppsV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeAppsV1alpha1) ClusterDefinitions() v1alpha1.ClusterDefinitionInterface {
	return &FakeClusterDefinitions{c}
}

func (c *FakeAppsV1alpha1) Components(namespace string) v1alpha1.ComponentInterface {
	return &FakeComponents{c, namespace}
}

func (c *FakeAppsV1alpha1) ComponentDefinitions() v1alpha1.ComponentDefinitionInterface {
	return &FakeComponentDefinitions{c}
}

func (c *FakeAppsV1alpha1) ComponentVersions() v1alpha1.ComponentVersionInterface {
	return &FakeComponentVersions{c}
}

func (c *FakeAppsV1alpha1) ConfigConstraints() v1alpha1.ConfigConstraintInterface {
	return &FakeConfigConstraints{c}
}

func (c *FakeAppsV1alpha1) OpsDefinitions() v1alpha1.OpsDefinitionInterface {
	return &FakeOpsDefinitions{c}
}

func (c *FakeAppsV1alpha1) OpsRequests(namespace string) v1alpha1.OpsRequestInterface {
	return &FakeOpsRequests{c, namespace}
}

func (c *FakeAppsV1alpha1) ServiceDescriptors(namespace string) v1alpha1.ServiceDescriptorInterface {
	return &FakeServiceDescriptors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
