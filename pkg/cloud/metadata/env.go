package metadata

import "os"

type ENVMetadata struct{}

var MetadataEnv = map[MetadataKey]string{
	RegionID:      "REGION_ID",
	AccountID:     "ALIBABA_CLOUD_ACCOUNT_ID",
	ClusterID:     "CLUSTER_ID",
	RRSATokenFile: "RRSA_TOKEN_FILE",
}

func (m *ENVMetadata) Get(key MetadataKey) (string, error) {
	if v, ok := MetadataEnv[key]; ok {
		if value, ok := os.LookupEnv(v); ok {
			return value, nil
		}
	}
	return "", ErrUnknownMetadataKey
}
