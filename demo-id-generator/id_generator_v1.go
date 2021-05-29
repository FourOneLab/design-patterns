package demo_id_generator

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type IdGeneratorV1 struct {
	*zap.Logger
}

func NewIdGeneratorV1(logger *zap.Logger) *IdGeneratorV1 {
	return &IdGeneratorV1{Logger: logger}
}

func (g *IdGeneratorV1) Generate() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return ""
	}

	return id.String()
}
