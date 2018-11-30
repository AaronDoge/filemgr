package filemgr

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ufilesdk-dev/ufile-gosdk"
)

const (
    base_format = "Jan 2 15:04:05 2006"
    std_format = "2006-01-02 15:04:05"
)

type FileInfo struct {
	FName 	string
	FSize 	int
	MTime 	string
	NSpace 	string
}

//获取指定命名空间的文件名（或者详细）列表
func (c *Client) List(files ...string) (flist []FileInfo, err error) {
	file := ""
	if len(files) != 0 {
		file = files[0]
	}
	file = strings.TrimSpace(file)
	//if len(args) == 0 {}

	req, err := ufsdk.NewFileRequest(c.Conf, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fileName := c.Namespace + "@" + file

	log.Println("正在获取文件列表...")
	list, err := req.PrefixFileList(fileName, "", 1000)
	if err != nil {
		log.Println("获取文件列表失败，错误信息为：", err.Error())
		return nil, err
	}

	var count int
	for _, v := range list.DataSet {
		fi := FileInfo{}
		tmp := strings.Split(v.FileName, "@")
		ns := tmp[0]
		if !strings.HasPrefix(v.FileName, fileName) {
			continue
		}
		fi.NSpace = ns
		fi.FName = tmp[len(tmp)-1]
		fi.FSize = v.Size
		fi.MTime =  time.Unix(int64(v.ModifyTime), 0).Format(base_format)
		fmt.Printf("%s	%d	%s	%s\n", ns, fi.FSize, fi.MTime, fi.FName)

		flist = append(flist, fi)
		count++
	}
	fmt.Printf("total %d \n", count)

	// log.Printf("获取文件列表返回的信息是：\n%s\n", flist)

	return flist,nil
}


//获取指定命名空间的文件名（或者详细）列表
func (c *Client) ListAll() (flist []FileInfo, err error) {

	req, err := ufsdk.NewFileRequest(c.Conf, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	log.Println("正在获取文件列表...")
	list, err := req.PrefixFileList("", "", 1000)
	if err != nil {
		log.Println("获取文件列表失败，错误信息为：", err.Error())
		return nil, err
	}

	var count int
	for _, v := range list.DataSet {
		fi := FileInfo{}
		tmp := strings.Split(v.FileName, "@")
		ns := tmp[0]

		fi.NSpace = ns
		fi.FName = tmp[len(tmp)-1]
		fi.FSize = v.Size
		MTimeOS :=  time.Unix(int64(v.ModifyTime), 0).Format(base_format)
		fi.MTime =  time.Unix(int64(v.ModifyTime), 0).Format(std_format)
		fmt.Printf("%s	%d	%s	%s\n", ns, fi.FSize, MTimeOS, fi.FName)

		flist = append(flist, fi)
		count++
	}
	fmt.Printf("total %d \n", count)

	// log.Printf("获取文件列表返回的信息是：\n%s\n", flist)

	return flist,nil
}
