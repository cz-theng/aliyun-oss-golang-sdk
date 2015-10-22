/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"fmt"
)

func signature(req *Request) (sig string, err error) {
	sigStr := "GET\n"
	sigStr += "" + "\n"
	sigStr += "\n"
	sigStr += req.Date + "\n"
	sigStr += "/"
	sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr))
	if err != nil {
		return
	}
	return
}

type Bucket struct {
	Name         string
	CreationDate string
	Location     string
}

type Buckets struct {
	Bucket []Bucket
}

type Owner struct {
	ID          string
	DisplayName string
}

type BucketsInfo struct {
	XMLName     xml.Name `xml:"ListAllMyBucketsResult"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	NextMarker  string   `xml:"NextMarker"`
	Owner       Owner    `xml:"Owner"`
	Buckets     Buckets  `xml:"Buckets"`
}

func GetServiceWith(prefix, marker string, maxKeys int) (bucketsInfo *BucketsInfo, err error) {
	args := ""
	path := "/"
	if "" != prefix {
		args += "prefix=" + prefix
	}
	if "" != marker {
		args += "marker=" + marker
	}
	if 0 < maxKeys && maxKeys <= 1000 {
		args += "maxkeys=" + string(maxKeys)
	}

	if "" != args {
		path += "?" + args
	}
	req := &Request{Host: "oss.aliyuncs.com", Path: "/"}
	rsp, err := req.Send(signature)
	if rsp.Result != ESUCC {
		return
	}
	body := make([]byte, 10000)
	rsp.httpRsp.Body.Read(body)
	fmt.Println(string(body))
	bucketsInfo = new(BucketsInfo)
	xml.Unmarshal(body, bucketsInfo)
	return
}

func GetService() (bucketsInfo *BucketsInfo, err error) {
	bucketsInfo, err = GetServiceWith("", "", 0)
	return
}
