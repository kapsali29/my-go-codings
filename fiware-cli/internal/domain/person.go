package domain

type Person struct {
	Id       string            `json:"id"`
	Type     string            `json:"type"`
	Name     Attribute[string] `json:"name"`
	Room     Attribute[string] `json:"room"`
	Presence Attribute[bool]   `json:"presence"`
}

func CreatePerson(
	Id, Type, Name, Room string,
	Presence bool,
) *Person {
	return &Person{
		Id:       Id,
		Type:     Type,
		Name:     Attribute[string]{Type: "Text", Value: Name},
		Room:     Attribute[string]{Type: "Text", Value: Room},
		Presence: Attribute[bool]{Type: "Boolean", Value: Presence},
	}
}
