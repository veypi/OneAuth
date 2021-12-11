package models

import (
	"fmt"
	"github.com/veypi/utils"
	"time"
)

// File es model
type File struct {
	_id       string
	CreatedAt time.Time
	OwnerID   string
	Path      string
	Size      uint
	MD5       string
	Count     uint
	Tag       string
}

func (f *File) ID() string {
	if f._id != "" {
		return f._id
	}
	f._id = utils.HashMd5(fmt.Sprintf("%s@%s", f.OwnerID, f.Path))
	return f._id
}

type Action string

const (
	ActGet    Action = "read"
	ActPut    Action = "put"
	ActDelete Action = "delete"
	ActRename Action = "rename"
)

type History struct {
	CreatedAt time.Time
	ActorID   string
	OwnerID   string
	FileID    string
	Action    Action
	Tag       string
	IP        string
}
