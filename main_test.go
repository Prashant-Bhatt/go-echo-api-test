package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

type test struct {
	name    string
	wantErr bool
}

var tt = []test{
	{
		name: "Coforge Ltd",
	},
	{
		name: "Capegemini",
	},
}

func Test_createOrganisation(t *testing.T) {

	for _, v := range tt {

		paramName := map[string]string{
			"name": v.name,
		}
		apitest.New().
			Handler(newApp()).
			Post("/orgs").
			JSON(paramName). // pass json param
			Expect(t).
			Status(http.StatusCreated).
			End()
	}
}

func Test_getAllOrganisation(t *testing.T) {

	v := &Organisation{
		Id:   "ss1234",
		Name: "CoforgeLtd",
	}
	orgMap["ss1234"] = *v
	b, _ := json.Marshal(orgMap)
	apitest.New().
		Handler(newApp()).
		Get("/orgs").
		Expect(t).
		Body(string(b)).
		Status(http.StatusOK).
		End()

}
