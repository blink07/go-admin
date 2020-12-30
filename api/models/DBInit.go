package models

import (
	"fmt"
	"go-admin/conf/settings"
	"log"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

// 启动时连接上数据库
func SetUp()  {
	log.Println("DB init ....")

	var err error

	db, err = gorm.Open(settings.DataBaseSettings.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DataBaseSettings.User, settings.DataBaseSettings.Password, settings.DataBaseSettings.Host, settings.DataBaseSettings.Name))

	if err!= nil{
		log.Fatalf("model.Setup err:%v", err)
	}

	// 表处理
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return settings.DataBaseSettings.TablePrefix + defaultTableName
	}
	db.SingularTable(true)

	// 打印数据库操作日志
	db.LogMode(true)

	//设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(10)
	//设置到数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(100)


	// 给model注册回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

}


func CloseDB()  {
	defer db.Close()
}

// 定义钩子方法，对不同类中相同字段进行统一处理
// 对创建时间（CreatedOn和ModifiedOn进行统一处理）
func updateTimeStampForCreateCallback(scope *gorm.Scope)  {

	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank{
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField,ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope)  {
	if !scope.HasError(){
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok{
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeleteOnField := scope.FieldByName("DeletedOn")
		if scope.Search.Unscoped && hasDeleteOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExTraSpaceIfExist(scope.CombinedConditionSql()),
				addExTraSpaceIfExist(extraOption),
			)).Exec()
		}else{
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExTraSpaceIfExist(scope.CombinedConditionSql()),
				addExTraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}


func addExTraSpaceIfExist(str string) string{
	if str !="" {
		return " "+ str
	}
	return ""
}