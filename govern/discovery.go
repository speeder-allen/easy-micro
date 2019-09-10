package govern

import (
	"io"
	"sync"
)

type Discoveryer interface {
	Discovery(name string, filters ...FilterOption) (Service, error)
}

type discovery struct {
	repo Repo
	once sync.Once
}

func (d *discovery) Discovery(name string, filters ...FilterOption) (Service, error) {
	if d.repo == nil {
		return emptyService, ErrorNilRepo
	}
	filter := Filter{}
	for _, f := range filters {
		f(&filter)
	}
	srv, err := d.repo.GetService(name)
	if err != nil {
		return emptyService, err
	}
	return FilterNodes(srv, filter), nil
}

func (d *discovery) Close() error {
	d.once.Do(func() {
		if d.repo != nil {
			if r, ok := d.repo.(io.Closer); ok {
				r.Close()
			}
		}
	})
	return nil
}

type Filter struct {
	Version  string
	Ssl      bool
	Protocol ServiceProtocol
}

type FilterOption func(f *Filter)

func OnlyVersion(version string) FilterOption {
	return func(f *Filter) {
		f.Version = version
	}
}

func GreaterThanVersion(version string) FilterOption {
	return func(f *Filter) {
		f.Version = ">" + version
	}
}

func GreaterOrEqualVersion(version string) FilterOption {
	return func(f *Filter) {
		f.Version = ">=" + version
	}
}
func LessThanVersion(version string) FilterOption {
	return func(f *Filter) {
		f.Version = "<" + version
	}
}

func LessOrEqualVersion(version string) FilterOption {
	return func(f *Filter) {
		f.Version = "<=" + version
	}
}

func Secure() FilterOption {
	return func(f *Filter) {
		f.Ssl = true
	}
}

func Insecure() FilterOption {
	return func(f *Filter) {
		f.Ssl = false
	}
}

func OnlyProtocol(protocol ServiceProtocol) FilterOption {
	return func(f *Filter) {
		f.Protocol = protocol
	}
}

func FilterNodes(srv Service, filter Filter) Service {
	var filterNodes []ServiceNode
	for _, node := range srv.Nodes {
		if node.SSL != filter.Ssl {
			continue
		}
		if node.Protocol != filter.Protocol {
			continue
		}
		if !matchVersion(node.Version, filter.Version) {
			continue
		}
		filterNodes = append(filterNodes, node)
	}
	return Service{
		Name:  srv.Name,
		Nodes: filterNodes,
	}
}

func matchVersion(current, target string) bool {
	ver1 := ConvertVersion(current)
	var ver2 Version
	var symbol string
	switch target[0:2] {
	case ">=":
		ver2 = ConvertVersion(target[2:])
		symbol = ">="
	case "<=":
		ver2 = ConvertVersion(target[2:])
		symbol = "<="
	default:
		switch target[0:1] {
		case ">":
			ver2 = ConvertVersion(target[0:1])
			symbol = ">"
		case "<":
			ver2 = ConvertVersion(target[0:1])
			symbol = "<"
		default:
			ver2 = ConvertVersion(target)
			symbol = "="
		}
	}
	if ver1 == emptyVersion || ver2 == emptyVersion {
		return false
	}
	switch symbol {
	case ">":
		return ver1.GreaterThan(ver2)
	case ">=":
		return ver1.GreaterOrEqualThan(ver2)
	case "=":
		return ver1.Equal(ver2)
	case "<=":
		return ver1.LessOrEqualThan(ver2)
	case "<":
		return ver1.LessThan(ver2)
	}
	return current == target
}

func NewDiscovery(repo Repo) Discoveryer {
	return &discovery{repo: repo}
}
