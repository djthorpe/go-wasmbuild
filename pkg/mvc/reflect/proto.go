package reflect

import "reflect"

type Proto struct {
	typ reflect.Type
}

func NewProto(proto any) *Proto {
	t := reflect.TypeOf(proto)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	return &Proto{typ: t}
}
