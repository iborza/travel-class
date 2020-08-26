package schema

// AdminJWT represents a default ADMIN token to support the training code.
// It's based on the private/public key provided and was generated by using
// the Travel-Auth tooling. The Auth tooling needs this to manage users.
var AdminJWT = `eyJhbGciOiJSUzI1NiIsImtpZCI6IjU0YmIyMTY1LTcxZTEtNDFhNi1hZjNlLTdkYTRhMGUxZTJjMSIsInR5cCI6IkpXVCJ9.eyJBdXRoIjp7IlJPTEUiOiJBRE1JTiJ9LCJleHAiOjE2MjMzNDI3MTQsImlhdCI6MTU5MTgwNjcxNCwiaXNzIjoidHJhdmVsIHByb2plY3QiLCJzdWIiOiIweDUifQ.dxZsiE9WSXBHB-WenJlSK6zqgXs7ykKpQM3BfrTd_WYvfjIo26FhlPxN-Fr_3dR5-U4aMAw61dTNxMMBNPbD4qs8-CnJ0xfSOl8Xa5Y3p-aKpYvTPL_rPZdjcfqTua2t_sOPmZ3d8_VWkKWmdK-42ab751tmXOCrM6kYXoS1_APQwXKfE_q5eBUlTfrIBR29vtrBfWnpN54wR4i-Uk6DalMOduUmUNuZnYGP9ocIU4Ao1RQ8TsZjo6iIsLGM3r86KYypBWsiRAZPMIZjoZAxqhjRBEOaqNUpq6X3vdhQcRYLgh_36_R1QPlhofAaNKrTMvcZNHkBrBsjOB5pwf6IMQ`

// _publicKey represents the public key generated for the project for training
// purposes. This is the same key found in the publickey.pem file in the root
// of the project. A new key can be generated with the Travel-Auth tooling. In
// a production system this data should be embedded using a tool like pakcr from
// the Buffalo project (https://github.com/gobuffalo/packr).
var _publicKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnZ/BW/tuLr0uxZFw1Q5m
P1JpIksU46o+kIaqIXZjSAduma18m+oSgd1L19Fs9otAjfAlkyU8HF1hJNj/PVv8
MY72vhIWv60xBB4caXuLmflAiJEtvxHfw3WtVR9npQqEowcwrsf7MSSfdHwM4S+F
bMmcl/mE9c7DUrYJBUgu1IbdI7vrEoPE65GFafjZQHkPLUX8OaRXOt4rkT6HfYv+
XqaCs6Ie+dt6xL5HiQpO90/89CAJhi2q8AXvhfxqCVVfLxxd3jNJVq2olkCOLJRE
uJ29Bb460yKOAiDigEUobUpmvT6ggUZNrX71yP0GZxQFBhq9j1IRgPVg4CDA0Pw5
FQIDAQAB
-----END RSA PUBLIC KEY-----`

var _document = `enum Role {
	ADMIN
	EMAIL
	MUTATE
	QUERY
}

type User @auth(
	query: { rule: "{$ROLE: { eq: \"ADMIN\" } }" },
	add: { rule: "{$ROLE: { eq: \"ADMIN\" } }" },
    update: { rule: "{$ROLE: { eq: \"ADMIN\" } }" },
    delete: { rule: "{$ROLE: { eq: \"ADMIN\" } }" },
){
	id: ID!
	email: String! @search(by: [exact])
	name: String!
	role: Role!
	password_hash: String!
	date_created: DateTime!
	date_updated: DateTime!
}

type City {
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
}

type UploadFeedResponse @remote {
	country_code: String
	city_name: String
	lat: Float
	lng: Float
	message: String
}

type Query {
	uploadFeed(countryCode: String!, cityName: String!, lat: Float!, lng: Float!): UploadFeedResponse @custom(http:{
		url: "http://service:3000/upload",
		method: "POST",
		body: "{countrycode: $countryCode, cityname: $cityName, lat: $lat, lng: $lng}"
	})
}

# Dgraph.Authorization {"header":"X-Travel-Auth", "namespace":"Auth", "algo":"RS256", "verificationkey":"-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnZ/BW/tuLr0uxZFw1Q5m\nP1JpIksU46o+kIaqIXZjSAduma18m+oSgd1L19Fs9otAjfAlkyU8HF1hJNj/PVv8\nMY72vhIWv60xBB4caXuLmflAiJEtvxHfw3WtVR9npQqEowcwrsf7MSSfdHwM4S+F\nbMmcl/mE9c7DUrYJBUgu1IbdI7vrEoPE65GFafjZQHkPLUX8OaRXOt4rkT6HfYv+\nXqaCs6Ie+dt6xL5HiQpO90/89CAJhi2q8AXvhfxqCVVfLxxd3jNJVq2olkCOLJRE\nuJ29Bb460yKOAiDigEUobUpmvT6ggUZNrX71yP0GZxQFBhq9j1IRgPVg4CDA0Pw5\nFQIDAQAB\n-----END RSA PUBLIC KEY-----"}
`
