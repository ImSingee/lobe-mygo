package dropdownmenu

import "github.com/lobehub/lobe-mygo/catalog/registry"

func init() {
	registry.Add("dropdown-menu/basic", Basic)
	registry.Add("dropdown-menu/action-icon", ActionIconTrigger)
}
