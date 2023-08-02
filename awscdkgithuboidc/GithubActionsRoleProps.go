package awscdkgithuboidc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props that define the IAM Role that can be assumed by Github Actions workflow via Github OpenID Connect Identity Provider.
//
// Besides `GithubConfiguration`, you may pass in any `iam.RoleProps` except `assumedBy`
// which will be defined by this construct (CDK will fail if you do).
//
// Example:
//   {
//     provider: GithubActionsIdentityProvider.fromAccount(scope, "GithubProvider"),
//     owner: 'octo-org',
//     repo: 'octo-repo',
//     filter: 'ref:refs/tags/v*',
//     roleName: 'MyDeployRole',
//   }
//
// Experimental.
type GithubActionsRoleProps struct {
	// Repository owner (organization or username).
	//
	// Example:
	//   'octo-org'
	//
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Reference to Github OpenID Connect Provider configured in AWS IAM.
	//
	// Either pass an construct defined by `new GithubActionsIdentityProvider`
	// or a retrieved reference from `GithubActionsIdentityProvider.fromAccount`.
	// There can be only one (per AWS Account).
	// Experimental.
	Provider IGithubActionsIdentityProvider `field:"required" json:"provider" yaml:"provider"`
	// Repository name (slug) without the owner.
	//
	// Example:
	//   'octo-repo'
	//
	// Experimental.
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// Subject condition filter, appended after `repo:${owner}/${repo}:` string in IAM Role trust relationship.
	//
	// Example:
	//   'ref:refs/tags/v*'
	//   'ref:refs/heads/demo-branch'
	//   'pull_request'
	//   'environment:Production'
	//
	// See: https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#examples
	//
	// Default: '*'
	//
	// You may use this value to only allow Github to assume the role on specific branches, tags, environments, pull requests etc.
	//
	// Experimental.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// A description of the role.
	//
	// It can be up to 1000 characters long.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of IDs that the role assumer needs to provide one of when assuming this role.
	//
	// If the configured and provided external IDs do not match, the
	// AssumeRole operation will fail.
	// Default: No external ID required.
	//
	// Experimental.
	ExternalIds *[]*string `field:"optional" json:"externalIds" yaml:"externalIds"`
	// A list of named policies to inline into this role.
	//
	// These policies will be
	// created with the role, whereas those added by ``addToPolicy`` are added
	// using a separate CloudFormation resource (allowing a way around circular
	// dependencies that could otherwise be introduced).
	// Default: - No policy is inlined in the Role resource.
	//
	// Experimental.
	InlinePolicies *map[string]awsiam.PolicyDocument `field:"optional" json:"inlinePolicies" yaml:"inlinePolicies"`
	// A list of managed policies associated with this role.
	//
	// You can add managed policies later using
	// `addManagedPolicy(ManagedPolicy.fromAwsManagedPolicyName(policyName))`.
	// Default: - No managed policies.
	//
	// Experimental.
	ManagedPolicies *[]awsiam.IManagedPolicy `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// The maximum session duration that you want to set for the specified role.
	//
	// This setting can have a value from 1 hour (3600sec) to 12 (43200sec) hours.
	//
	// Anyone who assumes the role from the AWS CLI or API can use the
	// DurationSeconds API parameter or the duration-seconds CLI parameter to
	// request a longer session. The MaxSessionDuration setting determines the
	// maximum duration that can be requested using the DurationSeconds
	// parameter.
	//
	// If users don't specify a value for the DurationSeconds parameter, their
	// security credentials are valid for one hour by default. This applies when
	// you use the AssumeRole* API operations or the assume-role* CLI operations
	// but does not apply when you use those operations to create a console URL.
	// Default: Duration.hours(1)
	//
	// Experimental.
	MaxSessionDuration awscdk.Duration `field:"optional" json:"maxSessionDuration" yaml:"maxSessionDuration"`
	// The path associated with this role.
	//
	// For information about IAM paths, see
	// Friendly Names and Paths in IAM User Guide.
	// Default: /.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// AWS supports permissions boundaries for IAM entities (users or roles).
	//
	// A permissions boundary is an advanced feature for using a managed policy
	// to set the maximum permissions that an identity-based policy can grant to
	// an IAM entity. An entity's permissions boundary allows it to perform only
	// the actions that are allowed by both its identity-based policies and its
	// permissions boundaries.
	// Default: - No permissions boundary.
	//
	// Experimental.
	PermissionsBoundary awsiam.IManagedPolicy `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// A name for the IAM role.
	//
	// For valid values, see the RoleName parameter for
	// the CreateRole action in the IAM API Reference.
	//
	// IMPORTANT: If you specify a name, you cannot perform updates that require
	// replacement of this resource. You can perform updates that require no or
	// some interruption. If you must replace the resource, specify a new name.
	//
	// If you specify a name, you must specify the CAPABILITY_NAMED_IAM value to
	// acknowledge your template's capabilities. For more information, see
	// Acknowledging IAM Resources in AWS CloudFormation Templates.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that ID
	// for the role name.
	//
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

