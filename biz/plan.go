package biz

import (
	"log"
	"paas-spider/entity"
	"paas-spider/toolkit"

	"github.com/androidsr/sc-go/mapper"
	"github.com/androidsr/sc-go/model"
	"github.com/androidsr/sc-go/sbuilder"
	"github.com/androidsr/sc-go/sno"
)

type PlanBiz struct {
	browser *toolkit.PageToolkit
	db      *mapper.Mapper[entity.Plan]
}

func NewPlanBiz() *PlanBiz {
	db := mapper.NewHelper[entity.Plan]()
	return &PlanBiz{db: db}
}

func (m *PlanBiz) Get(id string) model.HttpResult {
	data := new(entity.Plan)
	data.Id = id
	err := m.db.SelectOne(data)
	if err != nil {
		log.Printf("查询预案失败: %v\n", err)
		return model.NewFailDefault("查询预案失败")
	}
	return model.NewOK(data)
}

func (m *PlanBiz) Add(Plan *entity.Plan) model.HttpResult {
	Plan.Id = sno.GetString()
	err := m.db.Insert(Plan)
	if err != nil {
		log.Printf("添加预案失败: %v\n", err)
		return model.NewFailDefault("添加预案失败")
	}
	return model.NewOK(Plan)
}

func (m *PlanBiz) Edit(Plan *entity.Plan) model.HttpResult {
	err := m.db.Update(Plan, "id", Plan.Id).Error
	if err != nil {
		log.Printf("修改预案失败: %v\n", err)
		return model.NewFailDefault("修改预案失败")
	}
	return model.NewOK(Plan)
}

func (t *PlanBiz) Delete(id string) model.HttpResult {
	err := t.db.Delete("id", id).Error
	if err != nil {
		log.Printf("删除预案失败: %v\n", err)
		return model.NewFailDefault("删除预案失败")
	}
	return model.NewOK(true)
}

type PlanQuery struct {
	Page       *model.PageInfo `json:"page" column:"-"`
	Name       string          `json:"name" column:"a.name" keyword:"like"`
	ExecRemark string          `json:"execRemark" column:"a.exec_remark" keyword:"eq"`
}

func (m *PlanBiz) Page(query *PlanQuery) model.HttpResult {
	sql := `select * from plan a where 1=1 `
	data := make([]entity.Plan, 0)
	b := sbuilder.StructToBuilder(query, sql)
	sql, values := b.Build()
	sql += " order by name,id desc "
	result := m.db.SelectPage(&data, query.Page, sql, values...)
	return model.NewOK(result)
}

func (m *PlanBiz) GetList() model.HttpResult {
	sql := `select id as value,name as label from plan a where 1=1 `
	data := make([]model.SelectVO, 0)
	b := sbuilder.Builder(sql)
	sql, values := b.Build()
	sql += " order by name,id desc "
	err := m.db.SelectSQL(&data, sql, values...)
	if err != nil {
		log.Printf("查询预案列表失败: %v\n", err)
		return model.NewFailDefault("查询预案列表失败")
	}
	return model.NewOK(data)
}

func (m *PlanBiz) GetGroupList() model.HttpResult {
	sql := `select DISTINCT exec_remark as label ,exec_remark as value from plan a where 1=1 `
	data := make([]model.SelectVO, 0)
	b := sbuilder.Builder(sql)
	sql, values := b.Build()
	err := m.db.SelectSQL(&data, sql, values...)
	if err != nil {
		log.Printf("查询任务列表失败: %v\n", err)
		return model.NewFailDefault("查询任务列表失败")
	}
	return model.NewOK(data)
}
