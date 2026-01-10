package models

type Result struct {
	Value any
	Err   error
}

func Ok(v any) Result {
	return Result{Value: v}
}

func Err(e error) Result {
	return Result{Err: e}
}

func (r Result) IsOk() bool {
	return r.Err == nil
}
