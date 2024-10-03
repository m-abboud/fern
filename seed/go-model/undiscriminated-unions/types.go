// This file was auto-generated by Fern from our API Definition.

package undiscriminatedunions

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/undiscriminated-unions/fern/core"
)

type Key struct {
	KeyType              KeyType
	defaultStringLiteral string
}

func NewKeyFromKeyType(value KeyType) *Key {
	return &Key{KeyType: value}
}

func NewKeyWithDefaultStringLiteral() *Key {
	return &Key{defaultStringLiteral: "default"}
}

func (k *Key) DefaultStringLiteral() string {
	return k.defaultStringLiteral
}

func (k *Key) UnmarshalJSON(data []byte) error {
	var valueKeyType KeyType
	if err := json.Unmarshal(data, &valueKeyType); err == nil {
		k.KeyType = valueKeyType
		return nil
	}
	var valueDefaultStringLiteral string
	if err := json.Unmarshal(data, &valueDefaultStringLiteral); err == nil {
		k.defaultStringLiteral = valueDefaultStringLiteral
		if k.defaultStringLiteral != "default" {
			return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", k, "default", valueDefaultStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k Key) MarshalJSON() ([]byte, error) {
	if k.KeyType != "" {
		return json.Marshal(k.KeyType)
	}
	if k.defaultStringLiteral != "" {
		return json.Marshal("default")
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KeyVisitor interface {
	VisitKeyType(KeyType) error
	VisitDefaultStringLiteral(string) error
}

func (k *Key) Accept(visitor KeyVisitor) error {
	if k.KeyType != "" {
		return visitor.VisitKeyType(k.KeyType)
	}
	if k.defaultStringLiteral != "" {
		return visitor.VisitDefaultStringLiteral(k.defaultStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KeyType string

const (
	KeyTypeName  KeyType = "name"
	KeyTypeValue KeyType = "value"
)

func NewKeyTypeFromString(s string) (KeyType, error) {
	switch s {
	case "name":
		return KeyTypeName, nil
	case "value":
		return KeyTypeValue, nil
	}
	var t KeyType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (k KeyType) Ptr() *KeyType {
	return &k
}

type TypeWithOptionalUnion struct {
	MyUnion *MyUnion `json:"myUnion,omitempty" url:"myUnion,omitempty"`

	extraProperties map[string]interface{}
}

func (t *TypeWithOptionalUnion) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TypeWithOptionalUnion) UnmarshalJSON(data []byte) error {
	type unmarshaler TypeWithOptionalUnion
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TypeWithOptionalUnion(value)

	extraProperties, err := core.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties

	return nil
}

func (t *TypeWithOptionalUnion) String() string {
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
