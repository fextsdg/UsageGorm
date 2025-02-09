package main

import (
	"UsageGorm/test2/dbUtil"
	"UsageGorm/test2/utilClass"
	"fmt"
)

func main() {
	db := dbUtil.OpenConnect()
	err := db.AutoMigrate(utilClass.Course{}, utilClass.Major{}, utilClass.Student{}, utilClass.Teacher{}, utilClass.Tcourse{}, utilClass.StuCourse{})
	if err != nil {
		panic(fmt.Sprintf("自动迁移失败！err:=%v", err))
	}
	fmt.Println("自动迁移成功！")

}
