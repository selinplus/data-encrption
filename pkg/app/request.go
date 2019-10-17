package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/selinplus/data-encrption/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
