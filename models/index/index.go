package index

import (
	"github.com/olivere/elastic/v7"
	"github.com/veypi/OneAuth/cfg"
)

type index string

const (
	File    index = "file"
	History index = "history"
)

func (i index) String() string {
	return string(i)
}

func (i index) Count() *elastic.CountService {
	return cfg.ES().Count().Index(string(i))
}

func (i index) Index() *elastic.IndexService {
	return cfg.ES().Index().Index(string(i))
}
func (i index) Get() *elastic.GetService {
	return cfg.ES().Get().Index(string(i))
}

func (i index) Update() *elastic.UpdateService {
	return cfg.ES().Update().Index(string(i))
}

func (i index) Search() *elastic.SearchService {
	return cfg.ES().Search().Index(string(i))
}

func (i index) Delete() *elastic.DeleteService {
	return cfg.ES().Delete().Index(string(i))
}
func (i index) DeleteByQuery() *elastic.DeleteByQueryService {
	return cfg.ES().DeleteByQuery().Index(string(i))
}
