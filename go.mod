module blogv2

go 1.23.2

replace github.com/siddontang/go-mysql v1.9.1 => github.com/go-mysql-org/go-mysql v1.9.1

require (
	github.com/PuerkitoBio/goquery v1.10.0
	github.com/cloudflare/ahocorasick v0.0.0-20240916140611-054963ec9396
	github.com/davecgh/go-spew v1.1.1
	github.com/gin-gonic/gin v1.10.0
	github.com/go-ego/gse v0.80.3
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.20.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/gomarkdown/markdown v0.0.0-20241205020045-f7e15b2f3e62
	github.com/google/uuid v1.6.0
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/lionsoul2014/ip2region/binding/golang v0.0.0-20240510055607-89e20ab7b6c6
	github.com/mojocn/base64Captcha v1.3.6
	github.com/mssola/user_agent v0.6.0
	github.com/olivere/elastic/v7 v7.0.32
	github.com/pkg/errors v0.9.1
	github.com/qiniu/go-sdk/v7 v7.25.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/schollz/progressbar/v3 v3.17.1
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.31.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
	gorm.io/plugin/dbresolver v1.5.3
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/alex-ant/gomath v0.0.0-20160516115720-89013a210a82 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/elastic/go-sysinfo v1.15.0 // indirect
	github.com/elastic/go-windows v1.0.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/goph/emperror v0.17.2 // indirect
	github.com/huandu/xstrings v1.3.3 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nikolalohinski/gonja v1.5.3 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.36.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pkoukk/tiktoken-go v0.1.6 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/redis/rueidis v1.0.34 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/tmc/langchaingo v0.1.13-pre.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/vcaesar/cedar v0.20.2 // indirect
	github.com/yargevad/filepathx v1.0.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	gitlab.com/golang-commonmark/html v0.0.0-20191124015941-a22733972181 // indirect
	gitlab.com/golang-commonmark/linkify v0.0.0-20191026162114-a0c2df6c8f82 // indirect
	gitlab.com/golang-commonmark/markdown v0.0.0-20211110145824-bf3e522c626a // indirect
	gitlab.com/golang-commonmark/mdurl v0.0.0-20191124015652-932350d1cb84 // indirect
	gitlab.com/golang-commonmark/puny v0.0.0-20191124015043-9f83538fa04f // indirect
	go.starlark.net v0.0.0-20230302034142-4b1e35fe2254 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
	golang.org/x/image v0.23.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	modernc.org/fileutil v1.0.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
