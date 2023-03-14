package flag

import (
	"errors"
	"flag"
	"github.com/imdario/mergo"
	source2 "go-admin/common/core/config/sdk/source"
	"strings"
	"time"
)

type flagsrc struct {
	opts source2.Options
}

func (fs *flagsrc) Read() (*source2.ChangeSet, error) {
	if !flag.Parsed() {
		return nil, errors.New("flags not parsed")
	}

	var changes map[string]interface{}

	visitFn := func(f *flag.Flag) {
		n := strings.ToLower(f.Name)
		keys := strings.FieldsFunc(n, split)
		reverse(keys)

		tmp := make(map[string]interface{})
		for i, k := range keys {
			if i == 0 {
				tmp[k] = f.Value
				continue
			}

			tmp = map[string]interface{}{k: tmp}
		}

		_ = mergo.Map(&changes, tmp) // need to sort error handling
		return
	}

	unset, ok := fs.opts.Context.Value(includeUnsetKey{}).(bool)
	if ok && unset {
		flag.VisitAll(visitFn)
	} else {
		flag.Visit(visitFn)
	}

	b, err := fs.opts.Encoder.Encode(changes)
	if err != nil {
		return nil, err
	}

	cs := &source2.ChangeSet{
		Format:    fs.opts.Encoder.String(),
		Data:      b,
		Timestamp: time.Now(),
		Source:    fs.String(),
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func split(r rune) bool {
	return r == '-' || r == '_'
}

func reverse(ss []string) {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
}

func (fs *flagsrc) Watch() (source2.Watcher, error) {
	return source2.NewNoopWatcher()
}

func (fs *flagsrc) Write(cs *source2.ChangeSet) error {
	return nil
}

func (fs *flagsrc) String() string {
	return "flag"
}

// NewSource returns a config source for integrating parsed flags.
// Hyphens are delimiters for nesting, and all keys are lowercased.
//
// Example:
//      dbhost := flag.String("database-host", "localhost", "the db host name")
//
//      {
//          "database": {
//              "host": "localhost"
//          }
//      }
func NewSource(opts ...source2.Option) source2.Source {
	return &flagsrc{opts: source2.NewOptions(opts...)}
}
