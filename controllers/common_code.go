package controllers

/*
 * 定义API返回的错误码(ERROR CODE)规范
 * API框架逻辑会自动从错误信息中分析ERROR CODE(比如返回的error信息中有[477]errormsg..，则认为477即为状态码)
 * 如果有分析出有效的ERROR CODE，则使用错误信息中的ERROR CODE覆盖默认指定的
 * 4xx为客户端错误
 * 5xx为服务端错误
 * 1000+ 为具体逻辑定义错误，这里不做定义
 */
const (
	EC_DEFAULT_ERR = 0 //默认错误码

	EC_JSON_INVALID  = 400
	EC_ACCESS_DENIED = 401
	EC_ACCESS_EXCEED = 402
	EC_DATA_INVALID  = 403
	EC_DATA_LOST     = 405

	EC_INNER_ERR = 500
)
