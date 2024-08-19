package provider

import "context"

type Service interface {
	SendCard(ctx context.Context, url string, data []byte) []byte
}

/*0. 公共集群告警统一处理
  - robotName: "devopsRobot"
    webhook: "https://open.rwork.crc.com.cn/open-apis/bot/v2/hook/327a89a7-cce9-42f9-8dfe-2deb8b4413b7"
    key: "l2Dgny5nI48ZMcqDl8wbrh"
    extraLabels:
      - name: "cluster_name"
        desc: "集群名称"
      - name: "level"
        desc: "告警级别"

*/
// CardTemplate json模板，需按序提供如下参数
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
                "tag": "hr"
            },
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
