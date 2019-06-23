package config

import (
	uuid "github.com/satori/go.uuid"
	"k8s.io/klog"
	"os"
	"path/filepath"
)

var (
	HomeDir = filepath.Join(os.Getenv("HOME"), ".minion")
	UtilDir = filepath.Join(HomeDir, "util")
	LogDir  = filepath.Join(HomeDir, "log")

	MinionId string

	TaskPool = 10

	APITask     = "https://api.8ops.top/minion/task/default.js"
	APICallback = "https://api.8ops.top/minion/callback.js"
)

func init() {
	buildHome()
}

func buildHome() {
	//HomeDir
	if !exist(HomeDir) {
		os.Mkdir(HomeDir, 0700)
	}

	//UtilDir
	if !exist(UtilDir) {
		os.Mkdir(UtilDir, 0700)
	}

	//Logdir
	if !exist(LogDir) {
		os.Mkdir(LogDir, 0700)
	}

	//MinionId
	minionIdFile := filepath.Join(HomeDir, "id")
	if exist(minionIdFile) {
		if f, e1 := os.OpenFile(minionIdFile, os.O_RDONLY, 0400); e1 != nil {
			klog.Fatalln(e1)
		} else {
			if f != nil {
				defer f.Close()
			}
			b := make([]byte, 8)
			if _, e2 := f.Read(b); e2 != nil {
				klog.Fatalln(e2)
			}
			MinionId = string(b)

		}
	} else {
		MinionId = uuid.NewV4().String()[:8]
		if f, e1 := os.OpenFile(minionIdFile, os.O_CREATE|os.O_WRONLY, 0400); e1 != nil {
			klog.Fatal(e1)
		} else {
			if f != nil {
				defer f.Close()
			}
			f.WriteString(MinionId)
		}
	}
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
