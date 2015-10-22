/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/log"
	"net/http"
)

//LogCat
var Logger *log.Logger

func init() {
	var err error
	Logger, err = log.NewFileLogger(".ossapilog", "ossapi")
	if err != nil {
		fmt.Errorf("Create Logger Error\n")
		return
	}
	Logger.SetMaxFileSize(1024 * 1024 * 100) //100MB
	Logger.SetLevel(log.LDEBUG)
}

var (
	// global Access Key ID
	accessKeyID string
	// global Access Key Secret
	accessKeySecret string
)

const (
	// GMT Format Date
	DATE_FMT = "Mon, 02 Jan 2006 15:04:05 GMT"
)

// http client for http request
var httpClient http.Client

/**
* Init ossapi with Access Key's ID and secret
* @param ID : Access Key's ID
* @param secret : Access Key's secret
 */
func Init(ID string, secret string) error {
	if "" == ID || "" == secret {
		return EARG
	}
	accessKeyID = ID
	accessKeySecret = secret
	return nil
}
