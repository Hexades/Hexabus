package bus

import "github.com/google/uuid"

func ID(repoType string, instanceID uuid.UUID) RepositoryID {
	return &repository{id: instanceID, repoType: repoType}
}

type RepositoryID interface {
	GetID() uuid.UUID
	GetRepoType() string
}
type repository struct {
	id       uuid.UUID
	repoType string
}

// GetID implements Repository.
func (r *repository) GetID() uuid.UUID {
	return r.id
}

// GetRepoType implements Repository.
func (r *repository) GetRepoType() string {
	return r.repoType
}

type RepositoryEventListener interface {
	OnRepositoryEvent(event <-chan RepositoryEvent)
}

type RepositoryEvent interface {
	HexabusEvent
}

type RepositoryRequestEvent interface {
	RepositoryEvent
}
type RepositoryResponseEvent interface {
	RepositoryEvent
}
