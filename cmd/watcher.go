package main

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

func main() {
	w := watcher.New()
	w.FilterOps(watcher.Create, watcher.Remove)

	go func() {

		for {
			select {
			case event := <-w.Event:
				if event.IsDir() {
					continue
				}
				log.Println(event.Path)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	watchPath := "C:/watch"
	if err := w.AddRecursive(watchPath); err != nil {
		log.Fatalln(err)
	}

	log.Println("Watching:", watchPath)

	if err := w.Start(time.Duration(time.Second * 1)); err != nil {
		log.Fatalln(err)
	}
}
