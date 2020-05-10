package pkg

import (
	"github.com/gogf/gf/util/gvalid"
)


func ValidateParams(data interface{}) string {
	if e := gvalid.CheckStruct(data, nil); e != nil {
		s := e.String()
		return s
	}
    return ""
}
