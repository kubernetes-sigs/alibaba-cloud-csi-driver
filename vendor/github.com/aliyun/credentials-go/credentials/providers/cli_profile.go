package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"syscall"

	"github.com/aliyun/credentials-go/credentials/internal/utils"
)

type CLIProfileCredentialsProvider struct {
	profileFile   string
	profileName   string
	innerProvider CredentialsProvider
	// 文件锁，用于并发安全
	fileMutex sync.RWMutex
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
	Name                   string `json:"name"`
	Mode                   string `json:"mode"`
	AccessKeyID            string `json:"access_key_id"`
	AccessKeySecret        string `json:"access_key_secret"`
	SecurityToken          string `json:"sts_token"`
	RegionID               string `json:"region_id"`
	RoleArn                string `json:"ram_role_arn"`
	RoleSessionName        string `json:"ram_session_name"`
	DurationSeconds        int    `json:"expired_seconds"`
	StsRegion              string `json:"sts_region"`
	EnableVpc              bool   `json:"enable_vpc"`
	SourceProfile          string `json:"source_profile"`
	RoleName               string `json:"ram_role_name"`
	OIDCTokenFile          string `json:"oidc_token_file"`
	OIDCProviderARN        string `json:"oidc_provider_arn"`
	Policy                 string `json:"policy"`
	ExternalId             string `json:"external_id"`
	SignInUrl              string `json:"cloud_sso_sign_in_url"`
	AccountId              string `json:"cloud_sso_account_id"`
	AccessConfig           string `json:"cloud_sso_access_config"`
	AccessToken            string `json:"access_token"`
	AccessTokenExpire      int64  `json:"cloud_sso_access_token_expire"`
	OauthSiteType          string `json:"oauth_site_type"`
	OauthRefreshToken      string `json:"oauth_refresh_token"`
	OauthAccessToken       string `json:"oauth_access_token"`
	OauthAccessTokenExpire int64  `json:"oauth_access_token_expire"`
	StsExpire              int64  `json:"sts_expiration"`
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

var oauthBaseUrlMap = map[string]string{
	"CN":   "https://oauth.aliyun.com",
	"INTL": "https://oauth.alibabacloud.com",
}

var oauthClientMap = map[string]string{
	"CN":   "4038181954557748008",
	"INTL": "4103531455503354461",
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
	case "CloudSSO":
		credentialsProvider, err = NewCloudSSOCredentialsProviderBuilder().
			WithSignInUrl(p.SignInUrl).
			WithAccountId(p.AccountId).
			WithAccessConfig(p.AccessConfig).
			WithAccessToken(p.AccessToken).
			WithAccessTokenExpire(p.AccessTokenExpire).
			Build()
	case "OAuth":
		siteType := strings.ToUpper(p.OauthSiteType)
		signInUrl := oauthBaseUrlMap[siteType]
		if signInUrl == "" {
			err = fmt.Errorf("invalid site type, support CN or INTL")
			return
		}
		clientId := oauthClientMap[siteType]

		credentialsProvider, err = NewOAuthCredentialsProviderBuilder().
			WithSignInUrl(signInUrl).
			WithClientId(clientId).
			WithRefreshToken(p.OauthRefreshToken).
			WithAccessToken(p.OauthAccessToken).
			WithAccessTokenExpire(p.OauthAccessTokenExpire).
			WithTokenUpdateCallback(provider.getOAuthTokenUpdateCallback()).
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
			provider.profileFile = cfgPath
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

// updateOAuthTokens 更新OAuth令牌并写回配置文件
func (provider *CLIProfileCredentialsProvider) updateOAuthTokens(refreshToken, accessToken, accessKey, secret, securityToken string, accessTokenExpire, stsExpire int64) error {
	provider.fileMutex.Lock()
	defer provider.fileMutex.Unlock()

	cfgPath := provider.profileFile
	conf, err := newConfigurationFromPath(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	profileName := provider.profileName
	profile, err := conf.getProfile(profileName)
	if err != nil {
		return fmt.Errorf("failed to get profile %s: %v", profileName, err)
	}

	// update
	profile.OauthRefreshToken = refreshToken
	profile.OauthAccessToken = accessToken
	profile.OauthAccessTokenExpire = accessTokenExpire
	profile.AccessKeyID = accessKey
	profile.AccessKeySecret = secret
	profile.SecurityToken = securityToken
	profile.StsExpire = stsExpire

	// write back with file lock
	return provider.writeConfigurationToFileWithLock(cfgPath, conf)
}

// writeConfigurationToFile 将配置写入文件，使用原子写入确保数据完整性
func (provider *CLIProfileCredentialsProvider) writeConfigurationToFile(cfgPath string, conf *configuration) error {
	// 创建临时文件
	tempFile := cfgPath + ".tmp"

	// 序列化配置
	data, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	// 写入临时文件
	err = ioutil.WriteFile(tempFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write temp file: %v", err)
	}

	// 原子性重命名，确保文件完整性
	err = os.Rename(tempFile, cfgPath)
	if err != nil {
		// 清理临时文件
		os.Remove(tempFile)
		return fmt.Errorf("failed to rename temp file: %v", err)
	}

	return nil
}

// writeConfigurationToFileWithLock 使用操作系统级别的文件锁写入配置文件
func (provider *CLIProfileCredentialsProvider) writeConfigurationToFileWithLock(cfgPath string, conf *configuration) error {
	// 打开文件用于锁定
	file, err := os.OpenFile(cfgPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	// 获取独占锁（阻塞其他进程）
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if err != nil {
		return fmt.Errorf("failed to acquire file lock: %v", err)
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	// 序列化配置
	data, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	// 创建临时文件
	tempFile := cfgPath + ".tmp"
	err = ioutil.WriteFile(tempFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write temp file: %v", err)
	}

	// 原子性重命名
	err = os.Rename(tempFile, cfgPath)
	if err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to rename temp file: %v", err)
	}

	return nil
}

// getOAuthTokenUpdateCallback 获取OAuth令牌更新回调函数
func (provider *CLIProfileCredentialsProvider) getOAuthTokenUpdateCallback() OAuthTokenUpdateCallback {
	return func(refreshToken, accessToken, accessKey, secret, securityToken string, accessTokenExpire, stsExpire int64) error {
		return provider.updateOAuthTokens(refreshToken, accessToken, accessKey, secret, securityToken, accessTokenExpire, stsExpire)
	}
}
