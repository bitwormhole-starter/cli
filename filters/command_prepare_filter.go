package filters

import (
	"strings"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/arguments"
)

// CommandPrepareFilter ...
type CommandPrepareFilter struct {
}

func (inst *CommandPrepareFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *CommandPrepareFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "CommandPrepareFilter",
		Order:  OrderCommandPrepare,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *CommandPrepareFilter) Pass(task *cli.Task, chain cli.FilterChain) error {

	a1 := arguments.Parse(task.Command)
	a2 := task.Arguments

	// 混合在一起
	ax := make([]string, 0)
	ax = append(ax, a1...)
	ax = append(ax, a2...)

	// 去除空字符串
	dst := make([]string, 0)
	for _, item := range ax {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		dst = append(dst, item)
	}

	// 回填到 Task
	if len(dst) > 0 {
		task.Command = dst[0]
		task.Arguments = dst[1:]
	} else {
		task.Command = ""
		task.Arguments = []string{}
	}

	return chain.Pass(task)
}
