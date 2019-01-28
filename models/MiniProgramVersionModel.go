package models

import (
	"MPMS/helper"
	"MPMS/structure"
	"encoding/json"
	"fmt"
)

/**
MiniProgramVersion 模型
*/
type MiniProgramVersion struct {
	MiniProgramId int64
	Code          string
	Remark        string
	ShareWords    string
	Status        uint8
	Type          uint8
	Content       string
	Model
}

const (
	MiniProgramVersionStatusInit      = 0
	MiniProgramVersionStatusApproving = 1 //审核中
	MiniProgramVersionStatusApproved  = 2 //已审核
	MiniProgramVersionStatusOnline    = 3 //已上线
	MiniProgramVersionStatusOffline   = 4 //已下线

	MiniProgramVersionBusinessCard = 1 //名片
)

func MiniProgramVersionTypeToNameMap() structure.Uint8ToStringMap {
	return structure.Uint8ToStringMap{
		MiniProgramVersionBusinessCard: "名片展示",
	}
}

func GetMiniProgramVersionTypeNameByType(_type uint8) (string, error) {
	if name := MiniProgramVersionTypeToNameMap()[_type]; name != "" {
		return name, nil
	}
	return "", helper.CreateNewError(fmt.Sprintf("invalid MiniProgram Type : %d", _type))
}

func (mpv *MiniProgramVersion) GetTypeName() (string, error) {
	return GetMiniProgramVersionTypeNameByType(mpv.Type)
}

func MiniProgramVersionStatusToNameMap() structure.Uint8ToStringMap {
	return structure.Uint8ToStringMap{
		MiniProgramVersionStatusInit:      "编辑中",
		MiniProgramVersionStatusApproving: "审核中",
		MiniProgramVersionStatusApproved:  "已审核",
		MiniProgramVersionStatusOnline:    "已上线",
		MiniProgramVersionStatusOffline:   "已下线",
	}
}

func GetMiniProgramVersionStatusNameByStatus(status uint8) (string, error) {
	if name := MiniProgramVersionStatusToNameMap()[status]; name != "" {
		return name, nil
	}
	return "", helper.CreateNewError(fmt.Sprintf("invalid MiniProgramVersion status : %d", status))
}

func (mpv *MiniProgramVersion) GetStatusName() (string, error) {
	return GetMiniProgramVersionStatusNameByStatus(mpv.Status)
}

func (mpv *MiniProgramVersion) SelectOne(fields []string, where structure.StringToObjectMap) (version MiniProgramVersion, err error) {
	rows, fieldsAddr, err := mpv.quickQuery(fields, mpv.getFieldsMap, where, MiniProgramVersionTableName)
	if err != nil {
		return version, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return version, err
		}
		return *mpv, err
	}

	return version, err
}

type MpvWithRelatedInfo struct {
	Id         int64
	Code       string
	ShareWords string
	Status     uint8
	StatusName string
	MpId       int64
	MpName     string
	CId        int64
	CShortName string
}

func (mpv *MiniProgramVersion) GetList(where structure.StringToObjectMap) (list []MpvWithRelatedInfo, err error) {
	fields := "mpv.`id`,mpv.`code`,mpv.`share_words`,mpv.`status`,mp.id AS mp_id,mp.`name` AS mp_name,c.short_name AS c_short_name,c.id AS c_id"
	sql := "SELECT %s FROM mini_program_version mpv JOIN mini_program mp ON mpv.mini_program_id = mp.id JOIN company c ON mp.company_id = c.id WHERE %s %s"
	whereStr, whereValues := mpv.renderWhereDirectly(where)
	query := fmt.Sprintf(sql, fields, whereStr, "order by mpv.`id` desc")
	rows, err := mpv.query(query, whereValues...)
	if err != nil {
		return list, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		item := MpvWithRelatedInfo{}
		err = rows.Scan(&item.Id, &item.Code, &item.ShareWords, &item.Status, &item.MpId, &item.MpName, &item.CShortName, &item.CId)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, err
}

func (mpv *MiniProgramVersion) Select(fields []string, where structure.StringToObjectMap) ([]MiniProgramVersion, error) {
	rows, fieldsAddr, err := mpv.quickQuery(fields, mpv.getFieldsMap, where, MiniProgramVersionTableName)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	var MiniProgramVersions []MiniProgramVersion
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		MiniProgramVersions = append(MiniProgramVersions, *mpv)
	}

	return MiniProgramVersions, err
}

func (mpv *MiniProgramVersion) Count(where structure.StringToObjectMap) (int64, error) {
	return mpv.count(mpv.getFieldsMap, where, MiniProgramVersionTableName)
}

func (mpv *MiniProgramVersion) Insert(insMap structure.StringToObjectMap) (int64, error) {
	if content := insMap["content"]; content != nil {
		content, err := json.Marshal(content)
		if err != nil {
			return 0, err
		}
		insMap["content"] = content
	}
	return mpv.insertExec(insMap, mpv.getFieldsMap, MiniProgramVersionTableName)
}

func (mpv *MiniProgramVersion) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
	if content := toUpdate["content"]; content != nil {
		content, err := json.Marshal(content)
		if err != nil {
			return 0, err
		}
		toUpdate["content"] = content
	}
	return mpv.updateExec(toUpdate, where, mpv.getFieldsMap, MiniProgramVersionTableName)
}

/**
field与对应关系
*/
func (mpv *MiniProgramVersion) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":              &mpv.Id,
		"code":            &mpv.Code,
		"mini_program_id": &mpv.MiniProgramId,
		"share_words":     &mpv.ShareWords,
		"remark":          &mpv.Remark,
		"content":         &mpv.Content,
		"type":            &mpv.Type,
		"status":          &mpv.Status,
		"is_deleted":      &mpv.IsDeleted,
		"creator_id":      &mpv.CreatorId,
		"created_at":      &mpv.CreatedAt,
		"updated_at":      &mpv.UpdatedAt,
	}
}
