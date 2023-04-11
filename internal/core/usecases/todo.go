package usecases

import "github.com/len-nocnoc/go-hexagonal-architecture/internal/core/ports"
import "github.com/len-nocnoc/go-hexagonal-architecture/internal/core/domain"

type todoUseCase struct {
	todoRepo ports.TodoRepository
}

func NewToDoUseCase(todoRepo ports.TodoRepository) ports.TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

func (t *todoUseCase) Get(id string) (*domain.ToDo, error) {
	return &domain.ToDo{}, nil
}

func (t *todoUseCase) List() ([]domain.ToDo, error) {
	return []domain.ToDo{}, nil
}

func (t *todoUseCase) Create(title, description string) (*domain.ToDo, error) {
	return &domain.ToDo{}, nil
}
