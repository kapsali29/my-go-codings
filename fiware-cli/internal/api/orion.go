package api

import (
	"bytes"
	"encoding/json"
	"fiware-cli/internal/config"
	"fiware-cli/internal/domain"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OrionEntity interface {
	domain.Home | domain.Room | domain.SmartLight | domain.Thermostat | domain.MotionSensor | domain.DoorSensor | domain.SmartPlug | domain.Person | domain.Appliance
}

func getBaseUrl() string {
	return fmt.Sprintf(
		"%v/%v/%v",
		config.ORION_BASE_URL,
		config.ORION_API_VERSION,
		config.ENTITIES,
	)
}

func PostEntity[T OrionEntity](entity T) {
	url := getBaseUrl()
	jsonData, err := json.Marshal(entity)
	if err != nil {
		panic(err)
	}
	resp, errResp := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if errResp != nil {
		panic(errResp)
	} else {
		fmt.Println(resp)
	}
}

func GetEntitiesValues[T any](Type string) []T {
	var res []T
	fullUrl := fmt.Sprintf(
		"%v?type=%s&options=keyValues",
		getBaseUrl(),
		Type,
	)
	resp, resErr := http.Get(
		fullUrl,
	)
	if resErr != nil {
		panic(resErr)
	}
	defer resp.Body.Close()
	jsonErr := json.NewDecoder(resp.Body).Decode(&res)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return res
}

func UpdateMotionSensorAPI(Id string, Value bool) {
	fullUrl := fmt.Sprintf(
		"%v/%v/attrs",
		getBaseUrl(),
		Id,
	)
	payloadObject := domain.UpdateMotionSensor(Value)
	payload, err := json.Marshal(payloadObject)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(
		http.MethodPatch,
		fullUrl,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status", resp.Status)
	fmt.Println("Response:", string(body))
}

func UpdateSmartLightAPI(
	Id, Status string,
	Brightness int,
	Power float64,
) {
	fullUrl := fmt.Sprintf(
		"%v/%v/attrs",
		getBaseUrl(),
		Id,
	)
	payloadObject := domain.UpdateSmartLight(
		Status, Brightness, Power,
	)
	payload, err := json.Marshal(payloadObject)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(
		http.MethodPatch,
		fullUrl,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10*time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(*resp)
	fmt.Println(string(body))
}

func UpdateRoomApi(Id string, Occupancy int) {
	fullUrl := fmt.Sprintf(
		"%v/%v/attrs",
		getBaseUrl(),
		Id,
	)
	payloadObject := domain.UpdateRoom(Occupancy)
	payload, err := json.Marshal(payloadObject)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPatch, fullUrl, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10*time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(req.Body)
	fmt.Println(string(body))
	fmt.Println(*resp)
}