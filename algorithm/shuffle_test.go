package algorithm

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func Test_Shuffle(t *testing.T) {
	list := make([]int, 0)
	for i := 0; i < 10; i++ {
		list = append(list, rand.Intn(10000))
	}
	fmt.Println(list)
	Shuffle(list)
	fmt.Println(list)
	Shuffle(list)
	fmt.Println(list)
}

type People struct {
	Id   int32
	Name string
}

func Shuffle2(vals []*People) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func Test_Shuffle2(t *testing.T) {
	list := make([]*People, 0)
	for i := 1; i <= 10; i++ {
		list = append(list, &People{
			Id:   int32(i),
			Name: strconv.Itoa(i),
		})
	}

	mySlicePrint(list)
	Shuffle2(list)
	mySlicePrint(list)
	Shuffle2(list)
	mySlicePrint(list)
}

func mySlicePrint(list []*People) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("%v ", list[i])
	}
	fmt.Println("")
}

func Test_Reflect(t *testing.T) {
	list := make([]*People, 0)
	for i := 1; i <= 10; i++ {
		list = append(list, &People{
			Id:   int32(i),
			Name: strconv.Itoa(i),
		})
	}
	val := reflect.ValueOf(list)
	fmt.Println(val)
	fmt.Println("val.Len():", val.Len())
	val_1 := val.Index(1)
	fmt.Println("val.Index(1):", val.Index(1))
	val_1.Set(reflect.ValueOf(&People{Id: 11, Name: "11"}))
	fmt.Println(val)
}

func Test_Shuttle(t *testing.T) {
	list := make([]*People, 0)
	for i := 0; i < 10; i++ {
		list = append(list, &People{
			Id:   int32(i),
			Name: strconv.Itoa(i),
		})
	}
	mySlicePrint(list)

	rand.Shuffle(10, func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	mySlicePrint(list)
}
