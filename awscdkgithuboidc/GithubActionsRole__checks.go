//go:build !no_runtime_type_checking

package awscdkgithuboidc

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_GithubActionsRole) validateAddManagedPolicyParameters(policy awsiam.IManagedPolicy) error {
	if policy == nil {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateAddToPolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateAddToPrincipalPolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateAttachInlinePolicyParameters(policy awsiam.Policy) error {
	if policy == nil {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGrantParameters(grantee awsiam.IPrincipal) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGrantAssumeRoleParameters(identity awsiam.IPrincipal) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateGrantPassRoleParameters(identity awsiam.IPrincipal) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GithubActionsRole) validateWithoutPolicyUpdatesParameters(options *awsiam.WithoutPolicyUpdatesOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGithubActionsRole_CustomizeRolesParameters(scope constructs.Construct, options *awsiam.CustomizeRolesOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGithubActionsRole_FromRoleArnParameters(scope constructs.Construct, id *string, roleArn *string, options *awsiam.FromRoleArnOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if roleArn == nil {
		return fmt.Errorf("parameter roleArn is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGithubActionsRole_FromRoleNameParameters(scope constructs.Construct, id *string, roleName *string, options *awsiam.FromRoleNameOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if roleName == nil {
		return fmt.Errorf("parameter roleName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateGithubActionsRole_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGithubActionsRole_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateGithubActionsRole_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateGithubActionsRole_IsRoleParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewGithubActionsRoleParameters(scope constructs.Construct, id *string, props *GithubActionsRoleProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

