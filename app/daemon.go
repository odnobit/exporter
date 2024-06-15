package app

import (
	"log"
	"sync"
	"time"

	"github.com/odnobit/exporter/storage"
)

type AppDaemon struct{}

func (a *AppDaemon) doTaskOK() {
	log.Println("doing some job")
	time.Sleep(time.Second)
	log.Println("task done")
}

func (a *AppDaemon) doTaskFail() {
	log.Println("doing some job")
	time.Sleep(time.Second * 3)
	log.Println("task failed")
}

func (a *AppDaemon) Run() {
	var wg sync.WaitGroup

	// emulate work
	ticker := time.NewTicker(time.Second * 15)
	go func() {
		for ; true; <-ticker.C {
			wg.Add(2)
			go func(waitg *sync.WaitGroup) {
				defer waitg.Done()
				defer storage.Memory.IncSuccessMessages()
				defer storage.Memory.IncTotalMessages()

				a.doTaskOK()
			}(&wg)
			go func(waitg *sync.WaitGroup) {
				defer waitg.Done()
				defer storage.Memory.IncFailedMessages()
				defer storage.Memory.IncTotalMessages()

				a.doTaskFail()
			}(&wg)
			wg.Wait()
			log.Println("all tasks done, waiting for the tick")
		}
	}()

	select {}
}

func NewAppDaemon() *AppDaemon {
	return &AppDaemon{}
}
