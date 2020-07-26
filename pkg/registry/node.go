package registry

import (
	"time"
)

type State uint32

const (
	Down = 1
	UP   = 2
)

type Node struct {
	ID         string    `json:"id" description:"identifier of the node"`
	Hostname   string    `json:"hostname" description:"hostname of the node" default:"localhost"`
	CreateTime time.Time `json:"createTime" description:"The create time of the node" default:"now"`
	UpdateTime time.Time `json:"updateTime" description:"The update time of the node" default:"now"`
	State      State     `json:"state" description:"The state of the node" default:"1"`
}

func (n *Node) PutNode() {

}
