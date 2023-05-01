package ctx

import (
	"context"
	"github.com/xiaolongliu98/go-bootloader/common"
	"os"
)

// LoaderContext set/get from os.SetEnv/os.GetEnv
type LoaderContext struct {
	ctx context.Context
}

func (ctx *LoaderContext) SetContext(context context.Context) {
	ctx.ctx = context
}

func (ctx *LoaderContext) Context() context.Context {
	return ctx.ctx
}

func (ctx *LoaderContext) Get(key string) string {
	return os.Getenv(key)
}

func (ctx *LoaderContext) Set(key string, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic(err)
	}
}

func (ctx *LoaderContext) Del(key string) {
	err := os.Unsetenv(key)
	if err != nil {
		panic(err)
	}
}

func (ctx *LoaderContext) GetProjectRoot() string {
	return os.Getenv(common.ConfigProjectRoot)
}

// IsDisableSort return true if BOOTLOADER-DISABLE_SORT is set
func (ctx *LoaderContext) IsDisableSort() bool {
	return os.Getenv(common.ConfigDisableSort) != ""
}

// IsDisablePrint return true if BOOTLOADER-DISABLE_PRINT is set
func (ctx *LoaderContext) IsDisablePrint() bool {
	return os.Getenv(common.ConfigDisablePrint) != ""
}

// GetMode return BOOTLOADER-MODE
func (ctx *LoaderContext) GetMode() string {
	return os.Getenv(common.ConfigMode)
}

// GetProjectName return BOOTLOADER-PROJECT_NAME
func (ctx *LoaderContext) GetProjectName() string {
	return os.Getenv(common.ConfigProjectName)
}

// IsTestMode return true if BOOTLOADER-MODE is set to "test"
func (ctx *LoaderContext) IsTestMode() bool {
	return ctx.GetMode() == common.ModeTest
}

// IsProdMode return true if BOOTLOADER-MODE is set to "prod"
func (ctx *LoaderContext) IsProdMode() bool {
	return ctx.GetMode() == common.ModeProd
}

// IsDevMode return true if BOOTLOADER-MODE is set to "dev"
func (ctx *LoaderContext) IsDevMode() bool {
	return ctx.GetMode() == common.ModeDev
}
