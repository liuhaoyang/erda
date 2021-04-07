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

package pipelineymlv1

import (
	"github.com/erda-project/erda/apistructs"
)

const (
	BuiltinResourceTypeDefaultDockerImagePrefix = "registry.cn-hangzhou.aliyuncs.com/terminus"
	BuiltinResourceTypeDefaultDockerImageTag    = "latest"
)

type Option struct {
	builtinResourceTypeDockerImagePrefix string // builtin resource type 使用的 docker image prefix
	builtinResourceTypeDockerImageTag    string // builtin resource type 使用的 docker image tag

	containerResLimit containerResLimit // 容器资源限制，包括 cpu、mem 等

	nfsRealPath string // 网盘路径, 例如: /netdata/devops/ci，老版本是 /netdata/ci

	branch string // 分支

	clusterName string // 所属集群

	renderPlaceholder bool // 是否渲染占位符

	placeholders []apistructs.MetadataField // render pipeline.yml placeholder

	alreadyTransformed bool // 是否已经 transformed, 如果是, 则禁止自动插入某些节点

	agentHostPath      string //agent 外部地址
	agentContainerPath string //agent 容器映射地址
}

type containerResLimit struct {
	cpu float64
	mem float64
}

type OpOption func(*Option)

func WithBuiltinResourceTypeDockerImagePrefixAndTag(prefix, tag string) OpOption {
	return func(op *Option) {
		op.builtinResourceTypeDockerImagePrefix = prefix
		op.builtinResourceTypeDockerImageTag = tag
	}
}

func WithContainerResLimit(cpu, mem float64) OpOption {
	return func(op *Option) {
		op.containerResLimit = containerResLimit{cpu: cpu, mem: mem}
	}
}

func WithNFSRealPath(realPath string) OpOption {
	return func(op *Option) {
		op.nfsRealPath = realPath
	}
}

func WithBranch(branch string) OpOption {
	return func(op *Option) {
		op.branch = branch
	}
}

func WithClusterName(clusterName string) OpOption {
	return func(op *Option) {
		op.clusterName = clusterName
	}
}

func WithRenderPlaceholder(render bool) OpOption {
	return func(op *Option) {
		op.renderPlaceholder = render
	}
}

func WithPlaceholders(placeholders []apistructs.MetadataField) OpOption {
	return func(op *Option) {
		op.placeholders = placeholders
	}
}

func WithAlreadyTransformed(already bool) OpOption {
	return func(op *Option) {
		op.alreadyTransformed = already
	}
}

func WithAgentHostPath(agentHostPath string) OpOption {
	return func(op *Option) {
		op.agentHostPath = agentHostPath
	}
}

func WithAgentContainerPath(agentContainerPath string) OpOption {
	return func(op *Option) {
		op.agentContainerPath = agentContainerPath
	}
}
