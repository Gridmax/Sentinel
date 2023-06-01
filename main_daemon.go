package main

import (
  "log"
  "flag"
  "syscall"
  "os"
  "time"
  "fmt"
  "github.com/Gridmax/Sentinel/commun/client"	

//  "github.com/Gridmax/Sentinel/utility/logging"
	"github.com/sevlyar/go-daemon"
)

//func main() {
	// Start the client
//  client.Start("config.yaml")

//}

var (
	signal = flag.String("s", "", `Send signal to the daemon:
  quit — graceful shutdown
  stop — fast shutdown
  reload — reloading the configuration file`)
)
var (
	stop = make(chan struct{})
	done = make(chan struct{})
)
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

func main() {

	flag.Parse()


	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)
	
  fmt.Println("mok ", *signal)
  cntxt := &daemon.Context{
		PidFileName: "SentinelAgent.pid",
		PidFilePerm: 0644,
		LogFileName: "SentinelAgent.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

  if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
    fmt.Println("ok",d)
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %s", err.Error())
		}
	  daemon.SendCommands(d)
		return
	}
  fmt.Println("check", daemon.ActiveFlags())

	d, err := cntxt.Reborn()
	if err != nil {

		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

  log.Println("- - - - - - - - - - - - - - -")
	log.Println("Sentinel Agent started")
  client.Start("config.yaml")
  go worker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Println("daemon terminated")

}

func worker() {
LOOP:
	for {
		time.Sleep(time.Second) // this is work to be done by worker.
		select {
		case <-stop:
			break LOOP
		default:
		}
	}
	done <- struct{}{}
}
