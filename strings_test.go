package gostr_test

import (
	"testing"

	"github.com/takuoki/gostr"
)

type testStruct1 struct {
	b              bool
	i              int
	i8             int8
	i16            int16
	i32            int32
	i64            int64
	ui             uint
	ui8            uint8
	ui16           uint16
	ui32           uint32
	ui64           uint64
	uiptr          uintptr
	f32            float32
	f64            float64
	str            string
	s              testStruct2
	a              [3]string
	sl, esl        []string
	set            map[string]struct{}
	m, em          map[string]string
	im             map[int]string
	ptr, nptr      *string
	iface1, iface2 interface{}
}

type testStruct2 struct {
	foo string
	bar int
}

var (
	testStr = "test"

	ts = testStruct1{
		b:      true,
		i:      1,
		i8:     2,
		i16:    3,
		i32:    4,
		i64:    5,
		ui:     6,
		ui8:    7,
		ui16:   8,
		ui32:   9,
		ui64:   10,
		uiptr:  11,
		f32:    12.0,
		f64:    13.0,
		str:    "abc",
		s:      testStruct2{foo: "def", bar: 14},
		a:      [3]string{"g", "h", "i"},
		sl:     []string{"j", "k", "l"},
		esl:    []string{},
		set:    map[string]struct{}{"banana": struct{}{}, "cherry": struct{}{}, "apple": struct{}{}},
		m:      map[string]string{"japan": "JPY", "america": "USD", "france": "EUR"},
		em:     map[string]string{},
		im:     map[int]string{3: "x", 2: "y", 1: "z"},
		ptr:    &testStr,
		nptr:   nil,
		iface1: testStruct2{foo: "if1", bar: 15},
		iface2: &testStruct2{foo: "if2", bar: 16},
	}

	expected = `gostr_test.testStruct1{b:true, i:1, i8:2, i16:3, i32:4, i64:5, ui:6, ui8:7, ui16:8, ui32:9, ui64:10, uiptr:11, f32:12.000000, f64:13.000000, str:"abc", s:gostr_test.testStruct2{foo:"def", bar:14}, a:["g", "h", "i"], sl:["j", "k", "l"], esl:[], set:map["apple":{}, "banana":{}, "cherry":{}], m:map["america":"USD", "france":"EUR", "japan":"JPY"], em:map[], im:map[1:"z", 2:"y", 3:"x"], ptr:*"test", iface1:*gostr_test.testStruct2{foo:"if1", bar:15}, iface2:**gostr_test.testStruct2{foo:"if2", bar:16}}`
)

func TestStringify(t *testing.T) {

	// general test
	a := gostr.Stringify(ts)
	if a != expected {
		t.Errorf("result of general test is not match (expected=%s, actual=%s)", expected, a)
	}

	// nil test
	var nstr *string
	a = gostr.Stringify(nstr)
	if a != "<nil>" {
		t.Errorf("result of nil test is not match (expected=<nil>, actual=%s)", a)
	}
}
