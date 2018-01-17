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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	medusacontroller_v1alpha1 "github.com/danieloliveira079/medusa/pkg/apis/medusacontroller/v1alpha1"
	versioned "github.com/danieloliveira079/medusa/pkg/client/clientset/versioned"
	internalinterfaces "github.com/danieloliveira079/medusa/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/danieloliveira079/medusa/pkg/client/listers/medusacontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MedusaInformer provides access to a shared informer and lister for
// Medusas.
type MedusaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MedusaLister
}

type medusaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMedusaInformer constructs a new informer for Medusa type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMedusaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMedusaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMedusaInformer constructs a new informer for Medusa type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMedusaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MedusacontrollerV1alpha1().Medusas(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MedusacontrollerV1alpha1().Medusas(namespace).Watch(options)
			},
		},
		&medusacontroller_v1alpha1.Medusa{},
		resyncPeriod,
		indexers,
	)
}

func (f *medusaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMedusaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *medusaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&medusacontroller_v1alpha1.Medusa{}, f.defaultInformer)
}

func (f *medusaInformer) Lister() v1alpha1.MedusaLister {
	return v1alpha1.NewMedusaLister(f.Informer().GetIndexer())
}
