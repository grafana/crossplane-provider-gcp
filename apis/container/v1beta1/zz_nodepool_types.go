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

type NodeConfigGuestAcceleratorGpuSharingConfigObservation struct {
	GpuSharingStrategy *string `json:"gpuSharingStrategy,omitempty" tf:"gpu_sharing_strategy,omitempty"`

	MaxSharedClientsPerGpu *float64 `json:"maxSharedClientsPerGpu,omitempty" tf:"max_shared_clients_per_gpu,omitempty"`
}

type NodeConfigGuestAcceleratorGpuSharingConfigParameters struct {

	// +kubebuilder:validation:Optional
	GpuSharingStrategy *string `json:"gpuSharingStrategy,omitempty" tf:"gpu_sharing_strategy"`

	// +kubebuilder:validation:Optional
	MaxSharedClientsPerGpu *float64 `json:"maxSharedClientsPerGpu,omitempty" tf:"max_shared_clients_per_gpu"`
}

type NodePoolAutoscalingObservation struct {

	// Location policy specifies the algorithm used when
	// scaling-up the node pool. Location policy is supported only in 1.24.1+ clusters.
	LocationPolicy *string `json:"locationPolicy,omitempty" tf:"location_policy,omitempty"`

	// Maximum number of nodes per zone in the NodePool.
	// Must be >= min_node_count. Cannot be used with total limits.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// Minimum number of nodes per zone in the NodePool.
	// Must be >=0 and <= max_node_count. Cannot be used with total limits.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Total maximum number of nodes in the NodePool.
	// Must be >= total_min_node_count. Cannot be used with per zone limits.
	// Total size limits are supported only in 1.24.1+ clusters.
	TotalMaxNodeCount *float64 `json:"totalMaxNodeCount,omitempty" tf:"total_max_node_count,omitempty"`

	// Total minimum number of nodes in the NodePool.
	// Must be >=0 and <= total_max_node_count. Cannot be used with per zone limits.
	// Total size limits are supported only in 1.24.1+ clusters.
	TotalMinNodeCount *float64 `json:"totalMinNodeCount,omitempty" tf:"total_min_node_count,omitempty"`
}

type NodePoolAutoscalingParameters struct {

	// Location policy specifies the algorithm used when
	// scaling-up the node pool. Location policy is supported only in 1.24.1+ clusters.
	// +kubebuilder:validation:Optional
	LocationPolicy *string `json:"locationPolicy,omitempty" tf:"location_policy,omitempty"`

	// Maximum number of nodes per zone in the NodePool.
	// Must be >= min_node_count. Cannot be used with total limits.
	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// Minimum number of nodes per zone in the NodePool.
	// Must be >=0 and <= max_node_count. Cannot be used with total limits.
	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Total maximum number of nodes in the NodePool.
	// Must be >= total_min_node_count. Cannot be used with per zone limits.
	// Total size limits are supported only in 1.24.1+ clusters.
	// +kubebuilder:validation:Optional
	TotalMaxNodeCount *float64 `json:"totalMaxNodeCount,omitempty" tf:"total_max_node_count,omitempty"`

	// Total minimum number of nodes in the NodePool.
	// Must be >=0 and <= total_max_node_count. Cannot be used with per zone limits.
	// Total size limits are supported only in 1.24.1+ clusters.
	// +kubebuilder:validation:Optional
	TotalMinNodeCount *float64 `json:"totalMinNodeCount,omitempty" tf:"total_min_node_count,omitempty"`
}

type NodePoolManagementObservation_2 struct {

	// Whether the nodes will be automatically repaired.
	AutoRepair *bool `json:"autoRepair,omitempty" tf:"auto_repair,omitempty"`

	// Whether the nodes will be automatically upgraded.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`
}

type NodePoolManagementParameters_2 struct {

	// Whether the nodes will be automatically repaired.
	// +kubebuilder:validation:Optional
	AutoRepair *bool `json:"autoRepair,omitempty" tf:"auto_repair,omitempty"`

	// Whether the nodes will be automatically upgraded.
	// +kubebuilder:validation:Optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`
}

type NodePoolNetworkConfigObservation struct {

	// Whether to create a new range for pod IPs in this node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.
	CreatePodRange *bool `json:"createPodRange,omitempty" tf:"create_pod_range,omitempty"`

	// Whether nodes have internal IP addresses only.
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty" tf:"enable_private_nodes,omitempty"`

	// The IP address range for pod IPs in this node pool. Only applicable if createPodRange is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
	PodIPv4CidrBlock *string `json:"podIpv4CidrBlock,omitempty" tf:"pod_ipv4_cidr_block,omitempty"`

	// The ID of the secondary range for pod IPs. If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.
	PodRange *string `json:"podRange,omitempty" tf:"pod_range,omitempty"`
}

