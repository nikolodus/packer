package hcl2template

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	pkrversion "github.com/hashicorp/packer/version"
)

// CheckCoreVersionRequirements visits each of the block in the given
// configuration and verifies that any given Core version constraints match
// with the version of Packer Core that is being used.
//
// The returned diagnostics will contain errors if any constraints do not match.
// The returned diagnostics might also return warnings, which should be
// displayed to the user.
func (cfg *PackerConfig) CheckCoreVersionRequirements() hcl.Diagnostics {
	if cfg == nil {
		return nil
	}

	var diags hcl.Diagnostics

	for _, constraint := range cfg.Packer.VersionConstraints {
		if !constraint.Required.Check(pkrversion.SemVer) {
			diags = diags.Append(&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Unsupported Packer Core version",
				Detail: fmt.Sprintf(
					"This configuration does not support Packer version %s. To proceed, either choose another supported Packer version or update this version constraint. Version constraints are normally set for good reason, so updating the constraint may lead to other errors or unexpected behavior.",
					pkrversion.String(),
				),
				Subject: constraint.DeclRange.Ptr(),
			})
		}
	}

	return diags
}
