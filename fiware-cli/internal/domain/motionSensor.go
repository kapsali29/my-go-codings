package domain

type MotionSensor struct {
	Id             string             `json:"id"`
	Type           string             `json:"type"`
	Name           Attribute[string]  `json:"name"`
	Room           Attribute[string]  `json:"room"`
	MotionDetected Attribute[bool]    `json:"motionDetected"`
	BatteryLevel   Attribute[float64] `json:"batteryLevel"`
}

func CreateMotionSensor(
	Id, Type, Name, Room string,
	MotionDetected bool,
	BatteryLevel float64,
) *MotionSensor {
	return &MotionSensor{
		Id:             Id,
		Type:           Type,
		Name:           Attribute[string]{Type: "Text", Value: Name},
		Room:           Attribute[string]{Type: "Text", Value: Room},
		MotionDetected: Attribute[bool]{Type: "Boolean", Value: MotionDetected},
		BatteryLevel:   Attribute[float64]{Type: "Float", Value: BatteryLevel},
	}
}

type UpdateMotionSensorPayload struct {
	MotionDetected Attribute[bool] `json:"motionDetected"`
}

func UpdateMotionSensor(value bool) UpdateMotionSensorPayload {
	return UpdateMotionSensorPayload{
		MotionDetected: Attribute[bool]{Type: "Boolean", Value: value},
	}
}
