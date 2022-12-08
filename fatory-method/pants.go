package fatory_method

// 具体产品 裤子
// Pants
type Pants struct {
	OrderBase
}

// NewPants
func NewPants() Order {
	return &Pants{
		OrderBase{
			name:  "西装裤",
			price: "10000",
		},
	}
}
