/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/iwilltry42/klum/pkg/apis/klum.cattle.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubeconfigLister helps list Kubeconfigs.
type KubeconfigLister interface {
	// List lists all Kubeconfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Kubeconfig, err error)
	// Get retrieves the Kubeconfig from the index for a given name.
	Get(name string) (*v1alpha1.Kubeconfig, error)
	KubeconfigListerExpansion
}

// kubeconfigLister implements the KubeconfigLister interface.
type kubeconfigLister struct {
	indexer cache.Indexer
}

// NewKubeconfigLister returns a new KubeconfigLister.
func NewKubeconfigLister(indexer cache.Indexer) KubeconfigLister {
	return &kubeconfigLister{indexer: indexer}
}

// List lists all Kubeconfigs in the indexer.
func (s *kubeconfigLister) List(selector labels.Selector) (ret []*v1alpha1.Kubeconfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Kubeconfig))
	})
	return ret, err
}

// Get retrieves the Kubeconfig from the index for a given name.
func (s *kubeconfigLister) Get(name string) (*v1alpha1.Kubeconfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubeconfig"), name)
	}
	return obj.(*v1alpha1.Kubeconfig), nil
}
