package test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zippunov/protorpc"
	"github.com/zippunov/protorpc/test/generated/parking"
	"google.golang.org/protobuf/proto"
)

/*
type ParkingServiceService interface {
	CurrentCapacity(ctx context.Context, in *empty.Empty) (*CapacityResponse, error)

	TakeSlot(ctx context.Context, in *PlateRequest) (*TakeSlotResponse, error)

	CurrentBilling(ctx context.Context, in *PlateRequest) (*CurrentBillingResponse, error)

	FreeSlot(ctx context.Context, in *PlateRequest) (*FreeSlotResponse, error)
}
*/

type ParkingImpl struct {
}

func (i ParkingImpl) CurrentCapacity(ctx context.Context, in *empty.Empty) (*parking.CapacityResponse, protorpc.StatusError) {
	return &parking.CapacityResponse{
		AvailableSlots: 1,
		ReservedSlots:  1,
		TakenSlots:     20,
		TotalSlots:     22,
	}, nil
}

func (i ParkingImpl) TakeSlot(ctx context.Context, in *parking.PlateRequest) (*parking.TakeSlotResponse, protorpc.StatusError) {
	return &parking.TakeSlotResponse{
		Code:      true,
		Reason:    "",
		StartTime: time.Now().Unix(),
		Success:   true,
	}, nil
}

func (i ParkingImpl) CurrentBilling(ctx context.Context, in *parking.PlateRequest) (*parking.CurrentBillingResponse, protorpc.StatusError) {
	return nil, protorpc.RPCError{
		Msg:  "Not implemented",
		Code: 4,
	}
}

func (i ParkingImpl) FreeSlot(ctx context.Context, in *parking.PlateRequest) (*parking.FreeSlotResponse, protorpc.StatusError) {
	return nil, protorpc.RPCError{
		Msg:  "Not implemented",
		Code: 4,
	}
}

func TestRPCDispatcher(t *testing.T) {
	dispatcher, err := protorpc.BuildService(&parking.ParkingServiceServiceDescriptor, ParkingImpl{})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var serviceName string
	dispatcher.Use(func(d protorpc.RPCDispatcher, method string, ctx context.Context, in proto.Message) (context.Context, protorpc.StatusError) {
		serviceName = d.Name()
		return ctx, nil
	})
	rslt, err := dispatcher.RPC("CurrentCapacity", context.Background(), &empty.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}
	body, _ := json.Marshal(rslt)
	t.Log(rslt)
	t.Log(string(body))
	if serviceName != "parking.ParkingService" {
		t.Errorf("Expected serviceName be equal to %s", "parking.ParkingService")
	}
	if dispatcher.Name() != "parking.ParkingService" {
		t.Errorf("Expected dispatcher Name be equal to %s", "parking.ParkingService")
	}
}
