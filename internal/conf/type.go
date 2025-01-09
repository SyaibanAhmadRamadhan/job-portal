package conf

type OTELConfig struct {
	Host     string `json:"HOST" mapstructure:"HOST"`
	Username string `json:"USERNAME" mapstructure:"USERNAME"`
	Password string `json:"PASSWORD" mapstructure:"PASSWORD"`
}

type KafkaConfig struct {
	Host     string `json:"HOST" mapstructure:"HOST"`
	Username string `json:"USERNAME" mapstructure:"USERNAME"`
	Password string `json:"PASSWORD" mapstructure:"PASSWORD"`
}

type DatabaseConfig struct {
	COMMAND struct {
		DSN                       string `json:"DSN" mapstructure:"DSN"`
		DBMaxIdleConnection       int    `json:"DB_MAX_IDLE_CONNECTION" mapstructure:"DB_MAX_IDLE_CONNECTION"`
		DBMaxOpenConnection       int    `json:"DB_MAX_OPEN_CONNECTION" mapstructure:"DB_MAX_OPEN_CONNECTION"`
		ConnMaxLifetimeSecond     int    `json:"CONN_MAX_LIFETIME_SECOND" mapstructure:"CONN_MAX_LIFETIME_SECOND"`
		ConnIdleMaxLifetimeSecond int    `json:"CONN_IDLE_MAX_LIFETIME_SECOND" mapstructure:"CONN_IDLE_MAX_LIFETIME_SECOND"`
	} `json:"COMMAND" mapstructure:"COMMAND"`

	READER struct {
		DSN                       string `json:"DSN" mapstructure:"DSN"`
		DBMaxIdleConnection       int    `json:"DB_MAX_IDLE_CONNECTION" mapstructure:"DB_MAX_IDLE_CONNECTION"`
		DBMaxOpenConnection       int    `json:"DB_MAX_OPEN_CONNECTION" mapstructure:"DB_MAX_OPEN_CONNECTION"`
		ConnMaxLifetimeSecond     int    `json:"CONN_MAX_LIFETIME_SECOND" mapstructure:"CONN_MAX_LIFETIME_SECOND"`
		ConnIdleMaxLifetimeSecond int    `json:"CONN_IDLE_MAX_LIFETIME_SECOND" mapstructure:"CONN_IDLE_MAX_LIFETIME_SECOND"`
	} `json:"READER" mapstructure:"READER"`
}

type AppConfig struct {
	AppName  string         `json:"APP_NAME" mapstructure:"APP_NAME"`
	Port     int            `json:"PORT" mapstructure:"PORT"`
	Otel     OTELConfig     `json:"OTEL" mapstructure:"OTEL"`
	Kafka    KafkaConfig    `json:"KAFKA" mapstructure:"KAFKA"`
	Database DatabaseConfig `json:"DATABASE" mapstructure:"DATABASE"`
}
