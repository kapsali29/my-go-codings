package main

import (
	"fiware-cli/internal/api"
	"fiware-cli/internal/domain"
	"fmt"
)

func main() {
	fmt.Println("Hello from Fiware CLI")
	home := *domain.CreateHome("Home:001", "Home", "My Smart Home 1", "Main XX222X", "home")
	livingRoom := *domain.CreateRoom(
		"Room:LivingRoom", "Room", "Living Room", "living-room",
		"Home:001", 22.5, 24, 1,
	)
	kitchen := *domain.CreateRoom(
		"Room:Kitchen", "Room", "Kitchen", "kitchen",
		"Home:001", 21.0, 48, 1,
	)
	bedroom := *domain.CreateRoom(
		"Room:Bedroom", "Room", "Bedroom", "bedroom",
		"Home:001", 20.0, 41, 1,
	)
	sl1 := *domain.CreateSmartLight(
		"Device:Light_LivingRoom_01", "SmartLight", "Living Room Ceiling Light",
		"Room:LivingRoom", "off", 0, 0.0,
	)
	sl2 := *domain.CreateSmartLight(
		"Device:Light_Bedroom1_01", "SmartLight", "Bedroom Main Light",
		"Room:Bedroom", "off", 0, 0.0,
	)
	th := *domain.CreateThermostat(
		"Device:Thermostat_LivingRoom_01", "Thermostat",
		"Living Room Thermostat", "Room:LivingRoom",
		"auto", "idle", 21.5,
	)
	ms := *domain.CreateMotionSensor(
		"Device:Motion_LivingRoom_01", "MotionSensor",
		"Living Room Motion Sensor", "Room:LivingRoom",
		false, 87,
	)
	ms2 := *domain.CreateMotionSensor(
		"Device:Motion_Kitchen_01", "MotionSensor",
		"Kitchen Room Motion Sensor", "Room:Kitchen",
		false, 90,
	)
	ds := *domain.CreateDoorSensor(
		"Device:Door_Main_01", "DoorSensor", "Main Entrance Door Sensor",
		"Room:LivingRoom", false, 76,
	)
	sp := *domain.CreateSmartPlug(
		"Device:Plug_LivingRoom_01", "SmartPlug", "TV Smart Plug",
		"Room:LivingRoom", "on", 12.3,
	)
	p := *domain.CreatePerson(
		"Person:Panos",
		"Person",
		"Panos",
		"Room:LivingRoom",
		true,
	)
	a1 := *domain.CreateAppliance(
		"Appliance:TV01", "Appliance", "Television",
		"Room:LivingRoom", "standby", "Device:Plug_LivingRoom_01",
	)
	a2 := *domain.CreateAppliance(
		"Appliance:CoffeeMachine1", "Appliance", "Coffee Machine",
		"Room:Kitchen", "off", "",
	)
	api.PostEntity(home)
	api.PostEntity(livingRoom)
	api.PostEntity(kitchen)
	api.PostEntity(bedroom)
	api.PostEntity(sl1)
	api.PostEntity(sl2)
	api.PostEntity(th)
	api.PostEntity(ms)
	api.PostEntity(ms2)
	api.PostEntity(ds)
	api.PostEntity(sp)
	api.PostEntity(p)
	api.PostEntity(a1)
	api.PostEntity(a2)

	res := api.GetEntitiesValues[domain.ApplianceResponse]("Appliance")
	fmt.Println(res)
	api.UpdateMotionSensorAPI("Device:Motion_LivingRoom_01", true)
	api.UpdateSmartLightAPI("Device:Light_LivingRoom_01", "on", 80, 9.55)
	api.UpdateRoomApi("Room:LivingRoom", 2)
}
