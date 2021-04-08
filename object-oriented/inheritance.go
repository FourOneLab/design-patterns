package object_oriented

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Run() string {
	return "Running"
}

type Cat struct {
	Animal
	Color string
}

func (c Cat) Miaow() string {
	return "miaow"
}
