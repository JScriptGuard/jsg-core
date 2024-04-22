package kernel

import (
	"sync"
	"vnh1/core/consolecache"
	"vnh1/core/vmdb"
	"vnh1/types"

	v8 "rogchap.com/v8go"
)

type KernelConfig struct {
	Modules []types.KernelModuleInterface
}

type Kernel struct {
	*v8.Context
	id        types.KernelID
	config    *KernelConfig
	mutex     *sync.Mutex
	core      types.CoreInterface
	console   *consolecache.ConsoleOutputCache
	vmLink    types.CoreVMInterface
	register  map[string]interface{}
	vmImports map[string]*v8.Value
	dbEntry   *vmdb.VmDBEntry
}