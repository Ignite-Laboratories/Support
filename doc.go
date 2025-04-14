// Package support contains miscellaneous functions that are too general for a single project in JanOS.
package support

import "github.com/ignite-laboratories/core"

var ModuleName = "support"

func init() {
	core.ModuleReport(ModuleName)
}

func Report() {}
