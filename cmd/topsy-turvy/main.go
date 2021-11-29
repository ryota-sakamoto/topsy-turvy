package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/raft"

	"github.com/ryota-sakamoto/topsy-turvy/pkg/config"
	"github.com/ryota-sakamoto/topsy-turvy/pkg/fsm"
)

func main() {
	log.Println("init")

	apiConfig, err := config.New()
	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", apiConfig)

	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(apiConfig.ServerID)

	st := raft.NewInmemStore()
	sst := raft.NewInmemSnapshotStore()

	addr, tm := raft.NewInmemTransport(raft.NewInmemAddr())
	log.Println(addr)

	fsm := fsm.FSM{}
	r, err := raft.NewRaft(c, fsm, st, st, sst, tm)
	if err != nil {
		panic(err)
	}

	r.BootstrapCluster(raft.Configuration{
		Servers: []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       raft.ServerID(apiConfig.ServerID),
				Address:  raft.ServerAddress(addr),
			},
		},
	})

	ctx := context.Background()
	go func() {
		log.Println("run writer")

		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case a := <-ticker.C:
				r.Apply([]byte(a.String()), time.Second)
			case <-ctx.Done():
				log.Println("close writer")
				return
			}
		}
	}()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	<-ctx.Done()
	if err := r.Shutdown().Error(); err != nil {
		panic(err)
	}

	log.Println("shutdown")
}
