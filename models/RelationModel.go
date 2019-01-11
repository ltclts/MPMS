package models

import (
	"MPMS/structure"
)

/**
通用关系 模型
*/
type Relation struct {
	ReferId       int64
	ReferIdOthers int64
	ReferType     uint8
	Model
}

/**
获取关系信息
*/
func (r *Relation) Select(fields []string, where structure.StringToObjectMap) ([]Relation, error) {
	rows, fieldsAddr, err := r.QuickQuery(fields, r.getFieldsMap, where, RelationTableName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var relations []Relation
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		relations = append(relations, *r)
	}

	return relations, err
}

func (r *Relation) Insert(referType uint8, referId int64, referIdOthers int64, creatorId int64) (int64, error) {
	return r.InsertExec(structure.StringToObjectMap{
		"refer_type":      referType,
		"refer_id":        referId,
		"refer_id_others": referIdOthers,
		"creator_id":      creatorId,
	}, RelationTableName)
}

/**
field与对应关系
*/
func (r *Relation) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":              &r.Id,
		"refer_id":        &r.ReferId,
		"refer_id_others": &r.ReferIdOthers,
		"refer_type":      &r.ReferType,
		"is_deleted":      &r.IsDeleted,
		"creator_id":      &r.CreatorId,
		"created_at":      &r.CreatedAt,
		"updated_at":      &r.UpdatedAt,
	}
}
