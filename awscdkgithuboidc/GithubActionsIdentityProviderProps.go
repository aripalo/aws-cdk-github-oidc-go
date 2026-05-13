package awscdkgithuboidc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type GithubActionsIdentityProviderProps struct {
	// The removal policy for the provider.
	// Default: cdk.RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

