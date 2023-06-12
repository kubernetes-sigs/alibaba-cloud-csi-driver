package template

import (
	"fmt"
	"strings"
)

type PVProps struct {
	PVCName      string
	PVCNamespace string
	PVName       string
}

func Evaluate(tm []string, props interface{}, warnOnly bool) (map[string]string, error) {
	md := make(map[string]string)
	for _, s := range tm {
		st := strings.SplitN(s, "=", 2)
		if len(st) != 2 {
			return nil, fmt.Errorf("the key-value pair doesn't contain a value (string: %s)", s)
		}

		key, value := st[0], st[1]
		md[key] = value
	}
	return md, nil
}
