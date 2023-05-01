package common

import (
	"go-bootloader/ctx"
	"go-bootloader/util"
)

type Loader interface {
	Load(ctx ctx.LoaderContext) error // 加载方法
	Require() []string                // 需返回该加载器依赖的加载器名称列表
}

type LoaderList []Loader

func (ll LoaderList) Contains(name string) bool {
	for _, loader := range ll {
		if util.GetTypeName(loader) == name {
			return true
		}
	}
	return false
}

func (ll LoaderList) Get(name string) Loader {
	for _, loader := range ll {
		if util.GetTypeName(loader) == name {
			return loader
		}
	}
	return nil
}

// load mode
const (
	ModeTest = "0"
	ModeProd = "1"
	ModeDev  = "2" // default
)

// sort tag
const (
	Finally = "FINALLY"
	Firstly = "FIRSTLY"
)

// bootloader config
const (
	ConfigDisableLog  = "BOOTLOADER-DISABLE_PRINT"
	ConfigDisableSort = "BOOTLOADER-DISABLE_SORT"
	ConfigMode        = "BOOTLOADER-MODE"
	ConfigProjectName = "BOOTLOADER-PROJECT_NAME"
	ConfigProjectRoot = "BOOTLOADER-PROJECT_ROOT"
)
