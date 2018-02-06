package controllers

import (
	"fmt"

	"github.com/sunrisedo/daemon/datas"
	"github.com/sunrisedo/daemon/progress"
)

// Create your own business code
var Tasks = progress.NewTasks()

type Server struct {
	*Controller
}

func (c *Server) Start() {
	var ask datas.AskData
	c.RequestStruct(&ask)

	if task := Tasks.Get(ask.Key); task != nil {
		status, err := task.Status()
		if err != nil {
			c.ResultJson(105, err.Error())
			return
		}
		c.ResultJson(106, fmt.Sprintf("Task had start, %v", status))
		return
	}

	task := new(progress.Task)
	if err := task.Init(ask.Key); err != nil {
		c.ResultJson(101, err.Error())
		return
	}
	if err := task.Start(); err != nil {
		c.ResultJson(102, err.Error())
		return
	}

	Tasks.Set(ask.Key, task)
	c.ResultJson(0, "success")
}

func (c *Server) Restart() {
	var ask datas.AskData
	c.RequestStruct(&ask)

	if task := Tasks.Get(ask.Key); task != nil {
		if err := task.Stop(); err != nil {
			c.ResultJson(104, err.Error())
			return
		}
	}

	task := new(progress.Task)
	if err := task.Init(ask.Key); err != nil {
		c.ResultJson(101, err.Error())
		return
	}
	if err := task.Start(); err != nil {
		c.ResultJson(102, err.Error())
		return
	}
	Tasks.Set(ask.Key, task)
	c.ResultJson(0, "success")
}

func (c *Server) Stop() {
	var ask datas.AskData
	c.RequestStruct(&ask)
	task := Tasks.Get(ask.Key)
	if task == nil {
		c.ResultJson(107, fmt.Sprintf("Can not find %s task", ask.Key))
		return
	}

	if err := task.Stop(); err != nil {
		c.ResultJson(104, err.Error())
		return
	}
	Tasks.Set(ask.Key, nil)
	c.ResultJson(0, "success")
}

func (c *Server) Status() {
	var ask datas.AskData
	c.RequestStruct(&ask)
	task := Tasks.Get(ask.Key)
	if task == nil {
		c.ResultJson(107, fmt.Sprintf("Can not find %s task", ask.Key))
		return
	}

	status, err := task.Status()
	if err != nil {
		c.ResultJson(105, err.Error())
		return
	}
	c.ResultJson(0, status)
}
