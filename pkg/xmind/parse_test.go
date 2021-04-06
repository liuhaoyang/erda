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

package xmind

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestParseByXMindparser(t *testing.T) {
//	f, err := os.Open("./examples/content.json")
//	assert.NoError(t, err)
//	content, err := Parse(f)
//	assert.NoError(t, err)
//	b, _ := json.MarshalIndent(content, "", "  ")
//	fmt.Println(string(b))
//}

func TestParseByJSON(t *testing.T) {
	f, err := os.Open("./examples/content.json")
	assert.NoError(t, err)
	var content Content
	err = json.NewDecoder(f).Decode(&content)
	assert.NoError(t, err)

	b, _ := xml.MarshalIndent(&content, "", "  ")
	fmt.Println(string(b))
}
