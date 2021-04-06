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

import (
	"errors"

	"github.com/erda-project/erda/pkg/parser/diceyml"
)

type JobKind string

var (
	Metronome   JobKind = "Metronome"
	Flink       JobKind = "Flink"
	Spark       JobKind = "Spark"
	K8SSpark    JobKind = "K8SSpark"
	K8SFlink    JobKind = "K8SFlink"
	LocalDocker JobKind = "LocalDocker"
	LocalJob    JobKind = "LocalJob"
	Swarm       JobKind = "Swarm"
	Kubernetes  JobKind = "Kubernetes"
)

type JobVolume struct {
	Namespace string `json:"namespace"`
	// 用于生成volume id = <namespace>-<name:>
	Name string `json:"name"`
	// nfs | local
	Type string `json:"type"`

	Executor    string `json:"executor,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	Kind        string `json:"kind"`
}

type JobVolumeCreateResponse struct {
	ID    string `json:"id"`
	Error string `json:"error"`
}

// Job defines a Job.
type Job struct {
	CreatedTime    int64         `json:"created_time"`
	LastStartTime  int64         `json:"last_start_time"`
	LastFinishTime int64         `json:"last_finish_time"`
	LastModify     string        `json:"last_modify"`
	Result         interface{}   `json:"result,omitempty"`
	ScheduleInfo   ScheduleInfo  `json:"scheduleInfo,omitempty"`  // 根据集群配置以及 label 所计算出的调度规则
	ScheduleInfo2  ScheduleInfo2 `json:"scheduleInfo2,omitempty"` // 将会代替 ScheduleInfo
	BigdataConf    `json:"bigdataConf,omitempty"`
	JobFromUser
	StatusDesc
}

type JobFromUser struct {
	Name         string                 `json:"name"`
	Namespace    string                 `json:"namespace"`    // the default namespace is "default"
	ID           string                 `json:"id,omitempty"` // if Job has owner, e.g. jobflow, it's ID can be specified.
	CallBackUrls []string               `json:"callbackurls,omitempty"`
	Image        string                 `json:"image,omitempty"`
	Resource     string                 `json:"resource,omitempty"`  // Flink时，为jarId；Spark时，为jar url
	MainClass    string                 `json:"mainClass,omitempty"` // 入口类, 主要用于Flink/Spark
	MainArgs     []string               `json:"mainArgs"`            // 入口类参数, 主要用于Flink/Spark
	Cmd          string                 `json:"cmd,omitempty"`
	CPU          float64                `json:"cpu,omitempty"`
	Memory       float64                `json:"memory,omitempty"`
	Labels       map[string]string      `json:"labels,omitempty"`
	Extra        map[string]string      `json:"extra,omitempty"`
	Env          map[string]string      `json:"env,omitempty"`
	Binds        []Bind                 `json:"binds,omitempty"`
	Volumes      []diceyml.Volume       `json:"volumes,omitempty"`
	Executor     string                 `json:"executor,omitempty"`
	ClusterName  string                 `json:"clusterName,omitempty"`
	Kind         string                 `json:"kind"`              // Metronome/FLink/Spark/LocalDocker/Swarm/Kubernetes
	Depends      []string               `json:"depends,omitempty"` // JobName
	PreFetcher   *PreFetcher            `json:"preFetcher,omitempty"`
	BackoffLimit int                    `json:"backoffLimit,omitempty"`
	Params       map[string]interface{} `json:"params,omitempty"`
}

// PreFetcher 用于 job 下载功能
type PreFetcher struct {
	FileFromImage string `json:"fileFromImage,omitempty"` // 通过 k8s initcontainer 实现, fetch 的工作需要在 镜像 entrypoint 中做掉
	FileFromHost  string `json:"fileFromHost,omitempty"`  // 通过 bind 的方式实现, 兼容 metronome
	ContainerPath string `json:"containerPath"`
}

type JobCreateRequest JobFromUser

type JobCreateResponse struct {
	Name  string `json:"name"`
	Error string `json:"error"`
	Job   Job    `json:"job"`
}

type JobStartResponse struct {
	Name  string `json:"name"`
	Error string `json:"error"`
	Job   Job    `json:"job"`
}

type JobStopResponse struct {
	Name  string `json:"name"`
	Error string `json:"error"`
}

type JobDeleteResponse struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Error     string `json:"error"`
}

type JobsDeleteResponse []JobDeleteResponse

type JobBatchRequest struct {
	Names []string `json:"names"`
}

type JobBatchResponse struct {
	Names []string `json:"names"`
	Error string   `json:"error"`
	Jobs  []Job    `json:"jobs"`
}

var ErrJobIsRunning = errors.New("job is running")
