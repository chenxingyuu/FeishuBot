package services

import (
	"github.com/tietiexx/bot/code/backend/constant"
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/global"
	"github.com/tietiexx/bot/code/backend/models"
	"gorm.io/gorm"
	"sync"
)

func LarkBotById(lakeBotId int) (lakeBot *database.LarkBot, err error) {
	err = global.MySQLClient.First(&lakeBot, lakeBotId).Error
	return
}

func LarkBotByName(name string) (lakeBot *database.LarkBot, err error) {
	err = global.MySQLClient.Where("name = ?", name).First(&lakeBot).Error
	return
}

func LarkBotByUUID(uuid string) (lakeBot *database.LarkBot, err error) {
	err = global.MySQLClient.Where("uuid = ?", uuid).First(&lakeBot).Error
	return
}

func LarkBotCreate(larkAppID uint, name string, lakeBotType constant.LakeBotType) (err error) {
	var lakeBot = database.LarkBot{Name: name, BotType: lakeBotType, LarkAppID: larkAppID}
	err = global.MySQLClient.Create(&lakeBot).Error
	return
}

func LarkBotListQuery(params models.LarkBotListRequest, count bool) (query *gorm.DB) {
	query = global.MySQLClient.Model(&database.LarkBot{})
	return
}

func LarkBotList(params models.LarkBotListRequest) (list []*database.LarkBot, err error) {
	query := LarkBotListQuery(params, false)
	err = query.Find(&list).Error
	return
}

func LarkBotCount(params models.LarkBotListRequest) (count int64, err error) {
	query := LarkBotListQuery(params, true)
	if err = query.Count(&count).Error; err != nil {
		count = 0
		return
	}
	return
}

func LarkBotListPaginate(params models.LarkBotListRequest) (count int64, list []*database.LarkBot) {
	// 使用 WaitGroup 等待异步任务完成
	var wg sync.WaitGroup
	wg.Add(2) // 设置需要等待的异步任务数量

	// 异步执行 UserList 查询
	go func() {
		defer wg.Done() // 标记异步任务完成
		list, _ = LarkBotList(params)
	}()

	// 异步执行 UserCount 查询
	go func() {
		defer wg.Done() // 标记异步任务完成
		count, _ = LarkBotCount(params)
	}()

	// 等待所有异步任务完成
	wg.Wait()

	return
}
