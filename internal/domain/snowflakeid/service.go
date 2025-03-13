package snowflakeid

import "context"

var _ Service = &ServiceImplement{}

type Service interface {
	Generate(context.Context, TypeToGenerate) (Model, error)
	GenerateBatch(context.Context, int, TypeToGenerate) ([]Model, error)
}

type ServiceImplement struct {
	generator *Generator
}

func NewService(generator *Generator) *ServiceImplement {
	return &ServiceImplement{
		generator: generator,
	}
}

func (s *ServiceImplement) Generate(_ context.Context, typeToGenerate TypeToGenerate) (Model, error) {
	return s.generator.Generate(typeToGenerate)
}

func (s *ServiceImplement) GenerateBatch(_ context.Context, count int, typeToGenerate TypeToGenerate) ([]Model, error) {
	return nil, nil
}
