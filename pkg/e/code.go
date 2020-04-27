package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_DATA       = 10001
	ERROR_NOT_EXIST_DATA   = 10002
	ERROR_GET_DATA_FAIL    = 10003
	ERROR_CHECK_DATA_FAIL  = 10004
	ERROR_ADD_DATA_FAIL    = 10005
	ERROR_EDIT_DATA_FAIL   = 10006
	ERROR_DELETE_DATA_FAIL = 10006

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003

	ERROR_UESR_REGISTER_FAIL = 40000
	ERROR_UESR_LOGIN_FAIL    = 40001
	ERROR_USER_NOT_EXIST     = 40002
	ERROR_USER_GET_FAIL      = 40003
	ERROR_USERS_GET_FAIL     = 40004
)
