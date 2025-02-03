package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// interface{} có thể chứa bất kỳ kiểu dữ liệu nào.
// Khi dùng json.Marshal, Go tự động xử lý kiểu dữ liệu bên trong interface{} để chuyển thành JSON đúng định dạng.

//success response

func SuccessResponse(c *gin.Context, code int, data interface{}) {

	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})

}

// error response
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
