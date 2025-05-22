package oss

import (
	"net/url"
	"strings"
)

func isValidRegion(region string) bool {
	for _, v := range region {
		if !(('a' <= v && v <= 'z') || ('0' <= v && v <= '9') || v == '-') {
			return false
		}
	}
	return region != ""
}

func isValidEndpoint(endpoint *url.URL) bool {
	return (endpoint != nil)
}

func isValidBucketName(bucketName *string) bool {
	if bucketName == nil {
		return false
	}

	nameLen := len(*bucketName)
	if nameLen < 3 || nameLen > 63 {
		return false
	}

	if (*bucketName)[0] == '-' || (*bucketName)[nameLen-1] == '-' {
		return false
	}

	for _, v := range *bucketName {
		if !(('a' <= v && v <= 'z') || ('0' <= v && v <= '9') || v == '-') {
			return false
		}
	}
	return true
}

func isValidObjectName(objectName *string) bool {
	if objectName == nil || len(*objectName) == 0 {
		return false
	}
	return true
}

func isValidRange(r *string) bool {
	if _, err := ParseRange(*r); err != nil {
		return false
	}
	return true
}

var supportedMethod = map[string]struct{}{
	"GET":     {},
	"PUT":     {},
	"HEAD":    {},
	"POST":    {},
	"DELETE":  {},
	"OPTIONS": {},
}

func isValidMethod(method string) bool {
	if _, ok := supportedMethod[method]; ok {
		return true
	}
	return false
}

var supportedCopyDirective = map[string]struct{}{
	"COPY":    {},
	"REPLACE": {},
}

func isValidCopyDirective(value string) bool {
	upper := strings.ToUpper(value)
	if _, ok := supportedCopyDirective[upper]; ok {
		return true
	}
	return false
}
