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

package config_manage

import "github.com/erda-project/erda/modules/openapi/api/apis"

var CONFIG_MANAGE_CONFIG_EXPORT = apis.ApiSpec{
	Path:        "/api/config/actions/export",
	BackendPath: "/api/config/actions/export",
	Host:        "cmdb.marathon.l4lb.thisdcos.directory:9093",
	Scheme:      "http",
	Method:      "GET",
	CheckLogin:  true,
	Doc:         "导出配置",
}
