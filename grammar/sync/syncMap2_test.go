package sync

//import (
//	"github.com/YunCaiCaoYuan/test_golang/grammar/sync/conMap"
//	"reflect"
//	"strconv"
//	"sync"
//	"testing"
//)
//
//func BenchmarkMap(b *testing.B) {
//	rwLock := sync.RWMutex{}
//	mapA := make(map[int]int)
//	for i := 0; i < b.N; i++ {
//		rwLock.Lock()
//		mapA[i] = i
//		rwLock.Unlock()
//		rwLock.RLock()
//		_ = mapA[i]
//		rwLock.RUnlock()
//	}
//}
//
//func BenchmarkSyncMap(b *testing.B) {
//	mapB := sync.Map{}
//	for i := 0; i < b.N; i++ {
//		mapB.Store(i, i)
//		iX, ok := mapB.Load(i)
//		if ok {
//			_ = iX.(int)
//		}
//	}
//}
//
//func BenchmarkMySyncMap(b *testing.B) {
//	myMap := conMap.ConcurrentMap{
//		KeyType :reflect.TypeOf(123),
//		ValueType : reflect.TypeOf(123),
//	}
//	for i := 0; i < b.N; i++ {
//		myMap.Store(i, i)
//		iX, ok := myMap.Load(i)
//		if ok {
//			_ = iX.(int)
//		}
//	}
//}
//
//func BenchmarkMySyncMap2(b *testing.B) {
//	myMap := conMap.ConcurrentMap{
//		KeyType :reflect.TypeOf(123),
//		ValueType : reflect.TypeOf("123"),
//	}
//	for i := 0; i < b.N; i++ {
//		myMap.Store(i, strconv.Itoa(i))
//		iX, ok := myMap.Load(i)
//		if ok {
//			_ = iX.(string)
//		}
//	}
//}
