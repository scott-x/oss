# oss
utils based on [aliyun-oss-go-sdk](https://github.com/scott-x/aliyun-oss-go-sdk/tree/master/oss)

### usage
#### 1. create a file named `oss-conf.json` in your project

```json
{
    "endpoint" : "",
    "accessKeyId" : "",
    "accessKeySecret":"",
    "bucketName": ""
}
```
Then fill in the necessary data

#### 2. use the tool, `main.go`
```
package main

import (
	"github.com/scott-x/oss"
)

func main(){
	oss.Upload("a.png",true)
	oss.ListAll()
}
```
