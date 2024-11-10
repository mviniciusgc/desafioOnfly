package logs

import "github.com/mviniciusgc/onfly/src/enum"

type IRepository interface {
	CreateLog(logType enum.LogType, message string) (err error)
}
