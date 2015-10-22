/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestGetService(t *testing.T) {
	if buckets, err := GetService(); err != nil {
		if err != nil {
			t.Log("Error :", err.Error())
			fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
		}
	} else {
		fmt.Println(buckets)
	}
}
