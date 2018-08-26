package gostr

import (
	"bytes"
	"fmt"
	"io"
	"sort"

	"reflect"
)

// Stringify was heavily inspired by the google/go-github & goprotobuf library.

// Stringify attempts to create a reasonable string representation of types.
// It does things like resolve pointers to their values and omits struct fields
// with nil values.
func Stringify(message interface{}) string {
	var buf bytes.Buffer
	stringifyValue(&buf, reflect.ValueOf(message))
	return buf.String()
}

func stringifyValue(w io.Writer, v reflect.Value) {

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if v.IsNil() {
			w.Write([]byte("<nil>"))
			return
		}
	}

	switch v.Kind() {
	case reflect.Bool:
		fmt.Fprintf(w, "%t", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(w, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(w, "%f", v.Float())

	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)

	case reflect.Ptr, reflect.Interface:
		w.Write([]byte{'*'})
		stringifyValue(w, v.Elem())

	case reflect.Array, reflect.Slice:
		w.Write([]byte{'['})

		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte(", "))
			}
			stringifyValue(w, v.Index(i))
		}

		w.Write([]byte{']'})

	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}

		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if nullable(fv) && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			stringifyValue(w, fv)
		}

		w.Write([]byte{'}'})

	case reflect.Map:
		w.Write([]byte("map["))

		keys := v.MapKeys()
		if len(keys) != 0 {

			sort.SliceStable(keys, lessFunc(keys))
			for i, k := range keys {
				if i > 0 {
					w.Write([]byte(", "))
				}
				switch k.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					fmt.Fprintf(w, "%d:", k.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					fmt.Fprintf(w, "%d:", k.Uint())
				case reflect.String:
					fmt.Fprintf(w, `"%s":`, k)
				}
				stringifyValue(w, v.MapIndex(k))
			}
		}

		w.Write([]byte{']'})

	default:
		panic(fmt.Sprintf("unsupported data type: %v", v.Kind()))
	}
}

func lessFunc(vs []reflect.Value) func(i, j int) bool {

	if len(vs) == 0 {
		return nil
	}

	switch vs[0].Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return func(i, j int) bool { return vs[i].Int() < vs[j].Int() }
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return func(i, j int) bool { return vs[i].Uint() < vs[j].Uint() }
	case reflect.String:
		return func(i, j int) bool { return vs[i].String() < vs[j].String() }
	default:
		panic(fmt.Sprintf("unsupported map key type: %v", vs[0].Kind()))
	}
}

func nullable(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return true
	}
	return false
}
