package models

import (
	"MPMS/structure"
)

/**
Menu 模型
*/
type Menu struct {
	ParentId int64
	Type     uint8
	Name     string
	NameEn   string
	Uri      string
	Sort     int64
	Model
}

const (
	MenuTypeFirst  = 1 //一级菜单
	MenuTypeSecond = 2 //二级菜单
)

/**
获取菜单信息
*/
func (m *Menu) Select(fields []string, where structure.StringToObjectMap) ([]Menu, error) {
	rows, fieldsAddr, err := m.quickQueryWithExtra(fields, m.getFieldsMap, where, MenuTableName, "order by sort")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var menus []Menu
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		menus = append(menus, *m)
	}

	return menus, err
}

/**
field与对应关系
*/
func (m *Menu) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":         &m.Id,
		"parent_id":  &m.ParentId,
		"type":       &m.Type,
		"name":       &m.Name,
		"name_en":    &m.NameEn,
		"uri":        &m.Uri,
		"sort":       &m.Sort,
		"is_deleted": &m.IsDeleted,
		"creator_id": &m.CreatorId,
		"created_at": &m.CreatedAt,
		"updated_at": &m.UpdatedAt,
	}
}
