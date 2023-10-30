// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "acme.io/sdk/core"
	json "encoding/json"
	fmt "fmt"
)

type SetNameRequest struct {
	UserName string `json:"userName"`
}

type SetNameRequestV3 struct {
	XEndpointHeader string                `json:"-"`
	Body            *SetNameRequestV3Body `json:"-"`
}

func (s *SetNameRequestV3) UnmarshalJSON(data []byte) error {
	body := new(SetNameRequestV3Body)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.Body = body
	return nil
}

func (s *SetNameRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

type SetNameRequestV3Optional struct {
	XEndpointHeader string                `json:"-"`
	Body            *SetNameRequestV3Body `json:"-"`
}

func (s *SetNameRequestV3Optional) UnmarshalJSON(data []byte) error {
	body := new(SetNameRequestV3Body)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.Body = body
	return nil
}

func (s *SetNameRequestV3Optional) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

type SetNameRequestV4 struct {
	XEndpointHeader string   `json:"-"`
	Body            []string `json:"-"`
}

func (s *SetNameRequestV4) UnmarshalJSON(data []byte) error {
	var body []string
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	s.Body = body
	return nil
}

func (s *SetNameRequestV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Body)
}

type SetNameRequestV5 struct {
	XEndpointHeader string `json:"-"`
	Body            string `json:"-"`
}

func (s *SetNameRequestV5) UnmarshalJSON(data []byte) error {
	var body string
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	if body != "fern" {
		return fmt.Errorf("expected literal %q, but found %q", "fern", body)
	}
	s.Body = body
	return nil
}

func (s *SetNameRequestV5) MarshalJSON() ([]byte, error) {
	return json.Marshal("fern")
}

type Filter struct {
	Tag string `json:"tag"`

	_rawJSON json.RawMessage
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	type unmarshaler Filter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = Filter(value)
	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *Filter) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type SetNameRequestV3Body struct {
	UserName string `json:"userName"`

	_rawJSON json.RawMessage
}

func (s *SetNameRequestV3Body) UnmarshalJSON(data []byte) error {
	type unmarshaler SetNameRequestV3Body
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SetNameRequestV3Body(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SetNameRequestV3Body) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type Union struct {
	Type string
	Foo  *Foo
	Bar  *Bar
}

func NewUnionFromFoo(value *Foo) *Union {
	return &Union{Type: "foo", Foo: value}
}

func NewUnionFromBar(value *Bar) *Union {
	return &Union{Type: "bar", Bar: value}
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
	case "foo":
		value := new(Foo)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Foo = value
	case "bar":
		value := new(Bar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Bar = value
	}
	return nil
}

func (u Union) MarshalJSON() ([]byte, error) {
	switch u.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.Type, u)
	case "foo":
		var marshaler = struct {
			Type string `json:"type"`
			*Foo
		}{
			Type: u.Type,
			Foo:  u.Foo,
		}
		return json.Marshal(marshaler)
	case "bar":
		var marshaler = struct {
			Type string `json:"type"`
			*Bar
		}{
			Type: u.Type,
			Bar:  u.Bar,
		}
		return json.Marshal(marshaler)
	}
}

type UnionVisitor interface {
	VisitFoo(*Foo) error
	VisitBar(*Bar) error
}

func (u *Union) Accept(visitor UnionVisitor) error {
	switch u.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", u.Type, u)
	case "foo":
		return visitor.VisitFoo(u.Foo)
	case "bar":
		return visitor.VisitBar(u.Bar)
	}
}

type UpdateRequest struct {
	Tag            string                   `json:"-"`
	Extra          *string                  `json:"-"`
	Union          *Union                   `json:"union,omitempty"`
	Filter         *Filter                  `json:"filter,omitempty"`
	OptionalUnion  *core.Optional[Union]    `json:"optionalUnion,omitempty"`
	OptionalFilter *core.Optional[Filter]   `json:"optionalFilter,omitempty"`
	OptionalTags   *core.Optional[[]string] `json:"optionalTags,omitempty"`
}
