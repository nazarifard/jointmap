## JointMap
Jointmap package joins two maps simply and provides a new virtual map view without taking any extra memory. It is very usefull when we have a two or more big heavy maps and have to have a new map.

# Usage
```go
 package main
    
 import (
    "github.com/nazarifard/gomap"
    "github.com/nazarifard/jointmap"
 )

 func main() {
    a := gomap.New[string, float32]()
	a.Set("pi", 3.14159)
	a.Set("gamma", 0.57721)

	b := gomap.New[float32, rune]()
	b.Set(3.14159, 'π')
	b.Set(0.57721, 'γ')

	ab := jointmap.New(a, b)  //map[string][rune]
    
	for it := ab.Iterator(); it.Next(); {
		print(it.Key(), ":", string(it.Value()), "  ")        
	}
    //output pi:π  gamma:γ
 }
```
