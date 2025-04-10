package graphqlclient

import "github.com/bigmikesolutions/wingman/graphql/model"

type Errors []struct {
	Message   string
	Locations []struct {
		Line   int
		Column int
	}
}

type EnvironmentResponse struct {
	Data struct {
		Environment *model.Environment `json:"environment"`
	} `json:"data"`
	Errors Errors `json:"errors"`
}
