// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/DataDog/datadog-operator/pkg/apis/datadoghq/v1alpha1"
	scheme "github.com/DataDog/datadog-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatadogAgentDeploymentsGetter has a method to return a DatadogAgentDeploymentInterface.
// A group's client should implement this interface.
type DatadogAgentDeploymentsGetter interface {
	DatadogAgentDeployments(namespace string) DatadogAgentDeploymentInterface
}

// DatadogAgentDeploymentInterface has methods to work with DatadogAgentDeployment resources.
type DatadogAgentDeploymentInterface interface {
	Create(*v1alpha1.DatadogAgentDeployment) (*v1alpha1.DatadogAgentDeployment, error)
	Update(*v1alpha1.DatadogAgentDeployment) (*v1alpha1.DatadogAgentDeployment, error)
	UpdateStatus(*v1alpha1.DatadogAgentDeployment) (*v1alpha1.DatadogAgentDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DatadogAgentDeployment, error)
	List(opts v1.ListOptions) (*v1alpha1.DatadogAgentDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatadogAgentDeployment, err error)
	DatadogAgentDeploymentExpansion
}

// datadogAgentDeployments implements DatadogAgentDeploymentInterface
type datadogAgentDeployments struct {
	client rest.Interface
	ns     string
}

// newDatadogAgentDeployments returns a DatadogAgentDeployments
func newDatadogAgentDeployments(c *DatadoghqV1alpha1Client, namespace string) *datadogAgentDeployments {
	return &datadogAgentDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the datadogAgentDeployment, and returns the corresponding datadogAgentDeployment object, and an error if there is any.
func (c *datadogAgentDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.DatadogAgentDeployment, err error) {
	result = &v1alpha1.DatadogAgentDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatadogAgentDeployments that match those selectors.
func (c *datadogAgentDeployments) List(opts v1.ListOptions) (result *v1alpha1.DatadogAgentDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatadogAgentDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datadogAgentDeployments.
func (c *datadogAgentDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a datadogAgentDeployment and creates it.  Returns the server's representation of the datadogAgentDeployment, and an error, if there is any.
func (c *datadogAgentDeployments) Create(datadogAgentDeployment *v1alpha1.DatadogAgentDeployment) (result *v1alpha1.DatadogAgentDeployment, err error) {
	result = &v1alpha1.DatadogAgentDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		Body(datadogAgentDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a datadogAgentDeployment and updates it. Returns the server's representation of the datadogAgentDeployment, and an error, if there is any.
func (c *datadogAgentDeployments) Update(datadogAgentDeployment *v1alpha1.DatadogAgentDeployment) (result *v1alpha1.DatadogAgentDeployment, err error) {
	result = &v1alpha1.DatadogAgentDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		Name(datadogAgentDeployment.Name).
		Body(datadogAgentDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *datadogAgentDeployments) UpdateStatus(datadogAgentDeployment *v1alpha1.DatadogAgentDeployment) (result *v1alpha1.DatadogAgentDeployment, err error) {
	result = &v1alpha1.DatadogAgentDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		Name(datadogAgentDeployment.Name).
		SubResource("status").
		Body(datadogAgentDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the datadogAgentDeployment and deletes it. Returns an error if one occurs.
func (c *datadogAgentDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datadogAgentDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched datadogAgentDeployment.
func (c *datadogAgentDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatadogAgentDeployment, err error) {
	result = &v1alpha1.DatadogAgentDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datadogagentdeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}