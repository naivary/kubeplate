package plugin

type Kind int

const (
	Inputer = iota + 1
	Outputer
	Funcs
)

func (k Kind) String() string {
	switch k {
	case Inputer:
		return "inputer"
	case Outputer:
		return "outputer"
	case Funcs:
		return "funcs"
	}
	return "UNDEFINED"
}

