package demo_id_generator

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// IdGenerator used to generate random IDs
type IdGenerator interface {
	Generate() string
}

type LogTraceGenerator interface {
	IdGenerator
	Log()
}

type RandomGenerator struct {
	*zap.Logger
}

func NewRandomGenerator(logger *zap.Logger) *RandomGenerator {
	return &RandomGenerator{Logger: logger}
}

// Generate an random ID
func (g *RandomGenerator) Generate() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return ""
	}

	return id.String()
}

func (g *RandomGenerator) Log() {
	panic("implement me")
}
