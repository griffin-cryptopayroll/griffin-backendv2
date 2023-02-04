package common

import (
	"errors"
	"fmt"
	api_base "griffin-dao/api/base"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleOptionalQueryParam
//
//	Args: {
//	 	"necessary_query_argument": true,
//	     "unnecessary_query_argument": false,
//	}
//
// If there's no necessary query argument in `args` response with StatusBadRequest
// returns: map[string]string | should be turned into appropriate type
func HandleOptionalQueryParam(c *gin.Context, args map[string]bool) (map[string]string, error) {
	result := map[string]string{}
	for a, n := range args {
		i, ok := c.GetQuery(a)
		if n && !ok {
			msg := api_base.CommonResponse{
				Status:  false,
				Message: fmt.Sprintf("Request missing param => no %s", a),
			}
			c.JSON(http.StatusBadRequest, msg)
			return nil, errors.New("REQUEST_MISSING_PARAM")
		}
		result[a] = i
	}
	return result, nil
}

func HandleParamTypeCastFloat(args map[string]string, castKeys map[string]bool) (map[string]float64, error) {
	result := map[string]float64{}
	for k, v := range args {
		if ok := castKeys[k]; !ok {
			continue
		}
		// Change value of keys to float64 if `key` in `castKeys`
		castV, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("json key %s not a float %s", k, v))
		}
		result[k] = castV
	}
	return result, nil
}

func HandleParamTypeCastInt(args map[string]string, castKeys map[string]bool) (map[string]int, error) {
	result := map[string]int{}
	for k, v := range args {
		if ok := castKeys[k]; !ok {
			continue
		}
		// Change value of keys to int if `key` in `castKeys`
		castV, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("json key %s not an int %s", k, v))
		}
		result[k] = castV
	}
	return result, nil
}
