/**
* Author: CZ cz.theng@gmail.com
 */

package mupload

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"path"
	"strconv"
)

type UploadPartInfo struct {
	ObjectName string
	BucketName string
	Location   string
	PartNumber int
	UploadID   string
	Data       []byte
	CntType    string
}

func UploadPart(partInfo *UploadPartInfo) (ossapiError *ossapi.Error) {
	if partInfo == nil {
		ossapiError = ossapi.ArgError
		return
	}
	resource := path.Join("/", partInfo.BucketName, partInfo.ObjectName)
	host := partInfo.BucketName + "." + partInfo.Location + ".aliyuncs.com"
	req := &ossapi.Request{
		Host:     host,
		Path:     "/" + partInfo.ObjectName + "?partNumber=" + strconv.FormatUint(uint64(partInfo.PartNumber), 10) + "uploadId=" + partInfo.UploadID,
		Method:   "PUT",
		Body:     partInfo.Data,
		CntType:  partInfo.CntType,
		SubRes:   []string{"partNumber=" + strconv.FormatUint(uint64(partInfo.PartNumber), 10) + "uploadId=" + partInfo.UploadID},
		Resource: resource}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*ossapi.Error); !ok {
			ossapi.Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = ossapi.OSSAPIError
			return
		}
	}
	if rsp.Result != ossapi.ESUCC {
		ossapiError = err.(*ossapi.Error)
		return
	}
	fmt.Println("rsp:", rsp.HttpRsp)
	return
}