package domain

type NoFitPositionError struct{}

func (e *NoFitPositionError) Error() string {
	return "no fit position found"
}

type PlacementFailedError struct{}

func (e *PlacementFailedError) Error() string {
	return "placement failed"
}
