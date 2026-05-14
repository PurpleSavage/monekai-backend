package config 

import(
	"os"
	"github.com/joho/godotenv"
)
type ConfigEnvs struct{
	SecretJwt string
	Port string
	ClerkKey string
	Host string
	User string
	DbPassword string
	DbPort string
	DbName string 
	SslMode string
}

var Envs *ConfigEnvs
func LoadEnvs()  {
	// Intentamos cargar el archivo .env local.
	// Usamos "_" para ignorar el error, ya que en Railway/Prod 
	// el archivo no existirá y no queremos que la app se detenga.
	_ = godotenv.Load()

	Envs= &ConfigEnvs{
		SecretJwt: getEnv("JWT_SECRET", "default_secret_key"),
		Port:      getEnv("PORT", "8080"),
		ClerkKey:getEnv("CLERK_KEY", "default"),
		Host:getEnv("HOST", "default"),
		User:getEnv("USER:", "default"),
		DbPassword:getEnv("DB_PASSWORD", "default"),
		DbName:getEnv("DB_NAME", "default"),
		DbPort:getEnv("DB_PORT", "default"),
		SslMode:getEnv("SSLMODE", "disabled"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}