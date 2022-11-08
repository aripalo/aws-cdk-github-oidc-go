//go:build no_runtime_type_checking

// CDK constructs to use OpenID Connect for authenticating your Github Action workflow with AWS IAM
package awscdkgithuboidc

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubActionsRole) validateAddManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateAddToPrincipalPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateAttachInlinePolicyParameters(policy awsiam.Policy) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGrantParameters(grantee awsiam.IPrincipal) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGrantPassRoleParameters(identity awsiam.IPrincipal) error {
	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateWithoutPolicyUpdatesParameters(options *awsiam.WithoutPolicyUpdatesOptions) error {
	return nil
}

func validateGithubActionsRole_FromRoleArnParameters(scope constructs.Construct, id *string, roleArn *string, options *awsiam.FromRoleArnOptions) error {
	return nil
}

func validateGithubActionsRole_FromRoleNameParameters(scope constructs.Construct, id *string, roleName *string) error {
	return nil
}

func validateGithubActionsRole_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGithubActionsRole_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGithubActionsRoleParameters(scope constructs.Construct, id *string, props *GithubActionsRoleProps) error {
	return nil
}

