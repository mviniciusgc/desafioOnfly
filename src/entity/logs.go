package entity

import (
	"time"

	"github.com/mviniciusgc/onfly/src/enum"
)

type LogEntity struct {
	ID        int64
	LogType   enum.LogType
	Message   string
	CreatedAt time.Time
}
