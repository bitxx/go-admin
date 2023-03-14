package memory

import (
	"go-admin/common/core/config/sdk/loader"
	"go-admin/common/core/config/sdk/reader"
	"go-admin/common/core/config/sdk/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
