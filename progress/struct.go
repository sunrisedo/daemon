package progress

import (
	"sync"
)

type Tasks struct {
	*TaskMap
}

func NewTasks() *Tasks {
	return &Tasks{NewTaskMap()}
}

type TaskMap struct {
	value map[string]*Task
	lock  *sync.RWMutex
}

func NewTaskMap() *TaskMap {
	return &TaskMap{make(map[string]*Task), new(sync.RWMutex)}
}

func (c *TaskMap) Get(key string) *Task {
	defer c.lock.RUnlock()
	c.lock.RLock()
	return c.value[key]
}

func (c *TaskMap) Set(key string, value *Task) {
	defer c.lock.Unlock()
	c.lock.Lock()
	c.value[key] = value
}

// func (c *CommandClient) InitTask(key string) (*Task, error) {

// 	if value := c.task.Get(key); value != nil {
// 		return value.(*Task), nil
// 	}
// 	task := new(Task)
// 	task.Name = key
// 	log.Printf("Start Init:%v", task.Name)
// 	lf, err := os.OpenFile(fmt.Sprintf("log/%s.log", task.Name), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
// 	if err != nil {
// 		return nil, errors.New("Can not create log file.")
// 	}
// 	// defer lf.Close()
// 	task.Log = log.New(lf, "", os.O_APPEND)
// 	task.Sout = bytes.NewBuffer(nil) //fmt
// 	task.Serr = bytes.NewBuffer(nil) //log
// 	task.Cmd = exec.Command("nohup", fmt.Sprintf("./%s", task.Name), "&")
// 	task.Cmd.Stdout = task.Sout
// 	task.Cmd.Stderr = task.Serr
// 	c.task.Set(key, task)
// 	return task, nil
// }

// func (c *CommandClient) UpdatePid(key string) error {
// 	value := c.task.Get(key)
// 	if value == nil {
// 		return errors.New("Can not find the task.")
// 	}
// 	task := value.(*Task)
// 	cmd := exec.Command("pgrep", task.Name)
// 	data, err := cmd.Output()
// 	if err != nil {
// 		return errors.New("Find pid error," + err.Error())
// 	}

// 	pids := strings.Split(string(data), "\n")
// 	if len(pids) == 0 {
// 		return errors.New("Can not find pid.")
// 	}

// 	task.Pid, _ = strconv.Atoi(pids[0])
// 	c.task.Set(key, task)
// 	return nil
// }
