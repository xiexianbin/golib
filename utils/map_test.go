// Copyright 2020 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// https://github.com/xiexianbin/gsync/blob/7431fdacf3/utils/map_utils_test.go

package utils

import (
	"testing"
)

func TestUnionMap(t *testing.T) {
	m1 := map[string]string{"a": "a", "b": "b", "c": "c"}
	m2 := map[string]string{"a": "a", "b": "b", "d": "d"}
	unionMap := UnionMap(m1, m2)
	for k, v := range unionMap {
		t.Log(k, v)
	}
}

func TestDiffMap(t *testing.T) {
	type M map[string]interface{}
	m1 := M{"a": "a", "b": "b", "c": "c", "e": "e"}
	m2 := M{"a": "a", "b": "b", "d": "d", "e": "f"}
	justM1, justM2, diffM1AndM2 := DiffMap(m1, m2)
	t.Log("justM1")
	for k, v := range justM1 {
		t.Log(k, ":", v)
	}

	t.Log("justM2")
	for k, v := range justM2 {
		t.Log(k, ":", v)
	}

	t.Log("diffM1AndM2")
	for k, v := range diffM1AndM2 {
		t.Log(k, ":", v)
	}
}
