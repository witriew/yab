// Code generated by thriftrw v0.4.0
// @generated

package foo

import (
	"errors"
	"fmt"
	"github.com/yarpc/yab/testdata/yarpc/integration"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type BarArgs struct {
	Arg *int32 `json:"arg,omitempty"`
}

func (v *BarArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg != nil {
		w, err = wire.NewValueI32(*(v.Arg)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *BarArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.Arg = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *BarArgs) String() string {
	var fields [1]string
	i := 0
	if v.Arg != nil {
		fields[i] = fmt.Sprintf("Arg: %v", *(v.Arg))
		i++
	}
	return fmt.Sprintf("BarArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *BarArgs) MethodName() string {
	return "bar"
}

func (v *BarArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var BarHelper = struct {
	Args           func(arg *int32) *BarArgs
	IsException    func(error) bool
	WrapResponse   func(int32, error) (*BarResult, error)
	UnwrapResponse func(*BarResult) (int32, error)
}{}

func init() {
	BarHelper.Args = func(arg *int32) *BarArgs {
		return &BarArgs{Arg: arg}
	}
	BarHelper.IsException = func(err error) bool {
		switch err.(type) {
		case *integration.NotFound:
			return true
		default:
			return false
		}
	}
	BarHelper.WrapResponse = func(success int32, err error) (*BarResult, error) {
		if err == nil {
			return &BarResult{Success: &success}, nil
		}
		switch e := err.(type) {
		case *integration.NotFound:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for BarResult.NotFound")
			}
			return &BarResult{NotFound: e}, nil
		}
		return nil, err
	}
	BarHelper.UnwrapResponse = func(result *BarResult) (success int32, err error) {
		if result.NotFound != nil {
			err = result.NotFound
			return
		}
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type BarResult struct {
	Success  *int32                `json:"success,omitempty"`
	NotFound *integration.NotFound `json:"notFound,omitempty"`
}

func (v *BarResult) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueI32(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.NotFound != nil {
		w, err = v.NotFound.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("BarResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _NotFound_Read(w wire.Value) (*integration.NotFound, error) {
	var v integration.NotFound
	err := v.FromWire(w)
	return &v, err
}

func (v *BarResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.NotFound, err = _NotFound_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.NotFound != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("BarResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *BarResult) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.NotFound != nil {
		fields[i] = fmt.Sprintf("NotFound: %v", v.NotFound)
		i++
	}
	return fmt.Sprintf("BarResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *BarResult) MethodName() string {
	return "bar"
}

func (v *BarResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
