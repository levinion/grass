package app

import (
	"os"
	"os/exec"
	"strings"
)

type App struct {
	name   string
	args   []string
	cmd    *exec.Cmd
	status int
}

// eg: use CreateApp("waybar") to create an app.
func CreateApp(command string) *App {
	initWaylandEnvironment()
	args := strings.Split(command, " ")
	return &App{args[0], args[1:], nil, Stop}
}

func (app *App) Start() {
	if app.status == Stop {
		app.cmd = exec.Command(app.name, app.args...)
		app.cmd.Start()
		app.status = Running
	}
}

func (app *App) Stop() {
	if app.status == Running {
		app.cmd.Process.Signal(os.Interrupt)
		app.cmd.Process.Release()
		app.status = Stop
	}
}

func initWaylandEnvironment() {
	os.Setenv("GDK_BACKEND", "wayland")
}
