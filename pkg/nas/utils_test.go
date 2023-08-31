/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNfsDetails(t *testing.T) {

	nfsServerString := "0.0.0.0"
	nfsServer, nfsPath := GetNfsDetails(nfsServerString)
	assert.Equal(t, "0.0.0.0", nfsServer)
	assert.Equal(t, "/", nfsPath)

	nfsServerString = "0.0.0.0:/test"
	nfsServer, nfsPath = GetNfsDetails(nfsServerString)
	assert.Equal(t, "0.0.0.0", nfsServer)
	assert.Equal(t, "/test", nfsPath)

	nfsServerString = "0.0.0.0,0.0.0.1:/test/"
	nfsServer, nfsPath = GetNfsDetails(nfsServerString)
	assert.Equal(t, "0.0.0.0", nfsServer)
	assert.Equal(t, "/", nfsPath)

}

func TestParseMountFlags(t *testing.T) {
	type args struct {
		mntOptions []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			"vers=3",
			args{[]string{"mnt=/test", "vers=3.0"}},
			"3", "mnt=/test",
		},
		{
			"vers=3.0",
			args{[]string{"mnt=/test", "vers=3.0"}},
			"3", "mnt=/test",
		},
		{
			"vers=4",
			args{[]string{"mnt=/test", "vers=4"}},
			"4", "mnt=/test",
		},
		{
			"vers=4.0",
			args{[]string{"mnt=/test", "vers=4.0"}},
			"4.0", "mnt=/test",
		},
		{
			"vers=4.1",
			args{[]string{"mnt=/test", "vers=4.1"}},
			"4.1", "mnt=/test",
		},
		{
			"no vers",
			args{[]string{"mnt=/test", "a=b,,c=d"}},
			"", "mnt=/test,a=b,c=d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseMountFlags(tt.args.mntOptions)
			if got != tt.want {
				t.Errorf("ParseMountFlags() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseMountFlags() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
