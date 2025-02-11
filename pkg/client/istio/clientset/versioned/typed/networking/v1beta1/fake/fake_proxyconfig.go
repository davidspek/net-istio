/*
Copyright 2020 The Knative Authors

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

	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProxyConfigs implements ProxyConfigInterface
type FakeProxyConfigs struct {
	Fake *FakeNetworkingV1beta1
	ns   string
}

var proxyconfigsResource = schema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "proxyconfigs"}

var proxyconfigsKind = schema.GroupVersionKind{Group: "networking.istio.io", Version: "v1beta1", Kind: "ProxyConfig"}

// Get takes name of the proxyConfig, and returns the corresponding proxyConfig object, and an error if there is any.
func (c *FakeProxyConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ProxyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(proxyconfigsResource, c.ns, name), &v1beta1.ProxyConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ProxyConfig), err
}

// List takes label and field selectors, and returns the list of ProxyConfigs that match those selectors.
func (c *FakeProxyConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ProxyConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(proxyconfigsResource, proxyconfigsKind, c.ns, opts), &v1beta1.ProxyConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ProxyConfigList{ListMeta: obj.(*v1beta1.ProxyConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.ProxyConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested proxyConfigs.
func (c *FakeProxyConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(proxyconfigsResource, c.ns, opts))

}

// Create takes the representation of a proxyConfig and creates it.  Returns the server's representation of the proxyConfig, and an error, if there is any.
func (c *FakeProxyConfigs) Create(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.CreateOptions) (result *v1beta1.ProxyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(proxyconfigsResource, c.ns, proxyConfig), &v1beta1.ProxyConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ProxyConfig), err
}

// Update takes the representation of a proxyConfig and updates it. Returns the server's representation of the proxyConfig, and an error, if there is any.
func (c *FakeProxyConfigs) Update(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.UpdateOptions) (result *v1beta1.ProxyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(proxyconfigsResource, c.ns, proxyConfig), &v1beta1.ProxyConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ProxyConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProxyConfigs) UpdateStatus(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.UpdateOptions) (*v1beta1.ProxyConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(proxyconfigsResource, "status", c.ns, proxyConfig), &v1beta1.ProxyConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ProxyConfig), err
}

// Delete takes name of the proxyConfig and deletes it. Returns an error if one occurs.
func (c *FakeProxyConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(proxyconfigsResource, c.ns, name, opts), &v1beta1.ProxyConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProxyConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(proxyconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ProxyConfigList{})
	return err
}

// Patch applies the patch and returns the patched proxyConfig.
func (c *FakeProxyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ProxyConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(proxyconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.ProxyConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ProxyConfig), err
}
