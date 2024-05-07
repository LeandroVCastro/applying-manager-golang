package route_handlers

import (
	"encoding/json"
	"net/http"
)

func GetRequestBody(request *http.Request, body interface{}) error {
	err := json.NewDecoder(request.Body).Decode(body)
	// if err != nil {
	// 	return err
	// }
	// validatedBody := ValidateBody(body)
	// fmt.Println("validatedBody", validatedBody)
	return err
}
