// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSM":       schema_pkg_apis_nsm_v1alpha1_NSM(ref),
		"github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSMSpec":   schema_pkg_apis_nsm_v1alpha1_NSMSpec(ref),
		"github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSMStatus": schema_pkg_apis_nsm_v1alpha1_NSMStatus(ref),
	}
}

func schema_pkg_apis_nsm_v1alpha1_NSM(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSM is the Schema for the nsms API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSMSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSMStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSMSpec", "github.com/acmenezes/nsm-operator/pkg/apis/nsm/v1alpha1.NSMStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nsm_v1alpha1_NSMSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSMSpec defines the desired state of NSM",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_nsm_v1alpha1_NSMStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NSMStatus defines the observed state of NSM",
				Type:        []string{"object"},
			},
		},
	}
}
