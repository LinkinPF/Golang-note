package builder


type Builder interface {
	Build()
}

type  Director struct {
	builder Builder
}

func NewDirector(b Builder) Director {
	return Director{
		builder : b,
	}
}

func (d *Director)Construct() {
	//这里就把接口和具体实现分开了
	d.builder.Build()
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{
		built : false,
	}
}

func (b *ConcreteBuilder)Build() {
	b.built = true
}

type Product struct {
	built bool
}

//GetResult会创建一个product并返回
func (b *ConcreteBuilder)GetResult() Product {
	return Product{b.built}
}

