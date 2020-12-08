package gnet

import (
	"fmt"
	"testing"
)

func TestReadConfigFromCSV(t *testing.T) {

	err, data := ReadConfigFromCSV("e:\\heroexp.csv", 4, 1, &GenConfigObj{
		GenObjFun:   nil,
		ParseObjFun: nil,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}