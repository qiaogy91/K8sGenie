module gitee.com/qiaogy91/K8sGenie

go 1.23.0

require (
	gitee.com/qiaogy91/UserCenter v0.0.0-20240902070712-850914e686d5
	github.com/BurntSushi/toml v1.4.0
	github.com/caarlos0/env/v11 v11.2.2
	github.com/emicklei/go-restful/v3 v3.12.1
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gorilla/websocket v1.5.3
	github.com/larksuite/oapi-sdk-go/v3 v3.3.7
	github.com/rancher/norman v0.0.0-20240708202514-a0127673d1b9
	github.com/rancher/types v0.0.0-20220328215343-4370ff10ecd5
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.5.0
	github.com/rs/zerolog v1.33.0
	github.com/spf13/cobra v1.8.1
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.11
	k8s.io/api v0.30.1
	k8s.io/apimachinery v0.30.1
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/metrics v0.30.1
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rancher/wrangler/v3 v3.0.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/oauth2 v0.20.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/term v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240730163845-b1a4ccb954bf // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.30.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.30.1
	k8s.io/client-go => k8s.io/client-go v0.30.1
)
