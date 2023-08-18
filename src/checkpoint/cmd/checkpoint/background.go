package main

import (
	"context"
	"log"
	"strconv"
	"time"

	chckp "bitbucket.org/veldrane/golibs/checkpoint"
)

func handleBackgroundGatherer(ctx context.Context, errb chan error, logger *log.Logger, debug bool, rules *chckp.RulesT, period int) {

	var err error

	go func() {
		var e chckp.Egress
		var i chckp.Ingress
		for true {
			s := chckp.Session()
			rules.Locks.Egress.Lock()
			rules.Egress, err = e.Rules(s, ctx) // Getting egress rules
			rules.Locks.Egress.Unlock()

			rules.Locks.Ingress.Lock()
			rules.Ingress, err = i.Rules(s, ctx) // Getting Ingress rules
			rules.Locks.Ingress.Unlock()

			if err != nil {
				logger.Fatalf("[ Scraper Thread ] -> %s", err)
			}
			logger.Printf("[ Scraper Thread ] -> New definition has been loaded from OCP")
			time.Sleep(time.Duration(period) * time.Second)
		}
	}()

	logger.Printf("[ Scraper Thread ] -> started sucessfully with period %s seconds", strconv.Itoa(period))
	return
}
