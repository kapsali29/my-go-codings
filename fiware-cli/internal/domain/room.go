package domain

type Room struct {
	Id          string             `json:"id"`
	Type        string             `json:"type"`
	Name        Attribute[string]  `json:"name"`
	Category    Attribute[string]  `json:"category"`
	Home        Attribute[string]  `json:"home"`
	Temperature Attribute[float64] `json:"temperature"`
	Humidity    Attribute[float64] `json:"humidity"`
	Occupancy   Attribute[int]     `json:"occupancy"`
}

func CreateRoom(
	Id, Type, Name, Category, Home string,
	Temperature, Humidity float64, Occupancy int,
) *Room {
	return &Room{
		Id:          Id,
		Type:        Type,
		Name:        Attribute[string]{Type: "Text", Value: Name},
		Category:    Attribute[string]{Type: "Text", Value: Category},
		Home:        Attribute[string]{Type: "Text", Value: Home},
		Temperature: Attribute[float64]{Type: "Float", Value: Temperature},
		Humidity:    Attribute[float64]{Type: "Float", Value: Humidity},
		Occupancy:   Attribute[int]{Type: "Integer", Value: Occupancy},
	}
}

type UpdateRoomObject struct {
	Occupancy   Attribute[int]     `json:"occupancy"`
}

func UpdateRoom(Occupancy int) UpdateRoomObject {
	return UpdateRoomObject{
		Occupancy: Attribute[int]{Type: "Integer", Value: Occupancy},
	}
}
