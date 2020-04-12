package chain

// Initializers describes the 
// HEAD and TAIL of the chain
type Initializers struct {
	HEAD string `json:"head"`
	TAIL string `json:"tail"`
}

// NodeStatus describes the status 
// of each node whether it 
// is healthy or not
type NodeStatus struct {
	Node    string `json:"node"`
	Healthy bool   `json:"healthy"`
}
