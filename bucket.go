/*
* @Author: scottxiong
* @Date:   2019-07-29 02:32:19
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-07-29 03:00:06
*/
package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)
var(
	err error
	bucket *oss.Bucket
	client *oss.Client
	// chmod =map[string]interface{}{"0":oss.ACLDefault,"1":oss.ACLPrivate,"2":oss.ACLPublicRead,"3":oss.PublicReadWrite,}
)

func Connect(){
	// 创建OSSClient实例。
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