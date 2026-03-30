package domain

type Appliance struct {
	Id            string             `json:"id"`
	Type          string             `json:"type"`
	Name          Attribute[string]  `json:"name"`
	Room          Attribute[string]  `json:"room"`
	Status        Attribute[string]  `json:"status"`
	ConnectedPlug *Attribute[string] `json:"connectedPlug,omitempty"`
}

func CreateAppliance(
	Id, Type, Name, Room, Status, ConnectedPlug string,
) *Appliance {
	var plug *Attribute[string]
	if ConnectedPlug != "" {
		plug = &Attribute[string]{Type: "Text", Value: ConnectedPlug}
	}
	return &Appliance{
		Id:            Id,
		Type:          Type,
		Name:          Attribute[string]{Type: "Text", Value: Name},
		Room:          Attribute[string]{Type: "Text", Value: Room},
		Status:        Attribute[string]{Type: "Text", Value: Status},
		ConnectedPlug: plug,
	}
}

type ApplianceResponse struct {
	Id            string  `json:"id"`
	Type          string  `json:"type"`
	Name          string  `json:"name"`
	Room          string  `json:"room"`
	Status        string  `json:"status"`
	ConnectedPlug *string `json:"connectedPlug,omitempty"`
}
