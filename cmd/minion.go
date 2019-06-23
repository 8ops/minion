package main

import (
	"flag"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sevlyar/go-daemon"
	"k8s.io/klog"
	"os"
	"syscall"
	"time"
)

var (
	version = "0.0.1-20190623"
	verbose = flag.Bool("v", false, "Show version number and quit")
	signal  = flag.String("s", "", `Send signal to the daemon:
quit — graceful shutdown
stop — fast shutdown
reload — reloading the configuration file`)

	stop = make(chan struct{})
	done = make(chan struct{})
)

func main() {
	klog.InitFlags(&flag.FlagSet{})
	defer klog.Flush()

	flag.Usage = func() {
		fmt.Printf("minion/%s\nUsage: minion [-vh] [-s signal] \n\n", version)
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	//show version
	if *verbose {
		fmt.Printf("impower/%s\n", version)
		os.Exit(0)
	}

	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	ctx := &daemon.Context{
		PidFileName: "/tmp/minion.pid",
		PidFilePerm: 0644,
		LogFileName: "/tmp/minion.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[minion]"},
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := ctx.Search()
		if err != nil {
			klog.Fatalf("Unable send signal to the daemon: %s\n", err.Error())
		}
		daemon.SendCommands(d)
		return
	}

	d, err := ctx.Reborn()
	if err != nil {
		klog.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer ctx.Release()

	klog.Info("Daemon started")

	go worker(uuid.NewV4().String()[:8])

	err = daemon.ServeSignals()
	if err != nil {
		klog.Errorf("Error: %s", err.Error())
	}

	klog.Info("Daemon terminated")
}

func worker(tid string) {
	//process
	go func() {
		//defer recover()
		for {
			klog.Infof("Worker[process],tid[%s],%s", tid, time.Now().String())
			time.Sleep(time.Second)
		}
	}()

	for {
		time.Sleep(time.Second)
		select {
		case <-stop:
			goto EXIT
		default:
		}
	}
EXIT:
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	klog.Info("Terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

func reloadHandler(sig os.Signal) error {
	klog.Info("Configuration reloaded")
	return nil
}
