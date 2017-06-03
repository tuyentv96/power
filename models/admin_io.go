package model

type AddDeviceInput struct {
	Uid string `json:"uid"`
	LDid []string `json:"list_device"`
}

