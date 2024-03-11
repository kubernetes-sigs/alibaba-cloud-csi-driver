/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nas

import (
	"testing"
)

func Test_rrMuxServerSelector_SelectNfsServer(t *testing.T) {
	s := &rrMuxServerSelector{
		servers: make(map[string]int),
	}
	tests := []struct {
		name  string
		input string
		want  string
		want1 string
	}{
		{
			"single server",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com:/subpath",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com",
			"/subpath",
		},
		{
			"clean path format",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com:/subpath/",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com",
			"/subpath",
		},
		{
			"default path /",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com",
			"/",
		},
		{
			"invalid server",
			"::/",
			"",
			"",
		},
		{
			"invalid multi server",
			"server:/path,:path",
			"",
			"",
		},

		// multi server using round robin selector
		{
			"multi server first select",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com,aaaa-bbbb.cccc.nas.aliyuncs.com:/share",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com",
			"/",
		},
		{
			"multi server second select",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com,aaaa-bbbb.cccc.nas.aliyuncs.com:/share",
			"aaaa-bbbb.cccc.nas.aliyuncs.com",
			"/share",
		},
		{
			"multi server third select",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com,aaaa-bbbb.cccc.nas.aliyuncs.com:/share",
			"xxxx-yyyy.zzzz.nas.aliyuncs.com",
			"/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := s.SelectNfsServer(tt.input)
			if got != tt.want {
				t.Errorf("rrMuxServerSelector.SelectNfsServer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("rrMuxServerSelector.SelectNfsServer() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
