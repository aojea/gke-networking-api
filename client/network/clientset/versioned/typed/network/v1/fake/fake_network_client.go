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

package fake

import (
	v1 "github.com/GoogleCloudPlatform/gke-networking-api/client/network/clientset/versioned/typed/network/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1) GKENetworkParamSets() v1.GKENetworkParamSetInterface {
	return &FakeGKENetworkParamSets{c}
}

func (c *FakeNetworkingV1) GKENetworkParamSetLists() v1.GKENetworkParamSetListInterface {
	return &FakeGKENetworkParamSetLists{c}
}

func (c *FakeNetworkingV1) Networks() v1.NetworkInterface {
	return &FakeNetworks{c}
}

func (c *FakeNetworkingV1) NetworkInterfaces(namespace string) v1.NetworkInterfaceInterface {
	return &FakeNetworkInterfaces{c, namespace}
}

func (c *FakeNetworkingV1) NetworkInterfaceLists(namespace string) v1.NetworkInterfaceListInterface {
	return &FakeNetworkInterfaceLists{c, namespace}
}

func (c *FakeNetworkingV1) NetworkLists() v1.NetworkListInterface {
	return &FakeNetworkLists{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
