# Create all Entities
## Home
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Home:001",
    "type": "Home",
    "name": { "type": "Text", "value": "My Smart Home" },
    "address": { "type": "Text", "value": "Main Street 10" },
    "mode": { "type": "Text", "value": "home" }
  }'
```
## Rooms
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Room:LivingRoom",
    "type": "Room",
    "name": { "type": "Text", "value": "Living Room" },
    "category": { "type": "Text", "value": "livingroom" },
    "home": { "type": "Text", "value": "Home:001" },
    "temperature": { "type": "Number", "value": 22.5 },
    "humidity": { "type": "Number", "value": 44 },
    "occupancy": { "type": "Integer", "value": 1 }
  }'


curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Room:Kitchen",
    "type": "Room",
    "name": { "type": "Text", "value": "Kitchen" },
    "category": { "type": "Text", "value": "kitchen" },
    "home": { "type": "Text", "value": "Home:001" },
    "temperature": { "type": "Number", "value": 21.0 },
    "humidity": { "type": "Number", "value": 48 },
    "occupancy": { "type": "Integer", "value": 0 }
  }'

curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Room:Bedroom1",
    "type": "Room",
    "name": { "type": "Text", "value": "Bedroom 1" },
    "category": { "type": "Text", "value": "bedroom" },
    "home": { "type": "Text", "value": "Home:001" },
    "temperature": { "type": "Number", "value": 20.0 },
    "humidity": { "type": "Number", "value": 41 },
    "occupancy": { "type": "Integer", "value": 0 }
  }'
```
## Smart Lights
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Light_LivingRoom_01",
    "type": "SmartLight",
    "name": { "type": "Text", "value": "Living Room Ceiling Light" },
    "room": { "type": "Text", "value": "Room:LivingRoom" },
    "status": { "type": "Text", "value": "off" },
    "brightness": { "type": "Integer", "value": 0 },
    "power": { "type": "Number", "value": 0.0 }
  }'

curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Light_Bedroom1_01",
    "type": "SmartLight",
    "name": { "type": "Text", "value": "Bedroom Main Light" },
    "room": { "type": "Text", "value": "Room:Bedroom1" },
    "status": { "type": "Text", "value": "off" },
    "brightness": { "type": "Integer", "value": 0 },
    "power": { "type": "Number", "value": 0.0 }
  }'

```

## Thermostats
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Thermostat_LivingRoom_01",
    "type": "Thermostat",
    "name": { "type": "Text", "value": "Living Room Thermostat" },
    "room": { "type": "Text", "value": "Room:LivingRoom" },
    "targetTemperature": { "type": "Number", "value": 21.5 },
    "mode": { "type": "Text", "value": "auto" },
    "status": { "type": "Text", "value": "idle" }
  }'
```
## Motion Sensors
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Motion_LivingRoom_01",
    "type": "MotionSensor",
    "name": { "type": "Text", "value": "Living Room Motion Sensor" },
    "room": { "type": "Text", "value": "Room:LivingRoom" },
    "motionDetected": { "type": "Boolean", "value": false },
    "batteryLevel": { "type": "Number", "value": 87 }
  }'

curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Motion_Kitchen_01",
    "type": "MotionSensor",
    "name": { "type": "Text", "value": "Kitchen Motion Sensor" },
    "room": { "type": "Text", "value": "Room:Kitchen" },
    "motionDetected": { "type": "Boolean", "value": false },
    "batteryLevel": { "type": "Number", "value": 90 }
  }'
```
## Door Sensor
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Door_Main_01",
    "type": "DoorSensor",
    "name": { "type": "Text", "value": "Main Entrance Door Sensor" },
    "room": { "type": "Text", "value": "Room:LivingRoom" },
    "doorOpen": { "type": "Boolean", "value": false },
    "batteryLevel": { "type": "Number", "value": 76 }
  }'
```
## Smart Plug
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Device:Plug_LivingRoom_01",
    "type": "SmartPlug",
    "name": { "type": "Text", "value": "TV Smart Plug" },
    "room": { "type": "Text", "value": "Room:LivingRoom" },
    "status": { "type": "Text", "value": "on" },
    "power": { "type": "Number", "value": 12.3 }
  }'
