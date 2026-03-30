package domain

type DoorSensor struct {
	Id           string             `json:"id"`
	Type         string             `json:"type"`
	Name         Attribute[string]  `json:"name"`
	Room         Attribute[string]  `json:"room"`
	DoorOpen     Attribute[bool]    `json:"doorOpen"`
	BatteryLevel Attribute[float64] `json:"batteryLevel"`
}

func CreateDoorSensor(
	Id, Type, Name, Room string,
	DoorOpen bool,
	BatteryLevel float64,
) *DoorSensor {
	return &DoorSensor{
		Id:           Id,
		Type:         Type,
		Name:         Attribute[string]{Type: "Text", Value: Name},
		Room:         Attribute[string]{Type: "Text", Value: Room},
		DoorOpen:     Attribute[bool]{Type: "Boolean", Value: DoorOpen},
		BatteryLevel: Attribute[float64]{Type: "Float", Value: BatteryLevel},
	}
}
