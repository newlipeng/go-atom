module go-atom

go 1.13

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.47.0
	golang.org/x/build => github.com/golang/build v0.0.0-20191106211638-a32819c83475
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/exp => github.com/golang/exp v0.0.0-20191030013958-a1ab85dbe136
	golang.org/x/image => github.com/golang/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20191031020345-0945064e013a
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190823172224-ecb187b06eb0
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190730183949-1393eb018365
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190730215328-ed3277de2799
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191011141410-1b5146add898
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.13.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc => github.com/grpc/grpc-go v1.25.0
	gopkg.in/asn1-ber.v1 => github.com/go-asn1-ber/asn1-ber v1.3.1
	gopkg.in/fsnotify.v1 => github.com/Jwsonic/recinotify v0.0.0-20151201212458-7389700f1b43
	gopkg.in/gorethink/gorethink.v4 => github.com/rethinkdb/rethinkdb-go v4.0.0+incompatible
	gopkg.in/ini.v1 => github.com/go-ini/ini v1.50.0
	gopkg.in/src-d/go-billy.v4 => github.com/src-d/go-billy v4.2.0+incompatible
	gopkg.in/src-d/go-git-fixtures.v3 => github.com/src-d/go-git-fixtures v3.5.0+incompatible
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20180328195020-5420a8b6744d
	k8s.io/api => github.com/kubernetes/api v0.0.0-20191107030003-665c8a257c1a
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.0.0-20191107025710-670e6d490571
	k8s.io/client-go => github.com/kubernetes/client-go v12.0.0+incompatible
	k8s.io/klog => github.com/simonpasquier/klog-gokit v0.2.0
	k8s.io/kube-openapi => github.com/kubernetes/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d
	k8s.io/utils => github.com/kubernetes/utils v0.0.0-20191030222137-2b95a09bc58d
	sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0
)

require (
	github.com/bilibili/kratos v0.3.1 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/go-redis/redis/v7 v7.0.0-beta.4 // indirect
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mholt/archiver/v3 v3.3.0 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/rs/zerolog v1.15.0 // indirect
	gopkg.in/guregu/null.v3 v3.4.0 // indirect
)
