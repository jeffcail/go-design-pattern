package fatory_method

import (
	"fmt"
	"testing"
)

func TestGetOrder(t *testing.T) {
	shoe, _ := GetOrder("shoe")
	fmt.Println(fmt.Sprintf("order: shoe %s", shoe.getName()))
	fmt.Println(fmt.Sprintf("order: shoe %s", shoe.getPrice()))

	pants, _ := GetOrder("pants")
	fmt.Println(fmt.Sprintf("order: pants %s", pants.getName()))
	fmt.Println(fmt.Sprintf("order: pants %s", pants.getPrice()))
}
