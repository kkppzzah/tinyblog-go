// Package common 公用功能。
package common

const (
	// UserInfoHeader 鉴权服务设置的在http请求头中的用户信息。
	UserInfoHeader = "X-TB-USER-INFO"
	// EnvVarNameListenAddress 监听地址环境变量名。
	EnvVarNameListenAddress = "LISTEN_ADDRESS"
	// EnvRecommendServiceAddress 推荐服务地址。
	EnvRecommendServiceAddress = "RECOMMEND_SERVICE_ADDRESS"
	// EnvArticleServiceAddress 文章服务地址。
	EnvArticleServiceAddress = "ARTICLE_SERVICE_ADDRESS"
	// EnvAuthServiceAddress 鉴权服务地址。
	EnvAuthServiceAddress = "AUTH_SERVICE_ADDRESS"
	// EnvUserServiceAddress 用户服务地址。
	EnvUserServiceAddress = "USER_SERVICE_ADDRESS"
	// EnvSearchServiceAddress 搜索服务地址。
	EnvSearchServiceAddress = "SEARCH_SERVICE_ADDRESS"
	// CookieNameSession 用来存放会话的Cookie名称。
	CookieNameSession = "tb_session"
)
