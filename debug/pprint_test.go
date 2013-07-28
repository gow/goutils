package debug

import(
"testing"
"fmt"
)

func TestPPrint(t *testing.T) {
  type input struct {
    //anInterface interface{}
    //aMap map[string]string
    //aSlice []int
    anInt int
  }
  var obj = input {
    //map[string]string{"foo": "bar", "qwerty": "qqq123"},
    //[]int{3,6,9},
    32,
  }
  res := Pprint(obj)
  fmt.Println("Result: ", res)
  t.Errorf("unexpected error: %s", res)
}

