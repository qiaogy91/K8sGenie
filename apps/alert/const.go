package alert

/*
集群级告警：
	webhook: "https://open.rwork.crc.com.cn/open-apis/bot/v2/hook/327a89a7-cce9-42f9-8dfe-2deb8b4413b7"
    key: "l2Dgny5nI48ZMcqDl8wbrh"

产线级告警：
	webhook: "https://open.rwork.crc.com.cn/open-apis/bot/v2/hook/395fb852-7884-45da-880c-e2290ad2698e"
	key: “w8mKUHSZi4nEoMRqvVtFCb”
*/

const CardTemplate = `{
    "timestamp": %d,
    "sign": "%s",
    "msg_type": "interactive",
    "card": {
        "config": {
            "wide_screen_mode": true
        },
        "elements": [
            {
                "tag": "markdown",
                "content": "%s"
            },
            {
                "tag": "markdown",
                "content": "%s"
            },
            {
                "tag": "note",
                "elements": [
                    {
                        "tag": "plain_text",
                        "content": "IT 公共服务部运维项目"
                    }
                ]
            }
        ],
        "header": {
            "template": "%s",
            "title": {
                "content": "%s",
                "tag": "plain_text"
            }
        }
    }
}`
