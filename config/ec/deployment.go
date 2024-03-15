package ec

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func ConfigureDeployment(p *config.Provider) {
	p.AddResourceConfigurator("ec_deployment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "elasticcloud"
		r.UseAsync = true
	})
}
