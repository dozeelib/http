package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/dozeelib/http/dozee/endpoint"
)

func TestBuilder(t *testing.T) {
	e := endpoint.EndPointBuilder().SetService("console").SetDomain("dozee").SetTld("int").SetRegion("us").SetStage("sit").Build()
	fmt.Println(e)
}