type NodePoolNetworkConfigParameters struct {

	// Whether to create a new range for pod IPs in this node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.
	// +kubebuilder:validation:Optional
	CreatePodRange *bool `json:"createPodRange,omitempty" tf:"create_pod_range,omitempty"`

	// Whether nodes have internal IP addresses only.
	// +kubebuilder:validation:Optional
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty" tf:"enable_private_nodes,omitempty"`

	// The IP address range for pod IPs in this node pool. Only applicable if createPodRange is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
	// +kubebuilder:validation:Optional
	PodIPv4CidrBlock *string `json:"podIpv4CidrBlock,omitempty" tf:"pod_ipv4_cidr_block,omitempty"`

	// The ID of the secondary range for pod IPs. If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.
	// +kubebuilder:validation:Optional
	PodRange *string `json:"podRange,omitempty" tf:"pod_range,omitempty"`
}

type NodePoolNodeConfigGcfsConfigObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type NodePoolNodeConfigGcfsConfigParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type NodePoolNodeConfigGuestAcceleratorObservation struct {
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty" tf:"gpu_partition_size,omitempty"`

	GpuSharingConfig []NodeConfigGuestAcceleratorGpuSharingConfigObservation `json:"gpuSharingConfig,omitempty" tf:"gpu_sharing_config,omitempty"`

	// The type of the policy. Supports a single value: COMPACT.
	// Specifying COMPACT placement policy type places node pool's nodes in a closer
	// physical proximity in order to reduce network latency between nodes.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NodePoolNodeConfigGuestAcceleratorParameters struct {

	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count"`

	// +kubebuilder:validation:Optional
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty" tf:"gpu_partition_size"`

	// +kubebuilder:validation:Optional
	GpuSharingConfig []NodeConfigGuestAcceleratorGpuSharingConfigParameters `json:"gpuSharingConfig,omitempty" tf:"gpu_sharing_config"`

	// The type of the policy. Supports a single value: COMPACT.
	// Specifying COMPACT placement policy type places node pool's nodes in a closer
	// physical proximity in order to reduce network latency between nodes.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type NodePoolNodeConfigGvnicObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type NodePoolNodeConfigGvnicParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type NodePoolNodeConfigKubeletConfigObservation struct {
	CPUCfsQuota *bool `json:"cpuCfsQuota,omitempty" tf:"cpu_cfs_quota,omitempty"`

	CPUCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty" tf:"cpu_cfs_quota_period,omitempty"`

	CPUManagerPolicy *string `json:"cpuManagerPolicy,omitempty" tf:"cpu_manager_policy,omitempty"`

	PodPidsLimit *float64 `json:"podPidsLimit,omitempty" tf:"pod_pids_limit,omitempty"`
}

