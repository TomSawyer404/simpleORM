package ydorm

import "fmt"

func Save(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("exception catched:", err)
		}
	}()

	strSQL, p := genInsertSQL(entity)
	fmt.Println(strSQL)
	fmt.Println(p)
	isOk = true
	return
}

func Update(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("exception catched:", err)
		}
	}()

	strSQL, p := genUpdateSQL(entity)
	fmt.Println(strSQL)
	fmt.Println(p)
	isOk = true
	return
}

func Delete(entity interface{}) (isOk bool, err error) {
	defer func() {
		if err := recover(); err != nil {
			isOk = false
			fmt.Println("exception catched:", err)
		}
	}()

	strSQL, p := genDeleteSQL(entity)
	fmt.Println(strSQL)
	fmt.Println(p)
	isOk = true
	return
}
