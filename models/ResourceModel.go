package models

import (
	"MPMS/structure"
)

type Resource struct {
	ReferType    uint8
	ReferId      int64
	Size         int64
	OriginName   string
	RelativePath string
	Ext          string
	StoreType    uint8
	Sort         int64
	Model
}

const (
	//小程序 名片展示 轮播图
	ResourceReferTypeMiniProgramVersionBusinessCardCarousel = 1

	ResourceStoreTypeAliYunOss = 1
)

func (r *Resource) Select(fields []string, where structure.StringToObjectMap) ([]Resource, error) {
	rows, fieldsAddr, err := r.quickQueryWithExtra(fields, r.getFieldsMap, where, UserTableName, "order by sort asc")
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	var resources []Resource
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		resources = append(resources, *r)
	}

	return resources, err
}

func (r *Resource) SelectOne(fields []string, where structure.StringToObjectMap) (resource Resource, err error) {
	rows, fieldsAddr, err := r.quickQueryWithExtra(fields, r.getFieldsMap, where, ResourceTableName, "limit 1")
	if err != nil {
		return resource, err
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return resource, err
		}
		return *r, err
	}

	return resource, err
}

func (r *Resource) Insert(insMap structure.StringToObjectMap) (int64, error) {
	return r.insertExec(insMap, r.getFieldsMap, ResourceTableName)
}

func (r *Resource) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
	return r.updateExec(toUpdate, where, r.getFieldsMap, ResourceTableName)
}

/**
field与对应关系
*/
func (r *Resource) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":            &r.Id,
		"refer_type":    &r.ReferType,
		"refer_id":      &r.ReferId,
		"origin_name":   &r.OriginName,
		"relative_path": &r.RelativePath,
		"store_type":    &r.StoreType,
		"ext":           &r.Ext,
		"sort":          &r.Sort,
		"size":          &r.Size,
		"is_deleted":    &r.IsDeleted,
		"creator_id":    &r.CreatorId,
		"created_at":    &r.CreatedAt,
		"updated_at":    &r.UpdatedAt,
	}
}
