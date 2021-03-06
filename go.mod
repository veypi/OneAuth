module github.com/veypi/OneAuth

go 1.16

require (
	github.com/json-iterator/go v1.1.10
	github.com/olivere/elastic/v7 v7.0.29
	github.com/urfave/cli/v2 v2.2.0
	github.com/veypi/OneBD v0.4.3
	github.com/veypi/utils v0.3.1
	gorm.io/driver/mysql v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)

replace github.com/veypi/OneBD v0.4.3 => ../OceanCurrent/OneBD
