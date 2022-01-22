package models

import "gorm.io/gorm"

type Development struct {
	gorm.Model
	ID       int
	BitrixID string
	Name     string
}

type ResponseDevelopments struct {
	Result BitrixDevelopmentList `json:"result"`
}

type BitrixDevelopmentList struct {
	DevelopmentList BitrixDevelopmentItems `json:"UF_CRM_5D12A1A9D28ED"`
}

type BitrixDevelopmentItems struct {
	Items []BitrixDevelopments `json:"items"`
}

type BitrixDevelopments struct {
	ID   string `json:"ID"`
	Name string `json:"VALUE"`
}
