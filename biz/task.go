package biz

import (
	"encoding/json"
	"log"
	"math/rand"
	"paas-spider/entity"
	"paas-spider/toolkit"
	"strings"
	"time"

	"github.com/androidsr/sc-go/mapper"
	"github.com/androidsr/sc-go/model"
	"github.com/androidsr/sc-go/sbuilder"
	"github.com/androidsr/sc-go/sno"
	"github.com/androidsr/sc-go/syaml"
)

func init() {
	mapper.Initdb(&syaml.GormInfo{
		Driver:  "sqlite",
		Url:     "sqlite.db",
		MaxOpen: 10,
		MaxIdle: 2,
		ShowSql: false,
	})
	sno.New(syaml.SnowflakeInfo{WorkerId: 1})
}

type TaskBiz struct {
	startInterval float64
	endInterval   float64
	browser       *toolkit.PageToolkit
	db            *mapper.Mapper[entity.Task]
}

func NewTaskBiz() *TaskBiz {
	db := mapper.NewHelper[entity.Task]()
	return &TaskBiz{db: db}
}

func (m *TaskBiz) Get(id string) model.HttpResult {
	data := new(entity.Task)
	data.Id = id
	err := m.db.SelectOne(data)
	if err != nil {
		log.Printf("查询任务失败: %v\n", err)
		return model.NewFailDefault("查询任务失败")
	}
	return model.NewOK(data)
}

func (m *TaskBiz) Add(task *entity.Task) model.HttpResult {
	task.Id = sno.GetString()
	err := m.db.Insert(task)
	if err != nil {
		log.Printf("添加任务失败: %v\n", err)
		return model.NewFailDefault("添加任务失败")
	}
	return model.NewOK(task)
}

func (m *TaskBiz) Edit(task *entity.Task) model.HttpResult {
	err := m.db.Update(task, "id", task.Id).Error
	if err != nil {
		log.Printf("修改任务失败: %v\n", err)
		return model.NewFailDefault("修改任务失败")
	}
	return model.NewOK(task)
}

func (t *TaskBiz) Delete(id string) model.HttpResult {
	err := t.db.Delete("id", id).Error
	if err != nil {
		log.Printf("删除任务失败: %v\n", err)
		return model.NewFailDefault("删除任务失败")
	}
	return model.NewOK(true)
}

type TaskQuery struct {
	Page   *model.PageInfo `json:"page" column:"-"`
	Name   string          `json:"name" column:"a.name" keyword:"like"`
	System string          `json:"system" column:"a.system" keyword:"eq"`
}

func (m *TaskBiz) Page(query *TaskQuery) model.HttpResult {
	sql := `select * from task a where 1=1 `
	data := make([]entity.Task, 0)
	b := sbuilder.StructToBuilder(query, sql)
	sql, values := b.Build()
	sql += " order by name,id desc "
	result := m.db.SelectPage(&data, query.Page, sql, values...)
	return model.NewOK(result)
}

func (m *TaskBiz) GetList() model.HttpResult {
	sql := `select id as value,name as label from task a where 1=1 `
	data := make([]model.SelectVO, 0)
	b := sbuilder.Builder(sql)
	sql, values := b.Build()
	sql += " order by name,id desc "
	err := m.db.SelectSQL(&data, sql, values...)
	if err != nil {
		log.Printf("查询任务列表失败: %v\n", err)
		return model.NewFailDefault("查询任务列表失败")
	}
	return model.NewOK(data)
}

func (m *TaskBiz) GetSystemList() model.HttpResult {
	sql := `select DISTINCT system from task a where 1=1 `
	data := make([]string, 0)
	b := sbuilder.Builder(sql)
	sql, values := b.Build()
	err := m.db.SelectSQL(&data, sql, values...)
	if err != nil {
		log.Printf("查询任务列表失败: %v\n", err)
		return model.NewFailDefault("查询任务列表失败")
	}
	return model.NewOK(data)
}

// 初始化浏览器
func (m *TaskBiz) InitBrowser() {
	db := mapper.NewHelper[entity.Config]()
	c := new(entity.Config)
	c.Id = "1"
	err := db.SelectOne(c)
	var pw *toolkit.PageToolkit
	if err == nil && c.Content != "" && c.Content != "{}" {
		conf := make(map[string]interface{}, 0)
		err := json.Unmarshal([]byte(c.Content), &conf)
		if err == nil {
			startInterval, ok := conf["startInterval"]
			if ok {
				m.startInterval = startInterval.(float64)
			} else {
				m.startInterval = 700
			}
			endInterval, ok := conf["endInterval"]
			if ok {
				m.endInterval = endInterval.(float64)
			} else {
				m.endInterval = 1000
			}
			pw, _ = toolkit.NewBrowser(conf)
		}
	}
	m.browser = pw
}

