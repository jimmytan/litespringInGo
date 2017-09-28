package model


type (
	Event struct {
		SourceService	string
		DestinationService string
		DestinationUrl	string
		Body	string
		Header map[string]string
	}
)

