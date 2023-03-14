package sdk

import (
	"fmt"
	"go-admin/common/core/config/sdk/source"
	"go-admin/common/core/config/sdk/source/env"
	file2 "go-admin/common/core/config/sdk/source/file"
	memory2 "go-admin/common/core/config/sdk/source/memory"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func createFileForIssue18(t *testing.T, content string) *os.File {
	data := []byte(content)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}

	return fh
}

func createFileForTest(t *testing.T) *os.File {
	data := []byte(`{"foo": "bar"}`)
	path := filepath.Join(os.TempDir(), fmt.Sprintf("file.%d", time.Now().UnixNano()))
	fh, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	_, err = fh.Write(data)
	if err != nil {
		t.Error(err)
	}

	return fh
}

func TestConfigLoadWithGoodFile(t *testing.T) {
	fh := createFileForTest(t)
	path := fh.Name()
	defer func() {
		fh.Close()
		os.Remove(path)
	}()

	// Create new config
	conf, err := NewConfig()
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	// Load file source
	if err := conf.Load(file2.NewSource(
		file2.WithPath(path),
	)); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
}

func TestConfigLoadWithInvalidFile(t *testing.T) {
	fh := createFileForTest(t)
	path := fh.Name()
	defer func() {
		fh.Close()
		os.Remove(path)
	}()

	// Create new config
	conf, err := NewConfig()
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	// Load file source
	err = conf.Load(file2.NewSource(
		file2.WithPath(path),
		file2.WithPath("/i/do/not/exists.json"),
	))

	if err == nil {
		t.Fatal("Expected error but none !")
	}
	if !strings.Contains(fmt.Sprintf("%v", err), "/i/do/not/exists.json") {
		t.Fatalf("Expected error to contain the unexisting file but got %v", err)
	}
}

func TestConfigMerge(t *testing.T) {
	fh := createFileForIssue18(t, `{
  "amqp": {
    "host": "rabbit.platform",
    "port": 80
  },
  "handler": {
    "exchange": "springCloudBus"
  }
}`)
	path := fh.Name()
	defer func() {
		fh.Close()
		os.Remove(path)
	}()
	os.Setenv("AMQP_HOST", "rabbit.testing.com")

	conf, err := NewConfig()
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if err := conf.Load(
		file2.NewSource(
			file2.WithPath(path),
		),
		env.NewSource(),
	); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	actualHost := conf.Get("amqp", "host").String("backup")
	if actualHost != "rabbit.testing.com" {
		t.Fatalf("Expected %v but got %v",
			"rabbit.testing.com",
			actualHost)
	}
}

func equalS(t *testing.T, actual, expect string) {
	if actual != expect {
		t.Errorf("Expected %s but got %s", actual, expect)
	}
}

func TestConfigWatcherDirtyOverrite(t *testing.T) {
	n := runtime.GOMAXPROCS(0)
	defer runtime.GOMAXPROCS(n)

	runtime.GOMAXPROCS(1)

	l := 100

	ss := make([]source.Source, l, l)

	for i := 0; i < l; i++ {
		ss[i] = memory2.NewSource(memory2.WithJSON([]byte(fmt.Sprintf(`{"key%d": "val%d"}`, i, i))))
	}

	conf, _ := NewConfig()

	for _, s := range ss {
		_ = conf.Load(s)
	}
	runtime.Gosched()

	for i, _ := range ss {
		k := fmt.Sprintf("key%d", i)
		v := fmt.Sprintf("val%d", i)
		equalS(t, conf.Get(k).String(""), v)
	}
}
