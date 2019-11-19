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

package v2

import (
	time "time"

	convoxv2 "github.com/convox/rack/pkg/atom/pkg/apis/convox/v2"
	versioned "github.com/convox/rack/pkg/atom/pkg/client/clientset/versioned"
	internalinterfaces "github.com/convox/rack/pkg/atom/pkg/client/informers/externalversions/internalinterfaces"
	v2 "github.com/convox/rack/pkg/atom/pkg/client/listers/convox/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AtomInformer provides access to a shared informer and lister for
// Atoms.
type AtomInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.AtomLister
}

type atomInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAtomInformer constructs a new informer for Atom type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAtomInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAtomInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAtomInformer constructs a new informer for Atom type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAtomInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConvoxV2().Atoms(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConvoxV2().Atoms(namespace).Watch(options)
			},
		},
		&convoxv2.Atom{},
		resyncPeriod,
		indexers,
	)
}

func (f *atomInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAtomInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *atomInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&convoxv2.Atom{}, f.defaultInformer)
}

func (f *atomInformer) Lister() v2.AtomLister {
	return v2.NewAtomLister(f.Informer().GetIndexer())
}
