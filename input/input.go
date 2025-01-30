package input

type Inputer interface {
	Read(path string) ([]byte, error)
}

// x/y/z/vars.yaml
