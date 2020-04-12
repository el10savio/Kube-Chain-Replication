package store

// Entry is a struct representation 
// of the Redis store
type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
