package tag

import "reflect"

func Parse(reflect.StructTag) GarTag {
	var garTag GarTag
	return garTag
}

type GarTag map[string]string

func (t GarTag) Get(key string) string {
	return t[key]
}
