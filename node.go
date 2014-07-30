package gojenkins

type Node struct {
	Name            string `json:"nodeName"`
	NumExecutors    int    `json:"numExecutors"`
	Mode            string `json:"mode"`
	NodeDescription string `json:"nodeDescription"`
}
