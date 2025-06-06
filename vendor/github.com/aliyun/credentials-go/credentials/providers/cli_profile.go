package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/aliyun/credentials-go/credentials/internal/utils"
)

type CLIProfileCredentialsProvider struct {
	profileFile   string
	profileName   string
	innerProvider CredentialsProvider
}

type CLIProfileCredentialsProviderBuilder struct {
	provider *CLIProfileCredentialsProvider
}

func (b *CLIProfileCredentialsProviderBuilder) WithProfileFile(profileFile string) *CLIProfileCredentialsProviderBuilder {
	b.provider.profileFile = profileFile
	return b
}

func (b *CLIProfileCredentialsProviderBuilder) WithProfileName(profileName string) *CLIProfileCredentialsProviderBuilder {
	b.provider.profileName = profileName
	return b
}

func (b *CLIProfileCredentialsProviderBuilder) Build() (provider *CLIProfileCredentialsProvider, err error) {
	// 优先级：
	// 1. 使用显示指定的 profileFile
	// 2. 使用环境变量（ALIBABA_CLOUD_CONFIG_FILE）指定的 profileFile
	// 3. 兜底使用 path.Join(homeDir, ".aliyun/config") 作为 profileFile
	if b.provider.profileFile == "" {
		b.provider.profileFile = os.Getenv("ALIBABA_CLOUD_CONFIG_FILE")
	}
	// 优先级：
	// 1. 使用显示指定的 profileName
	// 2. 使用环境变量（ALIBABA_CLOUD_PROFILE）制定的 profileName
	// 3. 使用 CLI 配置中的当前 profileName
	if b.provider.profileName == "" {
		b.provider.profileName = os.Getenv("ALIBABA_CLOUD_PROFILE")
	}

	if strings.ToLower(os.Getenv("ALIBABA_CLOUD_CLI_PROFILE_DISABLED")) == "true" {
		err = errors.New("the CLI profile is disabled")
		return
	}

	provider = b.provider
	return
}

func NewCLIProfileCredentialsProviderBuilder() *CLIProfileCredentialsProviderBuilder {
	return &CLIProfileCredentialsProviderBuilder{
		provider: &CLIProfileCredentialsProvider{},
	}
}

