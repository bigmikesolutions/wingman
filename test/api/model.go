package api

import "github.com/bigmikesolutions/wingman/graphql/model"

type errors []struct {
	Message   string
	Locations []struct {
		Line   int
		Column int
	}
}

type environmentResponse struct {
	Data struct {
		Environment *model.Environment `json:"environment"`
	} `json:"data"`
	Errors errors `json:"errors"`
}
