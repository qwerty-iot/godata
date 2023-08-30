package edmx

import "github.com/qwerty-iot/tox"

type Options struct {
	nullable  *bool
	open      bool
	maxLength int
}

func NewOptions() *Options {
	return &Options{}
}
func (o *Options) Nullable() *Options {
	o.nullable = tox.ToBoolPtr(true)
	return o
}

func (o *Options) NotNullable() *Options {
	o.nullable = tox.ToBoolPtr(false)
	return o
}

func (o *Options) OpenType() *Options {
	o.open = true
	return o
}

func (o *Options) MaxLength(length int) *Options {
	o.maxLength = length
	return o
}

func (o *Options) getNullable() string {
	if o == nil {
		return ""
	}
	if o.nullable == nil {
		return ""
	}
	if *o.nullable {
		return "true"
	}
	return "false"
}

func (o *Options) getOpen() string {
	if o == nil {
		return ""
	}
	if o.open {
		return "true"
	}
	return ""
}

func (o *Options) getMaxLength() string {
	if o == nil {
		return ""
	}
	if o.maxLength == 0 {
		return ""
	}
	return tox.ToString(o.maxLength)
}
