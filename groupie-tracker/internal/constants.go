package internal
const (
	root      = "/"
	trackerURL = "/tracker/"
	searchUrl = "/search/"

	style = "/style/"

	OK                  = "OK"
	MethodNotAllowed    = "Method Not Allowed"
	BadRequest          = "Bad Request"
	InternalServerError = "Internal Server Error"
	NotFound            = "Not Found"
)
const (
	ApiArtists   = "https://groupietrackers.herokuapp.com/api/artists"
	ApiRelation  = "https://groupietrackers.herokuapp.com/api/relation"
)
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    map[string][]string
}

type DateLocation struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}

type Location struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Result struct {
	Singer   Artist
	Relation Location
	Text     string
	Type     string
}


