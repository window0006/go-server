package entity

import (
	"github.com/window0006/go-server/bindings"
	"github.com/window0006/go-server/dao/dal"
)

func CreateFamily(name string) (dal.Family, error) {
	// 从 db 中获取 family 列表
	newFamily := dal.Family{
		Name: name,
	}
	// 创建新的 family，并返回 id
	db := DB.SelectDB(true)
	return newFamily, db.Create(&newFamily).Error
}

func GetFamilyList(query *bindings.FamilyListQuery) ([]dal.Family, error) {
	// 从 db 中获取 family 列表
	var familyList []dal.Family
	db := DB.SelectDB(false)
	// 应用 query 参数
	db.Table("family_tab").Where(
		"name LIKE ?", "%"+query.Name+"%",
	).Offset((query.Page - 1) * query.Size).Limit(query.Size).Order("id DESC").Find(&familyList)
	return familyList, db.Error
}
