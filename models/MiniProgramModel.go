package models

import (
	"MPMS/structure"
)

/**
MiniProgram 模型
*/
type MiniProgram struct {
	CompanyId int64
	Name      string
	Remark    string
	Appid     string
	Model
}

func (mp *MiniProgram) SelectOne(fields []string, where structure.StringToObjectMap) (mpIns MiniProgram, err error) {
	rows, fieldsAddr, err := mp.quickQuery(fields, mp.getFieldsMap, where, MiniProgramTableName)
	if err != nil {
		return mpIns, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return mpIns, err
		}
		return *mp, err
	}
	return mpIns, err
}

func (mp *MiniProgram) Select(fields []string, where structure.StringToObjectMap) ([]MiniProgram, error) {
	rows, fieldsAddr, err := mp.quickQuery(fields, mp.getFieldsMap, where, MiniProgramTableName)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	var miniPrograms []MiniProgram
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		miniPrograms = append(miniPrograms, *mp)
	}

	return miniPrograms, err
}

func (mp *MiniProgram) Count(where structure.StringToObjectMap) (int64, error) {
	return mp.count(mp.getFieldsMap, where, MiniProgramTableName)
}

func (mp *MiniProgram) Insert(insMap structure.StringToObjectMap) (int64, error) {
	return mp.insertExec(insMap, mp.getFieldsMap, MiniProgramTableName)
}

func (mp *MiniProgram) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
	return mp.updateExec(toUpdate, where, mp.getFieldsMap, MiniProgramTableName)
}

/**
field与对应关系
*/
func (mp *MiniProgram) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":         &mp.Id,
		"name":       &mp.Name,
		"remark":     &mp.Remark,
		"appid":      &mp.Appid,
		"company_id": &mp.CompanyId,
		"is_deleted": &mp.IsDeleted,
		"creator_id": &mp.CreatorId,
		"created_at": &mp.CreatedAt,
		"updated_at": &mp.UpdatedAt,
	}
}
