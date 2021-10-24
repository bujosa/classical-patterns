package abstract_factory

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}