package ec

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func ConfigureStack(p *config.Provider) {
	p.AddResourceConfigurator("ec_stack", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "elasticcloud"
		r.UseAsync = true
	})
}
