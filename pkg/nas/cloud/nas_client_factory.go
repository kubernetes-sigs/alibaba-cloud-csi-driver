package cloud

import (
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"
)

const (
	// defaultQps is the aggregate ceiling. The service meters per (region,
	// action, account), so the per-action buckets are what match the quota;
	// this cap only keeps one process from fanning out across every action at
	// once. NAS_LIMIT_PERSECOND keeps this meaning.
	defaultQps = 200
	// defaultActionQps is where an action starts before it has seen the
	// account's quota. The published table says 120 requests/60s (2 QPS) for
	// most actions, which is the floor for an account that has not asked for
	// more. Slow start doubles from there and stops on the first Throttling
	// response, so the driver finds the real quota in a handful of seconds
	// without a configured operating point.
	defaultActionQps = 2
)

// actionQps holds the actions whose published quota differs from the default.
// https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-quota
// These are starting points for slow start, not ceilings: an account that has
// raised them keeps growing until it throttles.
var actionQps = map[string]float64{
	"CreateDir":           20,  // 1200/60s
	"DescribeFileSystems": 100, // 100/s
}

// actionLimitersRegistered lists every NAS OpenAPI action the driver calls, so
// each gets its own bucket instead of competing with unrelated actions in the
// aggregate one.
var actionLimitersRegistered = []string{
	"CreateDir",
	"SetDirQuota",
	"CancelDirQuota",
	"GetRecycleBinAttribute",
	"CreateAccessPoint",
	"DeleteAccessPoint",
	"DescribeAccessPoint",
	"ListAccessPoints",
	"DescribeAgenticSpaces",
	"DescribeFileSystems",
	"CreateAgenticSpace",
	"GetAgenticSpace",
	"DeleteAgenticSpace",
	"SetAgenticSpaceQuota",
}

type NasClientFactory struct {
	// limiter is the aggregate bucket every call passes through.
	limiter *rate.Limiter
	// actionLimiters holds one adaptive bucket per OpenAPI action.
	actionLimiters map[string]*actionLimiter
}

// NewNasClientFactory reads:
//
//	NAS_LIMIT_PERSECOND          aggregate ceiling (default 200)
//	NAS_ACTION_LIMIT_PERSECOND   rate an action starts at, before it discovers
//	                             the account quota (default 2, the published
//	                             floor)
//	NAS_ACTION_MAX_PERSECOND     ceiling an action may grow to (default: the
//	                             aggregate ceiling)
func NewNasClientFactory() *NasClientFactory {
	globalQPS := envPositiveInt("NAS_LIMIT_PERSECOND", defaultQps)
	actionRate := envPositiveInt("NAS_ACTION_LIMIT_PERSECOND", defaultActionQps)
	maxRate := envPositiveInt("NAS_ACTION_MAX_PERSECOND", globalQPS)
	if maxRate < actionRate {
		maxRate = actionRate
	}
	klog.InfoS("NAS OpenAPI rate limits",
		"aggregateQps", globalQPS, "actionQps", actionRate, "actionMaxQps", maxRate)

	limiters := make(map[string]*actionLimiter, len(actionLimitersRegistered))
	for _, action := range actionLimitersRegistered {
		rate := float64(actionRate)
		if configured, ok := actionQps[action]; ok {
			rate = configured
		}
		limiters[action] = newActionLimiter(action, rate, float64(maxRate))
	}
	return &NasClientFactory{
		limiter:        rate.NewLimiter(rate.Limit(globalQPS), burstFor(float64(globalQPS))),
		actionLimiters: limiters,
	}
}

func envPositiveInt(name string, fallback int) int {
	value, ok := os.LookupEnv(name)
	if !ok {
		return fallback
	}
	v, err := strconv.Atoi(value)
	if err != nil || v <= 0 {
		klog.ErrorS(nil, "invalid value, keeping the default", "env", name, "value", value, "default", fallback)
		return fallback
	}
	return v
}

// V2 gets a NAS OpenAPI client sourced from github.com/alibabacloud-go/nas-20170626.
// As github.com/aliyun/alibaba-cloud-sdk-go/services/nas won't be updated with new NAS APIs (e.g., access points),
// we will fully migrate to github.com/alibabacloud-go/nas-20170626 in the future.
func (fac *NasClientFactory) V2(region string) (interfaces.NasClientV2Interface, error) {
	client, err := NewNasClientV2(region)
	if err != nil {
		return nil, err
	}
	return &NasClientV2{
		region:         region,
		limiter:        fac.limiter,
		actionLimiters: fac.actionLimiters,
		client:         wrap.NAS(client),
	}, nil
}

// Deprecated: NAS openapi client provided by github.com/aliyun/alibaba-cloud-sdk-go/services/nas.
func (fac *NasClientFactory) V1(region string) (interfaces.NasV1Interface, error) {
	return newNasClientV1(region)
}
