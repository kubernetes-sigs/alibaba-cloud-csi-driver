package testing

import (
	"os"
	"time"
)

const (
	FakeAccessKeyId     = "fake-accessKeyId"
	FakeAccessKeySecret = "fake-accessKeySecret"
	FakeSecurityToken   = "STS.a-loooooooooooog-and-random-aweifuhaskichvaehiwciuhiuhiu-FakeSecurityToken"
	FakeEncryptedToken  = `{"access.key.id":"Ik032ULdaqWth2fvyigEcJHGMNNMnATo04/d+c5Jxz+bJN0Jh8fKd8A5MlNnYRwf","access.key.secret":"IDjwuQU+hced63SeMrRVBkVLQBvJu4M4KvWnYq5ysmCSZJh11xKkvUcCca6ekQnE","security.token":"PMz8KVai+0zGFRKQgtxMzU45UWkTryLO7KHmYA5nMCCITptCT/WNx+DDYAvrYugMSorxfEAj635cr94qopN0Ehft/wfnfb0jQ+7rgSzu27aiv5I8MFptekbt4fc8vygI","expiration":"2021-01-01T14:22:45Z","keyring":"fake-keyring1234"}`
)

var Expiration = time.Date(2021, 1, 1, 14, 22, 45, 0, time.UTC)

func WriteFakeToken(path string) error {
	return os.WriteFile(path, []byte(FakeEncryptedToken), 0600)
}
