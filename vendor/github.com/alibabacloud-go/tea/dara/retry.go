package dara

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_DELAY_TIME  = 120 * 1000              // 120 seconds
	MIN_DELAY_TIME  = 100                     // 100 milliseconds
	DEFAULT_MAX_CAP = 3 * 24 * 60 * 60 * 1000 // 3 days in milliseconds
	MAX_ATTEMPTS    = 3
)

// RetryPolicyContext holds context for the retry operation
type RetryPolicyContext struct {
	Key              string
	RetriesAttempted int
	HttpRequest      *Request  // placeholder for actual http.Request type
	HttpResponse     *Response // placeholder for actual http.Response type
	Exception        error
}

// BackoffPolicy interface with a method to get delay time
type BackoffPolicy interface {
	GetDelayTime(ctx *RetryPolicyContext) int
}

// BackoffPolicyFactory creates a BackoffPolicy based on the option
func BackoffPolicyFactory(option map[string]interface{}) (BackoffPolicy, error) {

	switch option["policy"] {
	case "Fixed":
		return NewFixedBackoffPolicy(option), nil
	case "Random":
		return NewRandomBackoffPolicy(option), nil
	case "Exponential":
		return NewExponentialBackoffPolicy(option), nil
	case "EqualJitter", "ExponentialWithEqualJitter":
		return NewEqualJitterBackoffPolicy(option), nil
	case "FullJitter", "ExponentialWithFullJitter":
		return NewFullJitterBackoffPolicy(option), nil
	}
	return nil, fmt.Errorf("unknown policy type")
}

// FixedBackoffPolicy implementation
type FixedBackoffPolicy struct {
	Period int
}

func NewFixedBackoffPolicy(option map[string]interface{}) *FixedBackoffPolicy {
	var period int
	if v, ok := option["period"]; ok {
		period = v.(int)
	}
	return &FixedBackoffPolicy{
		Period: period,
	}
}

func (f *FixedBackoffPolicy) GetDelayTime(ctx *RetryPolicyContext) int {
	return f.Period
}

// RandomBackoffPolicy implementation
type RandomBackoffPolicy struct {
	Period int
	Cap    int
}

func NewRandomBackoffPolicy(option map[string]interface{}) *RandomBackoffPolicy {
	var capValue int
	var period int
	if v, ok := option["cap"]; ok {
		capValue = v.(int)
	} else {
		capValue = 20 * 1000
	}
	if v, ok := option["period"]; ok {
		period = v.(int)
	}
	return &RandomBackoffPolicy{
		Period: period,
		Cap:    capValue,
	}
}

func (r *RandomBackoffPolicy) GetDelayTime(ctx *RetryPolicyContext) int {
	randomSeed := int64(ctx.RetriesAttempted * r.Period)
	randomTime := int(rand.Int63n(randomSeed))
	if randomTime > r.Cap {
		return r.Cap
	}
	return randomTime
}

// ExponentialBackoffPolicy implementation
type ExponentialBackoffPolicy struct {
	Period int
	Cap    int
}

func NewExponentialBackoffPolicy(option map[string]interface{}) *ExponentialBackoffPolicy {
	var capValue int
	var period int
	if v, ok := option["cap"]; ok {
		capValue = v.(int)
	} else {
		capValue = DEFAULT_MAX_CAP
	}
	if v, ok := option["period"]; ok {
		period = v.(int)
	}
	return &ExponentialBackoffPolicy{
		Period: period,
		Cap:    capValue,
	}
}

func (e *ExponentialBackoffPolicy) GetDelayTime(ctx *RetryPolicyContext) int {
	randomTime := int(math.Pow(2, float64(ctx.RetriesAttempted)*float64(e.Period)))
	if randomTime > e.Cap {
		return e.Cap
	}
	return randomTime
}

// EqualJitterBackoffPolicy implementation
type EqualJitterBackoffPolicy struct {
	Period int
	Cap    int
}

func NewEqualJitterBackoffPolicy(option map[string]interface{}) *EqualJitterBackoffPolicy {
	var capValue int
	var period int
	if v, ok := option["cap"]; ok {
		capValue = v.(int)
	} else {
		capValue = DEFAULT_MAX_CAP
	}

	if v, ok := option["period"]; ok {
		period = v.(int)
	}
	return &EqualJitterBackoffPolicy{
		Period: period,
		Cap:    capValue,
	}
}

func (e *EqualJitterBackoffPolicy) GetDelayTime(ctx *RetryPolicyContext) int {
	ceil := int64(math.Min(float64(e.Cap), float64(math.Pow(2, float64(ctx.RetriesAttempted)*float64(e.Period)))))
	randNum := rand.Int63n(ceil/2 + 1)
	return int(ceil/2 + randNum)
}

// FullJitterBackoffPolicy implementation
type FullJitterBackoffPolicy struct {
	Period int
	Cap    int
}

func NewFullJitterBackoffPolicy(option map[string]interface{}) *FullJitterBackoffPolicy {
	var capValue int
	var period int
	if v, ok := option["cap"]; ok {
		capValue = v.(int)
	} else {
		capValue = DEFAULT_MAX_CAP
	}
	if v, ok := option["period"]; ok {
		period = v.(int)
	}
	return &FullJitterBackoffPolicy{
		Period: period,
		Cap:    capValue,
	}
}

