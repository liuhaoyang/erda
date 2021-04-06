// Copyright 2021 Terminus, Inc.
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

import "time"

// VClusterCreateRequest 创建虚拟集群请求结构
type VClusterCreateRequest struct {
	// 集群名称
	Name string `json:"name"`

	// 物理集群Id
	ClusterID int64 `json:"clusterId"`

	// 物理集群名称
	ClusterName string `json:"clusterName"`

	// 集群对应组织Id
	OrgID int64 `json:"orgId"`

	// 集群对应组织名称
	OrgName string `json:"orgName"`

	// 集群拥有者
	Owner string `json:"owner"`
}

// VClusterCreateResponse 创建集群响应结构
type VClusterCreateResponse struct {
	Header

	// 集群Id
	Data int64 `json:"data"`
}

// VClusterFetchResponse 集群详情响应结构
type VClusterFetchResponse struct {
	Header
	Data VClusterFetchResponseData `json:"data"`
}

// VClusterFetchResponseData 集群详情数据
type VClusterFetchResponseData struct {
	// 集群uuid
	UUID string `json:"uuid"`

	// 集群名称
	Name string `json:"name"`

	// 物理集群Id
	ClusterID int64 `json:"clusterId"`

	// 物理集群名称
	ClusterName string `json:"clusterName"`

	// 集群对应组织ID
	OrgID int64 `json:"orgId"`

	// 集群对应组织名称
	OrgName string `json:"orgName"`

	// 集群拥有者
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// VClusterListResponse 集群列表响应结构
type VClusterListResponse struct {
	Header
	Data VClusterListResponseData `json:"data"`
}

// VClusterListResponseData 集群列表数据
type VClusterListResponseData struct {
	Clusters []VClusterFetchResponseData `json:"clusters"`
}
