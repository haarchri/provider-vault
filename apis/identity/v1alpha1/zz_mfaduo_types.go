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

type MfaDuoObservation struct {

	// API hostname for Duo
	APIHostname *string `json:"apiHostname,omitempty" tf:"api_hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Method ID.
	MethodID *string `json:"methodId,omitempty" tf:"method_id,omitempty"`

	// Mount accessor.
	MountAccessor *string `json:"mountAccessor,omitempty" tf:"mount_accessor,omitempty"`

	// Method name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Method's namespace ID.
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Method's namespace path.
	NamespacePath *string `json:"namespacePath,omitempty" tf:"namespace_path,omitempty"`

	// Push information for Duo.
	PushInfo *string `json:"pushInfo,omitempty" tf:"push_info,omitempty"`

	// MFA type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Resource UUID.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// Require passcode upon MFA validation.
	UsePasscode *bool `json:"usePasscode,omitempty" tf:"use_passcode,omitempty"`

	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

type MfaDuoParameters struct {

	// API hostname for Duo
	// +kubebuilder:validation:Optional
	APIHostname *string `json:"apiHostname,omitempty" tf:"api_hostname,omitempty"`

	// Integration key for Duo
	// +kubebuilder:validation:Optional
	IntegrationKeySecretRef v1.SecretKeySelector `json:"integrationKeySecretRef" tf:"-"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Push information for Duo.
	// +kubebuilder:validation:Optional
	PushInfo *string `json:"pushInfo,omitempty" tf:"push_info,omitempty"`

	// Secret key for Duo
	// +kubebuilder:validation:Optional
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`

	// Require passcode upon MFA validation.
	// +kubebuilder:validation:Optional
	UsePasscode *bool `json:"usePasscode,omitempty" tf:"use_passcode,omitempty"`

	// A template string for mapping Identity names to MFA methods.
	// +kubebuilder:validation:Optional
	UsernameFormat *string `json:"usernameFormat,omitempty" tf:"username_format,omitempty"`
}

// MfaDuoSpec defines the desired state of MfaDuo
type MfaDuoSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MfaDuoParameters `json:"forProvider"`
}

// MfaDuoStatus defines the observed state of MfaDuo.
type MfaDuoStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MfaDuoObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MfaDuo is the Schema for the MfaDuos API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type MfaDuo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.apiHostname)",message="apiHostname is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.integrationKeySecretRef)",message="integrationKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.secretKeySecretRef)",message="secretKeySecretRef is a required parameter"
	Spec   MfaDuoSpec   `json:"spec"`
	Status MfaDuoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MfaDuoList contains a list of MfaDuos
type MfaDuoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MfaDuo `json:"items"`
}

// Repository type metadata.
var (
	MfaDuo_Kind             = "MfaDuo"
	MfaDuo_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MfaDuo_Kind}.String()
	MfaDuo_KindAPIVersion   = MfaDuo_Kind + "." + CRDGroupVersion.String()
	MfaDuo_GroupVersionKind = CRDGroupVersion.WithKind(MfaDuo_Kind)
)

func init() {
	SchemeBuilder.Register(&MfaDuo{}, &MfaDuoList{})
}
