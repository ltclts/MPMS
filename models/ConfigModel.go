package models

import (
	"MPMS/structure"
)

/**
config 模型
*/
type Config struct {
	Type    uint8
	Content string
	Desc    string
	Model
}

/**
获取配置信息
*/
func (c *Config) Select(fields []string, where structure.StringToObjectMap) ([]Config, error) {
	rows, fieldsAddr, err := c.quickQuery(fields, c.getFieldsMap, where, ConfigTableName)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	var configs []Config
	for rows.Next() {
		err = rows.Scan(fieldsAddr...)
		if err != nil {
			return nil, err
		}
		configs = append(configs, *c)
	}

	return configs, err
}

/**
field与对应关系
*/
func (c *Config) getFieldsMap() structure.StringToObjectMap {
	return structure.StringToObjectMap{
		"id":         &c.Id,
		"type":       &c.Type,
		"content":    &c.Content,
		"desc":       &c.Desc,
		"is_deleted": &c.IsDeleted,
		"creator_id": &c.CreatorId,
		"created_at": &c.CreatedAt,
		"updated_at": &c.UpdatedAt,
	}
}
