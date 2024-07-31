package entity

import (
	vo "service-hf-orch-p5/internal/core/domain/entity/valueObject"
)

type Client struct {
	ID        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	CPF       vo.Cpf       `json:"cpf,omitempty"`
	Email     string       `json:"email,omitempty"`
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	Address   string       `json:"address,omitempty"`
	CreatedAt vo.CreatedAt `json:"createdAt,omitempty"`
}
