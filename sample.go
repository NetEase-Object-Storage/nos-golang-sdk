package main

import (
	"fmt"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/config"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/logger"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/model"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/nosclient"
)

func init() *nosclient.NosClient {
	conf := &config.Config{
		Endpoint:  "<endpoint>",
		AccessKey: "<AccessKeyId>",
		SecretKey: "<SecretKey>",

		NosServiceConnectTimeout:    3,
		NosServiceReadWriteTimeout:  60,
		NosServiceMaxIdleConnection: 100,

		LogLevel: logger.LogLevel(logger.DEBUG),
		Logger:   logger.NewDefaultLogger(),
	}

	nosClient, _ := nosclient.New(conf)
	return nosClient
}

func main() {
	path := "<File Path>"
	objectRequest := &model.ObjectRequest{
		Bucket: "<my-bucket>",
		Object: "<my-object>",
	}

	nosClient := init()
	nosClient.PutObjectByFile(path)

	objectResult, err := nosClient.GetObject(objectRequest)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		objectResult.Body.Close()
	}

	err = nosClient.DeleteObject(objectRequest)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Simple samples completed")
}
