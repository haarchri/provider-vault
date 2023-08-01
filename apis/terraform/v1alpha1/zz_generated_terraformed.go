/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this CloudSecretBackend
func (mg *CloudSecretBackend) GetTerraformResourceType() string {
	return "vault_terraform_cloud_secret_backend"
}

// GetConnectionDetailsMapping for this CloudSecretBackend
func (tr *CloudSecretBackend) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"token": "spec.forProvider.tokenSecretRef"}
}

// GetObservation of this CloudSecretBackend
func (tr *CloudSecretBackend) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CloudSecretBackend
func (tr *CloudSecretBackend) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CloudSecretBackend
func (tr *CloudSecretBackend) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CloudSecretBackend
func (tr *CloudSecretBackend) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CloudSecretBackend
func (tr *CloudSecretBackend) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this CloudSecretBackend
func (tr *CloudSecretBackend) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this CloudSecretBackend using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CloudSecretBackend) LateInitialize(attrs []byte) (bool, error) {
	params := &CloudSecretBackendParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CloudSecretBackend) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CloudSecretCreds
func (mg *CloudSecretCreds) GetTerraformResourceType() string {
	return "vault_terraform_cloud_secret_creds"
}

// GetConnectionDetailsMapping for this CloudSecretCreds
func (tr *CloudSecretCreds) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"lease_id": "status.atProvider.leaseId", "token": "status.atProvider.token"}
}

// GetObservation of this CloudSecretCreds
func (tr *CloudSecretCreds) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CloudSecretCreds
func (tr *CloudSecretCreds) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CloudSecretCreds
func (tr *CloudSecretCreds) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CloudSecretCreds
func (tr *CloudSecretCreds) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CloudSecretCreds
func (tr *CloudSecretCreds) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this CloudSecretCreds
func (tr *CloudSecretCreds) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this CloudSecretCreds using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CloudSecretCreds) LateInitialize(attrs []byte) (bool, error) {
	params := &CloudSecretCredsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CloudSecretCreds) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CloudSecretRole
func (mg *CloudSecretRole) GetTerraformResourceType() string {
	return "vault_terraform_cloud_secret_role"
}

// GetConnectionDetailsMapping for this CloudSecretRole
func (tr *CloudSecretRole) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CloudSecretRole
func (tr *CloudSecretRole) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CloudSecretRole
func (tr *CloudSecretRole) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CloudSecretRole
func (tr *CloudSecretRole) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CloudSecretRole
func (tr *CloudSecretRole) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CloudSecretRole
func (tr *CloudSecretRole) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this CloudSecretRole
func (tr *CloudSecretRole) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this CloudSecretRole using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CloudSecretRole) LateInitialize(attrs []byte) (bool, error) {
	params := &CloudSecretRoleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CloudSecretRole) GetTerraformSchemaVersion() int {
	return 0
}
