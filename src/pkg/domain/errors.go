package domain

type NoFitPositionError struct{}

func (e *NoFitPositionError) Error() string {
	return "no fit position found"
}
