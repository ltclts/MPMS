package models

import (
	"MPMS/structure"
	"encoding/json"
)

/**
通用流水 模型
*/
type Flow struct {
	ReferId   int64
	ReferType uint8
	Status    uint8
	Content   string
	Model
}

const (
	FlowReferTypeMinProgram        = 1 //小程序相关
	FlowReferTypeCompany           = 2 //公司相关
	FlowReferTypeContactUser       = 3 //联系人相关
	FlowReferTypeMinProgramVersion = 4 //小程序版本管理
	FlowReferTypeResource          = 5 //资源管理
	FlowReferTypeMPMPV             = 6 //小程序的版本相关

	FlowStatusCreate = 1 //创建
	FlowStatusEdit   = 2 //编辑
	FlowStatusDelete = 3 //删除

)

/**
获取流水信息
*/
func (f *Flow) Select(fields []string, where structure.StringToObjectMap) ([]Flow, error) {
	rows, fieldsAddr, err := f.quickQuery(fields, f.getFieldsMap, where, FlowTableName)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	var flows []Flow
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		flows = append(flows, *f)
	}

	return flows, err
}

func (f *Flow) Insert(referId int64, referType uint8, status uint8, creatorId int64, contentMap structure.StringToObjectMap) (int64, error) {
	content, err := json.Marshal(contentMap)
	if err != nil {
		return 0, err
	}

	return f.insertExec(structure.StringToObjectMap{
		"refer_id":   referId,
		"refer_type": referType,
		"status":     status,
		"creator_id": creatorId,
		"content":    content,
	}, f.getFieldsMap, FlowTableName)
}

/**
field与对应关系
*/
func (f *Flow) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":         &f.Id,
		"refer_id":   &f.ReferId,
		"refer_type": &f.ReferType,
		"content":    &f.Content,
		"status":     &f.Status,
		"is_deleted": &f.IsDeleted,
		"creator_id": &f.CreatorId,
		"created_at": &f.CreatedAt,
		"updated_at": &f.UpdatedAt,
	}
}
