package jointmap

import (
	"fmt"
)

type MapIterator[K comparable, V any] interface {
	Next() bool
	Value() V
	Key() K
}

// type Map[K comparable, V any] interface {
// 	Get(key K) (value V, ok bool)
// 	Iterator() MapIterator[K, V]
// 	Len() int
// }

type Map[K comparable, V any, I MapIterator[K, V]] interface {
	Get(key K) (value V, ok bool)
	Iterator() I
	Len() int
}

type JointMap[K1, K2 comparable, V any, I1 MapIterator[K1, K2], I2 MapIterator[K2, V]] struct {
	AddrMap Map[K1, K2, I1]
	DataMap Map[K2, V, I2]
}

func New[K1, K2 comparable, V any, I1 MapIterator[K1, K2], I2 MapIterator[K2, V]](
	a Map[K1, K2, I1],
	b Map[K2, V, I2]) Map[K1, V, MapIterator[K1, V]] {
	return &JointMap[K1, K2, V, I1, I2]{
		AddrMap: a,
		DataMap: b,
	}
}

func (bm *JointMap[K1, K2, V, I1, I2]) Len() int {
	return bm.AddrMap.Len()
}

type jointMapIterator[K1, K2 comparable, V any, I1 MapIterator[K1, K2], I2 MapIterator[K2, V]] struct {
	AddrIterator MapIterator[K1, K2]
	DataMap      Map[K2, V, I2]
}

func (bm JointMap[K1, K2, V, I1, I2]) Iterator() MapIterator[K1, V] {
	return jointMapIterator[K1, K2, V, I1, I2]{
		AddrIterator: bm.AddrMap.Iterator(),
		DataMap:      bm.DataMap,
	}
}

func (bm JointMap[K1, K2, V, I1, I2]) Get(key K1) (value V, ok bool) {
	v, ok := bm.AddrMap.Get(key)
	if ok {
		return bm.DataMap.Get(v)
	}
	return
}

func (it jointMapIterator[K1, K2, V, I1, I2]) Value() V {
	k2 := it.AddrIterator.Value()
	v, ok := it.DataMap.Get(k2)
	if !ok {
		panic(fmt.Errorf("invalid key! can not find key:%v in DataMap", k2))
	}
	return v
}

func (it jointMapIterator[K1, K2, V, I1, I2]) Next() bool {
	return it.AddrIterator.Next()
}

func (it jointMapIterator[K1, K2, V, I1, I2]) Key() K1 {
	return it.AddrIterator.Key()
}
