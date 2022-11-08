package awscdkgithuboidc

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-github-oidc.GithubActionsIdentityProvider",
		reflect.TypeOf((*GithubActionsIdentityProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderArn", GoGetter: "OpenIdConnectProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderIssuer", GoGetter: "OpenIdConnectProviderIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GithubActionsIdentityProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamOpenIdConnectProvider)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGithubActionsIdentityProvider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-github-oidc.GithubActionsRole",
		reflect.TypeOf((*GithubActionsRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addManagedPolicy", GoMethod: "AddManagedPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleAction", GoGetter: "AssumeRoleAction"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRolePolicy", GoGetter: "AssumeRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "attachInlinePolicy", GoMethod: "AttachInlinePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantPassRole", GoMethod: "GrantPassRole"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsBoundary", GoGetter: "PermissionsBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyFragment", GoGetter: "PolicyFragment"},
			_jsii_.MemberProperty{JsiiProperty: "principalAccount", GoGetter: "PrincipalAccount"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleId", GoGetter: "RoleId"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "withoutPolicyUpdates", GoMethod: "WithoutPolicyUpdates"},
		},
		func() interface{} {
			j := jsiiProxy_GithubActionsRole{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamRole)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-github-oidc.GithubActionsRoleProps",
		reflect.TypeOf((*GithubActionsRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-github-oidc.GithubConfiguration",
		reflect.TypeOf((*GithubConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-github-oidc.IGithubActionsIdentityProvider",
		reflect.TypeOf((*IGithubActionsIdentityProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderArn", GoGetter: "OpenIdConnectProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderIssuer", GoGetter: "OpenIdConnectProviderIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IGithubActionsIdentityProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIOpenIdConnectProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-github-oidc.RoleProps",
		reflect.TypeOf((*RoleProps)(nil)).Elem(),
	)
}
