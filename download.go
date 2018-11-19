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
	ConfigFile    = "conf/config.json"
)

func (c *Client) Download(path, file string) (err error){
	path = strings.TrimSpace(path)
	file = strings.TrimSpace(file)

	if file == "" {
		fmt.Println("[ERR] File name should be specified. ")
		return errors.Errorf("File name should be specified. ")
	}

	req, err := ufsdk.NewFileRequest(c.Conf, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var fileName string

	//
	if path != "" && !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	fileName = path + file
	_, err = os.Stat(fileName)
	if err == nil {
		count := 1
		for {
			fileName = fmt.Sprintf("%s_%d", path + file, count)
			_, err := os.Stat(fileName)
			if err != nil {
				break
			} else {
				count++
			}
		}
		fmt.Printf("[WARN] %s has exist!! Rename %s. \n", path + file, fileName)
	}

	// 去掉路径
	tmp := strings.Split(file, "/")
	remoteFileKey := c.Namespace + "@" + tmp[len(tmp) - 1]

	log.Println("正在下载文件。。。。")
	ftmp, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println("创建文件失败，错误信息为：", err.Error())
		return err
	}

	err = req.DownloadFile(ftmp, remoteFileKey)
	if err != nil {
		log.Println("下载文件出错，出错信息为：", err.Error())
	}
	ftmp.Close() //提前关闭文件，防止 etag 计算不准。

	etagCheck := req.CompareFileEtag(remoteFileKey, fileName)
	if !etagCheck {
		log.Println("文件下载出错，etag 比对不一致。")
		// 错误，删除文件
		_ = os.Remove(fileName)
	} else {
		log.Println("文件下载成功")
	}

	return nil
}