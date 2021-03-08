/*
Copyright 2019 The OpenEBS Authors

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

	ndmv1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/ndm/v1alpha1"
	internalclientset "github.com/openebs/maya/pkg/client/generated/openebs.io/ndm/v1alpha1/clientset/internalclientset"
	internalinterfaces "github.com/openebs/maya/pkg/client/generated/openebs.io/ndm/v1alpha1/informer/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/maya/pkg/client/generated/openebs.io/ndm/v1alpha1/lister/ndm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BlockDeviceClaimInformer provides access to a shared informer and lister for
// BlockDeviceClaims.
type BlockDeviceClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BlockDeviceClaimLister
}

type blockDeviceClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlockDeviceClaimInformer constructs a new informer for BlockDeviceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlockDeviceClaimInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlockDeviceClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlockDeviceClaimInformer constructs a new informer for BlockDeviceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlockDeviceClaimInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().BlockDeviceClaims(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().BlockDeviceClaims(namespace).Watch(context.TODO(), options)
			},
		},
		&ndmv1alpha1.BlockDeviceClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *blockDeviceClaimInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlockDeviceClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blockDeviceClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ndmv1alpha1.BlockDeviceClaim{}, f.defaultInformer)
}

func (f *blockDeviceClaimInformer) Lister() v1alpha1.BlockDeviceClaimLister {
	return v1alpha1.NewBlockDeviceClaimLister(f.Informer().GetIndexer())
}
