package dto

import (
	"service-hf-orch-p5/internal/core/domain/entity"
	vo "service-hf-orch-p5/internal/core/domain/entity/valueObject"
)

type ClientDB struct {
	UUID      string `json:"uuid,omitempty"`
	Name      string `json:"name,omitempty"`
	CPF       string `json:"cpf,omitempty"`
	Email     string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Address string `json:"address,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type (
	RequestClient struct {
		UUID      string `json:"uuid,omitempty"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
		Address string `json:"address,omitempty"`
		CreatedAt string `json:"createdAt,omitempty"`
	}

	OutputClient struct {
		UUID      string `json:"uuid,omitempty"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
		Address string `json:"address,omitempty"`
		CreatedAt string `json:"createdAt,omitempty"`
	}
)

func (r RequestClient) Client() entity.Client {
	return entity.Client{
		Name: r.Name,
		CPF: vo.Cpf{
			Value: r.CPF,
		},
		Email: r.Email,
	}
}
