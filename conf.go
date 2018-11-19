package filemgr

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Config 配置文件序列化所需的全部字段
type Config struct {
	PublicKey       string
	PrivateKey      string
	BucketName      string
	FileHost        string
	BucketHost      string
	VerifyUploadMD5 bool
}

type JsonConfig struct {
	Host 	string `json:"fileproc_host"`
}

func LoadConf(jsonPath string) (*JsonConfig, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	configBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	jc := new(JsonConfig)
	err = json.Unmarshal(configBytes, jc)

	return jc, nil
}

//LoadConfig 从配置文件获取文件服务地址，并请求通过验证后获取配置信息
//func LoadConfig(jsonPath, namespace, key string) (*ufsdk.Config, error) {
//	file, err := openFile(jsonPath)
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//	configBytes, err := ioutil.ReadAll(file)
//	if err != nil {
//		return nil, err
//	}
//
//	fmt.Printf("namespace: %s, key: %s \n", namespace, key)
//	//
//	jc := new(JsonConfig)
//	err = json.Unmarshal(configBytes, jc)
//
//	bodyJson, err := json.Marshal(ReqParam{Namespace:namespace, Key:key})
//	if err != nil {
//		panic("json marshal error. ")
//	}
//	res, err := http.Post(jc.Host, "application/json", bytes.NewBuffer(bodyJson))
//	if err != nil {
//		fmt.Printf("send request error. url: %s. err: %s\n", jc.Host, err)
//		return nil, err
//	}
//	defer res.Body.Close()
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println("parse body error. ", err)
//		return nil, err
//	}
//
//	mp := make(map[string]interface{})
//	err = json.Unmarshal(body, &mp)
//	if err != nil {
//		fmt.Println("json unmarshal error. ", err.Error())
//		return nil, err
//	}
//
//	c:= new(ufsdk.Config)
//
//	if res.StatusCode != 200 {
//		return nil, errors.Errorf("%s",mp["data"])
//	} else {
//		tmp := mp["data"].(map[string]interface{})
//		c.FileHost = fmt.Sprint(tmp["FileHost"])
//		c.PublicKey = fmt.Sprint(tmp["PublicKey"])
//		c.PrivateKey = fmt.Sprint(tmp["PrivateKey"])
//		c.BucketName = fmt.Sprint(tmp["BucketName"])
//		c.BucketHost = fmt.Sprint(tmp["BucketHost"])
//	}
//
//	return c, nil
//}

