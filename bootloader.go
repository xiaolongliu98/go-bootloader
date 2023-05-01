package bootloader

import (
	"context"
	"fmt"
	"go-bootloader/common"
	"go-bootloader/ctx"
	"go-bootloader/loader"
	"go-bootloader/util"
	"log"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	once      sync.Once
	loaderCtx ctx.LoaderContext
)

func init() {
	once.Do(func() {
		DevMode()
		loaderCtx = ctx.LoaderContext{}
	})
}

func DisablePrint() {
	loaderCtx.Set(common.ConfigDisableLog, "yes")
}

func DisableSort() {
	loaderCtx.Set(common.ConfigDisableSort, "yes")
}

// DevMode set mode to dev, default
func DevMode() {
	loaderCtx.Set(common.ConfigMode, common.ModeDev)
}

// ProdMode set mode to prod
func ProdMode() {
	loaderCtx.Set(common.ConfigMode, common.ModeProd)
}

// TestMode set mode to test
func TestMode(projectName string) {
	loaderCtx.Set(common.ConfigMode, common.ModeTest)
	loaderCtx.Set(common.ConfigProjectName, projectName)
}

// Load 启动加载
func Load(ctx context.Context, list common.LoaderList) {
	// check list contains BaseLoader
	if !list.Contains("BaseLoader") {
		// insert first
		list = append(common.LoaderList{&loader.BaseLoader{}}, list...)
	}

	once.Do(func() {
		t := time.Now()

		if !loaderCtx.IsDisableSort() {
			list = topoSort(list)
		}

		if !loaderCtx.IsDisableLog() {
			log.Printf("[Bootloader] 预计加载模块：\n")
			for i, loader := range list {
				log.Printf("[Bootloader] Module(%v/%v): %v\n",
					i+1,
					len(list),
					util.GetTypeName(loader))
			}
		}

		if !loaderCtx.IsDisableLog() {
			log.Printf("[Bootloader] start (total: %v)...\n", len(list))
		}

		for i, loader := range list {
			//log.Printf("[Bootloaders] start to load: %v\n", util.GetTypeName(loader))
			t := time.Now()
			if err := loader.Load(loaderCtx); err != nil {
				panicInfo := fmt.Sprintf("[%s]%s", util.GetTypeName(loader), err.Error())
				panic(panicInfo)
			}

			if !loaderCtx.IsDisableLog() {
				log.Printf("[Bootloader] successfully loaded(%v/%v, time: %vms): %v\n",
					i+1,
					len(list),
					time.Since(t).Milliseconds(),
					util.GetTypeName(loader))
			}
		}

		if !loaderCtx.IsDisableLog() {
			log.Printf("[Bootloader] load finished (total: %v), time: %vms\n",
				len(list),
				time.Since(t).Milliseconds())
		}
	})
}

// GetMinLoaderList 获取目标最小的loader列表
func GetMinLoaderList(target common.Loader, all common.LoaderList) common.LoaderList {
	list := NewLoaderListWithBaseLoader()
	q := common.LoaderList{target}
	// requires is a DAG, so we can use BFS to get the minimal target list
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if !list.Contains(util.GetTypeName(cur)) {
				list = append(list, cur)
			}

			for _, require := range cur.Require() {
				if l := all.Get(require); l != nil && !q.Contains(util.GetTypeName(l)) {
					q = append(q, l)
				}
			}
		}
	}
	return list
}

// NewLoaderListWithBaseLoader 生成一个包含BaseLoader的LoaderList
func NewLoaderListWithBaseLoader() common.LoaderList {
	return common.LoaderList{
		&loader.BaseLoader{},
	}
}

// topoSort 根据require拓扑排序
func topoSort(loaderList common.LoaderList) common.LoaderList {
	name2Loader := map[string]common.Loader{}
	Q := make([]string, 0, len(loaderList))
	var finalLoader common.Loader
	var firstLoader common.Loader

	// 建立依赖图结构
	// loader的入边（表示被依赖的）
	beRequiredMap := map[string]map[string]struct{}{}
	// loader的出边（表示依赖的）
	requireMap := map[string]map[string]struct{}{}

	for _, loader := range loaderList {
		// 检查重复
		if _, ok := name2Loader[util.GetTypeName(loader)]; ok {
			panic("重复的Loader, name:" + util.GetTypeName(loader))
		}
		name2Loader[util.GetTypeName(loader)] = loader
		requireList := loader.Require()

		// 检查特殊标记[FIRSTLY]
		if len(requireList) != 0 && strings.ToUpper(requireList[0]) == common.Firstly {
			if firstLoader != nil {
				panic("Firstly标记不唯一，请检查：[" + util.GetTypeName(firstLoader) + "]以及[" + util.GetTypeName(loader) + "]")
			}
			firstLoader = loader
			continue
		}
		// 检查特殊标记[FINALLY]
		if len(requireList) != 0 && strings.ToUpper(requireList[0]) == common.Finally {
			if finalLoader != nil {
				panic("Finally标记不唯一，请检查：[" + util.GetTypeName(finalLoader) + "]以及[" + util.GetTypeName(loader) + "]")
			}
			finalLoader = loader
			continue
		}

		// 将没有依赖的loader加入队列
		if len(requireList) == 0 {
			Q = append(Q, util.GetTypeName(loader))
		}
	}

	for _, loader := range loaderList {
		// 检查特殊标记[FIRSTLY]
		if firstLoader != nil && util.GetTypeName(loader) == util.GetTypeName(firstLoader) {
			continue
		}
		// 检查特殊标记[FINALLY]
		if finalLoader != nil && util.GetTypeName(loader) == util.GetTypeName(finalLoader) {
			continue
		}

		requireList := loader.Require()
		// 处理出度
		require := requireMap[util.GetTypeName(loader)]
		if require == nil {
			require = map[string]struct{}{}
		}
		for _, targetLoaderName := range requireList {
			if l, ok := name2Loader[targetLoaderName]; !ok || l == nil {
				errMsg := fmt.Sprintf("缺少或未找到%v，因此无法加载%v", targetLoaderName, util.GetTypeName(loader))
				panic(errMsg)
			}
			require[targetLoaderName] = struct{}{}
		}
		requireMap[util.GetTypeName(loader)] = require

		// 处理入度
		for _, targetLoaderName := range requireList {
			beRequired := beRequiredMap[targetLoaderName]
			if beRequired == nil {
				beRequired = map[string]struct{}{}
			}
			beRequired[util.GetTypeName(loader)] = struct{}{}
			beRequiredMap[targetLoaderName] = beRequired
		}
	}

	// 拓扑排序BFS
	result := make([]string, 0, len(loaderList))
	for len(Q) != 0 {
		size := len(Q)
		var subResult []string
		for i := 0; i < size; i++ {
			cur := Q[0]
			Q = Q[1:]
			subResult = append(subResult, cur)

			for targetName, _ := range beRequiredMap[cur] {
				// 删除target loader的该出度
				delete(requireMap[targetName], cur)
				if len(requireMap[targetName]) == 0 {
					Q = append(Q, targetName)
				}
			}
		}
		// 排序是为了稳定
		sort.Strings(subResult)
		result = append(result, subResult...)
	}

	sorted := make(common.LoaderList, 0, len(result)+1)

	for _, loaderName := range result {
		loader := name2Loader[loaderName]
		sorted = append(sorted, loader)
	}
	// 加入firstLoader
	if firstLoader != nil {
		sorted = append([]common.Loader{firstLoader}, sorted...)
	}
	// 加入finalLoader
	if finalLoader != nil {
		sorted = append(sorted, finalLoader)
	}

	return sorted
}
