package broker

type Topic string

func (m Topic) String() string {
	return string(m)
}

const (
	TopicRegisterCar Topic = "cars/register"
	TopicGetCars     Topic = "cars"
)
