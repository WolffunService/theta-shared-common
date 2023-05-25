package watcherenum

type Service string

const (
	SvcThetanRivalService Service = "thetan-rival-service"
)

func (s Service) String() string {
	return string(s)
}

type Topic string

const (
	TpCreateCosmetic Topic = "create-cosmetic"
)

func (t Topic) String() string {
	return string(t)
}
