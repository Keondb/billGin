package main

import (
	"bill/bill_serve/models/res"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const file = "models/res/error_code.json"

//const file = "models/res/error_code.json"

type ErrMap map[res.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error("111111111111")
		logrus.Error(err)
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error("3333333333333333")
		logrus.Error(err)
	}
	fmt.Println(errMap)
	fmt.Println(errMap[res.SettingsError])
}
