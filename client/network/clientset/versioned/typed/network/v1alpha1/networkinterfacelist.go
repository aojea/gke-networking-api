/*
Copyright 2024 Google LLC
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/network/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/gke-networking-api/client/network/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// NetworkInterfaceListsGetter has a method to return a NetworkInterfaceListInterface.
// A group's client should implement this interface.
type NetworkInterfaceListsGetter interface {
	NetworkInterfaceLists(namespace string) NetworkInterfaceListInterface
}

// NetworkInterfaceListInterface has methods to work with NetworkInterfaceList resources.
type NetworkInterfaceListInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NetworkInterfaceList, error)
	NetworkInterfaceListExpansion
}

// networkInterfaceLists implements NetworkInterfaceListInterface
type networkInterfaceLists struct {
	client rest.Interface
	ns     string
}

// newNetworkInterfaceLists returns a NetworkInterfaceLists
func newNetworkInterfaceLists(c *NetworkingV1alpha1Client, namespace string) *networkInterfaceLists {
	return &networkInterfaceLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkInterfaceList, and returns the corresponding networkInterfaceList object, and an error if there is any.
func (c *networkInterfaceLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceList, err error) {
	result = &v1alpha1.NetworkInterfaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacelists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}
