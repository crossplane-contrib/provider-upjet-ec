package ec

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ConfigureDeploymentTrafficFilter configures the Deployment Traffic Filter resource.
func ConfigureDeploymentTrafficFilter(p *config.Provider) {
	p.AddResourceConfigurator("ec_deployment_traffic_filter", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "elasticcloud"
		r.UseAsync = true
	})
}
