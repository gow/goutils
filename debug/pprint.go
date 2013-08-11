package debug

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	INDENT_WIDTH = 2
)

func PP(obj interface{}) string {
	refObj := reflect.ValueOf(obj)
	return stringify(refObj, 1, true)
}

func stringify(refVal reflect.Value, depth int, printType bool) string {
	switch refVal.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%#v", refVal.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return fmt.Sprintf("%#v", refVal.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return fmt.Sprintf("%#v", refVal.Uint())
	case reflect.Ptr:
		return stringifyPointer(refVal)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%#v", refVal.Float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%#v", refVal.Complex())
	case reflect.String:
		return fmt.Sprintf("\"%s\"", refVal.String())
	case reflect.Interface:
		if refVal.IsNil() == true {
			return "interface{} <nil>"
		}
		return fmt.Sprintf("%s{}", refVal.Elem().Type())

	case reflect.Array, reflect.Slice:
		return stringifyList(refVal, depth)

	case reflect.Map:
		return stringifyMap(refVal, depth, printType)
	case reflect.Struct:
		return stringifyStruct(refVal, depth, printType)
	case reflect.Invalid:
		return "<nil>"
	default:
		return "UNKNOWN_VALUE(" + refVal.Type().String() + ")"
	}
	return "<nil>"
}

func stringifyMap(refVal reflect.Value, depth int, printType bool) string {
	spaces := strings.Repeat(" ", depth*INDENT_WIDTH)
	typeStr := refVal.Type().String()
	itemCount := refVal.Len()
	innerContent := ""
	for _, key := range refVal.MapKeys() {
		val := refVal.MapIndex(key)
		keyStr := stringify(key, depth+1, false)
		valStr := stringify(val, depth+1, false)
		innerContent += spaces + keyStr + ": " + valStr + ",\n"
	}
	return "(" + strconv.Itoa(itemCount) + ")" + typeStr + "{\n" + innerContent + spaces + "}"
}

func stringifyPointer(refVal reflect.Value) string {
	ptrStr := "<nil>"
	if refVal.Pointer() != 0 {
		ptrStr = fmt.Sprintf("%#x", refVal.Pointer())
	}
	return ptrStr + "(" + refVal.Type().String() + ")"
}

func stringifyList(refVal reflect.Value, depth int) string {
	var strContent string
	for i := 0; i < refVal.Len(); i++ {
		strContent += stringify(refVal.Index(i), depth, false)
		if i < (refVal.Len() - 1) {
			strContent += ", "
		}
	}
	return refVal.Type().String() + "{" + strContent + "}"
}

func stringifyStruct(refVal reflect.Value, depth int, printType bool) string {
	spaces := strings.Repeat(" ", depth*INDENT_WIDTH)
	openBraces := ""
	if printType {
		openBraces = fmt.Sprintf("%s {", refVal.Type().String())
	} else {
		openBraces = "{"
	}
	innerContent := ""
	for i := 0; i < refVal.NumField(); i++ {
		f := refVal.Field(i)
		fieldStr := fmt.Sprintf(
			"%s: %s",
			refVal.Type().Field(i).Name,
			stringify(f, depth+1, true))
		if i < (refVal.NumField() - 1) {
			fieldStr += ", "
		}
		if depth < 3 {
			innerContent += spaces + fieldStr + "\n"
		} else {
			innerContent += fieldStr
		}
	}
	closingBraces := "}"
	if depth < 3 { // Indent only at the top levels.
		return openBraces + "\n" + innerContent + spaces + closingBraces
	}
	return openBraces + innerContent + closingBraces
}
