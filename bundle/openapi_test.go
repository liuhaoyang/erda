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

package bundle

//import (
//	"fmt"
//	"net/http"
//	"os"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	"github.com/erda-project/erda/apistructs"
//)
//
//func TestBundle_OpenapiOAuth2Token(t *testing.T) {
//	os.Setenv("OPENAPI_ADDR", "localhost:9529")
//	bdl := New(WithOpenapi())
//
//	// get token
//	ti, err := bdl.GetOpenapiOAuth2Token(apistructs.OpenapiOAuth2TokenGetRequest{
//		ClientID:     "pipeline",
//		ClientSecret: "devops/pipeline",
//		Payload: apistructs.OpenapiOAuth2TokenPayload{
//			AccessibleAPIs: []apistructs.AccessibleAPI{
//				{
//					Path:   "/api/pipelines/<pipelineID>/tasks/<taskID>/actions/get-bootstrap-info",
//					Method: http.MethodGet,
//					Schema: "http",
//				},
//			},
//			Metadata: map[string]string{
//				"pipelineID": "10000001",
//				"taskID":     "2",
//			},
//		},
//	})
//	assert.NoError(t, err)
//	fmt.Printf("%+v\n", ti)
//
//	// invalidate token
//	ti, err = bdl.InvalidateOpenapiOAuth2Token(apistructs.OpenapiOAuth2TokenInvalidateRequest{
//		AccessToken: ti.AccessToken,
//	})
//	assert.NoError(t, err)
//	fmt.Printf("%+v\n", ti)
//}
