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

/**
获取流水信息
*/
func (f *Flow) Select(fields []string, where structure.Map) ([]Flow, error) {
	rows, fieldsAddr, err := f.QuickQuery(fields, f.getFieldsMap, where, FlowTableName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
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

func (f *Flow) Insert(referId int64, referType uint8, status uint8, creatorId int64, contentMap structure.Map) (int64, error) {
	content, err := json.Marshal(contentMap)
	if err != nil {
		return 0, nil
	}

	return f.InsertExec(structure.Map{
		"refer_id":   referId,
		"refer_type": referType,
		"status":     status,
		"creator_id": creatorId,
		"content":    content,
	}, FlowTableName)
}

/**
field与对应关系
*/
func (f *Flow) getFieldsMap() structure.Map {
	return structure.Map{
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
