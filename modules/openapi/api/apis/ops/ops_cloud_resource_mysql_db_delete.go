// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ops

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var OPS_CLOUD_RESOURCE_MYSQL_DB_DELETE = apis.ApiSpec{
	Path:         "/api/cloud-mysql/actions/delete-db",
	BackendPath:  "/api/cloud-mysql/actions/delete-db",
	Host:         "ops.marathon.l4lb.thisdcos.directory:9027",
	Scheme:       "http",
	Method:       "DELETE",
	CheckLogin:   true,
	RequestType:  apistructs.DeleteCloudResourceMysqlDBRequest{},
	ResponseType: apistructs.CloudAddonResourceDeleteRespnse{},
	Doc:          "删除 mysql database",
	Audit: func(ctx *spec.AuditContext) error {
		var request apistructs.DeleteCloudResourceMysqlDBRequest
		if err := ctx.BindRequestData(&request); err != nil {
			return err
		}

		project, err := ctx.GetProject(request.ProjectID)
		if err != nil {
			return err
		}

		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.ProjectScope,
			ScopeID:      project.ID,
			ProjectID:    project.ID,
			TemplateName: apistructs.DeleteMysqlDbTemplate,
			Context: map[string]interface{}{
				"projectName": project.Name,
				"addonID":     request.AddonID,
			},
		})
	},
}
