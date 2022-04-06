module xinhari.com/hari

go 1.18

// set your local xinhari framework directory
replace xinhari.com/xinhari => ../xinhari

require (
	github.com/aws/aws-sdk-go v1.23.0
	github.com/boltdb/bolt v1.3.1
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/cloudflare/cloudflare-go v0.10.9
	github.com/dustin/go-humanize v1.0.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-acme/lego/v3 v3.4.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/micro/cli/v2 v2.1.2
	github.com/miekg/dns v1.1.27
	github.com/netdata/go-orchestrator v0.0.0-20190905093727-c793edba0e8f
	github.com/olekukonko/tablewriter v0.0.4
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.5.1
	github.com/xlab/treeprint v0.0.0-20181112141820-a009c3971eca
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781
	golang.org/x/tools v0.1.1
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.26.0
	xinhari.com/xinhari v0.1.0-alpha
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bwmarrin/discordgo v0.20.2 // indirect
	github.com/caddyserver/certmagic v0.10.6 // indirect
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/evanphx/json-patch/v5 v5.0.0 // indirect
	github.com/forestgiant/sliceutil v0.0.0-20160425183142-94783f95db6c // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.0.0 // indirect
	github.com/go-git/go-git/v5 v5.1.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.3 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/kevinburke/ssh_config v0.0.0-20190725054713-01f96b0aa0cd // indirect
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/lucas-clemente/quic-go v0.27.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.1 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/nats-io/jwt v0.3.2 // indirect
	github.com/nats-io/nats.go v1.9.2 // indirect
	github.com/nats-io/nkeys v0.1.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nlopes/slack v0.6.1-0.20191106133607-d06c2a2b3249 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pelletier/go-toml v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/spf13/afero v1.1.2 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/xanzy/ssh-agent v0.2.1 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	go.uber.org/atomic v1.5.0 // indirect
	go.uber.org/multierr v1.3.0 // indirect
	go.uber.org/tools v0.0.0-20190618225709-2cfd321de3ee // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/square/go-jose.v2 v2.3.1 // indirect
	gopkg.in/telegram-bot-api.v4 v4.6.4 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.0.1-2019.2.3 // indirect
)
