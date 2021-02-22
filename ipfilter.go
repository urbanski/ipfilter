package ipfilter

import (
	"github.com/seiflotfy/cuckoofilter"
	"net"
)

type IPFilter struct {
	filter *cuckoo.Filter
}

func (i *IPFilter) Lookup(addr net.IP) bool {
	return i.filter.Lookup(addr)
}

func (i *IPFilter) Insert(addr net.IP) error {
	i.filter.InsertUnique(addr)
	return nil
}

func NewFilter() *IPFilter {
	return &IPFilter{filter: cuckoo.NewFilter(1000)}
}