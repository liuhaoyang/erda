package scene_after

import (
	"strconv"
	"strings"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/aop/aoptypes"
)

type Plugin struct {
	aoptypes.PipelineBaseTunePoint
}

func (p *Plugin) Name() string {
	return "scene_after"
}

func (p *Plugin) Handle(ctx aoptypes.TuneContext) error {
	// source = autotest
	if ctx.SDK.Pipeline.PipelineSource == apistructs.PipelineSourceAutoTest && !ctx.SDK.Pipeline.IsSnippet {
		if strings.HasPrefix(ctx.SDK.Pipeline.PipelineYmlName, apistructs.PipelineSourceAutoTestPlan.String()+"-") {
			return nil
		}
		sceneID, err := strconv.ParseUint(ctx.SDK.Pipeline.PipelineYmlName, 10, 64)
		if err != nil {
			return err
		}
		var status apistructs.SceneStatus
		if ctx.SDK.Pipeline.Status.IsSuccessStatus() {
			status = apistructs.SuccessSceneStatus
		}
		if ctx.SDK.Pipeline.Status.IsFailedStatus() {
			status = apistructs.ErrorSceneStatus
		}
		if ctx.SDK.Pipeline.Status.IsReconcilerRunningStatus() {
			status = apistructs.ProcessingSceneStatus
		}
		var req apistructs.AutotestSceneRequest
		req.SceneID = sceneID
		req.UserID = ctx.SDK.Pipeline.PipelineExtra.Snapshot.PlatformSecrets["dice.user.id"]
		scene, err := ctx.SDK.Bundle.GetAutoTestScene(req)
		if err != nil {
			return err
		}
		req2 := apistructs.AutotestSceneSceneUpdateRequest{
			SceneID:     scene.ID,
			Description: scene.Description,
			Status:      status,
			IsStatus:    true,
		}
		req2.UserID = req.UserID
		_, err = ctx.SDK.Bundle.UpdateAutoTestScene(req2)
		if err != nil {
			return err
		}
	}
	return nil
}

func New() *Plugin {
	var p Plugin
	return &p
}
