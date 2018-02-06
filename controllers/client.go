package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sunrisedo/daemon/datas"
)

var Commands = map[string]string{
	"start":   "start",
	"restart": "restart",
	"stop":    "stop",
	"status":  "status",
}

type Client struct {
	path string
}

func NewClient(path string) *Client {
	return &Client{path}
}

// Manage by daemon commands or run the daemon
func (c *Client) Listen() error {
	if len(os.Args) < 3 {
		return errors.New("The input task is empty. Usage: mydaemon [command] [task name]")
	}
	cmd, ok := Commands[os.Args[1]]
	if !ok {
		return errors.New("Usage: mydaemon install | remove | start | stop | status")
	}

	var obj datas.AskData
	obj.Key = os.Args[2]
	ask, err := json.Marshal(obj)
	if err != nil {
		log.Printf("result to json error:%v", err)
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://localhost%s/server/%s", c.path, cmd), "", bytes.NewBuffer(ask))
	if err != nil {
		// log.Printf("Get symbol init error:%s", err.Error())
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read body error:", err)
		return err
	}
	log.Println("json to struct error:", string(data))
	var ret Result
	if err := json.Unmarshal(data, &ret); err != nil {
		log.Println("json to struct error:", err, string(data))
		// log.Println("json to struct error:", err)
		return err
	}

	log.Println("json to struct error:", ret)

	return nil
}
