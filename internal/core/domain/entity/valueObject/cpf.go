package valueobject

import (
	"errors"
	"fmt"
	"regexp"
)

type Cpf struct {
	Value string `json:"value,omitempty"`
}

var regCPF = regexp.MustCompile(`[0-9-.]*`)

func (c Cpf) Validate() error {

	if len(c.Value) == 0 {
		return errors.New("cpf value is null or invalid")
	}

	cpfMatch := regCPF.FindStringSubmatch(c.Value)

	if cpfMatch == nil {
		return fmt.Errorf("cpf %s is not valid", c.Value)
	}

	return nil
}
