package gojenkins

import (
	"fmt"
	"net/url"
)

type Executable struct {
	Number int    `json:"number"`
	Url    string `json:"url"`
}

type Executor struct {
	Idle        bool `json:"idle"`
	Number      int  `json:"number"`
	Progress    int  `json:"progress"`
	LikelyStuck bool `json:"likelyStuck"`

	CurrentExecutable Executable `json:"currentExecutable"`
}

type Computer struct {
	Name string `json:"displayName"`

	NumExecutors       int    `json:"numExecutors"`
	Idle               bool   `json:"idle"`
	Offline            bool   `json:"offline"`
	OfflineCauseReason string `json:"offlineCauseReason"`
	TemporarilyOffline bool   `json:"temporarilyOffline"`

	Executors []Executor `json:"executors"`
}

func (jenkins *Jenkins) GetComputer(name string) (computer Computer, err error) {
	err = jenkins.get(fmt.Sprintf("/computer/%s", name), nil, &computer, DefaultOptions)
	return
}

func (jenkins *Jenkins) ComputerToggleOffline(name string, message string) error {
	params := url.Values{}
	params.Add("offlineMessage", message)
	return jenkins.post(fmt.Sprintf("/computer/%s/toggleOffline", name), params, nil, Options{useApiEndPoint: false})
}
