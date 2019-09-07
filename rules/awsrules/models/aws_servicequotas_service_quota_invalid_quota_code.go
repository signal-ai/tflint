// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsServicequotasServiceQuotaInvalidQuotaCodeRule checks the pattern is valid
type AwsServicequotasServiceQuotaInvalidQuotaCodeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsServicequotasServiceQuotaInvalidQuotaCodeRule returns new rule with default attributes
func NewAwsServicequotasServiceQuotaInvalidQuotaCodeRule() *AwsServicequotasServiceQuotaInvalidQuotaCodeRule {
	return &AwsServicequotasServiceQuotaInvalidQuotaCodeRule{
		resourceType:  "aws_servicequotas_service_quota",
		attributeName: "quota_code",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-]{1,128}$`),
	}
}

// Name returns the rule name
func (r *AwsServicequotasServiceQuotaInvalidQuotaCodeRule) Name() string {
	return "aws_servicequotas_service_quota_invalid_quota_code"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicequotasServiceQuotaInvalidQuotaCodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicequotasServiceQuotaInvalidQuotaCodeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicequotasServiceQuotaInvalidQuotaCodeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicequotasServiceQuotaInvalidQuotaCodeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"quota_code must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"quota_code must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`quota_code does not match valid pattern ^[a-zA-Z][a-zA-Z0-9-]{1,128}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
