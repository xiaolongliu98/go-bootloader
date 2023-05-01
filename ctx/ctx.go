package ctx

import (
	"go-bootloader/common"
	"os"
)

// LoaderContext set/get from os.SetEnv/os.GetEnv
type LoaderContext struct {
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

// IsDisableLog return true if BOOTLOADER-DISABLE_PRINT is set
func (ctx *LoaderContext) IsDisableLog() bool {
	return os.Getenv(common.ConfigDisableLog) != ""
}

// GetMode return BOOTLOADER-MODE
func (ctx *LoaderContext) GetMode() string {
	return os.Getenv(common.ConfigMode)
}

// GetProjectName return BOOTLOADER-PROJECT_NAME
func (ctx *LoaderContext) GetProjectName() string {
	return os.Getenv(common.ConfigProjectName)
}
