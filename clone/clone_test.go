package clone_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/qulia/go-qulia/v2/clone"
	"github.com/stretchr/testify/assert"
)

func TestStructlone(t *testing.T) {
	cl := clone.NewCloner[ComprehensiveType]()
	i := 42
	comprehensiveObject := ComprehensiveType{
		IntField:    10,
		FloatField:  20.22,
		StringField: "Test String",
		BoolField:   true,

		IntPtr:      &i,
		InnerStruct: &InnerStruct{Field1: 5, Field2: "Inner String"},

		SliceField:       []int{1, 2, 3, 4},
		SliceStructField: []InnerStruct{{Field1: 1, Field2: "Slice Inner 1"}, {Field1: 2, Field2: "Slice Inner 2"}},

		ArrayField: [4]int{9, 8, 7, 6},
	}

	comprehensiveObjectClone, err := cl.Clone(comprehensiveObject)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(comprehensiveObject, comprehensiveObjectClone))
}

func TestInterfaceCone(t *testing.T) {
	cl := clone.NewInterfaceCloner[Animal, Dog]()
	dog := Dog{Name: "Buddy"}
	dogClone, err := cl.Clone(dog)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(dog, dogClone))

	dogClone2, err := cl.Clone(&dog)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(dog, dogClone2))

	var gen Animal = &dog
	dogClone3, err := cl.Clone(gen)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(dog, dogClone3))
}

func TestCustomMarshalers(t *testing.T) {
	cl := clone.NewCloner[Vector]()
	v1 := Vector{3, 4, 5}
	v2, err := cl.Clone(v1)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(v1, v2))
}

func TestAttemptClonerWithInterface(t *testing.T) {
	cl := clone.NewCloner[Animal]()
	cat := Cat{Name: "Buddy"}
	_, err := cl.Clone(cat)
	assert.NotNil(t, err)
}

type InnerStruct struct {
	Field1 int
	Field2 string
}

type ComprehensiveType struct {
	// Primitive type
	IntField    int
	FloatField  float64
	StringField string
	BoolField   bool

	// Pointer types
	IntPtr      *int
	InnerStruct *InnerStruct

	// Slice types
	SliceField       []int
	SliceStructField []InnerStruct

	// Array types
	ArrayField [4]int
}

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

type Cat struct {
	Name string
}

func (d Cat) Speak() string {
	return "Meow! My name is " + d.Name
}

type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

func (v *Vector) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}
