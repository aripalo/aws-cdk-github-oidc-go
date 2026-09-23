package awscdkgithuboidc


// Github related configuration that forms the trust policy for this IAM Role.
// Experimental.
type GithubConfiguration struct {
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
	// With `ownerId` & `repoId` given, appended after
	// `repo:${owner}@${ownerId}/${repo}@${repoId}:` instead.
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
	// Numeric Github ID of the repository owner (organization or user), which makes the subject immutable.
	//
	// Must be given together with `repoId`.
	//
	// Example:
	//   '123456'
	//
	// See: https://docs.github.com/en/actions/reference/security/oidc#immutable-subject-claims
	//
	// Default: - subject refers to the owner and repository by name.
	//
	// Experimental.
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// Numeric Github ID of the repository, which makes the subject immutable.
	//
	// Must be given together with `ownerId`.
	//
	// Example:
	//   '456789'
	//
	// See: https://docs.github.com/en/actions/reference/security/oidc#immutable-subject-claims
	//
	// Default: - subject refers to the owner and repository by name.
	//
	// Experimental.
	RepoId *string `field:"optional" json:"repoId" yaml:"repoId"`
}

