# ipfilter

A golang [cuckoo filter](https://www.cs.cmu.edu/~dga/papers/cuckoo-conext2014.pdf) [helper library](https://github.com/seiflotfy/cuckoofilter) for use with IP addresses.

## example
```
router := net.IPv4(192,168,1,1)

filter := ipfilter.NewFilter()
filter.Insert(router)

filter.Lookup(router) // true
filter.Lookup(net.IPv4(8,8,8,8)) // false
```