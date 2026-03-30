package domain

type Thermostat struct {
	Id                string             `json:"id"`
	Type              string             `json:"type"`
	Name              Attribute[string]  `json:"name"`
	Room              Attribute[string]  `json:"room"`
	Mode              Attribute[string]  `json:"mode"`
	Status            Attribute[string]  `json:"status"`
	TargetTemperature Attribute[float64] `json:"targetTemperature"`
}

func CreateThermostat(
	Id, Type, Name, Room, Mode, Status string,
	TargetTemperature float64,
) *Thermostat {
	return &Thermostat{
		Id:                Id,
		Type:              Type,
		Name:              Attribute[string]{Type: "Text", Value: Name},
		Room:              Attribute[string]{Type: "Text", Value: Room},
		Mode:              Attribute[string]{Type: "Text", Value: Mode},
		Status:            Attribute[string]{Type: "Text", Value: Status},
		TargetTemperature: Attribute[float64]{Type: "Float", Value: TargetTemperature},
	}
}
