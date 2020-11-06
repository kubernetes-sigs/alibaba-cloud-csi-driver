package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2ASCII(t *testing.T) {
	testStrs := map[string]string{
		"test1s912": "11610111511649115574950",
		"$est2LIA":  "3610111511650767365",
		"%st3/$#":   "3711511651473635",
	}
	for k, v := range testStrs {
		assert.Equal(t, v, str2ASCII(k))
	}
}

func TestGetTotalLimitKBFromCSV(t *testing.T) {
	testSamples := map[string]int{
		`Project,BlockStatus,FileStatus,BlockUsed,BlockSoftLimit,BlockHardLimit,BlockGrace,FileUsed,FileSoftLimit,FileHardLimit,FileGrace
		#0,ok,ok,16,0,0,,1,0,0,
		#1234567,ok,ok,4,0,0,,1,0,0,
		#12345678,ok,ok,4,4,10,6days,1,0,0,`: 10,
		`Project,BlockStatus,FileStatus,BlockUsed,BlockSoftLimit,BlockHardLimit,BlockGrace,FileUsed,FileSoftLimit,FileHardLimit,FileGrace
		#0,ok,ok,16,0,0,,1,0,0,
		#1234567,ok,ok,4,5,0,,1,0,0,
		#12345678,ok,ok,4,4,10,6days,1,0,0,`: 15,
		`Project,BlockStatus,FileStatus,BlockUsed,BlockSoftLimit,BlockHardLimit,BlockGrace,FileUsed,FileSoftLimit,FileHardLimit,FileGrace
		#0,ok,ok,16,0,1,,1,0,0,
		#1234567,ok,ok,4,5,0,,1,0,0,
		#12345678,ok,ok,4,4,10,6days,1,0,0,`: 15,
	}
	for k, v := range testSamples {
		actualResult, err := getTotalLimitKBFromCSV(k)
		assert.Nil(t, err)
		assert.Equal(t, v, actualResult)
	}
}
