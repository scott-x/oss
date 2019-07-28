/*
* @Author: scottxiong
* @Date:   2019-07-29 01:28:21
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-07-29 03:18:00
*/
package oss

import (
	"os"
	"encoding/json"
	"fmt"
   "github.com/aliyun/aliyun-oss-go-sdk/oss"
)
type conf struct{
	Endpoint string `json:"endpoint"`
	AccessKeyId string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName string `json:"bucketName"`
}

var(
   err error
   bucket *oss.Bucket
   client *oss.Client
   configuration *conf
)

func init(){
   f, err:= os.Open("./oss-conf.json")
   if err!=nil{
   	fmt.Printf("Open configuration error: %s\n",err)
   	return
   }
   defer f.Close()
   configuration :=&conf{}
   decoder := json.NewDecoder(f)
   err =decoder.Decode(configuration)
   if err!=nil{
   	panic(err)
   }
   fmt.Println(configuration)

   endpoint := configuration.Endpoint
   accessKeyId := configuration.AccessKeyId
   accessKeySecret := configuration.AccessKeySecret
   bucketName := configuration.BucketName
   client, err = oss.New(endpoint, accessKeyId, accessKeySecret)
   if err != nil {
       handleError(err)
   }
   // 获取存储空间。
   bucket, err = client.Bucket(bucketName)
   if err != nil {
       handleError(err)
   }
}
