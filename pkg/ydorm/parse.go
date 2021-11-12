package ydorm

import (
	"reflect"
	"strings"
)

// If struct has no Tag, then we assume name of struct field and name of table field is the same;
// If struct has a Tag without `":"`, then we assume name of tag is the name of table;
func parseEntity(entity interface{}) (tInfo *TableInfo, err error) {
	tInfo = &TableInfo{}
	ret_type := reflect.TypeOf(entity)
	ret_val := reflect.ValueOf(entity)
	if ret_type.Kind() == reflect.Ptr {
		ret_type = ret_type.Elem()
		ret_val = ret_val.Elem()
	}

	for i, j := 0, ret_type.NumField(); i < j; i += 1 {
		retTypeField := ret_type.Field(i) // StructField
		retValField := ret_val.Field(i)   // Value

		var f FieldInfo
		// No Tag, means same between struct field and table field
		if retTypeField.Tag == "" {
			f = FieldInfo{Name: retTypeField.Name, IsPrimaryKey: false, refValue: retValField}
		} else {
			strTag := string(retTypeField.Tag)
			if strings.Index(strTag, ":") == -1 {
				tInfo.Name = strTag
				continue
			} else {
				field := retTypeField.Tag.Get("field")
				isKey := false
				strIsKey := retTypeField.Tag.Get("isKey")
				if strIsKey == "1" {
					isKey = true
				}
				f = FieldInfo{Name: field, IsPrimaryKey: isKey, refValue: retValField}
			}
		}

		tInfo.Fields = append(tInfo.Fields, f)
	}
	return
}
