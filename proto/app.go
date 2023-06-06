package proto

import "time"

const (
	StateUnknown State = iota
	StateCreated
	StateBuilding
	StateDeploying
	StateReady
	StateFailed
)

type State uint8

type App struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Repo     string    `json:"repo"`
	State    State     `json:"state"`
	Resource *Resource `json:"resource"`
	Template *Template `json:"template"`
	Command  string    `json:"command"`
	Env      []KV      `json:"env"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

type Template struct {
	Image string `json:"image"`
}

type Resource struct {
	CPU    int `json:"cpu"`
	GPU    int `json:"gpu"`
	Memory int `json:"memory"`
	Disk   int `json:"disk"`
}

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
