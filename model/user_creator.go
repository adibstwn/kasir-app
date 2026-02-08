package model

import "time"

type UserCreator struct {
	CreatedBy *string    `json:"created_by,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
