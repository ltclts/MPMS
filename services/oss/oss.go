package oss

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/services/log"
	"MPMS/structure"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"mime/multipart"
	"path"
	"time"
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

func Remove() {
	dealDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	log.Info(fmt.Sprintf("清理 %s 之前的未使用的资源数据", dealDate))
	resource := models.Resource{}
	unusedWhere := structure.StringToObjectMap{
		"refer_id":   0,
		"created_at": structure.Array{"<", dealDate},
		"is_deleted": models.UnDeleted,
	}
	deletedWhere := structure.StringToObjectMap{
		"created_at": structure.Array{"<", dealDate},
		"is_deleted": models.Deleted,
	}
	unusedResourceList, err := resource.Select([]string{}, unusedWhere)
	if err != nil {
		log.Err("获取未使用资源信息失败", err.Error())
		return
	}

	deletedResourceList, err := resource.Select([]string{}, deletedWhere)
	if err != nil {
		log.Err("获取已删除资源信息失败", err.Error())
		return
	}

	fmt.Println(unusedResourceList, deletedResourceList)
	//var bucket *oss.Bucket
	//var err error
	//
	//
	//
	//bucket, err = getBucket(ossBucketName)
	//
	// bucket.DeleteObject(url)

}

func getBucket(bucketName string) (*oss.Bucket, error) {
	// New Client
	client, err := oss.New(ossEndPoint, ossAccessId, ossAccessKey)
	if err != nil {
		return nil, err
	}

	// Get Bucket
	return client.Bucket(bucketName)
}
