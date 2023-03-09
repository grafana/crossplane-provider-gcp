/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdminUsersObservation struct {
}

type AdminUsersParameters struct {

	// The name of the user, e.g. my-gcp-id@gmail.com.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type AuthorizationObservation struct {
}

type AuthorizationParameters struct {

	// Users that can perform operations as a cluster admin. A new ClusterRoleBinding will be created to grant the cluster-admin ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kubebuilder:validation:Required
	AdminUsers []AdminUsersParameters `json:"adminUsers" tf:"admin_users,omitempty"`
}

type AzureServicesAuthenticationObservation struct {
}

type AzureServicesAuthenticationParameters struct {

	// The Azure Active Directory Application ID for Authentication configuration.
	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// The Azure Active Directory Tenant ID for Authentication configuration.
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type ClusterObservation struct {

	// Output only. The time at which this cluster was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Output only. The endpoint of the cluster's API server.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Fleet configuration.
	// +kubebuilder:validation:Required
	Fleet []FleetObservation `json:"fleet,omitempty" tf:"fleet,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/azureClusters/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling,omitempty"`

	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. A globally unique identifier for the cluster.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The time at which this cluster was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Output only. Workload Identity settings.
	WorkloadIdentityConfig []WorkloadIdentityConfigObservation `json:"workloadIdentityConfig,omitempty" tf:"workload_identity_config,omitempty"`
}

type ClusterParameters struct {

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix  and name , separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Configuration related to the cluster RBAC settings.
	// +kubebuilder:validation:Required
	Authorization []AuthorizationParameters `json:"authorization" tf:"authorization,omitempty"`

	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	// +kubebuilder:validation:Required
	AzureRegion *string `json:"azureRegion" tf:"azure_region,omitempty"`

	// Azure authentication configuration for management of Azure resources
	// +kubebuilder:validation:Optional
	AzureServicesAuthentication []AzureServicesAuthenticationParameters `json:"azureServicesAuthentication,omitempty" tf:"azure_services_authentication,omitempty"`

	// Name of the AzureClient. The AzureClient resource must reside on the same GCP project and region as the AzureCluster. AzureClient names are formatted as projects/<project-number>/locations/<region>/azureClients/<client-id>. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	// +kubebuilder:validation:Optional
	Client *string `json:"client,omitempty" tf:"client,omitempty"`

	// Configuration related to the cluster control plane.
	// +kubebuilder:validation:Required
	ControlPlane []ControlPlaneParameters `json:"controlPlane" tf:"control_plane,omitempty"`

	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Fleet configuration.
	// +kubebuilder:validation:Required
	Fleet []FleetParameters `json:"fleet" tf:"fleet,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Cluster-wide networking configuration.
	// +kubebuilder:validation:Required
	Networking []NetworkingParameters `json:"networking" tf:"networking,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The ARM ID of the resource group where the cluster resources are deployed. For example: /subscriptions/*/resourceGroups/*
	// +kubebuilder:validation:Required
	ResourceGroupID *string `json:"resourceGroupId" tf:"resource_group_id,omitempty"`
}

type ControlPlaneObservation struct {
}

type ControlPlaneParameters struct {

	// Optional. Configuration related to application-layer secrets encryption.
	// +kubebuilder:validation:Optional
	DatabaseEncryption []DatabaseEncryptionParameters `json:"databaseEncryption,omitempty" tf:"database_encryption,omitempty"`

	// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. When unspecified, it defaults to a 8-GiB Azure Disk.
	// +kubebuilder:validation:Optional
	MainVolume []MainVolumeParameters `json:"mainVolume,omitempty" tf:"main_volume,omitempty"`

	// Proxy configuration for outbound HTTP(S) traffic.
	// +kubebuilder:validation:Optional
	ProxyConfig []ProxyConfigParameters `json:"proxyConfig,omitempty" tf:"proxy_config,omitempty"`

	// Configuration for where to place the control plane replicas. Up to three replica placement instances can be specified. If replica_placements is set, the replica placement instances will be applied to the three control plane replicas as evenly as possible.
	// +kubebuilder:validation:Optional
	ReplicaPlacements []ReplicaPlacementsParameters `json:"replicaPlacements,omitempty" tf:"replica_placements,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each control plane replica. When unspecified, it defaults to 32-GiB Azure Disk.
	// +kubebuilder:validation:Optional
	RootVolume []RootVolumeParameters `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// SSH configuration for how to access the underlying control plane machines.
	// +kubebuilder:validation:Required
	SSHConfig []SSHConfigParameters `json:"sshConfig" tf:"ssh_config,omitempty"`

	// The ARM ID of the subnet where the control plane VMs are deployed. Example: /subscriptions//resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/default.
	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// Optional. A set of tags to apply to all underlying control plane Azure resources.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Optional. The Azure VM size name. Example: Standard_DS2_v2. For available VM sizes, see https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions. When unspecified, it defaults to Standard_DS2_v2.
	// +kubebuilder:validation:Optional
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`

	// The Kubernetes version to run on control plane replicas (e.g. 1.19.10-gke.1000). You can list all supported versions on a given Google Cloud region by calling GetAzureServerConfig.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type DatabaseEncryptionObservation struct {
}

type DatabaseEncryptionParameters struct {

	// The ARM ID of the Azure Key Vault key to encrypt / decrypt data. For example: /subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name> Encryption will always take the latest version of the key and hence specific version is not supported.
	// +kubebuilder:validation:Required
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`
}

type FleetObservation struct {

	// The name of the managed Hub Membership resource associated to this cluster. Membership names are formatted as projects//locations/global/membership/.
	Membership *string `json:"membership,omitempty" tf:"membership,omitempty"`
}

type FleetParameters struct {

	// The number of the Fleet host project where this cluster will be registered.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type MainVolumeObservation struct {
}

type MainVolumeParameters struct {

	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +kubebuilder:validation:Optional
	SizeGib *float64 `json:"sizeGib,omitempty" tf:"size_gib,omitempty"`
}

type NetworkingObservation struct {
}

type NetworkingParameters struct {

	// The IP address range of the pods in this cluster, in CIDR notation (e.g. 10.96.0.0/14). All pods in the cluster get assigned a unique RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	// +kubebuilder:validation:Required
	PodAddressCidrBlocks []*string `json:"podAddressCidrBlocks" tf:"pod_address_cidr_blocks,omitempty"`

	// The IP address range for services in this cluster, in CIDR notation (e.g. 10.96.0.0/14). All services in the cluster get assigned a unique RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creating a cluster.
	// +kubebuilder:validation:Required
	ServiceAddressCidrBlocks []*string `json:"serviceAddressCidrBlocks" tf:"service_address_cidr_blocks,omitempty"`

	// The Azure Resource Manager (ARM) ID of the VNet associated with your cluster. All components in the cluster (i.e. control plane and node pools) run on a single VNet. Example: /subscriptions/*/resourceGroups/*/providers/Microsoft.Network/virtualNetworks/* This field cannot be changed after creation.
	// +kubebuilder:validation:Required
	VirtualNetworkID *string `json:"virtualNetworkId" tf:"virtual_network_id,omitempty"`
}

type ProxyConfigObservation struct {
}

type ProxyConfigParameters struct {

	// The ARM ID of the resource group where the cluster resources are deployed. For example: /subscriptions/*/resourceGroups/*
	// +kubebuilder:validation:Required
	ResourceGroupID *string `json:"resourceGroupId" tf:"resource_group_id,omitempty"`

	// The URL the of the proxy setting secret with its version. Secret ids are formatted as https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>.
	// +kubebuilder:validation:Required
	SecretID *string `json:"secretId" tf:"secret_id,omitempty"`
}

type ReplicaPlacementsObservation struct {
}

type ReplicaPlacementsParameters struct {

	// For a given replica, the Azure availability zone where to provision the control plane VM and the ETCD disk.
	// +kubebuilder:validation:Required
	AzureAvailabilityZone *string `json:"azureAvailabilityZone" tf:"azure_availability_zone,omitempty"`

	// For a given replica, the ARM ID of the subnet where the control plane VM is deployed. Make sure it's a subnet under the virtual network in the cluster configuration.
	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

type RootVolumeObservation struct {
}

type RootVolumeParameters struct {

	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +kubebuilder:validation:Optional
	SizeGib *float64 `json:"sizeGib,omitempty" tf:"size_gib,omitempty"`
}

type SSHConfigObservation struct {
}

type SSHConfigParameters struct {

	// The SSH public key data for VMs managed by Anthos. This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	// +kubebuilder:validation:Required
	AuthorizedKey *string `json:"authorizedKey" tf:"authorized_key,omitempty"`
}

type WorkloadIdentityConfigObservation struct {
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	WorkloadPool *string `json:"workloadPool,omitempty" tf:"workload_pool,omitempty"`
}

type WorkloadIdentityConfigParameters struct {
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. An Anthos cluster running on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
