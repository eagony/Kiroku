package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ToDo struct {
	gorm.Model
	Summary  string `json:"summary"`
	Detail   string `json:"detail"`
	Priority uint8  `json:"priority"`
	DeadLine string `json:"deadline"`
}

func (td *ToDo) String() string {
	return fmt.Sprintf(
		"Summary: %v,\nDetail: %v,\nPriority: %v,\nDeadLine%v,\n",
		td.Summary,
		td.Detail,
		td.Priority,
		td.DeadLine,
	)
}

func init() {
	RegisterSingleton("ToDo", func() General {
		return new(ToDo)
	})
}
