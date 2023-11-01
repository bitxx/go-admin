package fileutils

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetType 获取文件类型
func GetType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		return "", err
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		return "", err
	}

	filetype := http.DetectContentType(buff)

	//ext := GetExt(p)
	//var list = strings.Split(filetype, "/")
	//filetype = list[0] + "/" + ext
	return filetype, nil
}

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// PathExist 判断目录是否存在
func PathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func FileCreate(content bytes.Buffer, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		return err
	}
	return file.Close()
}

func CreateDirFromFilePath(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func ZipFilCreate(writer *zip.Writer, content bytes.Buffer, path string) error {

	f, err := writer.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(content.Bytes())
	return err
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		return nil
	}

	//文件类型需要进行过滤

	buf, err := os.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	_ = os.WriteFile(path, []byte(newContent), 0)

	return err
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, []byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		if ctx.Err() != nil {
			break
		}
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			return err
		}
		go hookfn(ctx, id, group, line)
	}
	return nil
}

// GetFileSize get file size
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// GetCurrentPath get current dir
func GetCurrentPath() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}
