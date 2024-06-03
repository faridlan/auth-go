package web

type Enpoint struct {
	URL    string
	Method string
}

func Enpoints() []Enpoint {
	return []Enpoint{
		{
			URL:    "/api/roles",
			Method: "GET",
		},
		{
			URL:    "/api/roles",
			Method: "POST",
		},
	}
}
