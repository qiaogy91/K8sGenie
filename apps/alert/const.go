package alert

const (
	NamespaceContent        = `**告警名称：**%s\n**告警级别：**%s\n**告警集群：**%s\n**所属产线：**%s\n**项目名称：**%s\n**项目编码：**%s\n**名称空间：**%s\n**故障描述：**%s\n**故障详情：**%s\n`
	ClusterContent          = `**告警名称：**%s\n**告警级别：**%s\n**故障描述：**%s\n**故障详情：**%s\n`
	NamespaceWithoutProject = `**告警名称：**%s\n**告警级别：**%s\n**告警集群：**%s\n**名称空间：**%s\n**故障描述：**%s\n**故障详情：**%s\n`
)

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

const UrgentCardTemplate = `{
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
}`
