package main

import (
	"flag"
	"github.com/satori/go.uuid"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"syscall"
	"time"
)

var (
	signal = flag.String("s", "", `Send signal to the daemon:
quit — graceful shutdown
stop — fast shutdown
reload — reloading the configuration file`)
)

func main() {
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	cntxt := &daemon.Context{
		PidFileName: "/tmp/minion.pid",
		PidFilePerm: 0644,
		LogFileName: "/tmp/minion.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[minion]"},
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %s", err.Error())
		}
		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker(uuid.NewV4().String()[:8])

	err = daemon.ServeSignals()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Println("daemon terminated")
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func worker(tid string) {
	go func() { //业务逻辑在此处
		//defer recover()
		for {
			log.Printf("worker,tid[%s],%s", tid, time.Now().String())
			time.Sleep(time.Second)
		}
	}()

	//LOOP:
	for {
		log.Printf("loop,tid[%s],%s", tid, time.Now().String())
		time.Sleep(time.Second) // this is work to be done by worker.
		select {
		case <-stop:
			//break LOOP
			goto EXIT
		default:
		}
	}
EXIT:
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	log.Println("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

func reloadHandler(sig os.Signal) error {
	log.Println("configuration reloaded")
	return nil
}
