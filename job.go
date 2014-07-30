package gojenkins

import (
	"fmt"
	"net/url"
)

type Build struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
	Url    string `json:"url"`

	FullDisplayName string `json:"fullDisplayName"`
	Description     string `json:"description"`

	Timestamp         int `json:"timestamp"`
	Duration          int `json:"duration"`
	EstimatedDuration int `json:"estimatedDuration"`

	Building bool   `json:"building"`
	KeepLog  bool   `json:"keepLog"`
	Result   string `json:"result"`
}

type Job struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`

	Buildable   bool   `json:"buildable"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`

	LastCompletedBuild    Build `json:"lastCompletedBuild"`
	LastFailedBuild       Build `json:"lastFailedBuild"`
	LastStableBuild       Build `json:"lastStableBuild"`
	LastSuccessfulBuild   Build `json:"lastSuccessfulBuild"`
	LastUnstableBuild     Build `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild Build `json:"lastUnsuccessfulBuild"`
}

func (jenkins *Jenkins) GetJobs() (jobs []Job, err error) {
	var payload = struct {
		Jobs []Job `json:"jobs"`
	}{
		Jobs: jobs,
	}
	err = jenkins.get("", nil, &payload, DefaultOptions)
	return
}

func (jenkins *Jenkins) GetJob(name string) (job Job, err error) {
	err = jenkins.get(fmt.Sprintf("/job/%s", name), nil, &job, DefaultOptions)
	return
}

func (jenkins *Jenkins) GetBuild(job Job, number int) (build Build, err error) {
	err = jenkins.get(fmt.Sprintf("/job/%s/%d", job.Name, number), nil, &build, DefaultOptions)
	return
}

func (jenkins *Jenkins) Build(job Job, params url.Values) error {
	return jenkins.post(fmt.Sprintf("/job/%s/buildWithParameters", job.Name), params, nil, DefaultOptions)
}
