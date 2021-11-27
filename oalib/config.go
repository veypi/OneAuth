package oalib

/**
* @name: config
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 16:30
* @description：config
* @update: 2021-11-17 16:30
**/

type Config struct {
	Host string
	UUID string
	Key  []byte
}

func (c *Config) Valid() bool {
	if c != nil && c.Host != "" && c.UUID != "" && c.Key != nil {
		return true
	}
	return false
}
