package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleParamEmployerGid(c *gin.Context) (string, error) {

	q, ok := c.GetQuery(EMPLOYER_GID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYER_ID,
		})
		return "", errors.New(REQUEST_MISSING_PARAM)
	}
	return q, nil
}

func handleParamEmployerPw(c *gin.Context) (string, error) {
	p, ok := c.GetQuery(EMPLOYER_PW)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYER_PW,
		})
		return "", errors.New(REQUEST_MISSING_PARAM)
	}
	return p, nil
}

func handleParamEmployerId(c *gin.Context) (string, error) {
	i, ok := c.GetQuery(EMPLOYER_ID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_MISSING_PARAM + " " + EMPLOYER_GID,
		})
		return "", errors.New(REQUEST_MISSING_PARAM)
	}
	return i, nil
}

func handleQueryParam(c *gin.Context, args ...string) (map[string]string, error) {
	result := map[string]string{}

	for _, a := range args {
		i, ok := c.GetQuery(a)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": REQUEST_MISSING_PARAM + " " + a,
			})
			return map[string]string{}, errors.New(REQUEST_MISSING_PARAM)
		}
		result[a] = i
	}
	return result, nil
}
