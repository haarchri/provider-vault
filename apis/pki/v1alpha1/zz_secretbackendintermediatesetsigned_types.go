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

type SecretBackendIntermediateSetSignedObservation struct {

	// The PKI secret backend the resource belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The certificate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SecretBackendIntermediateSetSignedParameters struct {

	// The PKI secret backend the resource belongs to.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// The certificate.
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// SecretBackendIntermediateSetSignedSpec defines the desired state of SecretBackendIntermediateSetSigned
type SecretBackendIntermediateSetSignedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendIntermediateSetSignedParameters `json:"forProvider"`
}

// SecretBackendIntermediateSetSignedStatus defines the observed state of SecretBackendIntermediateSetSigned.
type SecretBackendIntermediateSetSignedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendIntermediateSetSignedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendIntermediateSetSigned is the Schema for the SecretBackendIntermediateSetSigneds API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendIntermediateSetSigned struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.certificate)",message="certificate is a required parameter"
	Spec   SecretBackendIntermediateSetSignedSpec   `json:"spec"`
	Status SecretBackendIntermediateSetSignedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendIntermediateSetSignedList contains a list of SecretBackendIntermediateSetSigneds
type SecretBackendIntermediateSetSignedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendIntermediateSetSigned `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendIntermediateSetSigned_Kind             = "SecretBackendIntermediateSetSigned"
	SecretBackendIntermediateSetSigned_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendIntermediateSetSigned_Kind}.String()
	SecretBackendIntermediateSetSigned_KindAPIVersion   = SecretBackendIntermediateSetSigned_Kind + "." + CRDGroupVersion.String()
	SecretBackendIntermediateSetSigned_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendIntermediateSetSigned_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendIntermediateSetSigned{}, &SecretBackendIntermediateSetSignedList{})
}
