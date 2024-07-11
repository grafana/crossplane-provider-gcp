/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	dataset "github.com/upbound/provider-gcp/internal/controller/bigquery/dataset"
	table "github.com/upbound/provider-gcp/internal/controller/bigquery/table"
	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
	address "github.com/upbound/provider-gcp/internal/controller/compute/address"
	backendservice "github.com/upbound/provider-gcp/internal/controller/compute/backendservice"
	forwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/forwardingrule"
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
	recordset "github.com/upbound/provider-gcp/internal/controller/dns/recordset"
	webbackendserviceiammember "github.com/upbound/provider-gcp/internal/controller/iap/webbackendserviceiammember"
	providerconfig "github.com/upbound/provider-gcp/internal/controller/providerconfig"
	database "github.com/upbound/provider-gcp/internal/controller/sql/database"
	databaseinstance "github.com/upbound/provider-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/provider-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/provider-gcp/internal/controller/sql/sslcert"
	user "github.com/upbound/provider-gcp/internal/controller/sql/user"
	bucket "github.com/upbound/provider-gcp/internal/controller/storage/bucket"
	bucketaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/bucketaccesscontrol"
	bucketacl "github.com/upbound/provider-gcp/internal/controller/storage/bucketacl"
	bucketiammember "github.com/upbound/provider-gcp/internal/controller/storage/bucketiammember"
	bucketobject "github.com/upbound/provider-gcp/internal/controller/storage/bucketobject"
	defaultobjectaccesscontrol "github.com/upbound/provider-gcp/internal/controller/storage/defaultobjectaccesscontrol"
	defaultobjectacl "github.com/upbound/provider-gcp/internal/controller/storage/defaultobjectacl"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.Setup,
		table.Setup,
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
		address.Setup,
		backendservice.Setup,
		forwardingrule.Setup,
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
		recordset.Setup,
		webbackendserviceiammember.Setup,
		providerconfig.Setup,
		database.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
		bucket.Setup,
		bucketaccesscontrol.Setup,
		bucketacl.Setup,
		bucketiammember.Setup,
		bucketobject.Setup,
		defaultobjectaccesscontrol.Setup,
		defaultobjectacl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
