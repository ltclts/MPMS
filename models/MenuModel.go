package models

import (
	"MPMS/structure"
)

/**
Menu 模型
*/
type Menu struct {
	Id       uint
	ParentId uint
	Type     uint8
	Name     string
	NameEn   string
	Uri      string
	Sort     uint
	Model
}

const (
	MenuTypeFirst  = 1
	MenuTypeSecond = 2
)

/**
获取菜单信息
*/
func (m *Menu) Select(fields []string, where structure.Map) ([]Menu, error) {
	rows, fieldsAddr, err := m.QuickQueryWithExtra(fields, m.getFieldsMap, where, MenuTableName, "order by sort")
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
func (m *Menu) getFieldsMap() structure.Map {
	return structure.Map{
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
