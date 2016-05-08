package person

import (
	"code.olipicus.com/go_rest_api/api/rest"
)

const collection string = "person"

// HandlerPerson struct
type HandlerPerson struct {
	rest.REST
}

//Handler ...
var Handler HandlerPerson = HandlerPerson{
	rest.REST{
		Collection: collection,
		OBJ:        Person{},
	},
}
