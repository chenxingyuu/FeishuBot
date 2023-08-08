package services

import (
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/global"
	"github.com/tietiexx/bot/code/backend/models"
	"gorm.io/gorm"
	"sync"
)

func LarkAppById(lakeBotId int) (lakeApp *database.LarkApp, err error) {
	err = global.MySQLClient.First(&lakeApp, lakeBotId).Error
	return
}

func LarkAppByName(name string) (lakeApp *database.LarkApp, err error) {
	err = global.MySQLClient.Where("name = ?", name).First(&lakeApp).Error
	return
}

func LarkAppByUUID(uuid string) (lakeApp *database.LarkApp, err error) {
	err = global.MySQLClient.Where("uuid = ?", uuid).First(&lakeApp).Error
	return
}

func LarkAppCreate(uuid, name string) (err error) {
	var lakeApp = database.LarkApp{UUID: uuid, Name: name}
	err = global.MySQLClient.Create(&lakeApp).Error
	return
}

func LarkAppListQuery(params models.LarkAppListRequest, count bool) (query *gorm.DB) {
	query = global.MySQLClient.Model(&database.LarkApp{})
	return
}

func LarkAppList(params models.LarkAppListRequest) (list []*database.LarkApp, err error) {
	query := LarkAppListQuery(params, false)
	err = query.Find(&list).Error
	return
}

func LarkAppCount(params models.LarkAppListRequest) (count int64, err error) {
	query := LarkAppListQuery(params, true)
	if err = query.Count(&count).Error; err != nil {
		count = 0
		return
	}
	return
}

func LarkAppListPaginate(params models.LarkAppListRequest) (count int64, list []*database.LarkApp) {
	// 使用 WaitGroup 等待异步任务完成
	var wg sync.WaitGroup
	wg.Add(2) // 设置需要等待的异步任务数量

	// 异步执行 UserList 查询
	go func() {
		defer wg.Done() // 标记异步任务完成
		list, _ = LarkAppList(params)
	}()

	// 异步执行 UserCount 查询
	go func() {
		defer wg.Done() // 标记异步任务完成
		count, _ = LarkAppCount(params)
	}()

	// 等待所有异步任务完成
	wg.Wait()

	return
}
