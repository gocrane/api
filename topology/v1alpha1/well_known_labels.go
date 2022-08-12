package v1alpha1

const (
	// LabelNodeTopologyAwarenessKey is the node label key of topology-awareness.
	// This label is used to determine the default topology awareness policy of
	// a node when pod does not specify topology-awareness.
	LabelNodeTopologyAwarenessKey = "topology.crane.io/topology-awareness"
)
