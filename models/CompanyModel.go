package models

import (
	"MPMS/helper"
	"MPMS/structure"
	"fmt"
)

/**
company 模型
*/
type Company struct {
	Name      string
	ShortName string
	Remark    string
	ExpireAt  string
	Status    uint8
	Model
}

const (
	CompanyStatusInit      = 0 //初始状态
	CompanyStatusInUse     = 1 //启用
	CompanyStatusForbidden = 2 //禁用
)

func CompanyStatusToNameMap() structure.Uint8ToStringMap {
	return structure.Uint8ToStringMap{
		CompanyStatusInit:      "已创建",
		CompanyStatusInUse:     "已启用",
		CompanyStatusForbidden: "已禁用",
	}
}

func GetCompanyStatusNameByStatus(status uint8) (string, error) {
	if name := CompanyStatusToNameMap()[status]; name != "" {
		return name, nil
	}
	return "", helper.CreateNewError(fmt.Sprintf("invalid company status : %d", status))
}

func (c *Company) GetStatusName() (string, error) {
	return GetCompanyStatusNameByStatus(c.Status)
}

func (c *Company) GetCompanyByContactUserId(userId int64) (company Company, err error) {
	relation := Relation{}
	relations, err := relation.Select([]string{"refer_id"}, structure.StringToObjectMap{
		"is_deleted":      UnDeleted,
		"refer_type":      RelationReferTypeCompanyContactUser,
		"refer_id_others": userId,
	})
	if err != nil {
		return company, err
	}

	if 1 != len(relations) {
		return company, helper.CreateNewError(fmt.Sprintf("所在公司信息不唯一！user_id:%d", userId))
	}

	companies, err := company.Select([]string{"id", "name", "short_name", "creator_id", "expire_at"}, structure.StringToObjectMap{
		"is_deleted": UnDeleted,
		"id":         relations[0].ReferId,
	})

	if err != nil {
		return company, err
	}
	if len(companies) >= 1 {
		return companies[0], nil
	} else {
		return company, helper.CreateNewError("没有相关的公司信息！")
	}
}

/**
获取公司信息
*/
func (c *Company) Select(fields []string, where structure.StringToObjectMap) ([]Company, error) {
	rows, fieldsAddr, err := c.quickQuery(fields, c.getFieldsMap, where, CompanyTableName)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	var companies []Company
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		companies = append(companies, *c)
	}

	return companies, err
}

func (c *Company) Insert(insMap structure.StringToObjectMap) (int64, error) {
	return c.insertExec(insMap, c.getFieldsMap, CompanyTableName)
}

func (c *Company) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
	return c.updateExec(toUpdate, where, c.getFieldsMap, CompanyTableName)
}

/**
field与对应关系
*/
func (c *Company) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":         &c.Id,
		"name":       &c.Name,
		"short_name": &c.ShortName,
		"remark":     &c.Remark,
		"status":     &c.Status,
		"expire_at":  &c.ExpireAt,
		"is_deleted": &c.IsDeleted,
		"creator_id": &c.CreatorId,
		"created_at": &c.CreatedAt,
		"updated_at": &c.UpdatedAt,
	}
}
