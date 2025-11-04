package metadata

import (
	"encoding/json"
	"testing"

	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const getCallerIdentityRespJson = `{
	"IdentityType": "Account",
	"AccountId": "112233445566",
	"RequestId": "5051F631-1599-5DBD-9C0A-3DD86092DA9D",
	"PrincipalId": "112233445566",
	"UserId": "112233445566",
	"Arn": "acs:ram::112233445566:root"
}`

func testStsClient(ctrl *gomock.Controller) cloud.STSInterface {
	res := &sts20150401.GetCallerIdentityResponse{}
	err := json.Unmarshal([]byte(getCallerIdentityRespJson), &res.Body)
	if err != nil {
		panic(err)
	}

	stsClient := cloud.NewMockSTSInterface(ctrl)
	stsClient.EXPECT().GetCallerIdentity().Return(res, nil)
	return stsClient
}

func testStsClientFactory(ctrl *gomock.Controller) func(string) (cloud.STSInterface, error) {
	return func(regionID string) (cloud.STSInterface, error) {
		return testStsClient(ctrl), nil
	}
}

func TestGetSts(t *testing.T) {
	ctrl := gomock.NewController(t)
	stsClient := testStsClient(ctrl)

	m, err := NewStsMetadata(stsClient)
	assert.NoError(t, err)
	assert.Equal(t, "112233445566", MustGet(m, AccountID))
}

func TestGetStsEmptyJson(t *testing.T) {
	ctrl := gomock.NewController(t)
	res := &sts20150401.GetCallerIdentityResponse{}
	err := json.Unmarshal([]byte(`{}`), &res.Body)
	if err != nil {
		panic(err)
	}
	stsClient := cloud.NewMockSTSInterface(ctrl)
	stsClient.EXPECT().GetCallerIdentity().Return(res, nil)

	m, err := NewStsMetadata(stsClient)
	require.NoError(t, err)
	_, err = m.Get(AccountID)
	assert.ErrorIs(t, err, ErrUnknownMetadataKey)
	_, err = m.Get(999) // anything else
	assert.ErrorIs(t, err, ErrUnknownMetadataKey)
}
