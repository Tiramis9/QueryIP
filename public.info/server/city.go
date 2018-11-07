package server

//Structural information，matching ip ok returns to client
type CityANDCountry struct {
	Addr         string `json:"addressing"`
	Continent    string `json:"continent"`
	Country      string `json:"country"`
	Subdivisions string `json:"province"`
	Locationcity string `json:"city"`
	TimeZone     string `json:"timezone"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}
