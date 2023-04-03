// Package common 公用功能。
package common

const (
	// EnvVarNameListenAddress 监听地址环境变量名。
	EnvVarNameListenAddress = "LISTEN_ADDRESS"
	// EnvJWTSecret 用于JWT。
	EnvJWTSecret = "JWT_SECRET"
	// EnvJWTSecretSecretFile 用于JWT（文件中读取）。
	EnvJWTSecretSecretFile = "JWT_SECRET_SECRET_FILE"
	// EnvUserDBConnStr 用来连接用户数据库的地址。
	EnvUserDBConnStr = "USER_DB_CONN_STR"
	// EnvUserDBConnStrSecretFile 用来连接用户数据库的地址（文件中读取）。
	EnvUserDBConnStrSecretFile = "USER_DB_CONN_STR_SECRET_FILE"
	// EnvRedisConnStr 用来连接Redis的地址。
	EnvRedisConnStr = "REDIS_CONN_STR"
	// EnvKafkaBrokers Kafka brokers 逗号分隔。
	EnvKafkaBrokers = "KAFA_BROKERS"
	// ErrorCodeNoFound 未找到相关的数据。
	ErrorCodeNoFound = 101
	// ErrorCodeAuthMethodNotImplemented 鉴权方法未实现。
	ErrorCodeAuthMethodNotImplemented = 102
	// ErrorCodeDBOpError 执行数据库操作错误。
	ErrorCodeDBOpError = 103
	// RedisSessionKeyPrefix 存放在redis中的会话
	RedisSessionKeyPrefix = "tb:session:"
	// KafkaTopicNameUserEvent 用户事件topic名。
	KafkaTopicNameUserEvent = "user-event"
)
