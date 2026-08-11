package datacache

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/resource"
)

type Mode string

const (
	Writethrough Mode = "writethrough"
	Writeback    Mode = "writeback"
)

type Opts struct {
	Size resource.Quantity
	Mode Mode
}

func (o *Opts) Enabled() bool {
	return !o.Size.IsZero()
}

const (
	modeKey = "dataCacheMode"
	sizeKey = "dataCacheSize"
)

func GetOpts(opts map[string]string, d *Opts) error {
	if s := opts[sizeKey]; s != "" {
		size, err := resource.ParseQuantity(s)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", sizeKey, err)
		}
		d.Size = size
	}

	switch m := Mode(opts[modeKey]); m {
	case "", Writeback, Writethrough:
		d.Mode = m
	default:
		return fmt.Errorf("unrecognized %s: %s", modeKey, m)
	}

	if d.Mode != "" || !d.Size.IsZero() {
		if d.Size.IsZero() {
			return fmt.Errorf("must specify non-zero %s for dataCache", sizeKey)
		}
		if d.Mode == "" {
			d.Mode = Writethrough
		}
	}
	return nil
}