type NodePoolNodeConfigKubeletConfigParameters struct {

	// +kubebuilder:validation:Optional
	CPUCfsQuota *bool `json:"cpuCfsQuota,omitempty" tf:"cpu_cfs_quota,omitempty"`

	// +kubebuilder:validation:Optional
	CPUCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty" tf:"cpu_cfs_quota_period,omitempty"`

	// +kubebuilder:validation:Required
	CPUManagerPolicy *string `json:"cpuManagerPolicy" tf:"cpu_manager_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PodPidsLimit *float64 `json:"podPidsLimit,omitempty" tf:"pod_pids_limit,omitempty"`
}

type NodePoolNodeConfigLinuxNodeConfigObservation struct {
	Sysctls map[string]*string `json:"sysctls,omitempty" tf:"sysctls,omitempty"`
}

type NodePoolNodeConfigLinuxNodeConfigParameters struct {

	// +kubebuilder:validation:Required
	Sysctls map[string]*string `json:"sysctls" tf:"sysctls,omitempty"`
}

type NodePoolNodeConfigObservation_2 struct {
	BootDiskKMSKey *string `json:"bootDiskKmsKey,omitempty" tf:"boot_disk_kms_key,omitempty"`

	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	GcfsConfig []NodePoolNodeConfigGcfsConfigObservation `json:"gcfsConfig,omitempty" tf:"gcfs_config,omitempty"`

	GuestAccelerator []NodePoolNodeConfigGuestAcceleratorObservation `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	Gvnic []NodePoolNodeConfigGvnicObservation `json:"gvnic,omitempty" tf:"gvnic,omitempty"`

	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	KubeletConfig []NodePoolNodeConfigKubeletConfigObservation `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Parameters used in creating the node pool. See
	// google_container_cluster for schema.
	LinuxNodeConfig []NodePoolNodeConfigLinuxNodeConfigObservation `json:"linuxNodeConfig,omitempty" tf:"linux_node_config,omitempty"`

	LocalSsdCount *float64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`

	LoggingVariant *string `json:"loggingVariant,omitempty" tf:"logging_variant,omitempty"`

	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	NodeGroup *string `json:"nodeGroup,omitempty" tf:"node_group,omitempty"`

	OAuthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`

	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	ReservationAffinity []NodePoolNodeConfigReservationAffinityObservation `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	ResourceLabels map[string]*string `json:"resourceLabels,omitempty" tf:"resource_labels,omitempty"`

	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	ShieldedInstanceConfig []NodePoolNodeConfigShieldedInstanceConfigObservation_2 `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	Taint []NodePoolNodeConfigTaintObservation `json:"taint,omitempty" tf:"taint,omitempty"`

	WorkloadMetadataConfig []NodePoolNodeConfigWorkloadMetadataConfigObservation `json:"workloadMetadataConfig,omitempty" tf:"workload_metadata_config,omitempty"`
}

type NodePoolNodeConfigParameters_2 struct {

	// +kubebuilder:validation:Optional
	BootDiskKMSKey *string `json:"bootDiskKmsKey,omitempty" tf:"boot_disk_kms_key,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	GcfsConfig []NodePoolNodeConfigGcfsConfigParameters `json:"gcfsConfig,omitempty" tf:"gcfs_config,omitempty"`

	// +kubebuilder:validation:Optional
	GuestAccelerator []NodePoolNodeConfigGuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// +kubebuilder:validation:Optional
	Gvnic []NodePoolNodeConfigGvnicParameters `json:"gvnic,omitempty" tf:"gvnic,omitempty"`

	// +kubebuilder:validation:Optional
	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	// +kubebuilder:validation:Optional
	KubeletConfig []NodePoolNodeConfigKubeletConfigParameters `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Parameters used in creating the node pool. See
	// google_container_cluster for schema.
	// +kubebuilder:validation:Optional
	LinuxNodeConfig []NodePoolNodeConfigLinuxNodeConfigParameters `json:"linuxNodeConfig,omitempty" tf:"linux_node_config,omitempty"`

	// +kubebuilder:validation:Optional
	LocalSsdCount *float64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingVariant *string `json:"loggingVariant,omitempty" tf:"logging_variant,omitempty"`

	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// +kubebuilder:validation:Optional
	NodeGroup *string `json:"nodeGroup,omitempty" tf:"node_group,omitempty"`

	// +kubebuilder:validation:Optional
	OAuthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`

	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	// +kubebuilder:validation:Optional
	ReservationAffinity []NodePoolNodeConfigReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceLabels map[string]*string `json:"resourceLabels,omitempty" tf:"resource_labels,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate serviceAccount.
	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate serviceAccount.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []NodePoolNodeConfigShieldedInstanceConfigParameters_2 `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// +kubebuilder:validation:Optional
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Taint []NodePoolNodeConfigTaintParameters `json:"taint,omitempty" tf:"taint,omitempty"`

	// +kubebuilder:validation:Optional
	WorkloadMetadataConfig []NodePoolNodeConfigWorkloadMetadataConfigParameters `json:"workloadMetadataConfig,omitempty" tf:"workload_metadata_config,omitempty"`
}

