package server

type ErrCode int

const (
	ErrBookWasNotFound ErrCode = iota
	ErrBooksWereNotFound
	ErrBookWasNotInserted
	ErrBookWasNotParsed
	ErrIdWasNotParsed
	ErrWrongCredentials
)

type ErrCodeName string

const (
	ErrNameBookWasNotFound    ErrCodeName = "BookWasNotFound"
	ErrNameBooksWereNotFound  ErrCodeName = "BooksWereNotFound"
	ErrNameBookWasNotInserted ErrCodeName = "BookWasNotInserted"
	ErrNameBookWasNotParsed   ErrCodeName = "BookWasNotParsed"
	ErrNameIdWasNotParsed     ErrCodeName = "IdWasNotParsed"
	ErrNameWrongCredentials   ErrCodeName = "WrongCredentials"
)

var ErrCodeToName map[ErrCode]ErrCodeName = map[ErrCode]ErrCodeName{
	ErrBookWasNotFound:    ErrNameBookWasNotFound,
	ErrBooksWereNotFound:  ErrNameBooksWereNotFound,
	ErrBookWasNotInserted: ErrNameBookWasNotInserted,
	ErrBookWasNotParsed:   ErrNameBookWasNotParsed,
	ErrIdWasNotParsed:     ErrNameIdWasNotParsed,
	ErrWrongCredentials:   ErrNameWrongCredentials,
}

type Error struct {
	Message  string      `json:"message"`
	Code     ErrCode     `json:"code"`
	CodeName ErrCodeName `json:"code_name"`
}

func NewError(message string, code ErrCode) Error {
	return Error{
		Message:  message,
		Code:     code,
		CodeName: ErrCodeToName[code],
	}
}
