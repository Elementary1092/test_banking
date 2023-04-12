package postgresql

type Config struct {
	Username    string
	Password    string
	Host        string
	Port        string
	DBName      string
	MaxAttempts uint8
}
