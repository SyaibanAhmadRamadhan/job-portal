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
	PostgreConfig       PostgreConfig       `json:"POSTGRE" mapstructure:"POSTGRE"`
	ElasticsearchConfig ElasticsearchConfig `json:"ES" mapstructure:"ES"`
}

type ElasticsearchConfig struct {
	Host string `json:"HOST" mapstructure:"HOST"`
}

type PostgreConfig struct {
	DSN                       string `json:"DSN" mapstructure:"DSN"`
	DBMaxIdleConnection       int    `json:"DB_MAX_IDLE_CONNECTION" mapstructure:"DB_MAX_IDLE_CONNECTION"`
	DBMaxOpenConnection       int    `json:"DB_MAX_OPEN_CONNECTION" mapstructure:"DB_MAX_OPEN_CONNECTION"`
	ConnMaxLifetimeSecond     int    `json:"CONN_MAX_LIFETIME_SECOND" mapstructure:"CONN_MAX_LIFETIME_SECOND"`
	ConnIdleMaxLifetimeSecond int    `json:"CONN_IDLE_MAX_LIFETIME_SECOND" mapstructure:"CONN_IDLE_MAX_LIFETIME_SECOND"`
}

type RedisConfig struct {
	Host     string `json:"HOST" mapstructure:"HOST"`
	Password string `json:"PASSWORD" mapstructure:"PASSWORD"`
}

type AppConfig struct {
	AppName  string         `json:"APP_NAME" mapstructure:"APP_NAME"`
	Port     int            `json:"PORT" mapstructure:"PORT"`
	Otel     OTELConfig     `json:"OTEL" mapstructure:"OTEL"`
	Kafka    KafkaConfig    `json:"KAFKA" mapstructure:"KAFKA"`
	Database DatabaseConfig `json:"DATABASE" mapstructure:"DATABASE"`
	Redis    RedisConfig    `json:"REDIS" mapstructure:"REDIS"`
}
