package vo

import "time"

type Redis struct {
	Id                 *int64     `json:"id"`
	Name               *string    `json:"name"`
	Host               *string    `json:"host"`
	Db                 string     `json:"db"`
	ProjectId          *int64     `json:"projectId"`
	Project            *string    `json:"project"`
	Mode               *string    `json:"mode"`
	EnableSshTunnel    *int8      `json:"enableSshTunnel"`    // 是否启用ssh隧道
	SshTunnelMachineId *uint64    `json:"sshTunnelMachineId"` // ssh隧道机器id
	Remark             *string    `json:"remark"`
	TagId              *uint64    `json:"tagId"`
	TagPath            *string    `json:"tagPath"`
	CreateTime         *time.Time `json:"createTime"`
	Creator            *string    `json:"creator"`
	CreatorId          *int64     `json:"creatorId"`
	UpdateTime         *time.Time `json:"updateTime"`
	Modifier           *string    `json:"modifier"`
	ModifierId         *int64     `json:"modifierId"`
}

type Keys struct {
	Cursor map[string]uint64 `json:"cursor"`
	Keys   []*KeyInfo        `json:"keys"`
	DbSize int64             `json:"dbSize"`
}

type KeyInfo struct {
	Key  string `json:"key"`
	Ttl  int64  `json:"ttl"`
	Type string `json:"type"`
}
