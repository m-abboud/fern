// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params-complex/fixtures/core"
	uuid "github.com/google/uuid"
	time "time"
)

type GetUsersRequest struct {
	Id               uuid.UUID  `json:"-"`
	Date             time.Time  `json:"-"`
	Deadline         time.Time  `json:"-"`
	Bytes            []byte     `json:"-"`
	OptionalId       *uuid.UUID `json:"-"`
	OptionalDate     *time.Time `json:"-"`
	OptionalDeadline *time.Time `json:"-"`
	OptionalBytes    *[]byte    `json:"-"`
}

type User struct {
	Name string   `json:"name"`
	Tags []string `json:"tags,omitempty"`

	_rawJSON json.RawMessage
}

func (u *User) UnmarshalJSON(data []byte) error {
	type unmarshaler User
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = User(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *User) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
