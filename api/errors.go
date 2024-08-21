package api

const (
	ErrorCodeUserCredentialInvalid = "UserCredentialInvalid"
)

type ApiException struct {
	Code    string `json:"code"`
	Details string `json:"details"`
}
