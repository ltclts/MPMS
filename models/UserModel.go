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

const UserTableName = "user"

const (
	UserTypeAdmin    = 0 //后台管理员
	UserTypeCustomer = 1 //用户
)

const (
	UserStatusInitial   = 0 //初始状态
	UserStatusInUse     = 1 //在用状态
	UserStatusForbidden = 2 //禁用状态
)

/**
获取用户状态与名称的映射关系
*/
func (u *User) GetUserStatusToNameMap() map[uint8]string {
	return map[uint8]string{
		UserStatusInitial:   "未激活",
		UserStatusInUse:     "已激活",
		UserStatusForbidden: "已禁用",
	}
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
	rows, fieldsAddr, err := u.QuickQuery(fields, u.getFieldsMap, where, UserTableName)
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
