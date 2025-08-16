package infra

import "errors"

var (
	MaxSizeAchievedErr = errors.New("max size achieved")
	RecordNotFoundErr  = errors.New("record not found")
	RecordExistsErr    = errors.New("record already exists")
)
