// This file was auto-generated by Fern from our API Definition.

package notification

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
)

type Notification struct {
	Id      string `json:"id"`
	Message string `json:"message"`

	_rawJSON json.RawMessage
}

func (n *Notification) UnmarshalJSON(data []byte) error {
	type unmarshaler Notification
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = Notification(value)
	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *Notification) String() string {
	if len(n._rawJSON) > 0 {
		if value, err := core.StringifyJSON(n._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}
