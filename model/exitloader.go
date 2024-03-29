package model

import "github.com/xiaolongliu98/go-bootloader/ctx"

type ExitLoader interface {
	Loader
	BeforeExit(ctx ctx.LoaderContext)
}
