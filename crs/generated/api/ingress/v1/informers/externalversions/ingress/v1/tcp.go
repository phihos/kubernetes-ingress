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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	ingressv1 "github.com/haproxytech/kubernetes-ingress/crs/api/ingress/v1"
	versioned "github.com/haproxytech/kubernetes-ingress/crs/generated/api/ingress/v1/clientset/versioned"
	internalinterfaces "github.com/haproxytech/kubernetes-ingress/crs/generated/api/ingress/v1/informers/externalversions/internalinterfaces"
	v1 "github.com/haproxytech/kubernetes-ingress/crs/generated/api/ingress/v1/listers/ingress/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TCPInformer provides access to a shared informer and lister for
// TCPs.
type TCPInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TCPLister
}

type tCPInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTCPInformer constructs a new informer for TCP type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTCPInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTCPInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTCPInformer constructs a new informer for TCP type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTCPInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IngressV1().TCPs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IngressV1().TCPs(namespace).Watch(context.TODO(), options)
			},
		},
		&ingressv1.TCP{},
		resyncPeriod,
		indexers,
	)
}

func (f *tCPInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTCPInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tCPInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ingressv1.TCP{}, f.defaultInformer)
}

func (f *tCPInformer) Lister() v1.TCPLister {
	return v1.NewTCPLister(f.Informer().GetIndexer())
}
