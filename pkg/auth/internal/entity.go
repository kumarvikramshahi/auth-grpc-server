package internal

type LoginSuccess struct {
	Token           string
	ExpiryTimestamp int64
}

type SignUpSuccess struct {
	Message string
}
