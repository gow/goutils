package debug

import (
	"fmt"
	"testing"
)

type myTestInterface interface {
	dummyFunc1()
	dummyFunc2()
}

type myTestStruct struct {
	dummyInt int
}

func (this myTestStruct) dummyFunc1() {
}
func (this myTestStruct) dummyFunc2() {
}

func TestPPrint(t *testing.T) {
	var someBool bool = true
	var dummyUint32 uint32 = 3232
	var dummyFloat64 float64 = 3.14
	type dummyStruct struct {
		someInt   int
		someFloat float32
	}
	type input struct {
		boolVal      bool
		intVal       int
		int32Val     int32
		uintVal      uint
		uint64Val    uint64
		boolValPtr   *bool
		structValPtr *dummyStruct
		uintValPtr   *uint32
		floatValPtr  *float64
		floatVal     float32
		complexVal   complex128
		arrayVal     [8]int
		sliceVal     []dummyStruct
		aMap         map[float64]string
		anInterface  myTestInterface
		//aSlice []int
		//aBool bool
		//anInt64 int64
	}
	var obj = input{
		false,
		-1,
		32,
		4,
		1234567890,
		&someBool,
		nil,
		&dummyUint32,
		&dummyFloat64,
		3.14,
		complex128(2 + 3i),
		[8]int{0, 1, 2, 3, 4, 5, 6, 7},
		[]dummyStruct{
			dummyStruct{1, 1.1},
			dummyStruct{2, 2.2},
		},
		map[float64]string{3.14: "bar", 2.5: "qqq123"},
		myTestStruct{32},
		//[]int{3, 6, 9},
	}
	res := PP(obj)
	LogMsg(obj)
	fmt.Println(res)
	t.Errorf("unexpected error: %s", res)
}
