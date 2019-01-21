package api

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/services/oss"
	"MPMS/structure"
	"fmt"
	"github.com/astaxie/beego"
	"mime/multipart"
	"strings"
)

type ResourceApiController struct {
	Controller
}

func Upload(f multipart.File, h *multipart.FileHeader, referType uint8, referId int64, creatorId int64) (resourceId int64, url string, err error) {
	ext := Ext(h.Filename)
	if !CheckImgExt(ext) {
		return resourceId, url, helper.CreateNewError(fmt.Sprintf("无效的扩展名 %s", ext))
	}

	url, err = oss.Upload(f, h, fmt.Sprintf("%d/", referType))
	if err != nil {
		return resourceId, url, err
	}
	resource := models.Resource{}
	resourceId, err = resource.Insert(structure.StringToObjectMap{
		"refer_type":    referType,
		"refer_id":      referId,
		"origin_name":   h.Filename,
		"relative_path": url,
		"ext":           ext,
		"store_type":    models.ResourceStoreTypeAliYunOss,
		"size":          h.Size,
		"creator_id":    creatorId,
	})
	return resourceId, fmt.Sprintf(beego.AppConfig.String("oss.pathurl")+"%s", url), err
}

func Ext(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}

func CheckImgExt(ext string) bool {
	allowed := false
	for _, extItem := range []string{"BMP", "JPG", "JPEG", "PNG", "GIF"} {
		if extItem == strings.ToUpper(ext) {
			allowed = true
			break
		}
	}
	return allowed
}
