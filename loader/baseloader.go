package loader

import (
	"fmt"
	"github.com/xiaolongliu98/go-bootloader/common"
	"github.com/xiaolongliu98/go-bootloader/ctx"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BaseLoader 相关依赖：globals.ProjectRoot
// 1.初始化项目根目录globals.ProjectRoot
type BaseLoader struct {
}

func (loader *BaseLoader) Require() []string {
	return []string{common.Firstly}
}

func (loader *BaseLoader) Load(ctx ctx.LoaderContext) error {
	rand.Seed(time.Now().UnixNano())

	// run || debug || build&./target.exe：wd得到的是main.go/target.exe所在目录
	// test时候的到的wd是 test文件所在目录
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if mode := os.Getenv(common.ConfigMode); mode == common.ModeTest {
		limits := 100

		projectName := ctx.GetProjectName()
		// check projectName
		if projectName == "" {
			return fmt.Errorf("Test模式运行必须设置项目目录名称")
		}

		for ; limits > 0 && !strings.HasSuffix(wd, projectName); limits-- {
			wd = filepath.Dir(wd)
		}
		if limits <= 0 {
			return fmt.Errorf("未找到项目根目录名：%s", projectName)
		}

		ctx.Set(common.ConfigProjectRoot, wd)
		return nil
	}
	ctx.Set(common.ConfigProjectRoot, wd)
	return nil
}
