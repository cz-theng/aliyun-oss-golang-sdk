/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestGetBucketWebsite(t *testing.T) {
	if nil != Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := GetBucketWebsite("test-put-bucket4", L_Beijing); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("GetBucketWebsite Success")
		fmt.Println(info)
	}
}