type NodePoolNodeConfigReservationAffinityObservation struct {
	ConsumeReservationType *string `json:"consumeReservationType,omitempty" tf:"consume_reservation_type,omitempty"`

	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NodePoolNodeConfigReservationAffinityParameters struct {

	// +kubebuilder:validation:Required
	ConsumeReservationType *string `json:"consumeReservationType" tf:"consume_reservation_type,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NodePoolNodeConfigShieldedInstanceConfigObservation_2 struct {
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`
}

type NodePoolNodeConfigShieldedInstanceConfigParameters_2 struct {

	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`
}

type NodePoolNodeConfigTaintObservation struct {
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NodePoolNodeConfigTaintParameters struct {

	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type NodePoolNodeConfigWorkloadMetadataConfigObservation struct {
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type NodePoolNodeConfigWorkloadMetadataConfigParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type NodePoolObservation_2 struct {

	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling []NodePoolAutoscalingObservation `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`

	// The cluster to create the node pool for. Cluster must be present in location provided for clusters. May be specified in the format projects/{{project}}/locations/{{location}}/clusters/{{cluster}} or as just the name of the cluster.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// an identifier for the resource with format {{project}}/{{location}}/{{cluster}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource.  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsequent changes to this field.
	InitialNodeCount *float64 `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`

	// The resource URLs of the managed instance groups associated with this node pool.
	InstanceGroupUrls []*string `json:"instanceGroupUrls,omitempty" tf:"instance_group_urls,omitempty"`

	// The location (region or zone) of the cluster.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// List of instance group URLs which have been assigned to this node pool.
	ManagedInstanceGroupUrls []*string `json:"managedInstanceGroupUrls,omitempty" tf:"managed_instance_group_urls,omitempty"`

	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management []NodePoolManagementObservation_2 `json:"management,omitempty" tf:"management,omitempty"`

	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the official documentation
	// for more information.
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// The network configuration of the pool. Such as
	// configuration for Adding Pod IP address ranges) to the node pool. Or enabling private nodes. Structure is
	// documented below
	NetworkConfig []NodePoolNetworkConfigObservation `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Parameters used in creating the node pool. See
	// google_container_cluster for schema.
	NodeConfig []NodePoolNodeConfigObservation_2 `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside autoscaling.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// node_locations will be used.
	NodeLocations []*string `json:"nodeLocations,omitempty" tf:"node_locations,omitempty"`

	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// Specifies a custom placement policy for the
	// nodes.
	PlacementPolicy []NodePoolPlacementPolicyObservation `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specify node upgrade settings to change how GKE upgrades nodes.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	UpgradeSettings []NodePoolUpgradeSettingsObservation_2 `json:"upgradeSettings,omitempty" tf:"upgrade_settings,omitempty"`

	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and auto_upgrade are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodePoolParameters_2 struct {

	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	// +kubebuilder:validation:Optional
	Autoscaling []NodePoolAutoscalingParameters `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`

	// The cluster to create the node pool for. Cluster must be present in location provided for clusters. May be specified in the format projects/{{project}}/locations/{{location}}/clusters/{{cluster}} or as just the name of the cluster.
	// +crossplane:generate:reference:type=Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Reference to a Cluster to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource.  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsequent changes to this field.
	// +kubebuilder:validation:Optional
	InitialNodeCount *float64 `json:"initialNodeCount,omitempty" tf:"initial_node_count,omitempty"`

	// The location (region or zone) of the cluster.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	// +kubebuilder:validation:Optional
	Management []NodePoolManagementParameters_2 `json:"management,omitempty" tf:"management,omitempty"`

	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the official documentation
	// for more information.
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// The network configuration of the pool. Such as
	// configuration for Adding Pod IP address ranges) to the node pool. Or enabling private nodes. Structure is
	// documented below
	// +kubebuilder:validation:Optional
	NetworkConfig []NodePoolNetworkConfigParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Parameters used in creating the node pool. See
	// google_container_cluster for schema.
	// +kubebuilder:validation:Optional
	NodeConfig []NodePoolNodeConfigParameters_2 `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside autoscaling.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// node_locations will be used.
	// +kubebuilder:validation:Optional
	NodeLocations []*string `json:"nodeLocations,omitempty" tf:"node_locations,omitempty"`

	// Specifies a custom placement policy for the
	// nodes.
	// +kubebuilder:validation:Optional
	PlacementPolicy []NodePoolPlacementPolicyParameters `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specify node upgrade settings to change how GKE upgrades nodes.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	// +kubebuilder:validation:Optional
	UpgradeSettings []NodePoolUpgradeSettingsParameters_2 `json:"upgradeSettings,omitempty" tf:"upgrade_settings,omitempty"`

	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and auto_upgrade are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodePoolPlacementPolicyObservation struct {

	// The type of the policy. Supports a single value: COMPACT.
	// Specifying COMPACT placement policy type places node pool's nodes in a closer
	// physical proximity in order to reduce network latency between nodes.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NodePoolPlacementPolicyParameters struct {

	// The type of the policy. Supports a single value: COMPACT.
	// Specifying COMPACT placement policy type places node pool's nodes in a closer
	// physical proximity in order to reduce network latency between nodes.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NodePoolUpgradeSettingsBlueGreenSettingsObservation struct {

	// Time needed after draining the entire blue pool.
	// After this period, the blue pool will be cleaned up.
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty" tf:"node_pool_soak_duration,omitempty"`

	// Specifies the standard policy settings for blue-green upgrades.
	StandardRolloutPolicy []UpgradeSettingsBlueGreenSettingsStandardRolloutPolicyObservation `json:"standardRolloutPolicy,omitempty" tf:"standard_rollout_policy,omitempty"`
}

type NodePoolUpgradeSettingsBlueGreenSettingsParameters struct {

	// Time needed after draining the entire blue pool.
	// After this period, the blue pool will be cleaned up.
	// +kubebuilder:validation:Optional
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty" tf:"node_pool_soak_duration,omitempty"`

	// Specifies the standard policy settings for blue-green upgrades.
	// +kubebuilder:validation:Required
	StandardRolloutPolicy []UpgradeSettingsBlueGreenSettingsStandardRolloutPolicyParameters `json:"standardRolloutPolicy" tf:"standard_rollout_policy,omitempty"`
}

type NodePoolUpgradeSettingsObservation_2 struct {

	// The settings to adjust blue green upgrades.
	// Structure is documented below
	BlueGreenSettings []NodePoolUpgradeSettingsBlueGreenSettingsObservation `json:"blueGreenSettings,omitempty" tf:"blue_green_settings,omitempty"`

	// The number of additional nodes that can be added to the node pool during
	// an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously.
	// Can be set to 0 or greater.
	MaxSurge *float64 `json:"maxSurge,omitempty" tf:"max_surge,omitempty"`

	// The number of nodes that can be simultaneously unavailable during
	// an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in
	// parallel. Can be set to 0 or greater.
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// (Default SURGE) The upgrade stragey to be used for upgrading the nodes.
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type NodePoolUpgradeSettingsParameters_2 struct {

	// The settings to adjust blue green upgrades.
	// Structure is documented below
	// +kubebuilder:validation:Optional
	BlueGreenSettings []NodePoolUpgradeSettingsBlueGreenSettingsParameters `json:"blueGreenSettings,omitempty" tf:"blue_green_settings,omitempty"`

	// The number of additional nodes that can be added to the node pool during
	// an upgrade. Increasing max_surge raises the number of nodes that can be upgraded simultaneously.
	// Can be set to 0 or greater.
	// +kubebuilder:validation:Optional
	MaxSurge *float64 `json:"maxSurge,omitempty" tf:"max_surge,omitempty"`

	// The number of nodes that can be simultaneously unavailable during
	// an upgrade. Increasing max_unavailable raises the number of nodes that can be upgraded in
	// parallel. Can be set to 0 or greater.
	// +kubebuilder:validation:Optional
	MaxUnavailable *float64 `json:"maxUnavailable,omitempty" tf:"max_unavailable,omitempty"`

	// (Default SURGE) The upgrade stragey to be used for upgrading the nodes.
	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type UpgradeSettingsBlueGreenSettingsStandardRolloutPolicyObservation struct {

	// Number of blue nodes to drain in a batch.
	BatchNodeCount *float64 `json:"batchNodeCount,omitempty" tf:"batch_node_count,omitempty"`

	// Percentage of the blue pool nodes to drain in a batch.
	BatchPercentage *float64 `json:"batchPercentage,omitempty" tf:"batch_percentage,omitempty"`

	// (Optionial) Soak time after each batch gets drained.
	BatchSoakDuration *string `json:"batchSoakDuration,omitempty" tf:"batch_soak_duration,omitempty"`
}

type UpgradeSettingsBlueGreenSettingsStandardRolloutPolicyParameters struct {

	// Number of blue nodes to drain in a batch.
	// +kubebuilder:validation:Optional
	BatchNodeCount *float64 `json:"batchNodeCount,omitempty" tf:"batch_node_count,omitempty"`

	// Percentage of the blue pool nodes to drain in a batch.
	// +kubebuilder:validation:Optional
	BatchPercentage *float64 `json:"batchPercentage,omitempty" tf:"batch_percentage,omitempty"`

	// (Optionial) Soak time after each batch gets drained.
	// +kubebuilder:validation:Optional
	BatchSoakDuration *string `json:"batchSoakDuration,omitempty" tf:"batch_soak_duration,omitempty"`
}

// NodePoolSpec defines the desired state of NodePool
type NodePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolParameters_2 `json:"forProvider"`
}

// NodePoolStatus defines the observed state of NodePool.
type NodePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePool is the Schema for the NodePools API. Manages a GKE NodePool resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodePoolSpec   `json:"spec"`
	Status            NodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolList contains a list of NodePools
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

// Repository type metadata.
var (
	NodePool_Kind             = "NodePool"
	NodePool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePool_Kind}.String()
	NodePool_KindAPIVersion   = NodePool_Kind + "." + CRDGroupVersion.String()
	NodePool_GroupVersionKind = CRDGroupVersion.WithKind(NodePool_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}