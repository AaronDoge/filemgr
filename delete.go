package filemgr

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.com/ufilesdk-dev/ufile-gosdk"
)

// 删除ufile文件

func (c *Client) Delete(file string) (err error) {
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

	err = req.DeleteFile(c.Namespace + "@" + file)
	if err != nil {
		fmt.Printf("删除文件%s失败，失败原因：%s \n", c.Namespace + "@" + file, err.Error())
		return err
	} else {
		fmt.Printf("删除文件%s成功。。。\n", c.Namespace + "@" + file)
	}

	return nil
}