//go:build no_runtime_type_checking

// CDK constructs to use OpenID Connect for authenticating your Github Action workflow with AWS IAM
package awscdkgithuboidc

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubActionsIdentityProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GithubActionsIdentityProvider) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GithubActionsIdentityProvider) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateGithubActionsIdentityProvider_FromAccountParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateGithubActionsIdentityProvider_FromOpenIdConnectProviderArnParameters(scope constructs.Construct, id *string, openIdConnectProviderArn *string) error {
	return nil
}

func validateGithubActionsIdentityProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGithubActionsIdentityProvider_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGithubActionsIdentityProviderParameters(scope constructs.Construct, id *string) error {
	return nil
}

