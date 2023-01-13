package v1alpha1

const (
	// AnnotationPodCPUPolicyKey is the pod annotation key of cpu-policy.
	AnnotationPodCPUPolicyKey = "topology.crane.io/cpu-policy"

	// AnnotationPodTopologyAwarenessKey is the pod annotation key of topology-awareness.
	AnnotationPodTopologyAwarenessKey = "topology.crane.io/topology-awareness"

	// AnnotationPodTopologyResultKey is the pod scheduling annotation key of topology-result.
	AnnotationPodTopologyResultKey = "topology.crane.io/topology-result"

	// AnnotationPodExcludeReservedCPUs is the pod annotation key of exclude reserved cpus
	AnnotationPodExcludeReservedCPUs = "topology.crane.io/exclude-reserved-cpus"
)

const (
	// AnnotationPodCPUPolicyNone specifies none cpu policy. If specified, pod
	// will use the default CPUSet.
	AnnotationPodCPUPolicyNone = "none"

	// AnnotationPodCPUPolicyExclusive specifies exclusive cpu policy.  If specified,
	// pod will never share CPUSet with others.
	AnnotationPodCPUPolicyExclusive = "exclusive"

	// AnnotationPodCPUPolicyNUMA specifies NUMA cpu policy. If specified, pod
	// will use the default CPUSet which belongs to single NUMA node.
	AnnotationPodCPUPolicyNUMA = "numa"

	// AnnotationPodCPUPolicyImmovable specifies immovable cpu policy. If specified,
	// pod will use part of the default CPUSet to avoid uncertain context switch.
	AnnotationPodCPUPolicyImmovable = "immovable"
)
