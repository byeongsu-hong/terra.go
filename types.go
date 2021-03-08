package terra

type Q map[string]interface{}

type BroadcastMode string

const (
	ModeBlock BroadcastMode = "block"
	ModeSync  BroadcastMode = "sync"
	ModeAsync BroadcastMode = "async"
)
