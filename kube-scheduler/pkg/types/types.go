type PredicateFailureReason interface {
	GetReason() string
	GetInfo() (ResourceName, int64, int64, int64)
}

// used by scheduler
type DeviceScheduler interface {
	// see if pod fits on node & return device score
	PodFitsDevice(nodeInfo *NodeInfo, podInfo *PodInfo, fillAllocateFrom bool, runGrpScheduler bool) (bool, []PredicateFailureReason, float64)
	// allocate resources
	PodAllocate(nodeInfo *NodeInfo, podInfo *PodInfo, runGrpScheduler bool) error
	// take resources from node
	TakePodResources(*NodeInfo, *PodInfo, bool) error
	// return resources to node
	ReturnPodResources(*NodeInfo, *PodInfo, bool) error
	// GetName returns the name of a device
	GetName() string
	// Tells whether group scheduler is being used?
	UsingGroupScheduler() bool
}