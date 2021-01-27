package crd

import (
	"context"

	"github.com/iwilltry42/klum/pkg/apis/klum.cattle.io/v1alpha1"
	"github.com/rancher/norman/v2/pkg/openapi"
	"github.com/rancher/wrangler/pkg/crd"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/rest"
)

func Create(ctx context.Context, config *rest.Config) error {
	factory, err := crd.NewFactoryFromClient(config)
	if err != nil {
		return err
	}

	return factory.BatchCreateCRDs(ctx,
		newCRD("User.klum.cattle.io/v1alpha1", v1alpha1.User{}),
		newCRD("Kubeconfig.klum.cattle.io/v1alpha1", v1alpha1.Kubeconfig{})).BatchWait()
}

func newCRD(name string, obj interface{}) crd.CRD {
	return crd.NonNamespacedType(name).
		WithStatus().
		WithSchema(mustSchema(obj))
}

func mustSchema(obj interface{}) *v1beta1.JSONSchemaProps {
	result, err := openapi.ToOpenAPIFromStruct(obj)
	if err != nil {
		panic(err)
	}
	return result
}
