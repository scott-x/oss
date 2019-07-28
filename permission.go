/*
* @Author: scottxiong
* @Date:   2019-07-29 02:30:15
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-29 02:41:52
*/
package oss

import (
	"os"
	"fmt"
)
//permission: 文件的访问权限优先级高于存储空间的访问权限。例如存储空间的访问权限是私有，而文件的访问权限是公共读写，则所有用户都有该文件的读写权限。如果某个文件没有设置过访问权限，则遵循存储空间的访问权限。
/*
1 继承Bucket	文件遵循存储空间的访问权限。oss.ACLDefault
2 私有	文件的拥有者和授权用户有该文件的读写权限，其他用户没有权限操作该文件。	oss.ACLPrivate
3 公共读	文件的拥有者和授权用户有该文件的读写权限，其他用户只有文件的读权限。请谨慎使用该权限。	oss.ACLPublicRead
公共读写	所有用户都有该文件的读写权限。请谨慎使用该权限。	oss.PublicReadWrite
*/

// func SetPermission(obj string,n int){
// 	// 设置文件的访问权限。
// 	var permission oss.ACLType= chmod[string(n-1)]
// 	err = bucket.SetObjectACL(obj,permission)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(-1)
// 	}
// }
func GetPermission(obj string){
	// 获取文件的访问权限。
	aclRes, err := bucket.GetObjectACL(obj)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Object ACL:", aclRes.ACL)
}