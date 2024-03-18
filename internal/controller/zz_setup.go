package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	deployment "github.com/crossplane-contrib/provider-upjet-ec/internal/controller/elasticcloud/deployment"
	extension "github.com/crossplane-contrib/provider-upjet-ec/internal/controller/elasticcloud/extension"
	trafficfilter "github.com/crossplane-contrib/provider-upjet-ec/internal/controller/elasticcloud/trafficfilter"
	trafficfilterassociation "github.com/crossplane-contrib/provider-upjet-ec/internal/controller/elasticcloud/trafficfilterassociation"
	providerconfig "github.com/crossplane-contrib/provider-upjet-ec/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		deployment.Setup,
		extension.Setup,
		trafficfilter.Setup,
		trafficfilterassociation.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
