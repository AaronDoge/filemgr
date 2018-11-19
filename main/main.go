package main

import (
	"filemgr"
	"flag"
	"fmt"
	"os"
)

/**
做一些输入参数解析，有效性合法性检验，判断操作，namespace的key验证
*/

func main() {

	var operation string
	var namespace string
	var token string
	var region string
	//var help string

	flag.StringVar(&namespace, "namespace", "errorName", "--namespace:	set namespace [needed]")
	flag.StringVar(&operation, "operation", "errorOper", "--operation: 	set operation [needed]")
	flag.StringVar(&token, "token", "errToken", "--token: set token [needed]")
	flag.StringVar(&region, "region", "errRegion", "--region: set region [needed]")
	flag.Parse()

	if operation == "errorName" || operation == "" {
		fmt.Println("[ERR] --operation is needed. Use \"--operation=example\" to specify operation.")
		return
	}
	if namespace == "errorName" || namespace == "" {
		fmt.Println("[ERR] --namespace is needed. Use \"--namespace=example\" to specify namespace.")
		return
	}
	if token == "errorName" || token == "" {
		fmt.Println("[ERR] --token is needed. Use \"--token=example\" to specify token. ")
		return
	}
	if region == "errorName" || region == "" {
		fmt.Println("[ERR] --region is needed. Use \"--region=example\" to specify region. ")
		return
	}

	path := ""
	fileName := ""

	fi := os.Args[5:]
	if len(fi) == 2 {
		path = fi[0]
		fileName = fi[1]
	} else if len(fi) == 1 {
		path = ""
		fileName = fi[0]
	} else if (len(fi) == 0 && operation != "list") || len(fi) > 2 {
		fmt.Println("Pls input right director or file name. ")
		return
	}

	//jc, err := filemgr.LoadConf(filemgr.ConfigFile)
	fileHost := os.Getenv("FILE_HOST")
	// TODO
	//fileHost = "http://192.168.152.186:8000/filemgr"
	if fileHost == "" {
		fmt.Println("[ERR] Pls set env FILE_HOST in advance.")
		return
	}

	client, err := filemgr.NewClient(namespace, token, fileHost, region)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch operation {
	case "upload":
		client.Upload(fileName)
	case "download":
		client.Download(path, fileName)
	case "list":
		client.List(fileName)
	case "delete":
		client.Delete(fileName)
	default:
		fmt.Println("[ERR] Pls use correct operation! Such as \"upload, download, list, delete etc\". ")
		return
	}

}
