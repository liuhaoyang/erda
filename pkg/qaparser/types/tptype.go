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

package types

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/erda-project/erda/apistructs"
)

func TestTypeValues() []apistructs.TestType {
	return []apistructs.TestType{apistructs.UT, apistructs.IT}
}

func TestTypeValueOf(tptype string) (apistructs.TestType, error) {
	switch strings.TrimSpace(tptype) {
	case string(apistructs.UT):
		return apistructs.UT, nil
	case string(apistructs.IT):
		return apistructs.IT, nil
	default:
		return "", errors.Errorf("not supported yet %s", tptype)
	}
}
