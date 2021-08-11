package server

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestStr2ASCII(t *testing.T) {
	testStrs := map[string]string{
		"test1s912": "1161011151161115912",
		"$est2LIA":  "361011151162767365",
		"%st3/$#":   "371151163473635",
	}
	for k, v := range testStrs {
		assert.Equal(t, v, str2ASCII(k))
	}
}

func TestReadCSV(t *testing.T) {
	originData := `Project,BlockStatus,FileStatus,BlockUsed,BlockSoftLimit,BlockHardLimit,BlockGrace,FileUsed,FileSoftLimit,FileHardLimit,FileGrace
	#0,ok,ok,20,0,0,,2,0,0,`
	totalLimit, _ := getTotalLimitKBFromCSV(originData)
	assert.Equal(t, int(totalLimit), 0)
}

func TestReadCSV2(t *testing.T) {

	originData := `Project,BlockStatus,FileStatus,BlockUsed,BlockSoftLimit,BlockHardLimit,BlockGrace,FileUsed,FileSoftLimit,FileHardLimit,FileGrace
	#0,ok,ok,20,0,0,,2,0,0,
	#197762324,ok,ok,4,2097152,2097152,,1,0,0,
	#197762325,ok,ok,4,2097152,2097152,,1,0,0,
	#197762326,ok,ok,4,2097152,2097152,,1,0,0,`
	totalLimit, _ := getTotalLimitKBFromCSV(originData)
	assert.Equal(t, int(totalLimit), 2097152*3)
}

func TestSetProjectID2PVSubpath(t *testing.T) {

	testCases := map[string]map[string]error{
		"test1,/mnt/quotapath.namespace1.0/test1":          {"198410201": nil},
		"testxxxx2,/data/quotapath.namespace0.0/testxxxx2": {"": errors.New("testerror")},
	}
	for namespaceSubpath, value := range testCases {
		for procedureID, err := range value {
			var mockRun utils.CommandRunFunc
			mockRun = func(cmd string) (string, error) {
				return procedureID, err
			}
			keys := strings.Split(namespaceSubpath, ",")
			calculatedID, testErr := SetProjectID2PVSubpath(keys[0], keys[1], filepath.Dir(keys[1]), "ext4", mockRun)

			assert.Equal(t, procedureID, calculatedID)

			if err != nil {
				assert.Equal(t, "failed to set projectID to subpath with error: "+err.Error(), testErr.Error())
			} else {
				assert.Nil(t, testErr)
			}
		}
	}
}
