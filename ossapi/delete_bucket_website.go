/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"path"
)

func DeleteBucketWebsite(name, location string) (ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name) + "/"
	req := &Request{
		Host:     host,
		Path:     "/?website",
		Method:   "DELETE",
		Resource: resource,
		SubRes:   []string{"website"}}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = OSSAPIError
			return
		}
	}
	fmt.Println("status:", rsp.httpRsp.Status)
	if rsp.Result != ESUCC {
		ossapiError = err.(*Error)
		return
	}
	return
}