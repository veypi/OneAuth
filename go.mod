module github.com/veypi/OneAuth

go 1.21

require (
	github.com/json-iterator/go v1.1.12
	github.com/olivere/elastic/v7 v7.0.29
	github.com/urfave/cli/v2 v2.2.0
	github.com/veypi/OneBD v0.4.3
	github.com/veypi/utils v0.3.1
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	gorm.io/driver/mysql v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kardianos/service v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.17.2 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/veypi/OneBD v0.4.3 => ../OceanCurrent/OneBD
