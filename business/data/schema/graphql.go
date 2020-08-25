package schema

var _document = `type City {
	id: ID!
	name: String! @search(by: [exact])
	lat: Float!
	lng: Float!
	weather: Weather
	places: [Place] @hasInverse(field: city)
}

type Place {
	id: ID!
	address: String
	avg_user_rating: Float
	category: String @search(by: [exact])
	city: City!
	name: String! @search(by: [exact])
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
