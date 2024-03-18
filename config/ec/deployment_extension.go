package ec

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ConfigureDeploymentExtension configures the Deployment Extension resource.
func ConfigureDeploymentExtension(p *config.Provider) {
	p.AddResourceConfigurator("ec_deployment_extension", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "elasticcloud"
		r.UseAsync = true
	})
}
