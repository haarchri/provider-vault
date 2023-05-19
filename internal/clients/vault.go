/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/upbound/provider-vault/apis/v1beta1"
)

const (
	// provider config
	// source: https://registry.terraform.io/providers/hashicorp/vault/latest/docs

	// required origin URL of the Vault server
	keyAddress = "address"

	// optional provider arguments
	// remove options that we do not yet want to implement
	keyAddAddressToEnv      = "add_address_to_env"
	keyToken                = "token"
	keyTokenName            = "token_name"
	keyCaCertFile           = "ca_cert_file"
	keyCaCertDir            = "ca_cert_dir"
	keyAuthLoginUserpass    = "auth_login_userpath"
	keyAuthLoginAWS         = "auth_login_aws"
	keyAuthLoginCert        = "auth_login_cert"
	keyAuthLoginGCP         = "auth_login_gcp"
	keyAuthLoginKerberos    = "auth_login_kerberos"
	keyAuthLoginRadius      = "auth_login_radius"
	keyAuthLoginOCI         = "auth_login_oci"
	keyAuthLoginOIDC        = "auth_login_oidc"
	keyAuthLoginJWT         = "auth_login_jwt"
	keyAuthLoginAzure       = "auth_login_azure"
	keyAuthLogin            = "auth_login"
	keyClientAuth           = "client_auth"
	keySkipTLSVerify        = "skip_tls_verify"
	keyTLSServerName        = "tls_server_name"
	keySkipChildToken       = "skip_child_token"
	keyMaxLeaseTTLSeconds   = "max_lease_ttl_seconds"
	keyMaxRetries           = "max_retries"
	keyMaxRetriesCcc        = "max_retries_ccc"
	keyNamespace            = "namespace"
	keySkipGetVaultVersion  = "skip_get_vault_version"
	keyVaultVersionOverride = "vault_version_override"
	keyHeaders              = "headers"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal vault credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		// set provider configuration
		ps.Configuration = map[string]any{}

		// Assign mandatory address parameter
		ps.Configuration[keyAddress] = pc.Spec.Address

		// Assign optional parameters
		ps.Configuration[keyAddAddressToEnv] = pc.Spec.AddAddressToEnv
		ps.Configuration[keySkipTLSVerify] = pc.Spec.SkipTLSVerify
		if len(pc.Spec.TLSServerName) > 0 {
			ps.Configuration[keyTLSServerName] = pc.Spec.TLSServerName
		}
		ps.Configuration[keySkipChildToken] = pc.Spec.SkipChildToken
		ps.Configuration[keyMaxLeaseTTLSeconds] = pc.Spec.MaxLeaseTTLSeconds
		ps.Configuration[keyMaxRetries] = pc.Spec.MaxRetries
		ps.Configuration[keyMaxRetriesCcc] = pc.Spec.MaxRetriesCcc
		if len(pc.Spec.Namespace) > 0 {
			ps.Configuration[keyNamespace] = pc.Spec.Namespace
		}
		ps.Configuration[keySkipGetVaultVersion] = pc.Spec.SkipGetVaultVersion
		if len(pc.Spec.VaultVersionOverride) > 0 {
			ps.Configuration[keyVaultVersionOverride] = pc.Spec.VaultVersionOverride
		}
		if pc.Spec.Headers != (v1beta1.ProviderHeaders{}) {
			ps.Configuration[keyHeaders] = pc.Spec.Headers
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform
		// provider configuration
		credsKeys := [...]string{keyToken, keyTokenName, keyCaCertFile,
			keyCaCertDir, keyAuthLoginUserpass, keyAuthLoginAWS,
			keyAuthLoginCert, keyAuthLoginGCP, keyAuthLoginKerberos,
			keyAuthLoginRadius, keyAuthLoginOCI, keyAuthLoginOIDC,
			keyAuthLoginJWT, keyAuthLoginAzure, keyAuthLogin, keyClientAuth}

		for _, key := range credsKeys {
			if v, ok := creds[key]; ok {
				ps.Configuration[key] = v
			}
		}
		return ps, nil
	}
}
