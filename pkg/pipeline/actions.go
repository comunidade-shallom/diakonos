package pipeline

import (
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
)

type ActionParams map[string]string

type ActionSource struct {
	Source
	Action string `yml:"action"`
}

type ActionDefinition struct {
	ID     string       `yml:"id"`
	Type   Action       `yml:"type"`
	Params ActionParams `yml:"params"`
	Source ActionSource `yml:"source"`
}

var ErrMissingAction = errors.Business("missing action value", "DP:001")

func (e *ActionSource) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// try to get raw value
	var val string

	err := unmarshal(&val)

	if err == nil {
		e.Value = val

		return nil
	}

	// extract from map (to be used in nested actions)
	var data map[string]string

	err = unmarshal(&data)

	if err != nil {
		return err
	}

	action := data["action"]

	if action == "" {
		return ErrMissingAction
	}

	e.Action = action

	return nil
}
