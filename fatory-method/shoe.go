package fatory_method

// 具体产品 鞋子
// Shoe
type Shoe struct {
	OrderBase
}

// NewShoe
func NewShoe() Order {
	return &Shoe{
		OrderBase{
			name:  "AJ",
			price: "10000",
		},
	}
}
