// Package common 公用功能。
package common

const (
	// EnvVarNameListenAddress 监听地址环境变量名。
	EnvVarNameListenAddress = "LISTEN_ADDRESS"
	// EnvArticleDBConnStr 用来连接文章数据库的地址。
	EnvArticleDBConnStr = "ARTICLE_DB_CONN_STR"
	// EnvKafkaBrokers Kafka brokers 逗号分隔。
	EnvKafkaBrokers = "KAFA_BROKERS"
	// ErrorCodeNoFound 未找到相关的数据。
	ErrorCodeNoFound = 101
	// ErrorCodeDBOpError 执行数据库操作错误。
	ErrorCodeDBOpError = 103
	// KafkaTopicNameUserEvent 用户事件topic名。
	KafkaTopicNameUserEvent = "user-event"
	// KafkaTopicNameArticleEvent 文章事件topic名。
	KafkaTopicNameArticleEvent = "article-event"
	// KafkaTopicNameUserEventConsumerGroup 用户事件topic consumer group。
	KafkaTopicNameUserEventConsumerGroup = "user-event-article"
)
