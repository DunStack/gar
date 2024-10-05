package schema

import (
	"reflect"
)

type BaseModel struct{}

var baseModelType = reflect.TypeOf(new(BaseModel)).Elem()

func isBaseModel(f reflect.StructField) bool {
	return f.Name == "BaseModel" && f.Type == baseModelType
}
