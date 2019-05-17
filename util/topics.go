package util

const (
	HW_SUB_TOPIC = "/speech/A001"
	HW_PUB_TOPIC = "/speech/A002"

	BS_SUB_TOPIC = "/speech/A003"
	BS_PUB_TOPIC = "/speech/A004"

	ZF_SUB_TOPIC = "/speech/A005"
	ZF_PUB_TOPIC = "/speech/A006"
)

const (
	HOST = "39.96.4.124"
	PORT = "1883"
)

var TopicMap  = map[string]string{
	HW_SUB_TOPIC:HW_PUB_TOPIC,
	BS_SUB_TOPIC:BS_PUB_TOPIC,
	ZF_SUB_TOPIC:ZF_PUB_TOPIC,
}
