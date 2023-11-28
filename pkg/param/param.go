package param

type Param struct {
	value string
}

func New(value string) *Param {
	return &Param{
		value: value,
	}
}
