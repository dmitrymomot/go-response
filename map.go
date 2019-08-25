package response

// M is a helper to use more short syntax
// e.g.: response.JSON(w, 200, response.M{"data": "some data value"})
type M map[string]interface{}
