package gojenkins

import (
	"fmt"
)

type Label struct {
	Name           string `json:"name"`
	BusyExecutors  int    `json:"busyExecutors"`
	IdleExecutors  int    `json:"idleExecutors"`
	Offline        bool   `json:"offline"`
	TotalExecutors int    `json:"totalExecutors"`

	TiedJobs []Job `json:"tiedJobs"`

	Nodes []Node `json:"nodes"`
}

func (jenkins *Jenkins) GetLabel(name string) (label Label, err error) {
	err = jenkins.get(fmt.Sprintf("/label/%s", name), nil, &label, DefaultOptions)
	return
}
