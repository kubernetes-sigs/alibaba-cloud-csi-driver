package credentials

import (
	"errors"
	"os"

	"github.com/alibabacloud-go/tea/tea"
	alicred "github.com/aliyun/credentials-go/credentials"
	"github.com/sirupsen/logrus"
	"k8s.io/utils/clock"
)

func MigrateEnv() {
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	accessSecret := os.Getenv("ACCESS_KEY_SECRET")
	if accessKeyID != "" || accessSecret != "" {
		logrus.Warnf("enviroment variable ACCESS_KEY_ID and ACCESS_KEY_SECRET is deprecated, please use %s and %s instead.", alicred.EnvVarAccessKeyIdNew, alicred.EnvVarAccessKeySecret)

		// if ALIBABA_CLOUD_* not set, use ACCESS_KEY_ID and ACCESS_KEY_SECRET
		if os.Getenv(alicred.EnvVarAccessKeyIdNew) == "" && accessKeyID != "" {
			os.Setenv(alicred.EnvVarAccessKeyIdNew, accessKeyID)
		}
		if os.Getenv(alicred.EnvVarAccessKeySecret) == "" && accessSecret != "" {
			os.Setenv(alicred.EnvVarAccessKeySecret, accessSecret)
		}
	}
}

func NewCredential() (alicred.Credential, error) {
	MigrateEnv()

	credential, err := alicred.NewCredential(nil)
	if err == nil {
		return credential, nil
	}
	if err.Error() != "No credential found" {
		return nil, err
	}

	// try managed token credential
	credential, err = NewManagedTokenCredential(clock.RealClock{}, GetManagedTokenPath())
	if err == nil {
		return credential, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	// try ecs ram role credential as fallback
	return alicred.NewCredential(&alicred.Config{
		Type: tea.String("ecs_ram_role"),
	})
}
