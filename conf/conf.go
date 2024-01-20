package conf

const (
	// app
	Host = "127.0.0.1"
	Port = "8080"

	// mysql
	MysqlHost               = "127.0.0.1"
	MysqlPort               = "3306"
	MysqlUsername           = "root"
	MysqlPassword           = "123456"
	MysqlDbname             = "demo"
	MysqlMaxOpenConns       = 1000
	MysqlMaxIdleConns       = 500
	MysqlConnMaxLifeMinutes = 1 // unit: minute

	// redis
	RedisUser         = "root"
	RedisPassword     = "123456"
	RedisAddr         = "127.0.0.1:6379"
	RedisConnTimeout  = 100
	RedisReadTimeout  = 50
	RedisWriteTimeout = 50
	RedisMaxIdle      = 500
	RedisMaxActive    = 1000
	RedisExpireSecond = 7000

	// todo kafka
)