type profile struct {
	Name            string `json:"name"`
	Mode            string `json:"mode"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	SecurityToken   string `json:"sts_token"`
	RegionID        string `json:"region_id"`
	RoleArn         string `json:"ram_role_arn"`
	RoleSessionName string `json:"ram_session_name"`
	DurationSeconds int    `json:"expired_seconds"`
	StsRegion       string `json:"sts_region"`
	EnableVpc       bool   `json:"enable_vpc"`
	SourceProfile   string `json:"source_profile"`
	RoleName        string `json:"ram_role_name"`
	OIDCTokenFile   string `json:"oidc_token_file"`
	OIDCProviderARN string `json:"oidc_provider_arn"`
	Policy          string `json:"policy"`
	ExternalId      string `json:"external_id"`
}

type configuration struct {
	Current  string     `json:"current"`
	Profiles []*profile `json:"profiles"`
}

func newConfigurationFromPath(cfgPath string) (conf *configuration, err error) {
	bytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		err = fmt.Errorf("reading aliyun cli config from '%s' failed %v", cfgPath, err)
		return
	}

	conf = &configuration{}

	err = json.Unmarshal(bytes, conf)
	if err != nil {
		err = fmt.Errorf("unmarshal aliyun cli config from '%s' failed: %s", cfgPath, string(bytes))
		return
	}

	if conf.Profiles == nil || len(conf.Profiles) == 0 {
		err = fmt.Errorf("no any configured profiles in '%s'", cfgPath)
		return
	}

	return
}

func (conf *configuration) getProfile(name string) (profile *profile, err error) {
	for _, p := range conf.Profiles {
		if p.Name == name {
			profile = p
			return
		}
	}

	err = fmt.Errorf("unable to get profile with '%s'", name)
	return
}

func (provider *CLIProfileCredentialsProvider) getCredentialsProvider(conf *configuration, profileName string) (credentialsProvider CredentialsProvider, err error) {
	p, err := conf.getProfile(profileName)
	if err != nil {
		return
	}

	switch p.Mode {
	case "AK":
		credentialsProvider, err = NewStaticAKCredentialsProviderBuilder().
			WithAccessKeyId(p.AccessKeyID).
			WithAccessKeySecret(p.AccessKeySecret).
			Build()
	case "StsToken":
		credentialsProvider, err = NewStaticSTSCredentialsProviderBuilder().
			WithAccessKeyId(p.AccessKeyID).
			WithAccessKeySecret(p.AccessKeySecret).
			WithSecurityToken(p.SecurityToken).
			Build()
	case "RamRoleArn":
		previousProvider, err1 := NewStaticAKCredentialsProviderBuilder().
			WithAccessKeyId(p.AccessKeyID).
			WithAccessKeySecret(p.AccessKeySecret).
			Build()
		if err1 != nil {
			return nil, err1
		}

		credentialsProvider, err = NewRAMRoleARNCredentialsProviderBuilder().
			WithCredentialsProvider(previousProvider).
			WithRoleArn(p.RoleArn).
			WithRoleSessionName(p.RoleSessionName).
			WithDurationSeconds(p.DurationSeconds).
			WithStsRegionId(p.StsRegion).
			WithEnableVpc(p.EnableVpc).
			WithPolicy(p.Policy).
			WithExternalId(p.ExternalId).
			Build()
	case "EcsRamRole":
		credentialsProvider, err = NewECSRAMRoleCredentialsProviderBuilder().WithRoleName(p.RoleName).Build()
	case "OIDC":
		credentialsProvider, err = NewOIDCCredentialsProviderBuilder().
			WithOIDCTokenFilePath(p.OIDCTokenFile).
			WithOIDCProviderARN(p.OIDCProviderARN).
			WithRoleArn(p.RoleArn).
			WithStsRegionId(p.StsRegion).
			WithEnableVpc(p.EnableVpc).
			WithDurationSeconds(p.DurationSeconds).
			WithRoleSessionName(p.RoleSessionName).
			WithPolicy(p.Policy).
			Build()
	case "ChainableRamRoleArn":
		previousProvider, err1 := provider.getCredentialsProvider(conf, p.SourceProfile)
		if err1 != nil {
			err = fmt.Errorf("get source profile failed: %s", err1.Error())
			return
		}
		credentialsProvider, err = NewRAMRoleARNCredentialsProviderBuilder().
			WithCredentialsProvider(previousProvider).
			WithRoleArn(p.RoleArn).
			WithRoleSessionName(p.RoleSessionName).
			WithDurationSeconds(p.DurationSeconds).
			WithStsRegionId(p.StsRegion).
			WithEnableVpc(p.EnableVpc).
			WithPolicy(p.Policy).
			WithExternalId(p.ExternalId).
			Build()
	default:
		err = fmt.Errorf("unsupported profile mode '%s'", p.Mode)
	}

	return
}

// 默认设置为 GetHomePath，测试时便于 mock
var getHomePath = utils.GetHomePath

func (provider *CLIProfileCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	if provider.innerProvider == nil {
		cfgPath := provider.profileFile
		if cfgPath == "" {
			homeDir := getHomePath()
			if homeDir == "" {
				err = fmt.Errorf("cannot found home dir")
				return
			}

			cfgPath = path.Join(homeDir, ".aliyun/config.json")
		}

		conf, err1 := newConfigurationFromPath(cfgPath)
		if err1 != nil {
			err = err1
			return
		}

		if provider.profileName == "" {
			provider.profileName = conf.Current
		}

		provider.innerProvider, err = provider.getCredentialsProvider(conf, provider.profileName)
		if err != nil {
			return
		}
	}

	innerCC, err := provider.innerProvider.GetCredentials()
	if err != nil {
		return
	}

	providerName := innerCC.ProviderName
	if providerName == "" {
		providerName = provider.innerProvider.GetProviderName()
	}

	cc = &Credentials{
		AccessKeyId:     innerCC.AccessKeyId,
		AccessKeySecret: innerCC.AccessKeySecret,
		SecurityToken:   innerCC.SecurityToken,
		ProviderName:    fmt.Sprintf("%s/%s", provider.GetProviderName(), providerName),
	}

	return
}

func (provider *CLIProfileCredentialsProvider) GetProviderName() string {
	return "cli_profile"
}
