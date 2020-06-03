package util_test

import (
	"blog-for-go/util"
	"fmt"
	"testing"
)

type TestCls struct {
	Name string
	Age int
}

func TestPrettyJSON(t *testing.T) {
	var testCls TestCls
	testCls.Name = "jt_hu"
	testCls.Age = 30

	fmt.Printf("%s\n", util.PrettyJSON(testCls))

}
