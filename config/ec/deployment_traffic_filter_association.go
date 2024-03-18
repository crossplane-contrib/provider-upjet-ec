package ec

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ConfigureDeploymentTrafficFilterAssociation configures the Deployment Traffic Filter Association resource.
func ConfigureDeploymentTrafficFilterAssociation(p *config.Provider) {
	p.AddResourceConfigurator("ec_deployment_traffic_filter_association", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "elasticcloud"
		r.UseAsync = true
	})
}
