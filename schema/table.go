package schema

import (
	"reflect"
)

const tagKey = "gar"

func TableOf(m any) Table {
	table := Table{
		t:    reflect.TypeOf(m),
		cols: []Column{},
	}

	for _, f := range reflect.VisibleFields(table.t) {
		if isBaseModel(f) {
			table.bm = newBaseModel(f.Tag.Get(tagKey))
		} else {
			table.cols = append(table.cols, Column{f: f})
		}
	}

	return table
}

type Table struct {
	t  reflect.Type
	bm baseModel

	cols []Column
}

func (table Table) Name() string {
	return table.t.Name()
}
