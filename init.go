package filemgr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ufilesdk-dev/ufile-gosdk"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Namespace 	string
	Key 		string
	Region 		string
	FileHost	string
	Conf 		*ufsdk.Config
}

func NewClient(namespace, key, fileHost string, regions ...string) (*Client, error) {
	region := "default"
	if len(regions) != 0 {
		region = regions[0]
	}
	if !strings.HasSuffix(fileHost, "/") {
		fileHost = fileHost + "/"
	}

	c := &Client{
		Namespace:namespace,
		Key:key,
		FileHost:fileHost,
		Region:region,
	}

	err := c.check()
	if err != nil {
		fmt.Println("namespace check error. ", err.Error())
		return nil, err
	}

	return c, nil
}


func (c *Client) check() error {

	// jc, err := LoadConf(configFile)

	bodyJson, err := json.Marshal(struct {
		Namespace 	string `json:"namespace"`
		Key 		string `json:"key"`
		Region 		string `json:"region"`
	}{
		Namespace: c.Namespace,
		Key:c.Key,
		Region: c.Region,
	})
	if err != nil {
		fmt.Println("json marshal error. ", err.Error())
		return err
	}

	res, err := http.Post(c.FileHost + "check", "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		fmt.Printf("send request error. url: %s. err: %s\n", c.FileHost, err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("parse body error. ", err)
		return err
	}

	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	if err != nil {
		fmt.Println("json unmarshal error. ", err.Error())
		return err
	}

	cf := new(ufsdk.Config)
	if res.StatusCode != 200 {
		return errors.Errorf("%s",mp["data"])
	} else {
		tmp := mp["data"].(map[string]interface{})
		cf.FileHost = fmt.Sprint(tmp["FileHost"])
		cf.PublicKey = fmt.Sprint(tmp["PublicKey"])
		cf.PrivateKey = fmt.Sprint(tmp["PrivateKey"])
		cf.BucketName = fmt.Sprint(tmp["BucketName"])
		cf.BucketHost = fmt.Sprint(tmp["BucketHost"])
	}
	c.Conf = cf

	return nil
}