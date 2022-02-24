package main

import (
	"fmt"

	logger "github.com/skyflowapi/skyflow-go/commonutils/logwrapper"
	saUtil "github.com/skyflowapi/skyflow-go/service-account/util"
)

var token = ""

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error : ", err)
		}
	}()
	logger.SetLogLevel(logger.INFO) //set loglevel to INFO
	filePath := "<file_path>"
	if !saUtil.IsValid(token) {
		newToken, err := saUtil.GenerateBearerToken(filePath)
		if err != nil {
			panic(err)
		} else {
			token = newToken.AccessToken
		}
		fmt.Println("%v", token)
	}
}
