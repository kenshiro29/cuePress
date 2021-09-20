package functions

import (
	"github.com/google/uuid"
)

type UUID struct {
	STATUS string `json:"status" xml:"status"`
	UUID   string `json:"uuid" xml:"uuid"`
}

func genUUID() interface{} {

	id := uuid.New()
	uuid := &UUID{
		"200",
		id.String(),
	}
	return uuid
}
