/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/danieloliveira079/medusa/pkg/apis/medusacontroller/v1alpha1"
	scheme "github.com/danieloliveira079/medusa/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MedusasGetter has a method to return a MedusaInterface.
// A group's client should implement this interface.
type MedusasGetter interface {
	Medusas(namespace string) MedusaInterface
}

// MedusaInterface has methods to work with Medusa resources.
type MedusaInterface interface {
	Create(*v1alpha1.Medusa) (*v1alpha1.Medusa, error)
	Update(*v1alpha1.Medusa) (*v1alpha1.Medusa, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Medusa, error)
	List(opts v1.ListOptions) (*v1alpha1.MedusaList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Medusa, err error)
	MedusaExpansion
}

// medusas implements MedusaInterface
type medusas struct {
	client rest.Interface
	ns     string
}

// newMedusas returns a Medusas
func newMedusas(c *MedusacontrollerV1alpha1Client, namespace string) *medusas {
	return &medusas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the medusa, and returns the corresponding medusa object, and an error if there is any.
func (c *medusas) Get(name string, options v1.GetOptions) (result *v1alpha1.Medusa, err error) {
	result = &v1alpha1.Medusa{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("medusas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Medusas that match those selectors.
func (c *medusas) List(opts v1.ListOptions) (result *v1alpha1.MedusaList, err error) {
	result = &v1alpha1.MedusaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("medusas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested medusas.
func (c *medusas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("medusas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a medusa and creates it.  Returns the server's representation of the medusa, and an error, if there is any.
func (c *medusas) Create(medusa *v1alpha1.Medusa) (result *v1alpha1.Medusa, err error) {
	result = &v1alpha1.Medusa{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("medusas").
		Body(medusa).
		Do().
		Into(result)
	return
}

// Update takes the representation of a medusa and updates it. Returns the server's representation of the medusa, and an error, if there is any.
func (c *medusas) Update(medusa *v1alpha1.Medusa) (result *v1alpha1.Medusa, err error) {
	result = &v1alpha1.Medusa{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("medusas").
		Name(medusa.Name).
		Body(medusa).
		Do().
		Into(result)
	return
}

// Delete takes name of the medusa and deletes it. Returns an error if one occurs.
func (c *medusas) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("medusas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *medusas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("medusas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched medusa.
func (c *medusas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Medusa, err error) {
	result = &v1alpha1.Medusa{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("medusas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
