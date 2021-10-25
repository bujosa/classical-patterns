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