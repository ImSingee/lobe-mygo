package agentavatar

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("agent-avatar/basic", Basic)
	registry.Add("agent-avatar/states", States)
}
