package testhelper

import (
	"encoding/json"
	"fmt"
)

// Result common output
type Result struct {
	Data  interface{}
	Error error
}

//CreateHttpRequestBodyMock create http body mock
func CreateHttpRequestBodyMock(structure interface{}) string {
	json, _ := json.Marshal(structure)
	result := string(json)

	return result
}

//SetMockerySharedResult set shared result mock
func SetMockerySharedResult(result interface{}) <-chan Result {
	sharedResult := result.(Result)

	// simulasiin untuk set channel shared result
	resultShared := func() <-chan Result {
		output := make(chan Result)
		go func() { output <- sharedResult }()
		return output
	}()

	return resultShared
}

//SetTestcaseName set testcase name to prevent tech debt
func SetTestcaseName(number int, description string) string {
	return fmt.Sprintf("Testcase #%v : %s", number, description)
}
