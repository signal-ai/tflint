// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIotRoleAliasInvalidAliasRule checks the pattern is valid
type AwsIotRoleAliasInvalidAliasRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIotRoleAliasInvalidAliasRule returns new rule with default attributes
func NewAwsIotRoleAliasInvalidAliasRule() *AwsIotRoleAliasInvalidAliasRule {
	return &AwsIotRoleAliasInvalidAliasRule{
		resourceType:  "aws_iot_role_alias",
		attributeName: "alias",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w=,@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIotRoleAliasInvalidAliasRule) Name() string {
	return "aws_iot_role_alias_invalid_alias"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIotRoleAliasInvalidAliasRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIotRoleAliasInvalidAliasRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIotRoleAliasInvalidAliasRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIotRoleAliasInvalidAliasRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"alias must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"alias must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w=,@-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
