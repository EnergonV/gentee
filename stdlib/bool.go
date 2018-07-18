// Copyright 2018 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package stdlib

import (
	"github.com/gentee/gentee/core"
)

// InitBool appends stdlib bool functions to the virtual machine
func InitBool(vm *core.VirtualMachine) {
	for _, item := range []interface{}{
		Not, // unary boolean not
	} {
		vm.Units[core.DefName].NewEmbed(item)
	}
}

// Not changes true to false or false to true
func Not(val bool) bool {
	return !val
}