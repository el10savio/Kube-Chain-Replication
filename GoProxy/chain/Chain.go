package chain

// Initializers ...
type Initializers struct {
	HEAD string `json:"head"`
	TAIL string `json:"tail"`
}

// NodeStatus ...
type NodeStatus struct {
	Node    string `json:"node"`
	Healthy bool   `json:"healthy"`
}
