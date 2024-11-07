package feishu

const (
	AppName        = "feishu"
	AccessTokenUrl = "https://open.rwork.crc.com.cn/open-apis/auth/v3/tenant_access_token/internal/"
	MessageSendUrl = "https://open.rwork.crc.com.cn/open-apis/im/v1/messages?receive_id_type=open_id"
	UserOpenIdUrl  = "https://open.rwork.crc.com.cn/open-apis/contact/v3/users/batch_get_id"
	UrgentCallUrl  = "https://open.rwork.crc.com.cn/open-apis/im/v1/messages/%s/urgent_phone?user_id_type=open_id"
)
