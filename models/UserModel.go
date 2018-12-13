package models

import (
	"MPMS/helper"
	"MPMS/structure"
)

/**
user 模型
*/
type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
	Phone    string
	Type     uint8
	Status   uint8
	Sort     int
	model
}

/**
检查密码是否正确
*/
func (u *User) CheckPwd(pwd string) bool {
	pwd = helper.Md5(pwd)
	return pwd == u.Password
}

/**
获取用户信息
*/
func (u *User) Select(fields []string, where structure.Map) ([]User, error) {
	whereStr, whereValue := u.renderWhere(where)
	fieldsStr, fieldsAddr, err := u.renderFields(fields, u.getFieldsMap)
	if err != nil {
		return nil, err
	}

	sqlStr := "SELECT " + fieldsStr + "  FROM `user` where " + whereStr
	rows, err := u.Query(sqlStr, whereValue...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	return users, err
}

/**
field与对应关系
*/
func (u *User) getFieldsMap() structure.Map {
	return structure.Map{
		"id":         &u.Id,
		"name":       &u.Name,
		"email":      &u.Email,
		"phone":      &u.Phone,
		"password":   &u.Password,
		"sort":       &u.Sort,
		"status":     &u.Status,
		"type":       &u.Type,
		"is_deleted": &u.IsDeleted,
		"creator_id": &u.CreatorId,
		"created_at": &u.CreatedAt,
		"updated_at": &u.UpdatedAt,
	}
}
