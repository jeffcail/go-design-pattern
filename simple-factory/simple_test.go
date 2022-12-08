package simple_factory

import (
	"fmt"
	"testing"
)

func TestParseIni_Parse(t *testing.T) {
	config := NewParseConfig("ini")
	parse := config.Parse("ini")
	fmt.Println(parse)
}

func TestParseYaml_Parse(t *testing.T) {
	config := NewParseConfig("yaml")
	parse := config.Parse("yaml")
	fmt.Println(parse)
}

func TestParseJson_Parse(t *testing.T) {
	config := NewParseConfig("json")
	parse := config.Parse("json")
	fmt.Println(parse)
}
