package fsm

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
)

type FSM struct {
}

func (FSM) Apply(l *raft.Log) interface{} {
	log.Printf("Apply: %+v\n", string(l.Data))
	return nil
}

func (FSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("Snapshot")
	return nil, nil
}

func (FSM) Restore(io.ReadCloser) error {
	log.Println("Restore")
	return nil
}
