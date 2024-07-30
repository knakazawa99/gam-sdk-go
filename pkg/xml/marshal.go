package xml

import (
	"fmt"
	"reflect"
	"strings"
)

func DeepMarshal(v interface{}, ignoreNil bool) ([]byte, error) {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return deepMarshal(value, ignoreNil)
}

func deepMarshal(v reflect.Value, ignoreNil bool) ([]byte, error) {
	var builder strings.Builder
	if v.Kind() != reflect.Struct {
		builder.Write([]byte(fmt.Sprintf("%v", v)))
		return []byte(builder.String()), nil
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)
		// If field is a pointer, get the element it points to
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				if ignoreNil {
					continue
				}
				value = reflect.Zero(field.Type.Elem())
			} else {
				value = value.Elem()
			}
		}
		// Retrieve the XML tag from the field
		xmlTag := field.Tag.Get("xml")

		// eliminate the omitempty tag
		tags := strings.Split(xmlTag, ",")
		xmlTag = tags[0]
		omitEmpty := len(tags) > 1 && tags[1] == "omitempty"
		// Add this code to check if the field should be omitted
		if omitEmpty {
			zeroVal := reflect.Zero(value.Type()).Interface()
			if value.Type().Comparable() && value.Interface() == zeroVal {
				continue
			}
		}

		// If field is a struct or a pointer to a struct, marshal it recursively
		if value.Kind() == reflect.Struct || (value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct) {
			fieldBytes, err := deepMarshal(value, ignoreNil)
			if err != nil {
				return nil, err
			}
			// If fieldBytes is not empty, generate XML tags
			if len(fieldBytes) != 0 {
				if xmlTag != "" {
					builder.WriteString("<" + xmlTag + ">")
				}
				builder.Write(fieldBytes)
				if xmlTag != "" {
					builder.WriteString("</" + xmlTag + ">")
				}
			}
		} else if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
			for j := 0; j < value.Len(); j++ {
				elem := value.Index(j)
				if elem.Kind() == reflect.Ptr && elem.IsNil() && ignoreNil {
					continue
				}
				if elem.Kind() == reflect.Ptr {
					elem = elem.Elem()
				}
				fieldBytes, err := deepMarshal(elem, ignoreNil)
				if err != nil {
					return nil, err
				}
				// If fieldBytes is not empty, generate XML tags
				if len(fieldBytes) != 0 {
					if xmlTag != "" {
						builder.WriteString("<" + xmlTag + ">")
					}
					builder.Write(fieldBytes)
					if xmlTag != "" {
						builder.WriteString("</" + xmlTag + ">")
					}
				}
			}
		} else {
			builder.WriteString("<" + xmlTag + ">")
			valueStr := fmt.Sprintf("%v", value)
			builder.WriteString(valueStr)
			builder.WriteString("</" + xmlTag + ">")
		}
	}
	return []byte(builder.String()), nil
}
