package ujson

import "encoding/json"

// UnmarshalJSON implements custom json unmarshal logic.
func (j *TestUnmarshalJSON) UnmarshalJSON(data []byte) error {
	type unmarshalCtr struct {
		JsonField json.RawMessage `json:"json_field,omitempty"`
	}
	ctr := &unmarshalCtr{}
	if err := json.Unmarshal(data, ctr); err != nil {
		return err
	}
	j.JsonField = string(ctr.JsonField)
	return nil
}

// _ is a type assertion
var _ json.Unmarshaler = ((*TestUnmarshalJSON)(nil))
