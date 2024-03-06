package utils

import (
	"github.com/Improwised/golang-api/models"
	"github.com/Improwised/golang-api/pkg/structs"
)

// swagger:parameters RequestCreateUser
type RequestCreateUser struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqRegisterUser
	}
}

// swagger:parameters RequestCreateTitle
type RequestCreateTitle struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqRegisterTitle
	}
}

// swagger:parameters RequestCreateCredit
type RequestCreateCredit struct {
	// in:path
	TitleId string `json:"titleId"`
	// in:body
	// required: true
	Body struct {
		structs.ReqRegisterCredit
	}
}

// swagger:parameters RequestDeleteCredit
type RequestDeleteCredit struct {
	// in:path
	TitleId string `json:"titleId"`
	// in:path
	CreditId string `json:"creditId"`
}

// swagger:parameters RequestGetCredit
type RequestGetCredit struct {
	// in:path
	TitleId string `json:"titleId"`
	// in:path
	CreditId string `json:"creditId"`
}

// swagger:parameters RequestUpdateTitle
type RequestUpdateTitle struct {
	// in:path
	TitleId string `json:"titleId"`
	// in:body
	// required: true
	Body struct {
		structs.ReqRegisterTitle
	}
}

// swagger:response ResponseCreateUser
type ResponseCreateUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.User
		} `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseTitle
type ResponseTitle struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.Title
		} `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseCredit
type ResponseCredit struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.Credit
		} `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseDeleteCredit
type ResponseDeleteCredit struct {
	// in:body
	Body struct {
		// enum: success
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseListCredit
type ResponseListCredit struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   []struct {
			models.Credit
		} `json:"data"`
	} `json:"body"`
}

// swagger:response ResponseTitleList
type ResponseTitleList struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   []struct {
			models.Title
		} `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestGetUser
type RequestGetUser struct {
	// in:path
	UserId string `json:"userId"`
}

// swagger:parameters RequestGetTitle
type RequestGetTitle struct {
	// in:path
	TitleId string `json:"titleId"`
}

// swagger:parameters RequestDeleteTitle
type RequestDeleteTitle struct {
	// in:path
	TitleId string `json:"titleId"`
}

// swagger:parameters RequestGetCreditTitle
type RequestGetCreditTitle struct {
	// in:path
	TitleId string `json:"titleId"`
}

// swagger:response ResponseGetUser
type ResponseGetUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.User
		} `json:"data"`
	} `json:"body"`
}

// swagger:parameters RequestAuthnUser
type RequestAuthnUser struct {
	// in:body
	// required: true
	Body struct {
		structs.ReqLoginUser
	}
}

// swagger:response ResponseAuthnUser
type ResponseAuthnUser struct {
	// in:body
	Body struct {
		// enum: success
		Status string `json:"status"`
		Data   struct {
			models.User
		} `json:"data"`
	} `json:"body"`
}

////////////////////
// --- GENERIC ---//
////////////////////

// Response is okay
// swagger:response GenericResOk
type ResOK struct {
	// in:body
	Body struct {
		// enum:success
		Status string `json:"status"`
	}
}

// Fail due to user invalid input
// swagger:response GenericResFailBadRequest
type ResFailBadRequest struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to user invalid input
// swagger:response ResForbiddenRequest
type ResForbiddenRequest struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Server understand request but refuse to authorize it
// swagger:response GenericResFailConflict
type ResFailConflict struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to server understand request but unable to process
// swagger:response GenericResFailUnprocessableEntity
type ResFailUnprocessableEntity struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Fail due to resource not exists
// swagger:response GenericResFailNotFound
type ResFailNotFound struct {
	// in: body
	Body struct {
		// enum: fail
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	} `json:"body"`
}

// Unexpected error occurred
// swagger:response GenericResError
type ResError struct {
	// in: body
	Body struct {
		// enum: error
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	} `json:"body"`
}
