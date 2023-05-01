package example

import (
	"context"
	"fmt"
	bootloader "go-bootloader"
	"go-bootloader/common"
	"go-bootloader/ctx"
	"go-bootloader/model"
	"testing"
)

type Loader1 struct {
}

func (l Loader1) Load(ctx ctx.LoaderContext) error {
	return nil
}

func (l Loader1) Require() []string {
	return []string{}
}

type Loader2 struct {
}

func (l Loader2) Load(ctx ctx.LoaderContext) error {
	return nil
}

func (l Loader2) Require() []string {
	return []string{"Loader1"}
}

type Loader3 struct {
}

func (l Loader3) Load(ctx ctx.LoaderContext) error {
	return nil
}

func (l Loader3) Require() []string {
	return []string{"Loader2"}
}

type Loader4 struct {
}

func (l Loader4) Load(ctx ctx.LoaderContext) error {
	return nil
}

func (l Loader4) Require() []string {
	return []string{"Loader2"}
}

type Loader5 struct {
}

func (l Loader5) Load(ctx ctx.LoaderContext) error {
	return nil
}

func (l Loader5) Require() []string {
	return []string{"Loader3", "Loader4"}
}

type Loader6 struct {
}

func (l Loader6) Load(ctx ctx.LoaderContext) error {
	// you can get project root dir or project name, like this:
	fmt.Println(ctx.GetProjectName()) // only test mode valid, or output ""
	fmt.Println(ctx.GetProjectRoot())

	// you can get load mode and judge by common.ModeXXX, like this:
	switch ctx.GetMode() {
	case common.ModeTest:
		fmt.Println("test mode")
	case common.ModeDev:
		fmt.Println("dev mode")
	case common.ModeProd:
		fmt.Println("prod mode")
	}
	// or like this:
	if ctx.IsTestMode() {
		fmt.Println("test mode 2")
	}

	return nil
}

func (l Loader6) Require() []string {
	return []string{common.Finally}
}

/**
 * BaseLoader(first)  Loader1 <- Loader2 <- Loader3
                               ^          ^
							   |          |
							Loader4 <- Loader5   Loader6(final)
*/

func TestLoadAll(t *testing.T) {
	loaderList := model.LoaderList{
		&Loader1{},
		&Loader2{},
		&Loader3{},
		&Loader4{},
		&Loader5{},
		&Loader6{},
	}
	bootloader.TestMode("go-bootloader")
	bootloader.Load(context.TODO(), loaderList)
}
