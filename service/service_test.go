package service

import (
	"context"
	pb "github.com/takasho-tutrial/dummy-api/grpc/dummy-api.pb.go"
	"go.uber.org/zap"
	"testing"
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
}

func GetData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	s, _ := NewService(logger)

	arg := pb.DataType{SensorName: "camera-module"}
	exp := []int32{0, 1, 0}
	ret, err := s.GetData(arg)

	if err != nil {
		t.Errorf("Service didn't work %v", err)
	}

	if ret.SensorName != "camera-module" {
		t.Errorf("GetData() exptected %v, got %v", "camera-module", ret.SensorName)
	}

	for idx, val := range ret.SensorValues {
		if exp[idx] != val {
			t.Errorf("GetData() exptected %v, got %v", exp[idx], val)
		}
	}
}

func GetStatsData(t *testing.T) {
}

func GetSeqData(t *testing.T) {
}

func GetSeqStatsData(t *testing.T) {
}
