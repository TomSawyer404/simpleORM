package ydorm

import "reflect"

type FieldInfo struct {
	Name         string
	IsPrimaryKey bool
	refValue     reflect.Value
}
