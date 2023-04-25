/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OidcScopeObservation struct {

	// The scope's description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the scope. The openid scope name is reserved.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type OidcScopeParameters struct {

	// The scope's description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the scope. The openid scope name is reserved.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

// OidcScopeSpec defines the desired state of OidcScope
type OidcScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OidcScopeParameters `json:"forProvider"`
}

// OidcScopeStatus defines the observed state of OidcScope.
type OidcScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OidcScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OidcScope is the Schema for the OidcScopes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type OidcScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   OidcScopeSpec   `json:"spec"`
	Status OidcScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OidcScopeList contains a list of OidcScopes
type OidcScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OidcScope `json:"items"`
}

// Repository type metadata.
var (
	OidcScope_Kind             = "OidcScope"
	OidcScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OidcScope_Kind}.String()
	OidcScope_KindAPIVersion   = OidcScope_Kind + "." + CRDGroupVersion.String()
	OidcScope_GroupVersionKind = CRDGroupVersion.WithKind(OidcScope_Kind)
)

func init() {
	SchemeBuilder.Register(&OidcScope{}, &OidcScopeList{})
}
