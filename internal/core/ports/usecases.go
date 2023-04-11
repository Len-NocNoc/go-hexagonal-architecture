package ports

import "github.com/len-nocnoc/go-hexagonal-architecture/internal/core/domain"

type TodoUseCase interface {
	Get(id string) (*domain.ToDo, error)
	List() ([]domain.ToDo, error)
	Create(title, description string) (*domain.ToDo, error)
}