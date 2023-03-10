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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/longhorn/apis/longhorn/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupTargetLister helps list BackupTargets.
type BackupTargetLister interface {
	// List lists all BackupTargets in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.BackupTarget, err error)
	// BackupTargets returns an object that can list and get BackupTargets.
	BackupTargets(namespace string) BackupTargetNamespaceLister
	BackupTargetListerExpansion
}

// backupTargetLister implements the BackupTargetLister interface.
type backupTargetLister struct {
	indexer cache.Indexer
}

// NewBackupTargetLister returns a new BackupTargetLister.
func NewBackupTargetLister(indexer cache.Indexer) BackupTargetLister {
	return &backupTargetLister{indexer: indexer}
}

// List lists all BackupTargets in the indexer.
func (s *backupTargetLister) List(selector labels.Selector) (ret []*v1beta1.BackupTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackupTarget))
	})
	return ret, err
}

// BackupTargets returns an object that can list and get BackupTargets.
func (s *backupTargetLister) BackupTargets(namespace string) BackupTargetNamespaceLister {
	return backupTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupTargetNamespaceLister helps list and get BackupTargets.
type BackupTargetNamespaceLister interface {
	// List lists all BackupTargets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.BackupTarget, err error)
	// Get retrieves the BackupTarget from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.BackupTarget, error)
	BackupTargetNamespaceListerExpansion
}

// backupTargetNamespaceLister implements the BackupTargetNamespaceLister
// interface.
type backupTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackupTargets in the indexer for a given namespace.
func (s backupTargetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.BackupTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackupTarget))
	})
	return ret, err
}

// Get retrieves the BackupTarget from the indexer for a given namespace and name.
func (s backupTargetNamespaceLister) Get(name string) (*v1beta1.BackupTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("backuptarget"), name)
	}
	return obj.(*v1beta1.BackupTarget), nil
}
