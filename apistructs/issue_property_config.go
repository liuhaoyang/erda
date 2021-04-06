// Copyright 2021 Terminus
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

package apistructs

// 配置字段请求
type IssuePropertyConfigCreateRequest struct {
	ScopeID           int64             `json:"scopeID"`           // 系统管理员(sys)/企业(org)/项目(project)/应用(app)
	ScopeType         ScopeType         `json:"scopeType"`         // 企业ID/项目ID/应用ID
	OrgID             int64             `json:"orgID"`             // 企业ID
	ProjectID         int64             `json:"projectID"`         // 项目ID
	PropertyName      string            `json:"propertyName"`      // 属性名称
	PropertyType      PropertyType      `json:"propertyType"`      // 属性类型
	Required          bool              `json:"required"`          // 是否必填
	PropertyIssueType PropertyIssueType `json:"propertyIssueType"` // 任务类型
	IdentityInfo
}

type IssuePropertyConfig struct {
	ID                int64             `json:"id"`
	PropertyID        int64             `json:"propertyID"`        // 字段ID
	ScopeID           int64             `json:"scopeID"`           // 系统管理员(sys)/企业(org)/项目(project)/应用(app)
	ScopeType         ScopeType         `json:"scopeType"`         // 企业ID/项目ID/应用ID
	OrgID             int64             `json:"orgID"`             // 企业ID
	ProjectID         int64             `json:"projectID"`         // 项目ID
	Index             int64             `json:"index"`             // 排序级
	PropertyIssueType PropertyIssueType `json:"propertyIssueType"` // 任务类型
}

// 更新配置字段请求
type IssuePropertyConfigUpdateRequest struct {
	Data              []IssuePropertyConfig
	OrgID             int64             `json:"orgID"`             // 企业ID
	ProjectID         int64             `json:"projectID"`         // 项目ID
	PropertyIssueType PropertyIssueType `json:"propertyIssueType"` // 任务类型
	IdentityInfo
}

// 删除配置字段请求
type IssuePropertyConfigDeleteRequest struct {
	ConfigID int64 `json:"configID"` // 字段ID
	IdentityInfo
}

// 查询项目配置字段请求
type IssuePropertyConfigsGetRequest struct {
	ProjectID         int64             `json:"projectID"`         // 项目ID
	OrgID             int64             `json:"orgID"`             // 企业ID
	PropertyIssueType PropertyIssueType `json:"propertyIssueType"` // 任务类型
	IdentityInfo
}
