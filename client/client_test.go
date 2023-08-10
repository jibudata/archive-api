package client

import (
	"log"
	"testing"
)

func TestArchivalClient(t *testing.T) {

	addr := "localhost:50051"
	tapeName := "Tape1"
	srcPath := "~/haha.txt"
	archiveClient, err := NewArchivalClient(addr)
	if err != nil {
		log.Fatalf("Create archiveClient API failed")
	}
	if tapeName != "" {
		info, err := archiveClient.GetMediaInfo(tapeName)
		if err != nil {
			log.Fatalf("Get tape info failed")
		}
		log.Printf("Tape info: %v", info)
	}

	if srcPath != "" {
		re, err := archiveClient.Migrate(srcPath, "pool0")
		if err != nil {
			log.Fatalf("Migration failed")
		}
		log.Printf("Dst path: %v", re)
	}
}

func TestArchivalAPI(t *testing.T) {

	addr := "localhost:50051"
	tapeName := "Tape1"
	srcPath := "~/haha.txt"
	archiveClient, err := NewArchivalClient(addr)
	if err != nil {
		log.Fatalf("Create archiveClient API failed")
	}
	if tapeName != "" {
		info, err := archiveClient.GetMediaInfo(tapeName)
		if err != nil {
			log.Fatalf("Get tape info failed")
		}
		log.Printf("Tape info: %v", info)
	}

	pool, err := archiveClient.GetPoolInfo(tapeName)
	if err != nil {
		log.Printf("Can not find the pool")
	} else {
		log.Printf("pool: %v", pool)
	}

	_, e := archiveClient.Migrate(srcPath, "pool0")
	if e != nil {
		log.Printf("Migrate success")
	} else {
		log.Printf("Migrate failed")
	}

	_, e1 := archiveClient.Recall(srcPath, true)
	if e1 != nil {
		log.Printf("Recall success")
	} else {
		log.Printf("Recall failed")
	}

	pool, e2 := archiveClient.GetPoolInfo("pool0")
	if e2 != nil {
		log.Printf("Can not find the pool")
	} else {
		log.Printf("pool: %v", pool)
	}
}
