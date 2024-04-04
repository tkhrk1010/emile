package nature

import ()

type Nature struct{
	IsNature bool
	status string
}

func NewNature() *Nature {
	n := &Nature{status: "自然"}
	n.IsNature = n.isNatural()
	return n
}

func (n *Nature) isNatural() bool {
	return n.status == "習性によって変質している"
}
