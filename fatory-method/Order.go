package fatory_method

// 商品接口
type Order interface {
	setName(name string)
	setPrice(price string)
	getName() string
	getPrice() string
}
