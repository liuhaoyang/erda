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

package bundle

//import (
//	"fmt"
//	"os"
//	"testing"
//	"time"
//
//	"github.com/erda-project/erda/apistructs"
//)
//
//func TestBundle_GetLog(*testing.T) {
//	os.Setenv("MONITOR_ADDR", "monitor.default.svc.cluster.local:7096")
//	b := New(WithMonitor())
//
//	fmt.Println(b.GetLog(apistructs.DashboardSpotLogRequest{
//		ID:     "pipeline-task-244",
//		Source: apistructs.DashboardSpotLogSourceJob,
//		Stream: apistructs.DashboardSpotLogStreamStdout,
//		Count:  -50,
//		Start:  0,
//		End:    time.Duration(1590047806647571944),
//	}))
//}
