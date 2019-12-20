package g

type msg string

const ()

type Res struct {
	RetCode int    `json:"retcode"`
	Desc    string `json:"desc"`
}

type WsClientMsg struct {
	SN string `json:"sn"`
}

type WsServerMsg struct {
	ServerPort uint16 `json:"server_port"`
	ClientPort uint16 `json:"client_port"`
	SshdPort   uint16 `json:"sshd_port"`
	TaskID     string `json:"task_id"`
}

type ProgressStatus int

var (
	INIT    ProgressStatus = 0
	SUCCESS ProgressStatus = 1
	FAILED  ProgressStatus = 2
)

type Progress struct {
	Status ProgressStatus `json:"status"` // 0开始执行；1执行成功；2执行失败
	TaskID string         `json:"task_id"`
}
