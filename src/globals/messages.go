// globals/messages.go
package globals

var (
	SuccessMessage string
	ErrorMessage   string
	InfoMessage    string
)

func GetMessages() (success, error, info string) {
	return SuccessMessage, ErrorMessage, InfoMessage
}

func SetSuccess(text string) {
	SuccessMessage = text
}

func SetError(text string) {
	ErrorMessage = text
}

func SetInfo(text string) {
	InfoMessage = text
}

func ClearMessages() {
	SuccessMessage = ""
	ErrorMessage = ""
	InfoMessage = ""
}
