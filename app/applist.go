package app

import (
	"fmt"
)

type AppList struct {
	list map[string]*App
}

func CreateAppList() *AppList {
	return &AppList{
		make(map[string]*App),
	}
}

func (al *AppList) Add(command string) string {
	app := CreateApp(command)
	if _, ok := al.list[app.name]; ok {
		return fmt.Sprintf("%s exists!", app.name)
	}
	al.list[app.name] = app
	return "ok"
}

func (al *AppList) AddThenStart(command string) string {
	app := CreateApp(command)
	if _, ok := al.list[app.name]; ok {
		return fmt.Sprintf("%s exists!", app.name)
	}
	al.list[app.name] = app
	app.Start()
	return "ok"
}

func (al *AppList) Stop(name string) string{
	al.list[name].Stop()
	return "ok"
}

func (al *AppList) Remove(name string)string {
	al.list[name].Stop()
	delete(al.list, name)
	return "ok"
}

func (al *AppList) Start(name string) string{
	al.list[name].Start()
	return "ok"
}

func (al *AppList) Reload(name string) string{
	al.list[name].Stop()
	al.list[name].Start()
	return "ok"
}

func (al *AppList) Show() string {

	if len(al.list)==0{
		return "Nothing is under control!"
	}

	// 打印表头
	s := fmt.Sprintf("%-15s %s\n", "App Name", "Status")
	s += fmt.Sprintln("---------------------------------")

	// 遍历每个App打印信息
	for name, app := range al.list {
		var status string
		if app.status == Stop {
			status = "STOPPED"
		} else {
			status = "RUNNING"
		}
		s += fmt.Sprintf("%-15s %s\n", name, status)
	}
	return s
}
