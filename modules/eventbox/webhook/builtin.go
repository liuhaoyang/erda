package webhook

import (
	"fmt"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/eventbox/conf"
)

// 如果不存在相同名字的 webhook 则创建
func createIfNotExist(impl *WebHookImpl, req *CreateHookRequest) error {
	if impl == nil {
		return fmt.Errorf("nil webhookImpl")
	}
	resp, err := impl.ListHooks(apistructs.HookLocation{
		Org:         req.Org,
		Project:     req.Project,
		Application: req.Application,
	})
	if err != nil {
		return err
	}
	hooks := []apistructs.Hook(resp)
	for i := range hooks {
		if hooks[i].Name == req.Name {
			return nil
		}
	}
	if _, err := impl.CreateHook(req.Org, *req); err != nil {
		return err
	}
	return nil
}

// MakeSureBuiltinHooks 创建默认 webhook (如果不存在)
func MakeSureBuiltinHooks(impl *WebHookImpl) error {
	domainSuffix := map[bool]string{true: ".default.svc.cluster.local", false: ".marathon.l4lb.thisdcos.directory"}
	hooks := []CreateHookRequest{
		{
			Name:   "scheduler-clusterhook",
			Events: []string{"cluster"},
			URL:    fmt.Sprintf("http://scheduler%s:9091/clusterhook", domainSuffix[conf.UseK8S()]),
			Active: true,
			HookLocation: apistructs.HookLocation{
				Org:         "-1",
				Project:     "-1",
				Application: "-1",
			},
		},
	}

	for i := range hooks {
		if err := createIfNotExist(impl, &hooks[i]); err != nil {
			return err
		}
	}
	return nil
}