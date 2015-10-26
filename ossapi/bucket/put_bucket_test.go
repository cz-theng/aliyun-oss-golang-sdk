/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"testing"
)

func TestPutBucketDefault(t *testing.T) {
	if nil != Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if p, err := PutBucketDefault("test-put-bucket2"); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println(p)
	}
	if p, err := PutBucket("test-put-bucket3", L_Beijing, P_Private); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		fmt.Println(p)
	}
}