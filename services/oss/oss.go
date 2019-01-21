package oss

import (
	"MPMS/helper"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"mime/multipart"
	"path"
)

var (
	ossBucketName = beego.AppConfig.String("oss.bucketname")
	ossEndPoint   = beego.AppConfig.String("oss.endpoint")
	ossAccessId   = beego.AppConfig.String("oss.accessid")
	ossAccessKey  = beego.AppConfig.String("oss.accesskey")
)

func Upload(f multipart.File, h *multipart.FileHeader, dir string) (url string, err error) {
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return url, err
	}
	uuidV, err := uuid.NewV4()
	if err != nil {
		return url, err
	}
	filename := helper.Md5(uuidV.String()+helper.GetRandomStrBy(6)) + path.Ext(h.Filename)
	bucket, err := getBucket(ossBucketName)
	if err != nil {
		return url, err
	}

	//相对路径
	url = dir + filename
	err = bucket.PutObject(url, f)
	if err != nil {
		return url, err
	}

	return url, err
}

func getBucket(bucketName string) (*oss.Bucket, error) {
	// New Client
	client, err := oss.New(ossEndPoint, ossAccessId, ossAccessKey)
	if err != nil {
		return nil, err
	}

	// Create Bucket
	err = client.CreateBucket(bucketName)
	if err != nil {
		return nil, err
	}

	// Get Bucket
	return client.Bucket(bucketName)
}
