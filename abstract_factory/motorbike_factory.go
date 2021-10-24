package abstract_factory

import "fmt"

const (
	SportMotorbikeType   = 1
	CruiserMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiserMotorbikeType:
		return new(CruiserMotorbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}