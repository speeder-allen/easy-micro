package govern

import "errors"

type Repo interface {
	GetServices() []Service
	GetService(name string) (Service, error)
	AddService(Service)
	WatchService(...Service) chan struct{}
}

var ErrorNilRepo = errors.New("error the repo is nil")
