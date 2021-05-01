package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}
