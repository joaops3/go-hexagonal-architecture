package application

import "fmt"


type ProductInterface interface {
	isValid() (bool error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}


const (
	DISABLED = "DISABLED"
	ENABLED = "ENABLED"
)

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}


func (p *Product) isValid() (bool, error){
	return true, nil
}

func (p *Product) Enable() ( error){
	if(p.Price > 0) {
		p.Status = ENABLED
		return nil
	}
	
	return fmt.Errorf("O Preço deve ser maior que zero")

}

func (p *Product) Disable() ( error){

	p.Status = DISABLED
	return  nil
}

func (p *Product) GetId() (string){
	return  p.Name
}

func (p *Product) GetStatus() (string){
	return  p.Status
}
func (p *Product) GetPrice() (float64){
	return  p.Price
}

func (p *Product) GetName() (string){
	return  p.Name
}