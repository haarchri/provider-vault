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

type TotpObservation struct {

	// Specifies the hashing algorithm used to generate the TOTP code. Options include 'SHA1', 'SHA256' and 'SHA512'.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The number of digits in the generated TOTP token. This value can either be 6 or 8.
	Digits *float64 `json:"digits,omitempty" tf:"digits,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the key's issuing organization.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Specifies the size in bytes of the generated key.
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Name of the MFA method.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The length of time used to generate a counter for the TOTP token calculation.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The pixel size of the generated square QR code.
	QrSize *float64 `json:"qrSize,omitempty" tf:"qr_size,omitempty"`

	// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
	Skew *float64 `json:"skew,omitempty" tf:"skew,omitempty"`
}

type TotpParameters struct {

	// Specifies the hashing algorithm used to generate the TOTP code. Options include 'SHA1', 'SHA256' and 'SHA512'.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The number of digits in the generated TOTP token. This value can either be 6 or 8.
	// +kubebuilder:validation:Optional
	Digits *float64 `json:"digits,omitempty" tf:"digits,omitempty"`

	// The name of the key's issuing organization.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Specifies the size in bytes of the generated key.
	// +kubebuilder:validation:Optional
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Name of the MFA method.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The length of time used to generate a counter for the TOTP token calculation.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The pixel size of the generated square QR code.
	// +kubebuilder:validation:Optional
	QrSize *float64 `json:"qrSize,omitempty" tf:"qr_size,omitempty"`

	// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
	// +kubebuilder:validation:Optional
	Skew *float64 `json:"skew,omitempty" tf:"skew,omitempty"`
}

// TotpSpec defines the desired state of Totp
type TotpSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TotpParameters `json:"forProvider"`
}

// TotpStatus defines the observed state of Totp.
type TotpStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TotpObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Totp is the Schema for the Totps API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Totp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.issuer)",message="issuer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   TotpSpec   `json:"spec"`
	Status TotpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TotpList contains a list of Totps
type TotpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Totp `json:"items"`
}

// Repository type metadata.
var (
	Totp_Kind             = "Totp"
	Totp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Totp_Kind}.String()
	Totp_KindAPIVersion   = Totp_Kind + "." + CRDGroupVersion.String()
	Totp_GroupVersionKind = CRDGroupVersion.WithKind(Totp_Kind)
)

func init() {
	SchemeBuilder.Register(&Totp{}, &TotpList{})
}