func (f *FullJitterBackoffPolicy) GetDelayTime(ctx *RetryPolicyContext) int {
	ceil := int64(math.Min(float64(f.Cap), float64(math.Pow(2, float64(ctx.RetriesAttempted)*float64(f.Period)))))
	return int(rand.Int63n(ceil))
}

// RetryCondition holds the retry conditions
type RetryCondition struct {
	MaxAttempts int
	Backoff     BackoffPolicy
	Exception   []string
	ErrorCode   []string
	MaxDelay    int
}

func NewRetryCondition(condition map[string]interface{}) *RetryCondition {
	var backoff BackoffPolicy
	if condition["backoff"] != nil {
		backoffOption := condition["backoff"].(map[string]interface{})
		backoff, _ = BackoffPolicyFactory(backoffOption)
	}
	maxAttempts, ok := condition["maxAttempts"].(int)
	if !ok {
		maxAttempts = MAX_ATTEMPTS
	}

	exception, ok := condition["exception"].([]string)
	if !ok {
		exception = []string{}
	}

	errorCode, ok := condition["errorCode"].([]string)
	if !ok {
		errorCode = []string{}
	}

	maxDelay, ok := condition["maxDelay"].(int)
	if !ok {
		maxDelay = MAX_DELAY_TIME
	}

	return &RetryCondition{
		MaxAttempts: maxAttempts,
		Backoff:     backoff,
		Exception:   exception,
		ErrorCode:   errorCode,
		MaxDelay:    maxDelay,
	}
}

// RetryOptions holds the retry options
type RetryOptions struct {
	Retryable        bool
	RetryCondition   []*RetryCondition
	NoRetryCondition []*RetryCondition
}

func NewRetryOptions(options map[string]interface{}) *RetryOptions {
	retryConditions := make([]*RetryCondition, 0)
	for _, cond := range options["retryCondition"].([]interface{}) {
		condition := NewRetryCondition(cond.(map[string]interface{}))
		retryConditions = append(retryConditions, condition)
	}

	noRetryConditions := make([]*RetryCondition, 0)
	for _, cond := range options["noRetryCondition"].([]interface{}) {
		condition := NewRetryCondition(cond.(map[string]interface{}))
		noRetryConditions = append(noRetryConditions, condition)
	}

	return &RetryOptions{
		Retryable:        options["retryable"].(bool),
		RetryCondition:   retryConditions,
		NoRetryCondition: noRetryConditions,
	}
}

// shouldRetry determines if a retry should be attempted
func ShouldRetry(options *RetryOptions, ctx *RetryPolicyContext) bool {
	if ctx.RetriesAttempted == 0 {
		return true
	}

	if options == nil || !options.Retryable {
		return false
	}

	retriesAttempted := ctx.RetriesAttempted
	ex := ctx.Exception
	if baseErr, ok := ex.(BaseError); ok {
		conditions := options.NoRetryCondition

		for _, condition := range conditions {
			for _, exc := range condition.Exception {
				if exc == StringValue(baseErr.GetName()) {
					return false
				}
			}
			for _, code := range condition.ErrorCode {
				if code == StringValue(baseErr.GetCode()) {
					return false
				}
			}
		}

		conditions = options.RetryCondition
		for _, condition := range conditions {
			for _, exc := range condition.Exception {
				if exc == StringValue(baseErr.GetName()) {
					if retriesAttempted >= condition.MaxAttempts {
						return false
					}
					return true
				}
			}
			for _, code := range condition.ErrorCode {
				if code == StringValue(baseErr.GetCode()) {
					if retriesAttempted >= condition.MaxAttempts {
						return false
					}
					return true
				}
			}
		}
	}

	return false
}

// getBackoffDelay calculates backoff delay
func GetBackoffDelay(options *RetryOptions, ctx *RetryPolicyContext) int {
	if ctx.RetriesAttempted == 0 {
		return 0
	}

	if options == nil || !options.Retryable {
		return MIN_DELAY_TIME
	}

	ex := ctx.Exception
	conditions := options.RetryCondition
	if baseErr, ok := ex.(BaseError); ok {
		for _, condition := range conditions {
			for _, exc := range condition.Exception {
				if exc == StringValue(baseErr.GetName()) {
					maxDelay := condition.MaxDelay
					// Simulated "retryAfter" from an error response
					if respErr, ok := ex.(ResponseError); ok {
						retryAfter := Int64Value(respErr.GetRetryAfter())
						if retryAfter != 0 {
							return min(int(retryAfter), maxDelay)
						}
					}
					// This would be set properly based on your error handling

					if condition.Backoff == nil {
						return MIN_DELAY_TIME
					}
					return min(condition.Backoff.GetDelayTime(ctx), maxDelay)
				}
			}

			for _, code := range condition.ErrorCode {
				if code == StringValue(baseErr.GetCode()) {
					maxDelay := condition.MaxDelay
					// Simulated "retryAfter" from an error response
					if respErr, ok := ex.(ResponseError); ok {
						retryAfter := Int64Value(respErr.GetRetryAfter())
						if retryAfter != 0 {
							return min(int(retryAfter), maxDelay)
						}
					}

					if condition.Backoff == nil {
						return MIN_DELAY_TIME
					}

					return min(condition.Backoff.GetDelayTime(ctx), maxDelay)
				}
			}
		}
	}
	return MIN_DELAY_TIME
}

// helper function to find the minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
