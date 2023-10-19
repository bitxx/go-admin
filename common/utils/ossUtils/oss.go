package ossUtils

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"time"
)

type ALiYunOSS struct {
	Client          *oss.Client
	Bucket          *oss.Bucket
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

func (e *ALiYunOSS) InitOssClient(key, secret, endpoint, bucketName string) error {
	e.AccessKeyId = key
	e.AccessKeySecret = secret
	e.Endpoint = endpoint
	e.BucketName = bucketName

	client, err := oss.New(e.Endpoint, e.AccessKeyId, e.AccessKeySecret)
	if err != nil {
		return errors.New(fmt.Sprintf("初始化Oss客户端异常：%s", err.Error()))
	}
	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return errors.New(fmt.Sprintf("初始化Oss Bucket异常：%s", err.Error()))
	}
	e.Client = client
	e.Bucket = bucket
	return nil
}

// GeneratePresignedUrl
// @Description: 获取下载链接
// @receiver e
// @param key
// @return string
// @return error
func (e *ALiYunOSS) GeneratePresignedUrl(key string) (string, error) {
	if e.AccessKeyId == "" || e.AccessKeySecret == "" || e.Endpoint == "" || e.BucketName == "" {
		return "", errors.New("please init OSSclient")
	}
	path, err := e.Bucket.SignURL(key, oss.HTTPGet, 86400)
	if err != nil {
		return "", errors.New(fmt.Sprintf("general oss url err：%s", err.Error()))
	}
	return path, nil
}

// Upload
// @Description: 上传文件
// @receiver e
// @param objectKey
// @param localPath
// @return error
func (e *ALiYunOSS) Upload(objectKey, localPath string) error {
	// 上传本地文件
	return e.Bucket.PutObjectFromFile(objectKey, localPath)
}

// UploadWithSpace
// @Description: 分片上传
// @receiver e
// @param objectKey
// @param localPath
// @return error
func (e *ALiYunOSS) UploadWithSpace(objectKey, localPath string) error {
	// 上传本地文件
	chunks, err := oss.SplitFileByPartNum(localPath, 5)
	if err != nil {
		return err
	}
	fd, err := os.Open(localPath)
	defer fd.Close()

	// 指定过期时间。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
		// 指定该Object被下载时的网页缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定该Object的内容编码格式。
		// oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}
	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
	imur, err := e.Bucket.InitiateMultipartUpload(objectKey, options...)
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := e.Bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			return err
		}
		parts = append(parts, part)
	}
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	_, err = e.Bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		return err
	}
	return nil
}

// UpLoad 文件上传
/*func (e *ALiYunOSS) UpLoad(yourObjectName string, localFile string) error {
	err := e.Setup()
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	// 获取存储空间。
	bucket, err := e.Client.(*oss.Client).Bucket(e.BucketName)
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	// 设置分片大小为100 KB，指定分片上传并发数为3，并开启断点续传上传。
	// 其中<yourObjectName>与objectKey是同一概念，表示断点续传上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// "LocalFile"为filePath，100*1024为partSize。
	err = bucket.UploadFile(yourObjectName, localFile, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}
*/
