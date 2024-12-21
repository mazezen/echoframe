package middles

import (
	"github.com/mazezen/itools"
	"github.com/spf13/cast"
	"net/http"
)

// defaultHandlerFunc
type defaultHandlerFunc func(c echo.Context)

var (
	res = &itools.Response
	rl  = &itools.ResponsePage
)

// HandlerMiddleware Handling error and success response formats in controllers
func HandlerMiddleware(next defaultHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// recover panic
		defer func() {
			if r := recover(); r != nil {
				// 统一错误响应
				_ = c.JSON(http.StatusOK, (*res).ResponseError("", 5000, r.(string), nil))
			}
		}()

		next(c)

		// success response
		if !c.Response().Committed {
			content := c.Get("response")
			count := c.Get("count")
			total := cast.ToInt(count)
			if count != nil {
				_ = c.JSON(http.StatusOK, (*res).ResponseSuccess("", "success", (*rl).Pagination(total, content)))
			} else {
				_ = c.JSON(http.StatusOK, (*res).ResponseSuccess("", "success", content))
			}
		}
		return nil
	}
}
