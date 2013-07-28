package debug

import(
"fmt"
"reflect"
"strings"
)

const(
INDENT_WIDTH = 2
)

func Pprint(obj interface{}) string {
  return recurPretty(obj, 1)
}

func recurPretty(obj interface{}, depth int) string {
  if obj == nil {
    return ""
  }
  refObj := reflect.ValueOf(obj)
  typeObj := refObj.Type()

  spaces := strings.Repeat(" ", depth * INDENT_WIDTH)
  openBraces := fmt.Sprintf("%s {\n", typeObj)
  innerContent := ""
  for i:=0; i<refObj.NumField(); i++ {
    f := refObj.Field(i)
    switch f.Kind() {
      case reflect.Struct:
        innerContent += spaces + recurPretty(f.Interface(), depth+1)
      default:
        innerContent += spaces + fmt.Sprintf("%s: %v\n", typeObj.Field(i).Name, f.Int())
    }
  }
  closingBraces := fmt.Sprintln("}\n")
  return openBraces + innerContent + closingBraces
}

