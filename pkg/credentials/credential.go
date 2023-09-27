package credentials

import (
	"errors"
	"os"

	"github.com/alibabacloud-go/tea/tea"
	alicred "github.com/aliyun/credentials-go/credentials"
	"k8s.io/utils/clock"
)

func NewCredential() (alicred.Credential, error) {
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
