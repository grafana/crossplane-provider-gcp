/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	serviceaccount "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccount"
	serviceaccountiammember "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountiammember"
	serviceaccountkey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/serviceaccountkey"
	servicenetworkingpeereddnsdomain "github.com/upbound/provider-gcp/internal/controller/cloudplatform/servicenetworkingpeereddnsdomain"
	storagehmackey "github.com/upbound/provider-gcp/internal/controller/cloudplatform/storagehmackey"
	globaladdress "github.com/upbound/provider-gcp/internal/controller/compute/globaladdress"
	globalforwardingrule "github.com/upbound/provider-gcp/internal/controller/compute/globalforwardingrule"
	network "github.com/upbound/provider-gcp/internal/controller/compute/network"
	regionsslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/regionsslcertificate"
	regiontargethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpproxy"
	regiontargethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/regiontargethttpsproxy"
	regionurlmap "github.com/upbound/provider-gcp/internal/controller/compute/regionurlmap"
	sslcertificate "github.com/upbound/provider-gcp/internal/controller/compute/sslcertificate"
	targethttpproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpproxy"
	targethttpsproxy "github.com/upbound/provider-gcp/internal/controller/compute/targethttpsproxy"
	recordset "github.com/upbound/provider-gcp/internal/controller/dns/recordset"
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

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		servicenetworkingpeereddnsdomain.Setup,
		storagehmackey.Setup,
		globaladdress.Setup,
		globalforwardingrule.Setup,
		network.Setup,
		regionsslcertificate.Setup,
		regiontargethttpproxy.Setup,
		regiontargethttpsproxy.Setup,
		regionurlmap.Setup,
		sslcertificate.Setup,
		targethttpproxy.Setup,
		targethttpsproxy.Setup,
		recordset.Setup,
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
