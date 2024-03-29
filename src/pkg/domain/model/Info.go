package domain

// Info is configurable information usually set at build time with ldflags.
type Info struct {
	Version     string
	BuildDate   string
	Description string
	CommitHash  string
	CommitDate  string
}
