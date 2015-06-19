package typegraph

import (
	"reflect"
)

//
//
type Types []*Type

//
//
type Type struct {
	Name     string
	Desc     string
	Proto    interface{}
	Parent   *Type
	Children map[string]Types
}

//
//
type Children map[string]Types

//
//
func (t *Type) Child(v string) *Type {
	if t.Name == v {
		return t
	}
	for _, ts := range t.Children {
		for _, t := range ts {
			if c := t.Child(v); c != nil {
				return c
			}
		}
	}
	return nil
}

//
//
func (t *Type) Parents() Types {
	ts := Types{}
	for {
		if t == nil {
			break
		}
		ts = append(ts, t)
		t = t.Parent
	}
	// reverse for the client
	for i, j := 0, len(ts)-1; i < j; i, j = i+1, j-1 {
		ts[i], ts[j] = ts[j], ts[i]
	}
	return ts
}

//
//
func (t *Type) Walk(cb func(*Type)) Types {
	if cb != nil {
		cb(t)
	}
	ts := Types{t}
	for _, cts := range t.Children {
		for _, ct := range cts {
			ts = append(ts, ct.Walk(cb)...)
		}
	}
	return ts
}

// NewInstance creates and returns a copy of the type's prototype.
//
func (t *Type) NewInstance() interface{} {
	return reflect.New(reflect.TypeOf(t.Proto)).Interface()
}

//
//
func (t *Type) NewSlice(n int) []interface{} {
	s := reflect.MakeSlice(
		reflect.SliceOf(reflect.TypeOf(t.Proto)), 0, n).Interface()
	if i, ok := s.([]interface{}); ok {
		return i
	} else {
		return []interface{}{}
	}
}
