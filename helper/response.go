package helper

import "net/http"

func AuthOK(id int, token string) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"message": "success",
		"id":      id,
		"token":   token,
	}
}

func SuccessCreateNoData() (int, interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"message": "success",
	}
}

func FailedBadRequest() (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"message": "bad request",
	}
}

func FailedBadRequestWithMSG(msg string) (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"message": msg,
	}
}

func FailedNotFound() (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"message": "data not found",
	}
}
