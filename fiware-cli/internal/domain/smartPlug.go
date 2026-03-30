package domain

type SmartPlug struct {
	Id     string             `json:"id"`
	Type   string             `json:"type"`
	Name   Attribute[string]  `json:"name"`
	Room   Attribute[string]  `json:"room"`
	Status Attribute[string]  `json:"status"`
	Power  Attribute[float64] `json:"power"`
}

func CreateSmartPlug(
	Id, Type, Name, Room, Status string,
	Power float64,
) *SmartPlug {
	return &SmartPlug{
		Id:     Id,
		Type:   Type,
		Name:   Attribute[string]{Type: "Text", Value: Name},
		Room:   Attribute[string]{Type: "Text", Value: Room},
		Status: Attribute[string]{Type: "Text", Value: Status},
		Power:  Attribute[float64]{Type: "Float", Value: Power},
	}
}
