/*
Copyright 2022 EscherCloud.

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
	"context"

	v1alpha1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubernetesClusters implements KubernetesClusterInterface
type FakeKubernetesClusters struct {
	Fake *FakeUnikornV1alpha1
	ns   string
}

var kubernetesclustersResource = schema.GroupVersionResource{Group: "unikorn.eschercloud.ai", Version: "v1alpha1", Resource: "kubernetesclusters"}

var kubernetesclustersKind = schema.GroupVersionKind{Group: "unikorn.eschercloud.ai", Version: "v1alpha1", Kind: "KubernetesCluster"}

// Get takes name of the kubernetesCluster, and returns the corresponding kubernetesCluster object, and an error if there is any.
func (c *FakeKubernetesClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubernetesclustersResource, c.ns, name), &v1alpha1.KubernetesCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesCluster), err
}

// List takes label and field selectors, and returns the list of KubernetesClusters that match those selectors.
func (c *FakeKubernetesClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubernetesClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubernetesclustersResource, kubernetesclustersKind, c.ns, opts), &v1alpha1.KubernetesClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubernetesClusterList{ListMeta: obj.(*v1alpha1.KubernetesClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubernetesClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesClusters.
func (c *FakeKubernetesClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubernetesclustersResource, c.ns, opts))

}

// Create takes the representation of a kubernetesCluster and creates it.  Returns the server's representation of the kubernetesCluster, and an error, if there is any.
func (c *FakeKubernetesClusters) Create(ctx context.Context, kubernetesCluster *v1alpha1.KubernetesCluster, opts v1.CreateOptions) (result *v1alpha1.KubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubernetesclustersResource, c.ns, kubernetesCluster), &v1alpha1.KubernetesCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesCluster), err
}

// Update takes the representation of a kubernetesCluster and updates it. Returns the server's representation of the kubernetesCluster, and an error, if there is any.
func (c *FakeKubernetesClusters) Update(ctx context.Context, kubernetesCluster *v1alpha1.KubernetesCluster, opts v1.UpdateOptions) (result *v1alpha1.KubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubernetesclustersResource, c.ns, kubernetesCluster), &v1alpha1.KubernetesCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubernetesClusters) UpdateStatus(ctx context.Context, kubernetesCluster *v1alpha1.KubernetesCluster, opts v1.UpdateOptions) (*v1alpha1.KubernetesCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubernetesclustersResource, "status", c.ns, kubernetesCluster), &v1alpha1.KubernetesCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesCluster), err
}

// Delete takes name of the kubernetesCluster and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kubernetesclustersResource, c.ns, name, opts), &v1alpha1.KubernetesCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubernetesclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubernetesClusterList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesCluster.
func (c *FakeKubernetesClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubernetesCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubernetesclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.KubernetesCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesCluster), err
}