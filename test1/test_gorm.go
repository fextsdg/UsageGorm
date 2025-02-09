package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "testing"
)

// 创建mysql数据库链接
func getMysqlConnect() *gorm.DB {
	//确保有数据库gorm
	url := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		//使用了 fmt.Sprintf 来格式化错误信息，使其成为一个单一的字符串传递给 panic 函数
		panic(fmt.Sprintf("连接数据库失败，错误为：%v\n", err))
	}
	return db
}

type UserInfo struct {
	ID       uint   `gorm:"primaryKey" `
	UserName string `gorm:"size:100"`
	PassWord string `gorm:"size:100"`
}

func (ui UserInfo) TableName() string {
	return "userinfo"
}

// 自动迁移测试，如果表不存在会自动创建表
// 创建的表名会默认带“s”,user_infos
// 如果不想要它建表时加个s，可以在相应结构体 实现 TableName() 方法来覆盖默认实现

/*
*
迁移的具体步骤：
检查表是否存在：GORM 会查询数据库，检查 userinfo 表是否存在。
比较表结构：如果表存在，GORM 会比较表结构与 UserInfo 结构体的定义。
执行迁移：
如果表不存在，GORM 会创建新表。
如果表结构不同，GORM 会根据需要添加或修改字段。
*/
func testMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&UserInfo{}) //
	if err != nil {
		fmt.Printf("自动迁移失败，%v", err)
		return
	}
	fmt.Printf("迁移成功！")
}

// 测试插入数据
func testCreate(db *gorm.DB) {
	fmt.Println("-------------This is Insert!--------------")
	//1.插入一条数据
	userinfo := UserInfo{UserName: "三三三", PassWord: "sss123"}
	result := db.Create(&userinfo)
	fmt.Printf("自动创建的ID:%d,err:%v,影响的行数:%d\n", userinfo.ID, result.Error, result.RowsAffected)
	//2.插入多条数据

	userinfos := []UserInfo{
		{UserName: "55531", PassWord: "55566"},
		{UserName: "老刘", PassWord: "666666"},
		{UserName: "老八", PassWord: "秘制小汉堡"},
	}
	results := db.Create(userinfos)
	fmt.Printf("userinfos:%d,err:%v,影响的行数:%d\n", userinfos, results.Error, results.RowsAffected)

}

// 测试查询数据
func testSelect(db *gorm.DB) {
	fmt.Println("-------------This is Select!--------------")
	//1.根据主键查询第一条数据（主键升序）
	var userinfo UserInfo
	db.First(&userinfo)
	fmt.Println(userinfo)
	var userinfo1 UserInfo
	db.First(&userinfo1, 5)
	fmt.Println(userinfo1)
	//2. 查询所有信息
	var userinfos []UserInfo
	db.Find(&userinfos)
	fmt.Println(userinfos)
	//3.条件查询
	var userinfo2 UserInfo
	db.Where("User_Name=?", "老刘").Find(&userinfo2)
	fmt.Println(userinfo2)

}

// 更新测试
func testUpdate(db *gorm.DB) {
	fmt.Println("-------------This is Update!--------------")
	result := db.Model(&UserInfo{}).Where("id=?", 1).Update("user_name", "王zhon")
	fmt.Printf("err:%v,rows:%d\n", result.Error, result.RowsAffected)
	userinfo := UserInfo{}
	db.Where("id=?", 1).Find(&userinfo)
	fmt.Println(userinfo)

	// 选择 Struct 的字段（会选中零值的字段）
	result = db.Model(&userinfo).Select("UserName").Updates(UserInfo{UserName: "new_name", PassWord: "234567"})
	fmt.Printf("err:%v,rows:%d\n", result.Error, result.RowsAffected)
	newuserinfo := UserInfo{}
	db.Where("id=?", 1).Find(&newuserinfo)
	fmt.Println(newuserinfo)
}

// 测试删除
func testDelete(db *gorm.DB) {
	fmt.Println("-------------This is delete!--------------")
	//1.根据主键删除
	result := db.Delete(UserInfo{}, 5)
	fmt.Printf("err:%v,rows:%d\n", result.Error, result.RowsAffected)

	//2.条件+批量删除
	result = db.Where("Pass_Word LIKE ?", "%ss%").Delete(&UserInfo{})
	fmt.Printf("err:%v,rows:%d\n", result.Error, result.RowsAffected)
	fmt.Printf("err:%v,rows:%d\n", result.Error, result.RowsAffected)
	var userinfos []UserInfo
	db.Find(&userinfos)
	fmt.Println("删除后，", userinfos)
}
func main() {
	db := getMysqlConnect()
	testMigrate(db)
	testCreate(db)
	testSelect(db)
	testUpdate(db)
	testDelete(db)
}
