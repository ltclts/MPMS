package models

import (
	"MPMS/structure"
)

/**
config 模型
*/
type Config struct {
	Id      uint
	Type    uint8
	Content string
	model
}

const ConfigTableName = "config"

/**
获取配置信息
*/
func (c *Config) Select(fields []string, where structure.Map) ([]Config, error) {
	rows, fieldsAddr, err := c.QuickQuery(fields, c.getFieldsMap, where, ConfigTableName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
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
func (c *Config) getFieldsMap() structure.Map {
	return structure.Map{
		"id":         &c.Id,
		"type":       &c.Type,
		"content":    &c.Content,
		"is_deleted": &c.IsDeleted,
		"creator_id": &c.CreatorId,
		"created_at": &c.CreatedAt,
		"updated_at": &c.UpdatedAt,
	}
}
