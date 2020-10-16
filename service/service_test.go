package service

import (
	"context"
	"testing"

	pb "github.com/takasho-tutorial/dummy-api/grpc"
	"go.uber.org/zap"
)

func TestNewService(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("Zap logging failed %v", err)
	}
	s, err := NewService(logger)

	if err != nil {
		t.Errorf("Service didn't work %v", err)
	}

	if _, ok := s.(pb.DummyApiServiceServer); ok == false {
		t.Errorf("Service didn't support DummyApiServer")
	}
}

func GetData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	s, _ := NewService(logger)

	arg := &pb.DataType{SensorName: "camera-module"}
	exp := &pb.SensingData{SensorName: "camera-module", SensorValues: []int32{0, 1, 0}}
	ret, err := s.GetData(context.Background(), arg)

	if err != nil {
		t.Errorf("Service didn't work %v", err)
	}

	if ret.SensorName != exp.SensorName {
		t.Errorf("GetData() exptected %v, got %v", "camera-module", ret.SensorName)
	}

	for idx, val := range ret.SensorValues {
		if exp.SensorValues[idx] != val {
			t.Errorf("GetData() exptected %v, got %v", exp.SensorValues[idx], val)
		}
	}
}

func GetStatsData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	s, _ := NewService(logger)

	arg := []*pb.DataType{
		{SensorName: "camera-module"},
		{SensorName: "voltage-module"},
		{SensorName: "spectrum-module"},
	}
	exp := &pb.StatsData{
		SensorName: "ML-module",
		StatsInfo:  0.8,
	}

	err := s.GetStatsData(context.Background())
	for _, val := range arg {
		if err := stream.Send(val); err != nil {
			t.Errorf("Error occured @ Sending data %v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Errorf("Error occured @ Recv data")
	}
	if (reply.SensorName != exp.SensorName) || (reply.StatsInfo != exp.StatsInfo) {
		t.Errorf("GetStatsData expects %v, but recieve %v", exp, reply)
	}
}

func GetSeqData(t *testing.T) {
}

func GetSeqStatsData(t *testing.T) {
}
