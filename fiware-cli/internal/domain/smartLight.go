package domain

type SmartLight struct {
	Id         string             `json:"id"`
	Type       string             `json:"type"`
	Name       Attribute[string]  `json:"name"`
	Room       Attribute[string]  `json:"room"`
	Status     Attribute[string]  `json:"status"`
	Brightness Attribute[int]     `json:"brightness"`
	Power      Attribute[float64] `json:"power"`
}

func CreateSmartLight(
	Id, Type, Name, Room, Status string,
	Brightness int,
	Power float64,
) *SmartLight {
	return &SmartLight{
		Id:         Id,
		Type:       Type,
		Name:       Attribute[string]{Type: "Text", Value: Name},
		Room:       Attribute[string]{Type: "Text", Value: Room},
		Status:     Attribute[string]{Type: "Text", Value: Status},
		Brightness: Attribute[int]{Type: "Integer", Value: Brightness},
		Power:      Attribute[float64]{Type: "Float", Value: Power},
	}
}

type UpdateSmartLightObject struct {
	Status     Attribute[string]  `json:"status"`
	Brightness Attribute[int]     `json:"brightness"`
	Power      Attribute[float64] `json:"power"`
}

func UpdateSmartLight(
	Status string,
	Brightness int,
	Power float64,
) UpdateSmartLightObject {
	return UpdateSmartLightObject{
		Status:     Attribute[string]{Type: "Text", Value: Status},
		Brightness: Attribute[int]{Type: "Integer", Value: Brightness},
		Power:      Attribute[float64]{Type: "Float", Value: Power},
	}
}