```
## Person
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Person:Alice",
    "type": "Person",
    "name": { "type": "Text", "value": "Alice" },
    "currentRoom": { "type": "Text", "value": "Room:LivingRoom" },
    "presence": { "type": "Boolean", "value": true }
  }'
```
## Appliances
```bash
curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Appliance:TV01",
    "type": "Appliance",
    "name": { "type": "Text", "value": "Television" },
    "room": { "type": "Text", "value": "Room:LivingRoom" },
    "status": { "type": "Text", "value": "standby" },
    "connectedPlug": { "type": "Text", "value": "Device:Plug_LivingRoom_01" }
  }'

curl -iX POST 'http://localhost:1026/v2/entities' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "Appliance:CoffeeMachine01",
    "type": "Appliance",
    "name": { "type": "Text", "value": "Coffee Machine" },
    "room": { "type": "Text", "value": "Room:Kitchen" },
    "status": { "type": "Text", "value": "off" }
  }'
```
# Read Examples

```bash
curl -X GET 'http://localhost:1026/v2/entities?options=keyValues'
curl -X GET 'http://localhost:1026/v2/entities?type=Room&options=keyValues'
curl -X GET 'http://localhost:1026/v2/entities?type=SmartLight&options=keyValues'
curl -X GET \
  'http://localhost:1026/v2/entities?q=room==Room:LivingRoom&options=keyValues'

curl -X GET \
  'http://localhost:1026/v2/entities/Device:Thermostat_LivingRoom_01'
```
# Interaction 1
These are realistic flows you can test immediately with Orion by updating attributes and then querying the affected entities. Orion’s normal update pattern is attribute update via PATCH /v2/entities/{id}/attrs.

Interaction A: person enters living room, motion detected, light turns on
Motion sensor detects movement

## Motion sensor detects movement
```bash
curl -iX PATCH \
  'http://localhost:1026/v2/entities/Device:Motion_LivingRoom_01/attrs' \
  -H 'Content-Type: application/json' \
  -d '{
    "motionDetected": { "type": "Boolean", "value": true }
  }'
```
## Occupancy in living room increases
```bash
curl -iX PATCH \
  'http://localhost:1026/v2/entities/Room:LivingRoom/attrs' \
  -H 'Content-Type: application/json' \
  -d '{
    "occupancy": { "type": "Integer", "value": 2 }
  }'
```
## Turn on the living room light
```bash
curl -iX PATCH \
  'http://localhost:1026/v2/entities/Device:Light_LivingRoom_01/attrs' \
  -H 'Content-Type: application/json' \
  -d '{
    "status": { "type": "Text", "value": "on" },
    "brightness": { "type": "Integer", "value": 80 },
    "power": { "type": "Number", "value": 9.5 }
  }'
```
## Verify state
```bash
curl -X GET \
  'http://localhost:1026/v2/entities/Device:Light_LivingRoom_01?options=keyValues'
```
# Interaction 2
## Room temperature falls
```bash
curl -iX PATCH \
  'http://localhost:1026/v2/entities/Room:LivingRoom/attrs' \
  -H 'Content-Type: application/json' \
  -d '{
    "temperature": { "type": "Number", "value": 19.0 }
  }'
```
## Thermostat reacts
```bash
curl -iX PATCH \
  'http://localhost:1026/v2/entities/Device:Thermostat_LivingRoom_01/attrs' \
  -H 'Content-Type: application/json' \
  -d '{
    "status": { "type": "Text", "value": "heating" },
    "targetTemperature": { "type": "Number", "value": 21.5 }
  }'
```
## Check room and thermostat together
```bash
curl -X GET \
  'http://localhost:1026/v2/entities?type=Room,Thermostat&options=keyValues'
```