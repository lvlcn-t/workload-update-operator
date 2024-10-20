package v1

import (
	"encoding/json"
	"time"

	"gopkg.in/yaml.v3"
)

// timeFormats is a list of formats that are supported by the [Time] type.
var timeFormats = []string{
	time.TimeOnly,
	time.Kitchen,
	"15:04:05Z07:00",
}

// Time is a wrapper around the [time.Time] type that supports JSON and YAML serialization.
//
// It supports the following formats:
//   - [time.TimeOnly]
//   - [time.Kitchen]
//   - "15:04:05Z07:00"
type Time struct {
	time.Time
	format string
}

// String returns the formatted time string.
func (t *Time) String() string {
	return t.Time.Format(t.format)
}

var _ json.Marshaler = (*Time)(nil)

// MarshalJSON implements the [json.Marshaler] interface.
func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

var _ json.Unmarshaler = (*Time)(nil)

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (t *Time) UnmarshalJSON(data []byte) error {
	var raw string
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	return t.parseFrom(raw)
}

var _ yaml.Marshaler = (*Time)(nil)

// MarshalYAML implements the [yaml.Marshaler] interface.
func (t *Time) MarshalYAML() (any, error) {
	return t.String(), nil
}

var _ yaml.Unmarshaler = (*Time)(nil)

// UnmarshalYAML implements the [yaml.Unmarshaler] interface.
func (t *Time) UnmarshalYAML(node *yaml.Node) error {
	var raw string
	if err := node.Decode(&raw); err != nil {
		return err
	}

	return t.parseFrom(raw)
}

// parseFrom parses the raw string into the [Time] type.
func (t *Time) parseFrom(raw string) error {
	var err error
	for _, format := range timeFormats {
		t.Time, err = time.Parse(format, raw)
		if err == nil {
			return nil
		}
	}
	return err
}
