package utils

import (
	"reflect"
	"strings"
)

const dotByte = 46
const commaByte = 44

type Changer func(interface{}) interface{}

func MakeMapFromStruct(m interface{}, kChanger map[string]Changer, pkey string, skipEmpty bool) interface{} {
	value := reflect.ValueOf(m)

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return struct{}{}
		}
		value = value.Elem()
	}

	if value.IsZero() {
		return struct{}{}
	}

	result := make(map[string]interface{})
	refType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if !field.CanInterface() || (field.IsZero() && skipEmpty) {
			continue
		}

		typeOfField := refType.Field(i)
		key := typeOfField.Name
		jsonTag := typeOfField.Tag.Get("json")
		if len(jsonTag) > 0 {
			if jsonTag == "-" {
				continue
			}
			key = readUpToByte(jsonTag, commaByte)
		}

		withPkey := key
		if len(pkey) > 0 {
			var sb strings.Builder
			sb.Grow(len(pkey) + len(key) + 1)
			sb.WriteString(pkey)
			sb.WriteByte(dotByte)
			sb.WriteString(key)
			withPkey = sb.String()
		}

		data := field.Interface()
		if field.Kind() == reflect.Ptr || field.Kind() == reflect.Struct {
			data = MakeMapFromStruct(data, kChanger, withPkey, skipEmpty)
		}

		if kM, ok := kChanger[withPkey]; ok {
			result[key] = kM(data)
		} else {
			result[key] = data
		}
	}

	return result
}

func readUpToByte(s string, stopByte byte) string {
	var sb strings.Builder
	i := 0
	for i = range s {
		if s[i] == stopByte {
			i--
			break
		}
	}
	sb.Grow(i + 1)
	for i = range s {
		if s[i] == stopByte {
			break
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}
