// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	bar "github.com/fern-api/fern-go/internal/testdata/model/packages/fixtures/bar"
	core "github.com/fern-api/fern-go/internal/testdata/model/packages/fixtures/core"
)

type Base struct {
	Name string `json:"name"`
}

func (b *Base) String() string {
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type Value struct {
	Name string `json:"name"`
}

func (v *Value) String() string {
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type Foo struct {
	Name  string   `json:"name"`
	Value *Value   `json:"value,omitempty"`
	Bar   *bar.Bar `json:"bar,omitempty"`
}

func (f *Foo) String() string {
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type Union struct {
	Type         string
	Value        *Value
	AnotherValue *Value
	Bar          *bar.Bar
	AnotherBar   *bar.Bar
}

func NewUnionFromValue(value *Value) *Union {
	return &Union{Type: "value", Value: value}
}

func NewUnionFromAnotherValue(value *Value) *Union {
	return &Union{Type: "anotherValue", AnotherValue: value}
}

func NewUnionFromBar(value *bar.Bar) *Union {
	return &Union{Type: "bar", Bar: value}
}

func NewUnionFromAnotherBar(value *bar.Bar) *Union {
	return &Union{Type: "anotherBar", AnotherBar: value}
}

func (u *Union) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "value":
		value := new(Value)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Value = value
	case "anotherValue":
		var valueUnmarshaler struct {
			AnotherValue *Value `json:"anotherValue,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		u.AnotherValue = valueUnmarshaler.AnotherValue
	case "bar":
		value := new(bar.Bar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Bar = value
	case "anotherBar":
		var valueUnmarshaler struct {
			AnotherBar *bar.Bar `json:"anotherBar,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		u.AnotherBar = valueUnmarshaler.AnotherBar
	}
	return nil
}

func (u Union) MarshalJSON() ([]byte, error) {
	switch u.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.Type, u)
	case "value":
		var marshaler = struct {
			Type string `json:"type"`
			*Value
		}{
			Type:  u.Type,
			Value: u.Value,
		}
		return json.Marshal(marshaler)
	case "anotherValue":
		var marshaler = struct {
			Type         string `json:"type"`
			AnotherValue *Value `json:"anotherValue,omitempty"`
		}{
			Type:         u.Type,
			AnotherValue: u.AnotherValue,
		}
		return json.Marshal(marshaler)
	case "bar":
		var marshaler = struct {
			Type string `json:"type"`
			*bar.Bar
		}{
			Type: u.Type,
			Bar:  u.Bar,
		}
		return json.Marshal(marshaler)
	case "anotherBar":
		var marshaler = struct {
			Type       string   `json:"type"`
			AnotherBar *bar.Bar `json:"anotherBar,omitempty"`
		}{
			Type:       u.Type,
			AnotherBar: u.AnotherBar,
		}
		return json.Marshal(marshaler)
	}
}

type UnionVisitor interface {
	VisitValue(*Value) error
	VisitAnotherValue(*Value) error
	VisitBar(*bar.Bar) error
	VisitAnotherBar(*bar.Bar) error
}

func (u *Union) Accept(visitor UnionVisitor) error {
	switch u.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", u.Type, u)
	case "value":
		return visitor.VisitValue(u.Value)
	case "anotherValue":
		return visitor.VisitAnotherValue(u.AnotherValue)
	case "bar":
		return visitor.VisitBar(u.Bar)
	case "anotherBar":
		return visitor.VisitAnotherBar(u.AnotherBar)
	}
}
