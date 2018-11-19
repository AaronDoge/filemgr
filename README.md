# Instructions
> 对`ufile-gosdk`进行简单封装。

## 文件操作

### NewClient
```go
func NewClient(namespace, key string, regions ...string) (*Client, error)
```
- namespace：命名空间，需要提前注册；
- key：命名空间的token；
- region：区域，这个参数不传或传入空字符串，会使用默认值。

### Upload
```go
func (c *Client) Upload(file string) (err error)
```
- file：待上传的本地文件名。

### Download
```go
func (c *Client) Download(path, file string) (err error)
```
- path：指定下载目录；
- file：要下载的远程文件名。

### List
```go
func (c *Client) List(file ...string) (flist []FileInfo, err error)
```
- file：需要获取文件的文件名或文件前缀，没有参数表示获取全部；
- flist：获取的文件信息列表。
    - FName 	string  文件名
    - FSize 	int     文件大小
    - MTime 	string  修改时间
    - NSpace 	string  所属命名空间

### Delete
```go
func (c *Client) Delete(file string) (err error)
```
- file：待删除的远程文件名。

## 示例代码
```go
import "fielmgr"

// ...

    client, _ := filemgr.NewClient(namespace, token, region)

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
// ...

```