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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	volumesnapshotv1 "github.com/kubernetes-incubator/external-storage/snapshot/pkg/apis/volumesnapshot/v1"
	versioned "github.com/kubernetes-incubator/external-storage/snapshot/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubernetes-incubator/external-storage/snapshot/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/kubernetes-incubator/external-storage/snapshot/pkg/client/listers/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeSnapshotDataInformer provides access to a shared informer and lister for
// VolumeSnapshotDatas.
type VolumeSnapshotDataInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VolumeSnapshotDataLister
}

type volumeSnapshotDataInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVolumeSnapshotDataInformer constructs a new informer for VolumeSnapshotData type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeSnapshotDataInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotDataInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeSnapshotDataInformer constructs a new informer for VolumeSnapshotData type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeSnapshotDataInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumesnapshotV1().VolumeSnapshotDatas().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumesnapshotV1().VolumeSnapshotDatas().Watch(options)
			},
		},
		&volumesnapshotv1.VolumeSnapshotData{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeSnapshotDataInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotDataInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeSnapshotDataInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumesnapshotv1.VolumeSnapshotData{}, f.defaultInformer)
}

func (f *volumeSnapshotDataInformer) Lister() v1.VolumeSnapshotDataLister {
	return v1.NewVolumeSnapshotDataLister(f.Informer().GetIndexer())
}
