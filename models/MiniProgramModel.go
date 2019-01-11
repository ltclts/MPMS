package models

import (
	"MPMS/helper"
	"MPMS/structure"
	"encoding/json"
	"fmt"
)

/**
MiniProgram 模型
*/
type MiniProgram struct {
	CompanyId int64
	Name      string
	Remark    string
	Appid     string
	Content   string
	Status    uint8
	Type      uint8
	Model
}

const (
	MiniProgramStatusInit      = 0
	MiniProgramStatusInUse     = 1 //启用
	MiniProgramStatusForbidden = 2 //禁用

	MiniProgramTypeBusinessCard = 1 //名片
)

func MiniProgramTypeToNameMap() structure.Uint8ToStringMap {
	return structure.Uint8ToStringMap{
		MiniProgramTypeBusinessCard: "名片展示",
	}
}

func GetMiniProgramTypeNameByType(_type uint8) (string, error) {
	if name := MiniProgramTypeToNameMap()[_type]; name != "" {
		return name, nil
	}
	return "", helper.CreateNewError(fmt.Sprintf("invalid MiniProgram Type : %d", _type))
}

func (mp *MiniProgram) GetTypeName() (string, error) {
	return GetMiniProgramTypeNameByType(mp.Type)
}

func MiniProgramStatusToNameMap() structure.Uint8ToStringMap {
	return structure.Uint8ToStringMap{
		MiniProgramStatusInit:      "编辑中",
		MiniProgramStatusInUse:     "已启用",
		MiniProgramStatusForbidden: "已禁用",
	}
}

func GetMiniProgramStatusNameByStatus(status uint8) (string, error) {
	if name := MiniProgramStatusToNameMap()[status]; name != "" {
		return name, nil
	}
	return "", helper.CreateNewError(fmt.Sprintf("invalid MiniProgram status : %d", status))
}

func (mp *MiniProgram) GetStatusName() (string, error) {
	return GetMiniProgramStatusNameByStatus(mp.Status)
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

func (mp *MiniProgram) Insert(insMap structure.StringToObjectMap) (int64, error) {
	if content := insMap["content"]; content != nil {
		content, err := json.Marshal(content)
		if err != nil {
			return 0, err
		}
		insMap["content"] = content
	}
	return mp.insertExec(insMap, mp.getFieldsMap, MiniProgramTableName)
}

func (mp *MiniProgram) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
	if content := toUpdate["content"]; content != nil {
		content, err := json.Marshal(content)
		if err != nil {
			return 0, err
		}
		toUpdate["content"] = content
	}
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
		"status":     &mp.Status,
		"content":    &mp.Content,
		"type":       &mp.Type,
		"is_deleted": &mp.IsDeleted,
		"creator_id": &mp.CreatorId,
		"created_at": &mp.CreatedAt,
		"updated_at": &mp.UpdatedAt,
	}
}
