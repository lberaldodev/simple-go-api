package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

func ParseId(id string) (ID, error) {
	return uuid.Parse(id)
}
