package ipfilter

import "net"
import "testing"
import "github.com/stretchr/testify/assert"

func Test_Basic(t *testing.T) {

	router := net.IPv4(192,168,1,1)

	filter := NewFilter()
	filter.Insert(router)

	assert.Equal(t, filter.Lookup(router), true)
	assert.Equal(t, filter.Lookup(net.IPv4(8,8,8,8)), false)
}