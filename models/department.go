package models

import "github.com/jinzhu/gorm"

type SysDepart struct {
	ID         string `json:"id"`
	ParentID   string `json:"parent_id"`
	DepartName string `json:"depart_name"`
}

func GetDepartByParentID(ParentID string) ([]*SysDepart, error) {
	var departments []*SysDepart
	err := db.Table("sys_depart").Where("parent_id=?", ParentID).Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func GetDepartByID(id string) (*SysDepart, error) {
	var dt SysDepart
	if err := db.Table("sys_depart").Where("id=?", id).First(&dt).
		Error; err != nil {
		return nil, err
	}
	if len(dt.ID) > 0 {
		return &dt, nil
	}
	return nil, nil
}

func IsLeafDepart(deptId string) bool {
	var dt SysDepart
	err := db.Select("id").Where("parent_id =?", deptId).First(&dt).Error
	if err == gorm.ErrRecordNotFound {
		return true
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	if len(dt.ID) > 0 {
		return false
	}
	return true
}
