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
	DbUser string
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

	Envs = &ConfigEnvs{
		SecretJwt: getEnv("JWT_SECRET", "default_secret_key"),
		Port:      getEnv("PORT", "8080"),

		ClerkKey: getEnv("CLERK_KEY", "default"),

		Host:       getEnv("HOST", "localhost"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", ""),
		DbName:     getEnv("DB_NAME", "postgres"),
		DbPort:     getEnv("DB_PORT", "5432"),
		SslMode:    getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}