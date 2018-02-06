# This is a daemon in linux.

### For exmple:


### build daemon file in window os
```
set goos=linux
go build 
```


### Run daemon in linux
```
nohup ./daemon &
```

### Run task in daemon. 

usage: mydaemon [command] [task name]

command: start | restart | stop | status
```
./daemon start taskname
```


### route:

```
var RouteMap = map[string]func(http.ResponseWriter, *http.Request){
	"/server/": ServerRoute,
	"/alert/":  AlertRoute,
}

func ServerRoute(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")             //返回数据格式是json

	client := controllers.NewController(w, r, cfg)
	url := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(url, "/")
	inMethod := strings.Title(url)
	if len(parts) >= 2 {
		inMethod = strings.Title(parts[1])
	}

	controller := reflect.ValueOf(&controllers.Server{Controller: client})
	method := controller.MethodByName(inMethod)
	if !method.IsValid() {
		client.Error()
		return
	}

	method.Call(nil)
}
```

### controller:

```
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

```