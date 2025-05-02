package rest_err

import "net/http"

// Objeto para dados contidos nos retornos de requisições
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

// Objeto para causa de retorno de erro
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função para retornar apenas a mensagem de erro.
func (r *RestErr) Error() string {
	return r.Message
}

// Criando um construtor
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// Retorno caso a requisição seja falha
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

// Retorno caso a requisição tenha erro de validação
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// Retorno caso a requisição cause erro interno
func NewInternalServerError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}

// Retorno caso a requisição não tenha rota
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

// Retorno caso a requisição não seja autorizada
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
