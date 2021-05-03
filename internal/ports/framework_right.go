package ports

// DbPort is the port for an db adapter
type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) error
}
