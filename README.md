# go-echo-api-test
Go server with api test with unit testing
Unit testing with mocking objects

Conatains One Post and one Get function and unit test functions for both of these

e.GET("/orgs", getAllOrganisation)


e.POST("/orgs", createOrganisation)

POST function uses a map to store organisation temporarily.
To achieve unit testing in get function I have mocked the output because of temporary storage in map.


