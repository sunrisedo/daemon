package progress

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Task struct {
	Command string
	Name    string
	Pid     int
	Done    bool
	Log     *log.Logger
	Serr    *bytes.Buffer
	Sout    *bytes.Buffer
	Cmd     *exec.Cmd
}

// // Install the service
// func (c *CommandClient) Install() error {
// 	return nil
// }

// // Remove the service
// func (c *CommandClient) Remove() error {
// 	return nil
// }

// func (c *CommandClient) Restart() error {

// 	log.Printf("Restart")
// 	c.Stop()
// 	return nil
// }

func (c *Task) Init(key string) error {
	c.Name = key
	log.Printf("Start Init:%v", c.Name)
	lf, err := os.OpenFile(fmt.Sprintf("log/%s.log", c.Name), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return errors.New("Can not create log file.")
	}
	// defer lf.Close()
	c.Done = false
	c.Log = log.New(lf, "", os.O_APPEND)
	c.Sout = bytes.NewBuffer(nil) //fmt
	c.Serr = bytes.NewBuffer(nil) //log
	c.Cmd = exec.Command("nohup", fmt.Sprintf("./%s", c.Name), "&")
	c.Cmd.Stdout = c.Sout
	c.Cmd.Stderr = c.Serr
	return nil
}

// Start the service
func (c *Task) Start() error {
	go func() {
		var soutbuff, serrbuff []byte
		ticker := time.NewTicker(time.Second * 1)
		for {
			if c.Done {
				break
			}
			select {
			case <-ticker.C:
				// log.Println("Soutlen:", task.Sout.Len())
				if c.Sout.String() != "" {
					soutbuff = make([]byte, c.Sout.Len())
					c.Sout.Read(soutbuff)
					c.Log.Println(string(soutbuff))
				}
				// log.Println("Serrlen:", task.Serr.Len())
				if c.Serr.String() != "" {
					serrbuff = make([]byte, c.Serr.Len())
					c.Serr.Read(serrbuff)
					c.Log.Println(string(serrbuff))
				}
			}
		}
		log.Printf("%s log exit.", c.Name)
	}()

	go func() {
		for {
			if c.Done {
				break
			}
			log.Printf("Start task:%v", c.Name)
			if err := c.Cmd.Start(); err != nil {
				log.Printf("Start error:%v", err)
				time.Sleep(time.Second * 2)
				continue
			}

			c.Pid = c.Cmd.Process.Pid
			if err := c.Cmd.Wait(); err != nil {
				log.Printf("%s exit error:%v", c.Name, err)
				time.Sleep(time.Second * 2)
			}
		}
		log.Printf("%s daemon exit.", c.Name)
	}()

	return nil
}

// Stop the service
func (c *Task) Stop() error {

	if err := c.Cmd.Process.Kill(); err != nil {
		return err
	}
	c.Done = true
	// process, err := os.FindProcess(c.Pid)
	// if err != nil {
	// 	log.Println("FindProcess:", err)
	// 	return err
	// }
	// log.Println("do---1")
	// if err := process.Kill(); err != nil {
	// 	return err
	// }
	return nil
}

// Status - Get service status
func (c *Task) Status() (interface{}, error) {
	// process, err := os.FindProcess(c.Pid)
	// if err != nil {
	// 	log.Println("FindProcess:", err)
	// 	return nil, err
	// }
	// log.Println("do---1", process)

	// log.Println("doSysUsage1", c.Cmd.ProcessState.SysUsage())

	// log.Printf("%s status:%s", c.nowtask.Name, processState.String())
	status := fmt.Sprintf("%s status: %v", c.Name, *c.Cmd.Process)
	return status, nil
}
