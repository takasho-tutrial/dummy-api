package service

import (
	"context"

	pb "github.com/takasho-tutorial/dummy-api/grpc"
	"go.uber.org/zap"
)

type service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) (pb.DummyApiServiceServer, error) {
	return &service{
		logger: logger,
	}, nil
}

func (s *service) GetData(ctx context.Context, dataType *pb.DataType) (*pb.SensingData, error) {
	return &pb.SensingData{
		SensorName:   dataType.SensorName,
		SensorValues: []int32{0, 1, 0},
	}, nil
	// switch dataType.SensorName {
	// case "camera-module":
	// 	return &pb.SensingData{
	// 		SensorName:   dataType.SensorName,
	// 		SensorValues: []int32{0, 1, 0},
	// 	}, nil
	// case "impedance-module":
	// 	return &pb.SensingData{
	// 		SensorName:   dataType.SensorName,
	// 		SensorValues: []int32{1, 2, 1},
	// 	}, nil
	// }
}

func (s *service) GetStatsData(stream pb.DummyApiService_GetStatsDataServer) error {

	return nil
}

func (s *service) GetSeqData(dataType *pb.DataType, stream pb.DummyApiService_GetSeqDataServer) error {
	return nil
}

func (s *service) GetSeqStatsData(stream pb.DummyApiService_GetSeqStatsDataServer) error {
	return nil
}
