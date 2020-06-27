package events

type Publisher interface {
	Publish([]byte) error
}
