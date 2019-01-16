package models

import (
	"MPMS/helper"
	"MPMS/structure"
	"fmt"
)

/**
user 模型
*/
type User struct {
	Name     string
	Email    string
	Password string
	Phone    string
	Type     uint8
	Status   uint8
	Sort     int64
	Model
}

const (
	//用户类型定义
	UserTypeAdmin    = 0 //后台管理员
	UserTypeCustomer = 1 //用户

	//用户状态定义
	UserStatusInitial   = 0 //初始状态
	UserStatusInUse     = 1 //在用状态
	UserStatusForbidden = 2 //禁用状态
)

/**
获取用户状态与名称的映射关系
*/
func UserStatusToNameMap() structure.Uint8ToStringMap {
	return structure.Uint8ToStringMap{
		UserStatusInitial:   "未激活",
		UserStatusInUse:     "已激活",
		UserStatusForbidden: "已禁用",
	}
}

func GetUserStatusNameByStatus(status uint8) (string, error) {
	if name := UserStatusToNameMap()[status]; name != "" {
		return name, nil
	}
	return "", helper.CreateNewError(fmt.Sprintf("invalid user status : %d", status))
}

func (u *User) GetStatusName() (string, error) {
	return GetUserStatusNameByStatus(u.Status)
}

/**
检查密码是否正确
*/
func (u *User) CheckPwd(pwd string) bool {
	pwd = helper.Md5(pwd)
	return pwd == u.Password
}

func (u *User) CreateContactUser(companyId int64, creatorId int64) (int64, error) {
	//获取6位随机密码
	password := helper.GetRandomStrBy(6)
	userId, err := u.Insert(structure.StringToObjectMap{
		"name":       u.Name,
		"email":      u.Email,
		"phone":      u.Phone,
		"password":   helper.Md5(password),
		"status":     UserStatusInUse,
		"type":       UserTypeCustomer,
		"creator_id": creatorId,
	})
	if err != nil {
		return userId, err
	}
	fmt.Println(fmt.Sprintf("user=%d password=%s", userId, password))
	relation := Relation{}
	_, err = relation.Insert(RelationReferTypeCompanyContactUser, companyId, userId, creatorId)
	if err != nil {
		return userId, err
	}
	flow := Flow{}
	_, err = flow.Insert(userId, FlowReferTypeContactUser, FlowStatusCreate, creatorId, structure.StringToObjectMap{})
	//todo 发送邮件
	return userId, err
}

func (u *User) GetContactUserByCompanyId(companyId int64) (user User, err error) {
	relation := Relation{}
	relations, err := relation.Select([]string{"refer_id_others"}, structure.StringToObjectMap{
		"is_deleted": UnDeleted,
		"refer_type": RelationReferTypeCompanyContactUser,
		"refer_id":   companyId,
	})
	if err != nil {
		return user, err
	}
	if len(relations) == 1 {
		//联系人
		users, err := u.Select([]string{"name", "phone", "id", "email"}, structure.StringToObjectMap{
			"is_deleted": UnDeleted,
			"id":         relations[0].ReferIdOthers,
		})
		if err != nil {
			return user, err
		}
		if len(users) != 1 {
			return user, helper.CreateNewError("没有获取到联系人信息")
		}
		return users[0], nil
	}
	return user, helper.CreateNewError("没有获取到联系人信息或者联系人信息不唯一！")
}

/**
获取用户信息
*/
func (u *User) Select(fields []string, where structure.StringToObjectMap) ([]User, error) {
	rows, fieldsAddr, err := u.quickQuery(fields, u.getFieldsMap, where, UserTableName)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
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
获取用户信息
*/
func (u *User) SelectOne(fields []string, where structure.StringToObjectMap) (user User, err error) {
	rows, fieldsAddr, err := u.quickQueryWithExtra(fields, u.getFieldsMap, where, UserTableName, "limit 1")
	if err != nil {
		return user, err
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return user, err
		}
		return *u, err
	}

	return user, err
}

func (u *User) Insert(insMap structure.StringToObjectMap) (int64, error) {
	return u.insertExec(insMap, u.getFieldsMap, UserTableName)
}

func (u *User) Update(toUpdate structure.StringToObjectMap, where structure.StringToObjectMap) (int64, error) {
	return u.updateExec(toUpdate, where, u.getFieldsMap, UserTableName)
}

/**
field与对应关系
*/
func (u *User) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
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
