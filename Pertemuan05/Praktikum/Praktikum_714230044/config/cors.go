package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"https://Fadhail.github.io",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
