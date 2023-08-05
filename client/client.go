package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/jibudata/archive-api"
)

type ArchivalClient struct {
	addr   string
	client pb.ActiveArchiveClient
}

type Pool struct {
	Poolname string
	Total    uint64
	Free     uint64
	Unref    uint64
	Numtapes uint64
}

type AsyncStatus struct {
	ReqNumber   int64
	Success     bool
	Done        bool
	Resident    int64
	Transferred int64
	Premigrated int64
	Migrated    int64
	Failed      int64
}

type FileInfo struct {
	MigState   string
	FileName   string
	Size       uint64
	Blocks     uint64
	Fsidh      uint64
	Fsidl      uint64
	Igen       uint32
	Inum       uint64
	TapeId     string
	StartBlock uint64
}

func NewArchivalClient(addr string) (front *ArchivalClient, e error) {
	front = &ArchivalClient{addr: addr}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		e = err
		return
	}
	// defer conn.Close()
	client := pb.NewActiveArchiveClient(conn)
	front.client = client
	return
}

func (f *ArchivalClient) GetMediaInfo(tapeName string) (tapeInfo string, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	r, err := f.client.GetMediaInfo(ctx, &pb.LibraryManagerResourceKey{Name: tapeName})
	if err != nil {
		log.Fatalf("could not get tape info: %v", err)
	}
	return r.GetInfo(), err
}

func (f *ArchivalClient) GetPoolsInfo() ([]*Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	reply, err := f.client.GetPoolsInfo(ctx, &pb.DefaultResourceRequest{})
	if err != nil {
		return []*Pool{}, err
	} else {
		poolInfos := []*Pool{}
		for _, pool := range reply.Pools {
			info := &Pool{Poolname: pool.PoolName, Total: pool.Total, Free: pool.Free, Unref: pool.Unref, Numtapes: pool.NumTapes}
			poolInfos = append(poolInfos, info)
		}
		return poolInfos, err
	}

}

func (f *ArchivalClient) GetPoolInfo(name string) (Pool, bool) {
	pools, err := f.GetPoolsInfo()
	if err != nil {
		return Pool{}, false
	} else {
		for _, pool := range pools {
			if name == pool.Poolname {
				return *pool, true
			}
		}
		return Pool{}, false
	}
}

func (f *ArchivalClient) Migrate(file string, poolName string) (success bool, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(6000)*time.Second)
	defer cancel()
	status, err := f.client.Migrate(ctx, &pb.MigrateRequest{
		Pool: &pb.LibraryManagerResourceKey{
			Name: poolName,
		},
		Files: []string{file}})
	return status.Success, err
}

func (f *ArchivalClient) Recall(file string, resident bool) (success bool, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3000)*time.Second)
	defer cancel()
	status, err := f.client.Recall(ctx, &pb.RecallRequest{Resident: resident, Files: []string{file}})
	return status.Success, err
}

func (f *ArchivalClient) Retrieve() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3000)*time.Second)
	defer cancel()
	_, err := f.client.Retrieve(ctx, &pb.DefaultResourceRequest{})
	return err
}

func (f *ArchivalClient) MigrateAsync(file string, poolName string) (reqNumber int64, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	status, err := f.client.MigrateAsync(ctx, &pb.MigrateRequest{
		Pool: &pb.LibraryManagerResourceKey{
			Name: poolName,
		}, Files: []string{file}})
	return status.RequestNumber, err
}

func (f *ArchivalClient) RecallAsync(file string, resident bool) (reqNumber int64, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	status, err := f.client.RecallAsync(ctx, &pb.RecallRequest{Resident: resident, Files: []string{file}})
	return status.RequestNumber, err
}

func (f *ArchivalClient) GetAsyncStatus(reqNumber int64) (AsyncStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	status, err := f.client.GetAsyncStatus(ctx, &pb.AsyncStatusRequest{RequestNumber: reqNumber})
	if err != nil {
		log.Fatalf("could not get async status: %v", err)
		return AsyncStatus{}, err
	}
	return AsyncStatus{
		ReqNumber:   reqNumber,
		Success:     status.Success,
		Resident:    status.Resident,
		Transferred: status.Resident,
		Premigrated: status.Premigrated,
		Migrated:    status.Migrated,
		Failed:      status.Failed,
		Done:        status.Done}, err
}

func (f *ArchivalClient) GetFileInfo(fileName string) (FileInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	fileInfo, err := f.client.GetFileInfo(ctx, &pb.FileInfoRequest{FileName: fileName})
	if err != nil {
		log.Fatalf("Get file info failed with err: %v", err)
		return FileInfo{}, err
	}
	return FileInfo{
		MigState:   fileInfo.MigrationState,
		FileName:   fileInfo.FileName,
		Size:       fileInfo.Size,
		Blocks:     fileInfo.Blocks,
		Fsidh:      fileInfo.Fsidh,
		Fsidl:      fileInfo.Fsidl,
		Igen:       fileInfo.Igen,
		Inum:       fileInfo.Inum,
		TapeId:     fileInfo.TapeId,
		StartBlock: fileInfo.StartBlock}, err
}
