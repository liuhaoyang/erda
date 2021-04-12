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

package orchestrator

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var ORCHESTRATOR_MIDDLEWARE_FETCH_RESOURCE = apis.ApiSpec{
	Path:         "/api/middlewares/<addonId>/actions/get-resource",
	BackendPath:  "/api/middlewares/<addonId>/actions/get-resource",
	Host:         "orchestrator.marathon.l4lb.thisdcos.directory:8081",
	Scheme:       "http",
	Method:       "GET",
	ResponseType: apistructs.MiddlewareResourceFetchResponse{},
	CheckLogin:   true,
	CheckToken:   true,
	IsOpenAPI:    true,
	Doc:          `summary: middleware 资源详情`,
	Group:        "addons",
}
