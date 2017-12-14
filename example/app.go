package app

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/anekonnect/decorator"
)

type User struct {
	Id      int64  `json:"_id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

func init() {

	var GetAll = decorator.Decorate(
		decorator.HandlerFunc(getUser),
		decorator.Search([] string{
			"Name",
		}),
		decorator.Paginate("Company"),
		decorator.PublicAuth(""),
	)

	router := httprouter.New()
	router.GET("/getuser", decorator.ToHandle(GetAll))

	http.Handle("/", router)
}

func getUser(r *http.Request, ps httprouter.Params, username string) (interface{}, *decorator.ServerError) {
	//
	//    resp, err: = http.Get("http://www.json-generator.com/api/json/get/cpIuGuaIya?indent=2");
	//
	//    if err != nil {
	//        return nil, err
	//    }
	//
	//    defer resp.Body.Close()
	//
	//    if err: = json.NewDecoder(resp.Body).Decode( & User);
	//    err != nil {
	//        return nil, err
	//    }
	//
	return nil, nil
}
