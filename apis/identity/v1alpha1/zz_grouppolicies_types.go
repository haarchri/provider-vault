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

type GroupPoliciesInitParameters struct {

	// Defaults to true.
	// Should the resource manage policies exclusively? Beware of race conditions when disabling exclusive management
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// Group ID to assign policies to.
	// ID of the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of policies to assign to the group
	// Policies to be tied to the group.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type GroupPoliciesObservation struct {

	// Defaults to true.
	// Should the resource manage policies exclusively? Beware of race conditions when disabling exclusive management
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// Group ID to assign policies to.
	// ID of the group.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The name of the group that are assigned the policies.
	// Name of the group.
	GroupName *string `json:"groupName,omitempty" tf:"group_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of policies to assign to the group
	// Policies to be tied to the group.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type GroupPoliciesParameters struct {

	// Defaults to true.
	// Should the resource manage policies exclusively? Beware of race conditions when disabling exclusive management
	// +kubebuilder:validation:Optional
	Exclusive *bool `json:"exclusive,omitempty" tf:"exclusive,omitempty"`

	// Group ID to assign policies to.
	// ID of the group.
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// List of policies to assign to the group
	// Policies to be tied to the group.
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// GroupPoliciesSpec defines the desired state of GroupPolicies
type GroupPoliciesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupPoliciesParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider GroupPoliciesInitParameters `json:"initProvider,omitempty"`
}

// GroupPoliciesStatus defines the observed state of GroupPolicies.
type GroupPoliciesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupPoliciesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicies is the Schema for the GroupPoliciess API. Manages policies for an Identity Group for Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type GroupPolicies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groupId) || has(self.initProvider.groupId)",message="groupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policies) || has(self.initProvider.policies)",message="policies is a required parameter"
	Spec   GroupPoliciesSpec   `json:"spec"`
	Status GroupPoliciesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPoliciesList contains a list of GroupPoliciess
type GroupPoliciesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupPolicies `json:"items"`
}

// Repository type metadata.
var (
	GroupPolicies_Kind             = "GroupPolicies"
	GroupPolicies_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupPolicies_Kind}.String()
	GroupPolicies_KindAPIVersion   = GroupPolicies_Kind + "." + CRDGroupVersion.String()
	GroupPolicies_GroupVersionKind = CRDGroupVersion.WithKind(GroupPolicies_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupPolicies{}, &GroupPoliciesList{})
}
