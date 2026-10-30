package redis

// key 以命名空间 进行区分

const (
	KeyPrefix          = "maxbell:"
	KeyPostTime        = "post:time"   // Zset;帖子及发帖时间
	KeyPostScore       = "post:score"  // Zset;帖子及投票分数
	KeyPostVotedPrefix = "post:voted:" // Zset;用户及投票情况 参数时post_id
)
