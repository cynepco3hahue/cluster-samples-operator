// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/cluster-samples-operator/pkg/apis/samplesresource/v1alpha1"
	scheme "github.com/openshift/cluster-samples-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SamplesResourcesGetter has a method to return a SamplesResourceInterface.
// A group's client should implement this interface.
type SamplesResourcesGetter interface {
	SamplesResources() SamplesResourceInterface
}

// SamplesResourceInterface has methods to work with SamplesResource resources.
type SamplesResourceInterface interface {
	Create(*v1alpha1.SamplesResource) (*v1alpha1.SamplesResource, error)
	Update(*v1alpha1.SamplesResource) (*v1alpha1.SamplesResource, error)
	UpdateStatus(*v1alpha1.SamplesResource) (*v1alpha1.SamplesResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SamplesResource, error)
	List(opts v1.ListOptions) (*v1alpha1.SamplesResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SamplesResource, err error)
	SamplesResourceExpansion
}

// samplesResources implements SamplesResourceInterface
type samplesResources struct {
	client rest.Interface
}

// newSamplesResources returns a SamplesResources
func newSamplesResources(c *SamplesresourceV1alpha1Client) *samplesResources {
	return &samplesResources{
		client: c.RESTClient(),
	}
}

// Get takes name of the samplesResource, and returns the corresponding samplesResource object, and an error if there is any.
func (c *samplesResources) Get(name string, options v1.GetOptions) (result *v1alpha1.SamplesResource, err error) {
	result = &v1alpha1.SamplesResource{}
	err = c.client.Get().
		Resource("samplesresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SamplesResources that match those selectors.
func (c *samplesResources) List(opts v1.ListOptions) (result *v1alpha1.SamplesResourceList, err error) {
	result = &v1alpha1.SamplesResourceList{}
	err = c.client.Get().
		Resource("samplesresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested samplesResources.
func (c *samplesResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("samplesresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a samplesResource and creates it.  Returns the server's representation of the samplesResource, and an error, if there is any.
func (c *samplesResources) Create(samplesResource *v1alpha1.SamplesResource) (result *v1alpha1.SamplesResource, err error) {
	result = &v1alpha1.SamplesResource{}
	err = c.client.Post().
		Resource("samplesresources").
		Body(samplesResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a samplesResource and updates it. Returns the server's representation of the samplesResource, and an error, if there is any.
func (c *samplesResources) Update(samplesResource *v1alpha1.SamplesResource) (result *v1alpha1.SamplesResource, err error) {
	result = &v1alpha1.SamplesResource{}
	err = c.client.Put().
		Resource("samplesresources").
		Name(samplesResource.Name).
		Body(samplesResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *samplesResources) UpdateStatus(samplesResource *v1alpha1.SamplesResource) (result *v1alpha1.SamplesResource, err error) {
	result = &v1alpha1.SamplesResource{}
	err = c.client.Put().
		Resource("samplesresources").
		Name(samplesResource.Name).
		SubResource("status").
		Body(samplesResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the samplesResource and deletes it. Returns an error if one occurs.
func (c *samplesResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("samplesresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *samplesResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("samplesresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched samplesResource.
func (c *samplesResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SamplesResource, err error) {
	result = &v1alpha1.SamplesResource{}
	err = c.client.Patch(pt).
		Resource("samplesresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}