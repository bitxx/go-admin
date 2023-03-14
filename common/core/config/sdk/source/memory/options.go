package memory

import (
	"context"
	source2 "go-admin/common/core/config/sdk/source"
)

type changeSetKey struct{}

func withData(d []byte, f string) source2.Option {
	return func(o *source2.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, changeSetKey{}, &source2.ChangeSet{
			Data:   d,
			Format: f,
		})
	}
}

// WithChangeSet allows a changeset to be set
func WithChangeSet(cs *source2.ChangeSet) source2.Option {
	return func(o *source2.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, changeSetKey{}, cs)
	}
}

// WithJSON allows the source data to be set to json
func WithJSON(d []byte) source2.Option {
	return withData(d, "json")
}

// WithYAML allows the source data to be set to yaml
func WithYAML(d []byte) source2.Option {
	return withData(d, "yaml")
}
