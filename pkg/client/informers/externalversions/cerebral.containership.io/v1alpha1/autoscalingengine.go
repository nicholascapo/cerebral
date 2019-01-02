/*
Copyright 2019 The Kubernetes Authors.

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
	time "time"

	cerebralcontainershipiov1alpha1 "github.com/containership/cerebral/pkg/apis/cerebral.containership.io/v1alpha1"
	versioned "github.com/containership/cerebral/pkg/client/clientset/versioned"
	internalinterfaces "github.com/containership/cerebral/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/containership/cerebral/pkg/client/listers/cerebral.containership.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AutoscalingEngineInformer provides access to a shared informer and lister for
// AutoscalingEngines.
type AutoscalingEngineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AutoscalingEngineLister
}

type autoscalingEngineInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAutoscalingEngineInformer constructs a new informer for AutoscalingEngine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAutoscalingEngineInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAutoscalingEngineInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAutoscalingEngineInformer constructs a new informer for AutoscalingEngine type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAutoscalingEngineInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CerebralV1alpha1().AutoscalingEngines().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CerebralV1alpha1().AutoscalingEngines().Watch(options)
			},
		},
		&cerebralcontainershipiov1alpha1.AutoscalingEngine{},
		resyncPeriod,
		indexers,
	)
}

func (f *autoscalingEngineInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAutoscalingEngineInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *autoscalingEngineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cerebralcontainershipiov1alpha1.AutoscalingEngine{}, f.defaultInformer)
}

func (f *autoscalingEngineInformer) Lister() v1alpha1.AutoscalingEngineLister {
	return v1alpha1.NewAutoscalingEngineLister(f.Informer().GetIndexer())
}
