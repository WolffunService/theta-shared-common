package notification

type Notification struct {
}

func New() *Notification {
	return &Notification{}
}

func (notify Notification) Send() error {
	return nil
}
