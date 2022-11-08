// CDK constructs to use OpenID Connect for authenticating your Github Action workflow with AWS IAM
package awscdkgithuboidc

import (
	"github.com/aripalo/aws-cdk-github-oidc-go/awscdkgithuboidc/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Describes a Github OpenID Connect Identity Provider for AWS IAM.
// Experimental.
type IGithubActionsIdentityProvider interface {
	awsiam.IOpenIdConnectProvider
}

// The jsii proxy for IGithubActionsIdentityProvider
type jsiiProxy_IGithubActionsIdentityProvider struct {
	internal.Type__awsiamIOpenIdConnectProvider
}

