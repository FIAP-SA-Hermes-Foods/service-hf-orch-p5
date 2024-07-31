package logger

import (
	"encoding/json"
	"testing"
)

type tStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Line string `json:"line"`
}

// go test -v -cover -count=1 -failfast -run ^Test_Logger$
func Test_Logger(t *testing.T) {

	marty := tStruct{Name: "Marty McFly", Age: 23, Line: "That's heavy!"}
	doc := tStruct{Name: "Doc Emmett Brown", Age: 65, Line: "Great Scott!"}
	einstein := tStruct{Name: "Einstein", Age: 14, Line: "Au au!"}
	biff := tStruct{Name: "Biff Tannen", Age: 18, Line: "I hate manure!"}

	t.Run("testing INFO", func(*testing.T) {
		b, err := json.Marshal(marty)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Info(string(b))
	})

	t.Run("testing INFOF", func(*testing.T) {
		b, err := json.Marshal(marty)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Infof("out: ", " | ", string(b), marty.Name, marty.Age)
	})

	t.Run("testing DEBUG", func(*testing.T) {
		b, err := json.Marshal(doc)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Debug(string(b))
	})

	t.Run("testing DEBUGF", func(*testing.T) {
		b, err := json.Marshal(doc)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Debugf("out: ", " | ", string(b), doc.Name, doc.Age, doc.Line)
	})

	t.Run("testing WARNING", func(*testing.T) {
		b, err := json.Marshal(einstein)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Warning(string(b))
	})

	t.Run("testing WARNINGF", func(*testing.T) {
		b, err := json.Marshal(einstein)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Warningf("out: ", " | ", string(b), einstein.Name, einstein.Age, einstein.Line)
	})

	t.Run("testing ERROR", func(*testing.T) {
		b, err := json.Marshal(biff)

		if err != nil {
			Errorf("error: ", " | ", err)
		}

		Error(string(b))
	})

	t.Run("testing ERRORF", func(*testing.T) {
		b, err := json.Marshal(biff)

		if err != nil {
			Errorf("error: ", " | ", err)
		}
		Errorf("out: ", " | ", string(b), biff.Name, biff.Age, biff.Line)
	})
}
