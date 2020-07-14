package copy

import "encoding/json"

func DeepCopy(t, s interface{}) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, t)
	return err
}
