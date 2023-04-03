// Package common 公用功能。
package common

const (
	// EnvVarNameListenAddress 监听地址环境变量名。
	EnvVarNameListenAddress = "LISTEN_ADDRESS"
	// EnvKafkaBrokers Kafka brokers 逗号分隔。
	EnvKafkaBrokers = "KAFKA_BROKERS"
	// EnvEsClusterAddresses Elasticsearch集群地址列表。
	EnvEsClusterAddresses = "ES_CLUSTER_ADDRESSES"
	// EnvArticleServiceAddress 文章服务名。
	EnvArticleServiceAddress = "ARTICLE_SERVICE_ADDRESS"
	// ErrorCodeNoFound 未找到相关的数据。
	ErrorCodeNoFound = 101
	// ErrorCodeDBOpError 执行数据库操作错误。
	ErrorCodeDBOpError = 103
	// ErrorCodeInternalError 内部错误。
	ErrorCodeInternalError = 110
	// KafkaTopicNameUserEvent 用户事件topic名。
	KafkaTopicNameUserEvent = "user-event"
	// KafkaTopicNameArticleEvent 文章事件topic名。
	KafkaTopicNameArticleEvent = "article-event"
	// KafkaTopicNameUserEventConsumerGroup 用户事件topic consumer group。
	KafkaTopicNameUserEventConsumerGroup = "user-event-search"
	// KafkaTopicNameArticleEventConsumerGroup 用户事件topic consumer group。
	KafkaTopicNameArticleEventConsumerGroup = "article-event-search"
	// EsIndexNameArticle Elasticsearch中文章索引名称。
	EsIndexNameArticle = "article"
)
