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

package v1alpha1

import (
	"context"
	time "time"

	armancomv1alpha1 "github.com/sheikh-arman/test-crd-controller/pkg/apis/arman.com/v1alpha1"
	versioned "github.com/sheikh-arman/test-crd-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/sheikh-arman/test-crd-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/sheikh-arman/test-crd-controller/pkg/client/listers/arman.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BlackmanInformer provides access to a shared informer and lister for
// Blackmans.
type BlackmanInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BlackmanLister
}

type blackmanInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlackmanInformer constructs a new informer for Blackman type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlackmanInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlackmanInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlackmanInformer constructs a new informer for Blackman type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlackmanInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArmanV1alpha1().Blackmans(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArmanV1alpha1().Blackmans(namespace).Watch(context.TODO(), options)
			},
		},
		&armancomv1alpha1.Blackman{},
		resyncPeriod,
		indexers,
	)
}

func (f *blackmanInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlackmanInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blackmanInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&armancomv1alpha1.Blackman{}, f.defaultInformer)
}

func (f *blackmanInformer) Lister() v1alpha1.BlackmanLister {
	return v1alpha1.NewBlackmanLister(f.Informer().GetIndexer())
}
