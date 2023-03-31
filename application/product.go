package application

type Product interface {
	isValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type product struct {
	ID     string
	Name   string
	Status string
	Price  float64
}

func (p *Product) IsValid() (bool, error) {

}

func (p *Product) Enable() error {

}

func (p *Product) Disable() error {

}

func (p *Product) GetID() string {

}

func (p *Product) GetName() string {

}

func (p *Product) GetStatus() string {

}

func (p *Product) GetPrice() float64 {

}
