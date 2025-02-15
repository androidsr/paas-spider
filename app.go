package main

import (
	"context"
	"os"
	"paas-spider/entity"

	"github.com/androidsr/sc-go/mapper"
	"github.com/androidsr/sc-go/model"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Close() {
	os.Exit(0)

}

func (a *App) SetConfig(config *entity.Config) model.HttpResult {
	db := mapper.NewHelper[entity.Config]()
	err := db.Exec(`
		CREATE TABLE IF NOT EXISTS "config" (
		"id" integer NOT NULL,
		"content" TEXT NOT NULL,
		"legal_statement" integer NOT NULL,
		PRIMARY KEY ("id")
		);

		CREATE TABLE IF NOT EXISTS "plan" (
		"id" INTEGER NOT NULL,
		"name" TEXT NOT NULL,
		"exec_cron" TEXT,
		"exec_remark" TEXT,
		"content" TEXT,
		"status" text NOT NULL,
		PRIMARY KEY ("id")
		);

		CREATE TABLE IF NOT EXISTS "record" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"task_id" TEXT NOT NULL,
		"selector" TEXT,
		"input_value" TEXT,
		"output_value" TEXT,
		"exec_time" TEXT,
		"source_type" TEXT
		);

		CREATE TABLE IF NOT EXISTS "task" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL,
		"system" TEXT NOT NULL,
		"public_id" TEXT NOT NULL,
		"exec_time" TEXT NOT NULL,
		"content" TEXT NOT NULL,
		"created_at" TEXT,
		"updated_at" TEXT
		);
	`).Error

	if err != nil {
		return model.NewFailDefault(err.Error())
	}
	err = db.SaveOrUpdate(config).Error
	if err != nil {
		return model.NewFailDefault("设置参数失败")
	}
	return model.NewOK(nil)
}

func (a *App) GetConfig() model.HttpResult {
	db := mapper.NewHelper[entity.Config]()
	data := new(entity.Config)
	data.Id = "1"
	err := db.SelectOne(data)
	if err != nil {
		return model.NewFailDefault("设置参数失败")
	}
	if data.LegalStatement == "1" {
		return model.NewOK(data)
	} else {
		return model.NewFailDefault("不同意法律声明")
	}
}
