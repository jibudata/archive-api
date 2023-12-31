package client

import (
	"context"
	"os"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/jibudata/archive-api"
)

const (
	DefaultMigrationTimeout = 5 * time.Minute
	DefaultRecallTimeout    = 15 * time.Minute
	DefaultRequestTimeout   = 30 * time.Second
)

type ArchivalClient struct {
	addr   string
	client pb.ActiveArchiveClient
	logger *slog.Logger
}

type Pool struct {
	PoolName    string
	Total       uint64
	Free        uint64
	Reclaimable uint64
	NumberTapes uint64
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
	MigState          string
	FileName          string
	Size              uint64
	Blocks            uint64
	FilesystemUidHigh uint64
	FilesystemUidLow  uint64
	InodeGeneration   uint32
	InodeNumber       uint64
	TapeId            string
	StartBlock        uint64
}

type TapeInfo struct {
	Id             string
	Slot           uint64
	TotalCapacity  uint64
	RemainCapacity uint64
	Reclaimable    uint64
	Status         string
	Inprogress     uint64
	Pool           string
	State          string
}

func NewArchivalClient(addr string) (*ArchivalClient, error) {
	archiveClient := &ArchivalClient{addr: addr}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewActiveArchiveClient(conn)
	archiveClient.client = client
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	archiveClient.logger = logger.With("client", addr)

	return archiveClient, nil
}

func (f *ArchivalClient) GetMediaInfo(tapeName string) (*TapeInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRequestTimeout))
	defer cancel()
	media, err := f.client.GetMediaInfo(ctx, &pb.LibraryManagerResourceKey{Name: tapeName})
	if err != nil {
		f.logger.Error("could not get tape info, error: ", err)
		return nil, err
	}

	tape := media.GetTape()
	tapeInfo := &TapeInfo{
		Id:             tape.GetId(),
		Slot:           tape.GetSlot(),
		TotalCapacity:  tape.GetTotalCapacity(),
		RemainCapacity: tape.GetRemainCapacity(),
		Reclaimable:    tape.GetReclaimable(),
		Status:         tape.GetStatus(),
		Inprogress:     tape.GetInprogress(),
		Pool:           tape.GetPool(),
		State:          tape.GetState(),
	}
	return tapeInfo, err
}

func (f *ArchivalClient) GetPoolsInfo() ([]*Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRequestTimeout))
	defer cancel()

	reply, err := f.client.GetPoolsInfo(ctx, &pb.DefaultResourceRequest{})
	if err != nil {
		return []*Pool{}, err
	}

	poolInfos := []*Pool{}
	for _, pool := range reply.Pools {
		info := &Pool{
			PoolName:    pool.PoolName,
			Total:       pool.Total,
			Free:        pool.Free,
			Reclaimable: pool.Reclaimable,
			NumberTapes: pool.NumberTapes}
		poolInfos = append(poolInfos, info)
	}
	return poolInfos, err
}

func (f *ArchivalClient) GetPoolInfo(name string) (*Pool, error) {
	pools, err := f.GetPoolsInfo()
	if err != nil {
		return nil, err
	}

	for _, pool := range pools {
		if name == pool.PoolName {
			return pool, nil
		}
	}
	return nil, nil
}

func (f *ArchivalClient) Migrate(file string, poolName string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultMigrationTimeout))
	defer cancel()

	status, err := f.client.Migrate(ctx,
		&pb.MigrateRequest{
			Pool: &pb.LibraryManagerResourceKey{
				Name: poolName,
			},
			Files: []string{file},
		})

	return status.Success, err
}

func (f *ArchivalClient) Recall(file string, resident bool) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRecallTimeout))
	defer cancel()

	status, err := f.client.Recall(ctx, &pb.RecallRequest{Resident: resident, Files: []string{file}})
	return status.Success, err
}

func (f *ArchivalClient) Retrieve() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRequestTimeout))
	defer cancel()
	_, err := f.client.Retrieve(ctx, &pb.DefaultResourceRequest{})
	return err
}

func (f *ArchivalClient) MigrateAsync(file string, poolName string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultMigrationTimeout))
	defer cancel()

	status, err := f.client.MigrateAsync(ctx,
		&pb.MigrateRequest{
			Pool: &pb.LibraryManagerResourceKey{
				Name: poolName,
			}, Files: []string{file},
		})

	return status.RequestNumber, err
}

func (f *ArchivalClient) RecallAsync(file string, resident bool) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRecallTimeout))
	defer cancel()

	status, err := f.client.RecallAsync(ctx,
		&pb.RecallRequest{
			Resident: resident,
			Files:    []string{file},
		})
	return status.RequestNumber, err
}

func (f *ArchivalClient) GetAsyncStatus(reqNumber int64) (*AsyncStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRequestTimeout))
	defer cancel()

	status, err := f.client.GetAsyncStatus(ctx, &pb.AsyncStatusRequest{RequestNumber: reqNumber})
	if err != nil {
		f.logger.Error("could not get async status, error: ", err)
		return nil, err
	}

	return &AsyncStatus{
		ReqNumber:   reqNumber,
		Success:     status.Success,
		Resident:    status.Resident,
		Transferred: status.Resident,
		Premigrated: status.Premigrated,
		Migrated:    status.Migrated,
		Failed:      status.Failed,
		Done:        status.Done}, nil
}

func (f *ArchivalClient) GetFileInfo(fileName string) (*FileInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(DefaultRequestTimeout))
	defer cancel()

	fileInfo, err := f.client.GetFileInfo(ctx, &pb.FileInfoRequest{FileName: fileName})
	if err != nil {
		f.logger.Error("Get file info failed with err: ", err)
		return nil, err
	}

	return &FileInfo{
		MigState:          fileInfo.MigrationState,
		FileName:          fileInfo.FileName,
		Size:              fileInfo.Size,
		Blocks:            fileInfo.Blocks,
		FilesystemUidHigh: fileInfo.FilesystemUidHigh,
		FilesystemUidLow:  fileInfo.FilesystemUidLow,
		InodeGeneration:   fileInfo.InodeGeneration,
		InodeNumber:       fileInfo.InodeNumber,
		TapeId:            fileInfo.TapeId,
		StartBlock:        fileInfo.StartBlock}, nil
}
