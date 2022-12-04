package v0

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// params
// args: true if the query key is necessary
// returns: map[string]string | should be turned into appropriate type
func handleOptionalQueryParam(c *gin.Context, args map[string]bool) (map[string]string, error) {
	result := map[string]string{}
	for a, nec := range args {
		i, ok := c.GetQuery(a)
		if nec && !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": REQUEST_MISSING_PARAM + " " + a,
			})
			return map[string]string{}, errors.New(REQUEST_MISSING_PARAM)
		}
		result[a] = i
	}
	return result, nil
}
