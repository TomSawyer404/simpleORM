package main

import (
	"04-simpleORM/entity"
	"04-simpleORM/pkg/ydorm"
	"fmt"
)

func main() {
	stu := entity.Student{
		Id:   "001",
		Name: "Jack",
		Age:  25,
		Sex:  "male",
	}

	if isOk, _ := ydorm.Save(&stu); isOk {
		fmt.Println("Insert successfully")
	}
	if isOk, _ := ydorm.Update(&stu); isOk {
		fmt.Println("Update successfully")
	}
	if isOk, _ := ydorm.Delete(&stu); isOk {
		fmt.Println("Delete successfully")
	}
}
