package utils

const (
	ONEMINUTE int64 = 60
	ONEHOUR int64 = 3600
	ONEDAY int64 = 24 * ONEHOUR
	TWOWEEKS int64 = 14 * ONEDAY

	UPLOAD_LINK_EXPIRED_TIME int64 = 2 * ONEHOUR
	UPLOAD_LINK_USED_LIMITS int = 5		//上传链接token使用次数限制
)