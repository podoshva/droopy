// Package pkg
package pkg

type S3Config struct {
	Endpoint  string `env:"ENDPOINT"`
	SecretKey string `env:"SECRET_KEY,required"`
	AccessKey string `env:"ACCESS_KEY,required"`
}

type DBConfig struct {
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Name     string `env:"NAME,required"`
}

type Config struct {
	S3Config S3Config `envPrefix:"S3_"`
	DBConfig DBConfig `envPrefix:"DB_"`
}
