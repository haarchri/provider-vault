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

type SecretBackendRoleObservation struct {

	// The path of the AWS Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Role credential type.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	DefaultStsTTL *float64 `json:"defaultStsTtl,omitempty" tf:"default_sts_ttl,omitempty"`

	// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	IAMGroups []*string `json:"iamGroups,omitempty" tf:"iam_groups,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is one of assumed_role or federation_token.
	MaxStsTTL *float64 `json:"maxStsTtl,omitempty" tf:"max_sts_ttl,omitempty"`

	// Unique name for the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
	PermissionsBoundaryArn *string `json:"permissionsBoundaryArn,omitempty" tf:"permissions_boundary_arn,omitempty"`

	// ARN for an existing IAM policy the role should use.
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// IAM policy the role should use in JSON format.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is 'assumed_role'
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// The path for the user name. Valid only when credential_type is iam_user. Default is /
	UserPath *string `json:"userPath,omitempty" tf:"user_path,omitempty"`
}

type SecretBackendRoleParameters struct {

	// The path of the AWS Secret Backend the role belongs to.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Role credential type.
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	// +kubebuilder:validation:Optional
	DefaultStsTTL *float64 `json:"defaultStsTtl,omitempty" tf:"default_sts_ttl,omitempty"`

	// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	// +kubebuilder:validation:Optional
	IAMGroups []*string `json:"iamGroups,omitempty" tf:"iam_groups,omitempty"`

	// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is one of assumed_role or federation_token.
	// +kubebuilder:validation:Optional
	MaxStsTTL *float64 `json:"maxStsTtl,omitempty" tf:"max_sts_ttl,omitempty"`

	// Unique name for the role.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
	// +kubebuilder:validation:Optional
	PermissionsBoundaryArn *string `json:"permissionsBoundaryArn,omitempty" tf:"permissions_boundary_arn,omitempty"`

	// ARN for an existing IAM policy the role should use.
	// +kubebuilder:validation:Optional
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// IAM policy the role should use in JSON format.
	// +kubebuilder:validation:Optional
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is 'assumed_role'
	// +kubebuilder:validation:Optional
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// The path for the user name. Valid only when credential_type is iam_user. Default is /
	// +kubebuilder:validation:Optional
	UserPath *string `json:"userPath,omitempty" tf:"user_path,omitempty"`
}

// SecretBackendRoleSpec defines the desired state of SecretBackendRole
type SecretBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendRoleParameters `json:"forProvider"`
}

// SecretBackendRoleStatus defines the observed state of SecretBackendRole.
type SecretBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRole is the Schema for the SecretBackendRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.credentialType)",message="credentialType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   SecretBackendRoleSpec   `json:"spec"`
	Status SecretBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRoleList contains a list of SecretBackendRoles
type SecretBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendRole `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendRole_Kind             = "SecretBackendRole"
	SecretBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendRole_Kind}.String()
	SecretBackendRole_KindAPIVersion   = SecretBackendRole_Kind + "." + CRDGroupVersion.String()
	SecretBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendRole{}, &SecretBackendRoleList{})
}
