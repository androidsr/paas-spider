package entity

type BaseEntity struct {
	Id string `json:"id" gorm:"primaryKey;column:id"`
}

type Task struct {
	BaseEntity
	Name     string `json:"name" gorm:"column:name"`
	System   string `json:"system" gorm:"column:system"`
	PublicId string `json:"publicId" gorm:"column:public_id"`
	ExecTime string `json:"execTime" gorm:"column:exec_time"`
	Content  string `json:"content" gorm:"column:content"`
}

type Plan struct {
	BaseEntity
	Name       string `gorm:"column:name" json:"name"`
	ExecCron   string `gorm:"column:exec_cron" json:"execCron"`     // 执行的Cron表达式
	ExecRemark string `gorm:"column:exec_remark" json:"execRemark"` // 执行备注
	Content    string `gorm:"column:content" json:"content"`        // 任务内容
	Status     string `gorm:"column:status" json:"status"`          // 任务状态
}

type Record struct {
	BaseEntity
	TaskId      string `json:"taskId" gorm:"column:task_id"`
	Selector    string `json:"selector" gorm:"column:selector"`
	InputValue  string `json:"inputValue" gorm:"column:input_value"`
	OutputValue string `json:"outputValue" gorm:"column:output_value"`
	ExecTime    string `json:"execTime" gorm:"column:exec_time"`
	SourceType  string `json:"sourceType" gorm:"column:source_type"`
}

type Step struct {
	Id         string `json:"id"`
	TaskId     string `json:"taskId"`
	Name       string `json:"name"`
	EventType  string `json:"eventType"`
	Selector   string `json:"selector"`
	InputValue string `json:"inputValue"`
	SleepTime  int    `json:"sleepTime"`
}

type Config struct {
	BaseEntity
	LegalStatement string `json:"legalStatement" gorm:"column:legal_statement"`
	Content        string `json:"content" gorm:"column:content"`
}
