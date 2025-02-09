package doSomething

import (
	"UsageGorm/test2/utilClass"
	"fmt"
	"gorm.io/gorm"
)

//定义了一些操作表函数

// 创建学生信息
func InsertStudent(db *gorm.DB, student utilClass.Student) error {
	result := db.Select("Name", "Password", "MajorNum").Create(&student)
	if result.Error != nil {
		fmt.Printf("创建失败！err:=%v\n", result.Error)
	} else {
		fmt.Printf("创建成功！影响行数：%v\n", result.RowsAffected)
	}
	return result.Error
}

// 删除学生课程
func DeleteSCById(db *gorm.DB, id int) error {
	result := db.Where("Id=?", id).Delete(&utilClass.StuCourse{})
	if result.Error != nil {
		fmt.Printf("删除失败！err:=%v\n", result.Error)
	} else {
		fmt.Printf("删除成功！影响行数：%v\n", result.RowsAffected)
	}
	return result.Error
}
