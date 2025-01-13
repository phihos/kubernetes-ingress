//
// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/haproxytech/kubernetes-ingress/crs/api/ingress/v3"
	scheme "github.com/haproxytech/kubernetes-ingress/crs/generated/api/ingress/v3/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TCPsGetter has a method to return a TCPInterface.
// A group's client should implement this interface.
type TCPsGetter interface {
	TCPs(namespace string) TCPInterface
}

// TCPInterface has methods to work with TCP resources.
type TCPInterface interface {
	Create(ctx context.Context, tCP *v3.TCP, opts v1.CreateOptions) (*v3.TCP, error)
	Update(ctx context.Context, tCP *v3.TCP, opts v1.UpdateOptions) (*v3.TCP, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.TCP, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.TCPList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.TCP, err error)
	TCPExpansion
}

// tCPs implements TCPInterface
type tCPs struct {
	client rest.Interface
	ns     string
}

// newTCPs returns a TCPs
func newTCPs(c *IngressV3Client, namespace string) *tCPs {
	return &tCPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tCP, and returns the corresponding tCP object, and an error if there is any.
func (c *tCPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.TCP, err error) {
	result = &v3.TCP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tcps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TCPs that match those selectors.
func (c *tCPs) List(ctx context.Context, opts v1.ListOptions) (result *v3.TCPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.TCPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tcps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tCPs.
func (c *tCPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tcps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tCP and creates it.  Returns the server's representation of the tCP, and an error, if there is any.
func (c *tCPs) Create(ctx context.Context, tCP *v3.TCP, opts v1.CreateOptions) (result *v3.TCP, err error) {
	result = &v3.TCP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tcps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tCP).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tCP and updates it. Returns the server's representation of the tCP, and an error, if there is any.
func (c *tCPs) Update(ctx context.Context, tCP *v3.TCP, opts v1.UpdateOptions) (result *v3.TCP, err error) {
	result = &v3.TCP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tcps").
		Name(tCP.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tCP).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tCP and deletes it. Returns an error if one occurs.
func (c *tCPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tcps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tCPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tcps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tCP.
func (c *tCPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.TCP, err error) {
	result = &v3.TCP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tcps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
