package helper

import "net/http"

func AuthOK(id int, token string) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"message": "success",
		"code":    200,
		"id":      id,
		"token":   token,
	}
}

func SuccessCreateNoData() (int, interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"message": "success",
		"code":    201,
	}
}
func SuccessGetData(data interface{}) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
		"code":    200,
	}
}

func FailedBadRequest() (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"message": "bad request",
		"code":    400,
	}
}

func FailedBadRequestWithMSG(msg string) (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"message": msg,
		"code":    400,
	}
}

func FailedNotFound() (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"message": "data not found",
		"code":    404,
	}
}
