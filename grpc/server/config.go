package server

type Config struct {
	Host             string `env:"HOST"              envDefault:"0.0.0.0"`         // GRPC_HOST
	Port             int    `env:"PORT"              envDefault:"9090"`            // GRPC_PORT
	Address          string `env:"ADDRESS,expand"    envDefault:"${HOST}:${PORT}"` // GRPC_ADDRESS
	AuthEnable       bool   `env:"AUTH_ENABLE"       envDefault:"false"`           // GRPC_AUTH_ENABLE
	AuthKey          string `env:"AUTH_KEY,notEmpty" envDefault:"testToken"`       // GRPC_AUTH_KEY
	ReflectionEnable bool   `env:"REFLECTION_ENABLE" envDefault:"false"`           // GRPC_REFLECTION_ENABLE
}
