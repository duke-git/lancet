package typemap

import (
	"fmt"
	"reflect"
	"strings"
)

var mapUtils = map[reflect.Kind]func(reflect.Value, reflect.Value) error{
	reflect.String:     convertNormal,
	reflect.Int:        convertNormal,
	reflect.Int16:      convertNormal,
	reflect.Int32:      convertNormal,
	reflect.Int64:      convertNormal,
	reflect.Uint:       convertNormal,
	reflect.Uint16:     convertNormal,
	reflect.Uint32:     convertNormal,
	reflect.Uint64:     convertNormal,
	reflect.Float32:    convertNormal,
	reflect.Float64:    convertNormal,
	reflect.Uint8:      convertNormal,
	reflect.Int8:       convertNormal,
	reflect.Struct:     convertNormal,
	reflect.Complex64:  convertNormal,
	reflect.Complex128: convertNormal,
}

var _ = func() struct{} {
	mapUtils[reflect.Map] = convertMap
	mapUtils[reflect.Array] = convertSlice
	mapUtils[reflect.Slice] = convertSlice
	return struct{}{}
}()

// MapTo try to map any interface to struct or base type
/*
	Eg:
		v := map[string]interface{}{
			"service":map[string]interface{}{
				"ip":"127.0.0.1",
				"port":1234,
			},
			version:"v1.0.01"
		}

		type Target struct {
			Service struct {
				Ip string `json:"ip"`
				Port int `json:"port"`
			} `json:"service"`
			Ver string `json:"version"`
		}

		var dist Target
		if err := typemap.MapTo(v,&dist); err != nil {
			log.Println(err)
			return
		}

		log.Println(dist)

*/
func MapTo(src any, dst any) error {

	dstRef := reflect.ValueOf(dst)
	if dstRef.Kind() != reflect.Ptr {
		return fmt.Errorf("dst is not ptr")
	}

	dstRef = reflect.Indirect(dstRef)

	srcRef := reflect.ValueOf(src)
	if srcRef.Kind() == reflect.Ptr || srcRef.Kind() == reflect.Interface {
		srcRef = srcRef.Elem()
	}
	if f, ok := mapUtils[srcRef.Kind()]; ok {
		return f(srcRef, dstRef)
	}

	return fmt.Errorf("no implemented:%s", srcRef.Type())
}

func convertNormal(src reflect.Value, dst reflect.Value) error {
	if dst.CanSet() {
		if src.Type() == dst.Type() {
			dst.Set(src)
		} else if src.CanConvert(dst.Type()) {
			dst.Set(src.Convert(dst.Type()))
		} else {
			return fmt.Errorf("can not convert:%s:%s", src.Type().String(), dst.Type().String())
		}
	}
	return nil
}

func convertSlice(src reflect.Value, dst reflect.Value) error {
	if dst.Kind() != reflect.Array && dst.Kind() != reflect.Slice {
		return fmt.Errorf("error type:%s", dst.Type().String())
	}
	l := src.Len()
	target := reflect.MakeSlice(dst.Type(), l, l)
	if dst.CanSet() {
		dst.Set(target)
	}
	for i := 0; i < l; i++ {
		srcValue := src.Index(i)
		if srcValue.Kind() == reflect.Ptr || srcValue.Kind() == reflect.Interface {
			srcValue = srcValue.Elem()
		}
		if f, ok := mapUtils[srcValue.Kind()]; ok {
			err := f(srcValue, dst.Index(i))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func convertMap(src reflect.Value, dst reflect.Value) error {
	if src.Kind() != reflect.Map || dst.Kind() != reflect.Struct {
		if src.Kind() == reflect.Interface {
			return convertMap(src.Elem(), dst)
		} else {
			return fmt.Errorf("src or dst type error,%s,%s", src.Type().String(), dst.Type().String())
		}
	}
	dstType := dst.Type()
	num := dstType.NumField()
	exist := map[string]int{}
	for i := 0; i < num; i++ {
		k := dstType.Field(i).Tag.Get("json")
		if k == "" {
			k = dstType.Field(i).Name
		}
		if strings.Contains(k, ",") {
			taglist := strings.Split(k, ",")
			if taglist[0] == "" {

				k = dstType.Field(i).Name
			} else {
				k = taglist[0]

			}

		}
		exist[k] = i
	}

	keys := src.MapKeys()
	for _, key := range keys {
		if index, ok := exist[key.String()]; ok {
			v := dst.Field(index)
			if v.Kind() == reflect.Struct {
				err := convertMap(src.MapIndex(key), v)
				if err != nil {
					return err
				}
			} else {
				if v.CanSet() {
					if v.Type() == src.MapIndex(key).Elem().Type() {
						v.Set(src.MapIndex(key).Elem())
					} else if src.MapIndex(key).Elem().CanConvert(v.Type()) {
						v.Set(src.MapIndex(key).Elem().Convert(v.Type()))
					} else if f, ok := mapUtils[src.MapIndex(key).Elem().Kind()]; ok && f != nil {
						err := f(src.MapIndex(key).Elem(), v)
						if err != nil {
							return err
						}
					} else {
						return fmt.Errorf("error type:d(%s)s(%s)", v.Type(), src.Type())
					}
				}
			}
		}
	}

	return nil
}
