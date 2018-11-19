package filemgr

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"

	"github.com/ufilesdk-dev/ufile-gosdk"
)

const (
	putUpload = iota
	postUpload
	mput
	asyncmput

)

func (c *Client) Upload(file string) (err error) {
	file = strings.TrimSpace(file)

	if file == "" {
		fmt.Println("[ERR] File name should be specified. ")
		return errors.Errorf("[ERR] File name should be specified. ")
	}

	req, err := ufsdk.NewFileRequest(c.Conf, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Printf("[ERR] %s not exist. \n", file)
		return err
	}

	tmp := strings.Split(file, "/")
	fileName := tmp[len(tmp) - 1]
	uploadKey := c.Namespace + "@" + fileName

	var uploadType int

	// 大于100M
	if fileInfo.Size() > 104857600 {
		uploadType = 2
	}
	scheduleUpload(file, uploadKey, uploadType, req)

	return nil
}

func scheduleUpload(filePath, keyName string, uploadType int, req *ufsdk.UFileRequest) {
	log.Println("上传的文件 Key 为：", keyName)
	var err error
	switch uploadType {
	case putUpload:
		// 文件小于100M
		log.Println("正在使用PUT接口上传文件...")
		err = req.PutFile(filePath, keyName, "")
		break
	case postUpload:
		// 文件小于100M
		log.Println("正在使用 POST 接口上传文件...")
		err = req.PostFile(filePath, keyName, "")
	case mput:
		// 文件大于100M
		log.Println("正在使用同步分片上传接口上传文件...")
		err = req.MPut(filePath, keyName, "")
	case asyncmput:
		// 文件大于100M
		log.Println("正在使用异步分片上传接口上传文件...")
		err = req.AsyncMPut(filePath, keyName, "")
	}
	if err != nil {
		log.Println("文件上传失败!!，错误信息为：", err.Error())
		//如果 err 给出的提示信息不够，你可 dump 整个 response 出来查看 http 的返回。
		log.Printf("%s\n", req.DumpResponse(true))
		return
	}
	log.Println("文件上传成功!!")
}

