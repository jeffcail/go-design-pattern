package fatory_method

type OrderBase struct {
	name  string
	price string
}

func (o *OrderBase) setName(name string) {
	o.name = name
}

func (o *OrderBase) getName() string {
	return o.name
}

func (o *OrderBase) setPrice(price string) {
	o.price = price
}

func (o *OrderBase) getPrice() string {
	return o.price
}
