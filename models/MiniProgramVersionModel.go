package models

import (
	"MPMS/helper"
	"MPMS/structure"
	"fmt"
)

/**
MiniProgramVersion 模型
*/
type MiniProgramVersion struct {
	MiniProgramId int64
	Code          string
	Status        uint8
	Model
}

//0-编辑中 1-审核中 2-审核通过 3-已上线 4-已下线

const (
	MiniProgramVersionStatusInit      = 0
	MiniProgramVersionStatusApproving = 1 //审核中
	MiniProgramVersionStatusApproved  = 2 //已审核
	MiniProgramVersionStatusOnline    = 3 //已上线
	MiniProgramVersionStatusOffline   = 4 //已下线
)

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
	return mpv.insertExec(insMap, mpv.getFieldsMap, MiniProgramVersionTableName)
}

func (mpv *MiniProgramVersion) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
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
		"status":          &mpv.Status,
		"is_deleted":      &mpv.IsDeleted,
		"creator_id":      &mpv.CreatorId,
		"created_at":      &mpv.CreatedAt,
		"updated_at":      &mpv.UpdatedAt,
	}
}
