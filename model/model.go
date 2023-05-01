package model

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
