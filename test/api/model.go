package api

type errors []struct {
	Message   string
	Locations []struct {
		Line   int
		Column int
	}
}
