package conf

type OTELConfig struct {
	Host     string `json:"HOST"`
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type KafkaConfig struct {
	Host     string `json:"HOST"`
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type DatabaseConfig struct {
	COMMAND struct {
		DSN                       string `json:"DSN"`
		DBMaxIdleConnection       int    `json:"DB_MAX_IDLE_CONNECTION"`
		DBMaxOpenConnection       int    `json:"DB_MAX_OPEN_CONNECTION"`
		ConnMaxLifetimeSecond     int    `json:"CONN_MAX_LIFETIME_SECOND"`
		ConnIdleMaxLifetimeSecond int    `json:"CONN_IDLE_MAX_LIFETIME_SECOND"`
	} `json:"COMMAND"`

	READER struct {
		DSN                       string `json:"DSN"`
		DBMaxIdleConnection       int    `json:"DB_MAX_IDLE_CONNECTION"`
		DBMaxOpenConnection       int    `json:"DB_MAX_OPEN_CONNECTION"`
		ConnMaxLifetimeSecond     int    `json:"CONN_MAX_LIFETIME_SECOND"`
		ConnIdleMaxLifetimeSecond int    `json:"CONN_IDLE_MAX_LIFETIME_SECOND"`
	} `json:"READER"`
}

type AppConfig struct {
	AppName  string         `json:"APP_NAME"`
	Port     int            `json:"PORT"`
	Otel     OTELConfig     `json:"OTEL"`
	Kafka    KafkaConfig    `json:"KAFKA"`
	Database DatabaseConfig `json:"DATABASE"`
}
