package bus

type hexabus struct {
	repositoryListenerChannels []chan RepositoryEvent
	serverListenerChannels []chan ServerEvent


}

type Hexabus interface {

	AddRepositoryListener(eventListener RepositoryEventListener)
	SendRepositoryEvent(event RepositoryEvent)

	AddServerListener(eventListener ServerEventListener)
	SendServerEvent(event ServerEvent)

}

var eb = hexabus{}
//TODO Come up with better mechanism to avoid the Get()
func Get() Hexabus {
	return &eb
}


func (bus *hexabus) AddRepositoryListener(eventListener RepositoryEventListener) {
	listenerChannel := make(chan RepositoryEvent, 10)
	bus.repositoryListenerChannels = append(bus.repositoryListenerChannels, listenerChannel)
	go eventListener.OnRepositoryEvent(listenerChannel)
}

func (bus *hexabus) SendRepositoryEvent(event RepositoryEvent) {
	for _, channel := range bus.repositoryListenerChannels {
		channel <- event
	}
}

func (bus *hexabus) AddServerListener(eventListener ServerEventListener) {
	listenerChannel := make(chan ServerEvent, 10)
	bus.serverListenerChannels = append(bus.serverListenerChannels, listenerChannel)
	go eventListener.OnServerEvent(listenerChannel)
}

func (bus *hexabus) SendServerEvent(event ServerEvent) {
	for _, channel := range bus.repositoryListenerChannels {
		channel <- event
	}
}