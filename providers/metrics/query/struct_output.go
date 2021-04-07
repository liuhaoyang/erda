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

package query

// 单点数据，对应API: {{scope}}?...
type Point struct {
	// Title string
	Name string
	Data []*PointData
}

type PointData struct {
	Name      string      `mapstructure:"name"`
	AggMethod string      `mapstructure:"agg"`
	Data      interface{} `mapstructure:"data"`
}

// 时序数据, 对应API：{{scope}}/histogram?...
type Series struct {
	Name       string
	Data       []*SeriesData
	TimeSeries []int // 毫秒
}

type SeriesData struct {
	Name      string    `mapstructure:"name"`
	AggMethod string    `mapstructure:"agg"`
	Data      []float64 `mapstructure:"data"`
	Tag       string    `mapstructure:"tag"`
}
