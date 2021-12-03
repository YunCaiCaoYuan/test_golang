package main

import (
	"errors"
	"fmt"
)

// ============ example 4 : 反转控制 ============

// 控制逻辑
type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

// 业务逻辑
type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func main() {
	obj := NewIntSet()
	obj.Add(1)
	fmt.Println("obj.data1:", obj.data)
	obj.Undo()
	fmt.Println("obj.data2:", obj.data)
}

// ============ example 4 ============
/*
type IntSet struct {
	data map[int]bool
}
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}
func (set *IntSet) Add(x int) {
	fmt.Println("IntSet-Add")
	set.data[x] = true
}
func (set *IntSet) Delete(x int) {
	fmt.Println("IntSet-Delete")
	delete(set.data, x)
}
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}


type UndoableIntSet struct { // Poor style
	IntSet    // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}


func (set *UndoableIntSet) Add(x int) { // Override
	fmt.Println("UndoableIntSet-Add")
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}


func (set *UndoableIntSet) Delete(x int) { // Override
	fmt.Println("UndoableIntSet-Delete")
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}

func main() {
	obj := NewUndoableIntSet()
	obj.Add(1)
	obj.Undo()
}*/

// ============ example 3 ============
/*
type IntSet struct {
	data map[int]bool
}
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}
func (set *IntSet) Add(x int) {
	set.data[x] = true
}
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func main() {
	obj := NewIntSet()
	obj.Add(1)
	fmt.Println("obj:", obj.data)
	fmt.Println("obj.Contains(1):", obj.Contains(1))
	obj.Delete(1)
	fmt.Println("obj:", obj.data)
}
*/

// ============ example 2 : 方法重写 ============
/*
type Widget struct {
	X, Y int
}
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}

func NewButton(x,y int, text string) Button {
	return Button{Label{Widget{x, y}, text}}
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}


type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}


func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中重载这个接口方法
func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}


func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

func main() {
	//obj := &Button{Label{Text: "hello"}}
	//obj.Paint()

	label := Label{Widget{10, 10}, "State:"}
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}
	button1 := Button{Label{Widget{10, 70}, "OK"}}
	//button2 := NewButton(50, 70, "Cancel")

	for _, painter := range []Painter{label, listBox, button1} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}*/

/* ============ example 1 : 结构体嵌入（委托） ============
type Widget struct {
	X, Y int
}
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

func main() {
	label := Label{Widget{10, 10}, "State:"}
	fmt.Println("label1:", label)
	label.X = 11
	label.Y = 12
	fmt.Println("label2:", label)
}
*/
