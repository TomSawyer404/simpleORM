package ydorm

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func genInsertSQL(entity interface{}) (string, []interface{}) {
	tbInfo, err := parseEntity(entity)
	if err != nil {
		log.Println(err)
	}
	strSQL := "insert into " + tbInfo.Name
	strFields := ""
	strValues := ""
	var params []interface{}

	for _, v := range tbInfo.Fields {
		strFields += v.Name + ","
		if v.refValue.CanInterface() {
			params = append(params, v.refValue.Interface())
			refValue := v.refValue
			if refValue.Kind() == reflect.Int {
				strValues += strconv.Itoa(int(v.refValue.Int())) + ","
			} else {
				strValues += v.refValue.String() + ","
			}
		} else {
			strValues += "?,"
		}
	}

	strFields = strings.TrimRight(strFields, ",")
	strValues = strings.TrimRight(strValues, ",")
	strSQL += " (" + strFields + ") values(" + strValues + ") "
	return strSQL, params
}

func genUpdateSQL(entity interface{}) (string, []interface{}) {
	tbInfo, err := parseEntity(entity)
	if err != nil {
		log.Println(err)
	}
	strSQL := "update " + tbInfo.Name + " set "
	var params []interface{}

	valuesMap := make(map[string]string)
	for _, v := range tbInfo.Fields {
		switch v.Name {
		case "id":
			valuesMap[v.Name] = v.refValue.String()
		case "cname":
			valuesMap[v.Name] = v.refValue.String()
		case "age":
			valuesMap[v.Name] = strconv.Itoa(int(v.refValue.Int()))
		case "Sex":
			valuesMap[v.Name] = v.refValue.String()
		}

		if v.refValue.CanInterface() {
			params = append(params, v.refValue.Interface())
		}
	}

	for k, v := range valuesMap {
		if k != "id" {
			strSQL += k + "=" + v + ","
		}
	}
	strSQL = strings.TrimRight(strSQL, ",")
	strSQL += " where id=" + valuesMap["id"]

	return strSQL, params
}

func genDeleteSQL(entity interface{}) (string, []interface{}) {
	tbInfo, err := parseEntity(entity)
	if err != nil {
		log.Println(err)
	}
	strSQL := "delete from " + tbInfo.Name + " where "
	var params []interface{}

	valuesMap := make(map[string]string)
	for _, v := range tbInfo.Fields {
		switch v.Name {
		case "id":
			valuesMap[v.Name] = v.refValue.String()
		case "cname":
			valuesMap[v.Name] = v.refValue.String()
		case "age":
			valuesMap[v.Name] = strconv.Itoa(int(v.refValue.Int()))
		case "Sex":
			valuesMap[v.Name] = v.refValue.String()
		}

		if v.refValue.CanInterface() {
			params = append(params, v.refValue.Interface())
		}
	}

	for k, v := range valuesMap {
		strSQL += k + "=" + v
		break
	}

	return strSQL, params
}
