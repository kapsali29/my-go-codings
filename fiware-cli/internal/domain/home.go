package domain

type AttrValueField interface {
	int | float64 | string | bool
}

type Attribute[T AttrValueField] struct {
	Type  string `json:"type"`
	Value T      `json:"value"`
}

type Home struct {
	Id      string            `json:"id"`
	Type    string            `json:"type"`
	Name    Attribute[string] `json:"name"`
	Address Attribute[string] `json:"address"`
	Mode    Attribute[string] `json:"mode"`
}

func CreateHome(Id, Type, Name, Address, Mode string) *Home {
	return &Home{
		Id:      Id,
		Type:    Type,
		Name:    Attribute[string]{Type: "Text", Value: Name},
		Address: Attribute[string]{Type: "Text", Value: Address},
		Mode:    Attribute[string]{Type: "Text", Value: Mode},
	}
}
