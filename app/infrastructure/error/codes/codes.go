package codes

type Code string

const (
	OK              Code = "OK"
	InitializeError Code = "initialize_error"
	IO              Code = "io_error"
	Database        Code = "database_error"
	Unknown         Code = "system_error"
)
