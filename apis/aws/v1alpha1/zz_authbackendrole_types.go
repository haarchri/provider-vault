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

type AuthBackendRoleObservation struct {

	// When true, allows migration of the underlying instance where the client resides. Use with caution.
	AllowInstanceMigration *bool `json:"allowInstanceMigration,omitempty" tf:"allow_instance_migration,omitempty"`

	// The auth type permitted for this role.
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// Unique name of the auth backend to configure.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Only EC2 instances using this AMI ID will be permitted to log in.
	BoundAMIIds []*string `json:"boundAmiIds,omitempty" tf:"bound_ami_ids,omitempty"`

	// Only EC2 instances with this account ID in their identity document will be permitted to log in.
	BoundAccountIds []*string `json:"boundAccountIds,omitempty" tf:"bound_account_ids,omitempty"`

	// Only EC2 instances that match this instance ID will be permitted to log in.
	BoundEC2InstanceIds []*string `json:"boundEc2InstanceIds,omitempty" tf:"bound_ec2_instance_ids,omitempty"`

	// Only EC2 instances associated with an IAM instance profile ARN that matches this value will be permitted to log in.
	BoundIAMInstanceProfileArns []*string `json:"boundIamInstanceProfileArns,omitempty" tf:"bound_iam_instance_profile_arns,omitempty"`

	// The IAM principal that must be authenticated using the iam auth method.
	BoundIAMPrincipalArns []*string `json:"boundIamPrincipalArns,omitempty" tf:"bound_iam_principal_arns,omitempty"`

	// Only EC2 instances that match this IAM role ARN will be permitted to log in.
	BoundIAMRoleArns []*string `json:"boundIamRoleArns,omitempty" tf:"bound_iam_role_arns,omitempty"`

	// Only EC2 instances in this region will be permitted to log in.
	BoundRegions []*string `json:"boundRegions,omitempty" tf:"bound_regions,omitempty"`

	// Only EC2 instances associated with this subnet ID will be permitted to log in.
	BoundSubnetIds []*string `json:"boundSubnetIds,omitempty" tf:"bound_subnet_ids,omitempty"`

	// Only EC2 instances associated with this VPC ID will be permitted to log in.
	BoundVPCIds []*string `json:"boundVpcIds,omitempty" tf:"bound_vpc_ids,omitempty"`

	// When true, only allows a single token to be granted per instance ID.
	DisallowReauthentication *bool `json:"disallowReauthentication,omitempty" tf:"disallow_reauthentication,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The region to search for the inferred entities in.
	InferredAwsRegion *string `json:"inferredAwsRegion,omitempty" tf:"inferred_aws_region,omitempty"`

	// The type of inferencing Vault should do.
	InferredEntityType *string `json:"inferredEntityType,omitempty" tf:"inferred_entity_type,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Whether or not Vault should resolve the bound_iam_principal_arn to an AWS Unique ID. When true, deleting a principal and recreating it with the same name won't automatically grant the new principal the same roles in Vault that the old principal had.
	ResolveAwsUniqueIds *bool `json:"resolveAwsUniqueIds,omitempty" tf:"resolve_aws_unique_ids,omitempty"`

	// Name of the role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The Vault generated role ID.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// The key of the tag on EC2 instance to use for role tags.
	RoleTag *string `json:"roleTag,omitempty" tf:"role_tag,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type AuthBackendRoleParameters struct {

	// When true, allows migration of the underlying instance where the client resides. Use with caution.
	// +kubebuilder:validation:Optional
	AllowInstanceMigration *bool `json:"allowInstanceMigration,omitempty" tf:"allow_instance_migration,omitempty"`

	// The auth type permitted for this role.
	// +kubebuilder:validation:Optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type,omitempty"`

	// Unique name of the auth backend to configure.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Only EC2 instances using this AMI ID will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundAMIIds []*string `json:"boundAmiIds,omitempty" tf:"bound_ami_ids,omitempty"`

	// Only EC2 instances with this account ID in their identity document will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundAccountIds []*string `json:"boundAccountIds,omitempty" tf:"bound_account_ids,omitempty"`

	// Only EC2 instances that match this instance ID will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundEC2InstanceIds []*string `json:"boundEc2InstanceIds,omitempty" tf:"bound_ec2_instance_ids,omitempty"`

	// Only EC2 instances associated with an IAM instance profile ARN that matches this value will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundIAMInstanceProfileArns []*string `json:"boundIamInstanceProfileArns,omitempty" tf:"bound_iam_instance_profile_arns,omitempty"`

	// The IAM principal that must be authenticated using the iam auth method.
	// +kubebuilder:validation:Optional
	BoundIAMPrincipalArns []*string `json:"boundIamPrincipalArns,omitempty" tf:"bound_iam_principal_arns,omitempty"`

	// Only EC2 instances that match this IAM role ARN will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundIAMRoleArns []*string `json:"boundIamRoleArns,omitempty" tf:"bound_iam_role_arns,omitempty"`

	// Only EC2 instances in this region will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundRegions []*string `json:"boundRegions,omitempty" tf:"bound_regions,omitempty"`

	// Only EC2 instances associated with this subnet ID will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundSubnetIds []*string `json:"boundSubnetIds,omitempty" tf:"bound_subnet_ids,omitempty"`

	// Only EC2 instances associated with this VPC ID will be permitted to log in.
	// +kubebuilder:validation:Optional
	BoundVPCIds []*string `json:"boundVpcIds,omitempty" tf:"bound_vpc_ids,omitempty"`

	// When true, only allows a single token to be granted per instance ID.
	// +kubebuilder:validation:Optional
	DisallowReauthentication *bool `json:"disallowReauthentication,omitempty" tf:"disallow_reauthentication,omitempty"`

	// The region to search for the inferred entities in.
	// +kubebuilder:validation:Optional
	InferredAwsRegion *string `json:"inferredAwsRegion,omitempty" tf:"inferred_aws_region,omitempty"`

	// The type of inferencing Vault should do.
	// +kubebuilder:validation:Optional
	InferredEntityType *string `json:"inferredEntityType,omitempty" tf:"inferred_entity_type,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Whether or not Vault should resolve the bound_iam_principal_arn to an AWS Unique ID. When true, deleting a principal and recreating it with the same name won't automatically grant the new principal the same roles in Vault that the old principal had.
	// +kubebuilder:validation:Optional
	ResolveAwsUniqueIds *bool `json:"resolveAwsUniqueIds,omitempty" tf:"resolve_aws_unique_ids,omitempty"`

	// Name of the role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The key of the tag on EC2 instance to use for role tags.
	// +kubebuilder:validation:Optional
	RoleTag *string `json:"roleTag,omitempty" tf:"role_tag,omitempty"`

	// Specifies the blocks of IP addresses which are allowed to use the generated token
	// +kubebuilder:validation:Optional
	TokenBoundCidrs []*string `json:"tokenBoundCidrs,omitempty" tf:"token_bound_cidrs,omitempty"`

	// Generated Token's Explicit Maximum TTL in seconds
	// +kubebuilder:validation:Optional
	TokenExplicitMaxTTL *float64 `json:"tokenExplicitMaxTtl,omitempty" tf:"token_explicit_max_ttl,omitempty"`

	// The maximum lifetime of the generated token
	// +kubebuilder:validation:Optional
	TokenMaxTTL *float64 `json:"tokenMaxTtl,omitempty" tf:"token_max_ttl,omitempty"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	// +kubebuilder:validation:Optional
	TokenNoDefaultPolicy *bool `json:"tokenNoDefaultPolicy,omitempty" tf:"token_no_default_policy,omitempty"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	// +kubebuilder:validation:Optional
	TokenNumUses *float64 `json:"tokenNumUses,omitempty" tf:"token_num_uses,omitempty"`

	// Generated Token's Period
	// +kubebuilder:validation:Optional
	TokenPeriod *float64 `json:"tokenPeriod,omitempty" tf:"token_period,omitempty"`

	// Generated Token's Policies
	// +kubebuilder:validation:Optional
	TokenPolicies []*string `json:"tokenPolicies,omitempty" tf:"token_policies,omitempty"`

	// The initial ttl of the token to generate in seconds
	// +kubebuilder:validation:Optional
	TokenTTL *float64 `json:"tokenTtl,omitempty" tf:"token_ttl,omitempty"`

	// The type of token to generate, service or batch
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// AuthBackendRoleSpec defines the desired state of AuthBackendRole
type AuthBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthBackendRoleParameters `json:"forProvider"`
}

// AuthBackendRoleStatus defines the observed state of AuthBackendRole.
type AuthBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRole is the Schema for the AuthBackendRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type AuthBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.role)",message="role is a required parameter"
	Spec   AuthBackendRoleSpec   `json:"spec"`
	Status AuthBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthBackendRoleList contains a list of AuthBackendRoles
type AuthBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthBackendRole `json:"items"`
}

// Repository type metadata.
var (
	AuthBackendRole_Kind             = "AuthBackendRole"
	AuthBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthBackendRole_Kind}.String()
	AuthBackendRole_KindAPIVersion   = AuthBackendRole_Kind + "." + CRDGroupVersion.String()
	AuthBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(AuthBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthBackendRole{}, &AuthBackendRoleList{})
}
