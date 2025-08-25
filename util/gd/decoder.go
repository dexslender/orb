package gd

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)


func structToURLValues(item any) url.Values {
	res := url.Values{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := range v.NumField() {
		tag := v.Field(i).Tag.Get("query")
		field := reflectValue.Field(i).String()
		if tag != "" && tag != "-" {
			res.Add(tag, field)
		}
	}
	return res
}

func DecodeGDUserData(data []byte) (map[int]string, error) {
	spl := strings.Split(string(data), "#")
	user_RobFormat := strings.Split(spl[0], ":")
	user := make(map[int]string)
	for i, d := range user_RobFormat {
		if i&1 == 0 {
			k, err := strconv.Atoi(d)
			if err != nil {
				return nil, err
			}
			user[k] = user_RobFormat[i+1]
		}
	}
	return user, nil
}

func DecodeGDData[T any](data string) *T {
	spl := strings.Split(data, "#")
	RobFormat := strings.Split(spl[0], ":")
	if len(RobFormat) <= 1 { return nil }

	gd := make(map[string]string)
	for i, key := range RobFormat {
		if i&1 == 0 {
			gd[key] = RobFormat[i+1]
		}
	}

	var target T
	tS := reflect.TypeOf(target)
	if tS.Kind() != reflect.Struct {
		return nil
	}
	vS := reflect.ValueOf(&target)
	updateStruct(gd, tS, vS.Elem())
	return &target
}

func updateStruct(data map[string]string, t reflect.Type, v reflect.Value) {
	for i := range t.NumField() {
		if tag, ok := t.Field(i).Tag.Lookup("prop"); ok {
			field := v.Field(i)
			if field.CanSet() {
				robvalue := data[tag]
				switch field.Kind() {
				case reflect.String:
					field.SetString(robvalue)
				case reflect.Int:
					vi, _ := strconv.Atoi(robvalue)
					field.Set(reflect.ValueOf(vi))
				case reflect.Bool:
					vb, _ := strconv.ParseBool(robvalue)
					field.SetBool(vb)
				case reflect.Struct:
					tfield := field.Type()
					updateStruct(data, tfield, field)
				}
			}
		}
	}
}
