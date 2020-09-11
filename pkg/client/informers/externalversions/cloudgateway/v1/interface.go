/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// EGateways returns a EGatewayInformer.
	EGateways() EGatewayInformer
	// EServices returns a EServiceInformer.
	EServices() EServiceInformer
	// ESites returns a ESiteInformer.
	ESites() ESiteInformer
	// ServiceExposes returns a ServiceExposeInformer.
	ServiceExposes() ServiceExposeInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	tenant           string
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, tenant: "system", namespace: namespace, tweakListOptions: tweakListOptions}
}

func NewWithMultiTenancy(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc, tenant string) Interface {

	return &version{factory: f, tenant: tenant, namespace: namespace, tweakListOptions: tweakListOptions}
}

// EGateways returns a EGatewayInformer.
func (v *version) EGateways() EGatewayInformer {
	return &eGatewayInformer{factory: v.factory, namespace: v.namespace, tenant: v.tenant, tweakListOptions: v.tweakListOptions}
}

// EServices returns a EServiceInformer.
func (v *version) EServices() EServiceInformer {
	return &eServiceInformer{factory: v.factory, namespace: v.namespace, tenant: v.tenant, tweakListOptions: v.tweakListOptions}
}

// ESites returns a ESiteInformer.
func (v *version) ESites() ESiteInformer {
	return &eSiteInformer{factory: v.factory, namespace: v.namespace, tenant: v.tenant, tweakListOptions: v.tweakListOptions}
}

// ServiceExposes returns a ServiceExposeInformer.
func (v *version) ServiceExposes() ServiceExposeInformer {
	return &serviceExposeInformer{factory: v.factory, namespace: v.namespace, tenant: v.tenant, tweakListOptions: v.tweakListOptions}
}
