package creational

import "fmt"

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	return nil, fmt.Errorf("not implemented yet")
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtCloner {
	var cache ShirtsCache
	return &cache
}

func whitePrototype() *Shirt {
	return &Shirt{
		Price: 15.00,
		SKU:   "empty",
		Color: White,
	}
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}