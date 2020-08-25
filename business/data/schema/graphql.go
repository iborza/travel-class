package schema

var _document = `type City {
	id: ID!
	name: String! @search(by: [exact])
	lat: Float!
	lng: Float!
	weather: Weather
}

type Weather {
	id: ID!
	city_name: String!
	description: String
	pressure: Int
	temp: Float
	temp_min: Float
	temp_max: Float
}`
