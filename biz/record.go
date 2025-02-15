package biz

import (
	"log"
	"paas-spider/entity"

	"github.com/androidsr/sc-go/mapper"
	"github.com/androidsr/sc-go/model"
	"github.com/androidsr/sc-go/sbuilder"
	"github.com/androidsr/sc-go/sno"
)

type RecordBiz struct {
	db *mapper.Mapper[entity.Record]
}

func NewRecordBiz() *RecordBiz {
	db := mapper.NewHelper[entity.Record]()
	return &RecordBiz{db: db}
}

func (m *RecordBiz) Get(id string) model.HttpResult {
	data := new(entity.Record)
	data.Id = id
	err := m.db.SelectOne(data)
	if err != nil {
		log.Printf("查询数据记录失败: %v\n", err)
		return model.NewFailDefault("查询数据记录失败")
	}
	return model.NewOK(data)
}

func (m *RecordBiz) Add(Record *entity.Record) model.HttpResult {
	Record.Id = sno.GetString()
	err := m.db.Insert(Record)
	if err != nil {
		log.Printf("添加数据记录失败: %v\n", err)
		return model.NewFailDefault("添加数据记录失败")
	}
	return model.NewOK(Record)
}

func (m *RecordBiz) Edit(Record *entity.Record) model.HttpResult {
	err := m.db.Update(Record, "id", Record.Id).Error
	if err != nil {
		log.Printf("修改数据记录失败: %v\n", err)
		return model.NewFailDefault("修改数据记录失败")
	}
	return model.NewOK(Record)
}

func (t *RecordBiz) Delete(id string) model.HttpResult {
	err := t.db.Delete("id", id).Error
	if err != nil {
		log.Printf("删除数据记录失败: %v\n", err)
		return model.NewFailDefault("删除数据记录失败")
	}
	return model.NewOK(true)
}

type RecordQuery struct {
	Page   *model.PageInfo `json:"page" column:"-"`
	TaskId string          `json:"taskId" column:"a.task_id" keyword:"eq"`
	Name   string          `json:"name" column:"b.name" keyword:"like"`
	System string          `json:"system" column:"b.system" keyword:"like"`
}

type RecordVo struct {
	entity.Record
	Name   string `json:"name" gorm:"column:name"`
	System string `json:"system" gorm:"column:system"`
}

func (m *RecordBiz) Page(query *RecordQuery) model.HttpResult {
	sql := `SELECT a.id,b.system,b.name,a.input_value,a.output_value,a.exec_time FROM record a left join task b on a.task_id = b.id where 1=1  `
	data := make([]RecordVo, 0)
	b := sbuilder.StructToBuilder(query, sql)
	sql, values := b.Build()
	sql += " order by strftime('%Y-%m-%d %H:%M:%S', a.exec_time),a.id desc"
	result := m.db.SelectPage(&data, query.Page, sql, values...)
	return model.NewOK(result)
}
