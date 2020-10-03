package main

type Product interface {
	Size() int
}

type Car struct {

}

type Bus struct {

}

func(c *Car) Size() int {
	return 1
}

func(b *Bus) Size() int {
	return 2
}


type Factory struct {

}

func NewFactory() *Factory {
	return &Factory{}
}

func(f *Factory) MakeProduct(kind int) Product {
	if kind == 1 {
		return &Car{}
	}

	return &Bus{}
}