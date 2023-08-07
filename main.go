package main

import (
	"github.com/TommyLemon/unitauto-go/unitauto"
	"unitauto.demo/unitauto/demo"
)

func Test() string {
	return "OK"
}

func main() {
	unitauto.DEFAULT_MODULE_PATH = "unitauto.demo" // TODO your package

	/*
		main.Test
		{
			"static": true // methodName = "Test"
		}
	*/
	unitauto.CLASS_MAP["main.Test"] = Test

	/*
		demo.Add.Add
		{
			"methodArgs": [
				"int:1",
				"int:2"
			]
		}
	*/
	unitauto.CLASS_MAP["demo.Add"] = demo.Add

	/*
		Test.SetName // unitauto.DEFAULT_MODULE_PATH = "unitauto.demo"
		{
			"this": {
				"type": "*unitauto.test.Test",
				"value": {
					"Id": 2,
					"Name": "hello"
				}
			},
			"methodArgs": [
				"UnitAuto@Go"
			]
		}
	*/
	unitauto.CLASS_MAP["unitauto.demo.Test"] = demo.Test{}
	unitauto.CLASS_MAP["*unitauto.demo.Test"] = &demo.Test{}

	println("test with http://apijson.cn/unit/go or Postman")
	unitauto.Start(8082)
}
