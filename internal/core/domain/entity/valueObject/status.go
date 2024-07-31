package valueobject

import (
	"errors"
	"strings"
)

type Status struct {
	Value string `json:"value,omitempty"`
}

const (
	PaidStatusKey       = "paid"
	CanceledStatusKey   = "canceled"
	ReceivedStatusKey   = "received"
	InProgressStatusKey = "in progress"
	DoneStatusKey       = "done"
	FinishedStatusKey   = "finished"
)

const (
	PaidStatus       = "Paid"
	CanceledStatus   = "Canceled"
	ReceivedStatus   = "Received"
	InProgressStatus = "In Progress"
	DoneStatus       = "Done"
	FinishedStatus   = "Finished"
)

var statusMap = map[string]string{
	PaidStatusKey:       PaidStatus,
	CanceledStatusKey:   CanceledStatus,
	ReceivedStatusKey:   ReceivedStatus,
	InProgressStatusKey: InProgressStatus,
	DoneStatusKey:       DoneStatus,
	FinishedStatusKey:   FinishedStatus,
}

func (v *Status) Validate() error {

	status, ok := statusMap[strings.ToLower(v.Value)]

	if !ok {
		return errors.New("status is not valid")
	}

	v.Value = status

	return nil
}
