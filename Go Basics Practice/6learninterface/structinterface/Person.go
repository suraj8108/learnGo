package structinterface

type Citizens interface {
	GetState() string
	GetAddress() string
	GetAdhar() string
}

type Person struct {
	Name    string
	Address string
	Age     int
}

func New(name string, address string, age int) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Age:     age,
	}
}

func (p *Person) GetState() string {
	return p.Address[5:]
}

func (p *Person) GetAddress() string {
	return p.Address
}

func (p *Person) GetAdhar() string {
	return "To be shared"
}
