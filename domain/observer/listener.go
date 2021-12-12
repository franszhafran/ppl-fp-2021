package observer

type Listener interface {
	Update(string, interface{}) error
}
