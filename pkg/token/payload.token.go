package token

import "github.com/google/uuid"

type Payload struct {
	UserId uuid.UUID
	RoleId string
	Start  int64
	Exp    int64 `json:"exp,omitempty"`
}
