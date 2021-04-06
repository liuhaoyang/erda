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

package pipelineymlvars

// envs
const (
	CLUSTER_NAME = "CLUSTER_NAME"
)

type YmlField string

func (f YmlField) String() string {
	return string(f)
}

const (
	FieldAggregate           YmlField = "aggregate"
	FieldGet                 YmlField = "get"
	FieldPut                 YmlField = "put"
	FieldTask                YmlField = "task"
	FieldDisable             YmlField = "disable"
	FieldPause               YmlField = "pause"
	FieldEnvs                YmlField = "envs"
	FieldParams              YmlField = "params"
	FieldParamForceBuildpack YmlField = "force_buildpack"
	FieldTaskConfig          YmlField = "config"
	FieldTaskConfigEnvs      YmlField = "envs"
)
