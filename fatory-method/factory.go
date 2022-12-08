package fatory_method

import "fmt"

// GetOrder
// 工厂
func GetOrder(t string) (Order, error) {
	if t == "shoe" {
		return NewShoe(), nil
	}
	if t == "pants" {
		return NewPants(), nil
	}
	return nil, fmt.Errorf("ORDER IS NOT EXISTS")
}
