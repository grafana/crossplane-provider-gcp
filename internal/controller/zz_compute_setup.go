/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backendservice "github.com/upbound/provider-gcp/internal/controller/compute/backendservice"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/globalforwardingrule"
	healthcheck "github.com/upbound/provider-gcp/internal/controller/compute/healthcheck"
	httphealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/httphealthcheck"
	httpshealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/httpshealthcheck"
	instancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/instancegroupmanager"
	instancetemplate "github.com/upbound/provider-gcp/internal/controller/compute/instancetemplate"
	network "github.com/upbound/provider-gcp/internal/controller/compute/network"
	regionbackendservice "github.com/upbound/provider-gcp/internal/controller/compute/regionbackendservice"
	regionhealthcheck "github.com/upbound/provider-gcp/internal/controller/compute/regionhealthcheck"
	regioninstancegroupmanager "github.com/upbound/provider-gcp/internal/controller/compute/regioninstancegroupmanager"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpsproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/compute/regionurlmap"
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/sslcertificate"
	subnetwork "github.com/upbound/provider-gcp/internal/controller/compute/subnetwork"
	targethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpproxy"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpsproxy"
	targetpool "github.com/upbound/provider-gcp/internal/controller/compute/targetpool"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backendservice.Setup,
		globaladdress.Setup,
		globalforwardingrule.Setup,
		healthcheck.Setup,
		httphealthcheck.Setup,
		httpshealthcheck.Setup,
		instancegroupmanager.Setup,
		instancetemplate.Setup,
		network.Setup,
		regionbackendservice.Setup,
		regionhealthcheck.Setup,
		regioninstancegroupmanager.Setup,
		regionsslcertificate.Setup,
		regiontargethttpproxy.Setup,
		regiontargethttpsproxy.Setup,
		regionurlmap.Setup,
		sslcertificate.Setup,
		subnetwork.Setup,
		targethttpproxy.Setup,
		targethttpsproxy.Setup,
		targetpool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
