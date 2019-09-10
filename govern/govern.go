package govern

import "github.com/google/uuid"

type Service struct {
	Name  string        `json:"name"`
	Nodes []ServiceNode `json:"nodes"`
}

type ServiceNode struct {
	UUID     string                 `json:"uuid"`
	Address  string                 `json:"address"`
	Port     uint32                 `json:"port"`
	Endpoint string                 `json:"endpoint"`
	Protocol ServiceProtocol        `json:"protocol"`
	SSL      bool                   `json:"ssl"`
	Version  string                 `json:"version"`
	MetaData map[string]interface{} `json:"meta_data"`
}

type NodeOption func(node *ServiceNode)

type ServiceProtocol uint8

const (
	REST ServiceProtocol = iota
	Grpc
)

var serviceTypes = map[ServiceProtocol]string{REST: "REST", Grpc: "grpc"}

func (t ServiceProtocol) String() string {
	if s, ok := serviceTypes[t]; ok {
		return s
	}
	return "Invalid Service Types"
}

func NewService(name string, nodes ...ServiceNode) Service {
	return Service{Name: name, Nodes: nodes}
}

var (
	defaultServiceProtocol = REST
	defaultServiceSSL      = false
	emptyService           = Service{}
)

func NewNode(address string, port uint32, version Version, opts ...NodeOption) ServiceNode {
	node := ServiceNode{
		UUID:     uuid.New().String(),
		Address:  address,
		Port:     port,
		SSL:      false,
		Version:  version.String(),
		Protocol: defaultServiceProtocol,
		MetaData: make(map[string]interface{}),
	}
	for _, o := range opts {
		o(&node)
	}
	return node
}
