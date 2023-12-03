package constants

type ErrorMessage string

const (
	ErrorNotFound           ErrorMessage = "element not found"
	ErrorBadRequest         ErrorMessage = "bad request"
	ErrorSomethingWentWrong ErrorMessage = "something went wrong"
)