func (m *TaskBiz) ExecStep(step *entity.Step) model.HttpResult {
	if m.browser == nil || (m.browser.Page == nil && step.EventType != "0") {
		return model.NewFailDefault("浏览器未打开")
	}
	if m.endInterval > 0 {
		randomNumber := rand.Intn(int(m.endInterval-m.startInterval+1)) + int(m.startInterval)
		time.Sleep(time.Duration(randomNumber) * time.Millisecond)
	}
	if step.SleepTime > 0 {
		time.Sleep(time.Duration(step.SleepTime) * time.Second)
	}
	switch step.EventType {
	case "0": //打开页面
		_, err := m.browser.NewPage(step.InputValue)
		if err != nil {
			return model.NewFailDefault("打开页面失败" + err.Error())
		}
	case "1": //点击操作
		err := m.browser.ClickForWait(step.Selector)
		if err != nil {
			return model.NewFailDefault("点击操作失败:" + err.Error())
		}
	case "2": //输入参数
		err := m.browser.SetInputValue(step.Selector, step.InputValue)
		if err != nil {
			return model.NewFailDefault("输入参数失败" + err.Error())
		}
	case "3": //点击键盘
		err := m.browser.Keyboard(step.InputValue)
		if err != nil {
			return model.NewFailDefault("点击键盘失败" + err.Error())
		}
	case "4": //截幕
		m.browser.Screenshot(step.InputValue)
	case "5": //等待
		err := m.browser.WaitForElement(step.InputValue)
		if err != nil {
			return model.NewFailDefault("等待元素失败" + err.Error())
		}
	case "6": //滚动
		err := m.browser.ScrollToBottom()
		if err != nil {
			return model.NewFailDefault("滚动页面失败" + err.Error())
		}
	case "7":
		ck, err := m.browser.GetCookie(step.InputValue)
		if err != nil {
			return model.NewFailDefault("获取Cookie失败" + err.Error())
		}
		return model.NewOK(ck)
	case "8":
		//获取header
	case "9":
		//获取链接文本
		result, err := m.browser.GetLinkContent(step.Selector, "text")
		if err != nil {
			return model.NewFailDefault("未找到有效的音频源")
		}
		if len(result) == 0 {
			return model.NewOK(nil)
		}
		bs, err := json.Marshal(result)
		if err == nil {
			NewRecordBiz().Add(&entity.Record{
				TaskId:      step.TaskId,
				Selector:    step.Selector,
				InputValue:  step.InputValue,
				OutputValue: string(bs),
				ExecTime:    time.Now().Format("2006-01-02 15:04:05"),
				SourceType:  "文本",
			})
		}
		return model.NewOK(string(bs))
	case "10":
		//获取图片
		result, err := m.browser.GetLinkContent(step.Selector, "img")
		if err != nil {
			return model.NewFailDefault("未找到有效的图片源")
		}
		if len(result) == 0 {
			return model.NewOK(nil)
		}
		bs, err := json.Marshal(result)
		if err == nil {
			NewRecordBiz().Add(&entity.Record{
				TaskId:      step.TaskId,
				Selector:    step.Selector,
				InputValue:  step.InputValue,
				OutputValue: string(bs),
				ExecTime:    time.Now().Format("2006-01-02 15:04:05"),
				SourceType:  "图片",
			})
		}
		return model.NewOK(string(bs))
	case "11":
		//获取音频
		result, err := m.browser.GetLinkContent(step.Selector, "audio")
		if err != nil {
			return model.NewFailDefault("未找到有效的音频源")
		}
		if len(result) == 0 {
			return model.NewOK(nil)
		}
		bs, err := json.Marshal(result)
		if err == nil {
			NewRecordBiz().db.Insert(&entity.Record{
				TaskId:      step.TaskId,
				Selector:    step.Selector,
				InputValue:  step.InputValue,
				OutputValue: string(bs),
				ExecTime:    time.Now().Format("2006-01-02 15:04:05"),
				SourceType:  "音频",
			})
		}
		return model.NewOK(string(bs))
	case "12":
		//获取视频
		result, err := m.browser.GetLinkContent(step.Selector, "video")
		if err != nil {
			return model.NewFailDefault("未找到有效的视频源")
		}
		if len(result) == 0 {
			return model.NewOK(nil)
		}
		bs, err := json.Marshal(result)
		if err == nil {
			NewRecordBiz().Add(&entity.Record{
				TaskId:      step.TaskId,
				Selector:    step.Selector,
				InputValue:  step.InputValue,
				OutputValue: string(bs),
				ExecTime:    time.Now().Format("2006-01-02 15:04:05"),
				SourceType:  "视频",
			})
		}
		return model.NewOK(string(bs))
	case "19":
		//获取链接
		urls, err := m.browser.GetLinks(step.Selector)
		if err != nil {
			return model.NewFailDefault("获取文本内容失败" + err.Error())
		}
		if len(urls) == 0 {
			return model.NewOK(nil)
		}
		keys := make([]string, 0, len(urls))
		for k := range urls {
			keys = append(keys, "《"+k+"》 ")
		}

		return model.NewOK(strings.Join(keys, ""))
	}
	return model.NewOK(nil)
}
