package schema

var _document = `
type City {
	id: ID!
	name: String! @search(by: [exact])
	lat: Float!
	lng: Float!
	places: [Place] @hasInverse(field: city)
	advisory: Advisory
	weather: Weather
}
`
