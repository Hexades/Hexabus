package bus

type RepositoryEventListener interface {
	OnRepositoryEvent(event <-chan RepositoryEvent)
}

type RepositoryEvent interface {
}

type ServerEventListener interface {
	OnServerEvent(event <-chan ServerEvent)
}

type ServerEvent interface {
}
