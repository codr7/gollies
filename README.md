## gollies

### order
Ordered collections share the same function signature for comparisons.

```
type Order = int

const (
	Lt = Order(-1)
	Eq = Order(0)
	Gt = Order(1)
)

type Compare = func(x, y interface{}) Order

func CompareInt(x, y interface{}) Order {
	xv, yv := x.(int), y.(int)
	
	if xv < yv {
		return Lt
	}

	if xv > yv {
		return Gt
	}

	return Eq	
}
```

### maps

```
type Map interface {
	Add(key interface{}, val interface{}) interface{}
	Remove(key interface{}) interface{}

	Find(key interface{}) interface{}

	Keys() []interface{}
	Values() []interface{}

	Len() int
}
```

#### hash maps
Hash maps simply wrap the regular native implementation.

```
func NewHashMap() *HashMap
```

#### slice maps
Slice maps are implemented as ordered slices of items, each map may be configured with a custom compare function.

```
func NewSliceMap(cmp Compare) *SliceMap
```

### tests & benchmarks
Slice maps break even with hash maps around 1000 items on my machine, being significantly cheaper to initialize is the main reason they're able to compete at all. 

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/codr7/gollies/m/v2
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceMap/SliceMap-12         	    2482	    469806 ns/op
BenchmarkSliceMap/HashMap-12          	    3806	    308998 ns/op
PASS
ok  	github.com/codr7/gollies/m/v2	2.526s
```
