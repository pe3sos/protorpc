package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zippunov/protorpc"
	"github.com/zippunov/protorpc/test/generated/parking"
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

func (i ParkingImpl) CurrentCapacity(ctx context.Context, in *empty.Empty) (*parking.CapacityResponse, error) {
	return &parking.CapacityResponse{
		AvailableSlots: 1,
		ReservedSlots:  1,
		TakenSlots:     20,
		TotalSlots:     22,
	}, nil
}

func (i ParkingImpl) TakeSlot(ctx context.Context, in *parking.PlateRequest) (*parking.TakeSlotResponse, error) {
	return &parking.TakeSlotResponse{
		Code:      true,
		Reason:    "",
		StartTime: time.Now().Unix(),
		Success:   true,
	}, nil
}

func (i ParkingImpl) CurrentBilling(ctx context.Context, in *parking.PlateRequest) (*parking.CurrentBillingResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (i ParkingImpl) FreeSlot(ctx context.Context, in *parking.PlateRequest) (*parking.FreeSlotResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

func TestRPCDispatcher(t *testing.T) {
	implementor, err := protorpc.BuildService(&parking.ParkingServiceServiceDescriptor, ParkingImpl{})
	if err != nil {
		t.Errorf(err.Error())
	}
	rslt, err := implementor.RPC("CurrentCapacity", context.Background(), &empty.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}
	body, _ := json.Marshal(rslt)
	t.Log(rslt)
	t.Log(string(body))
}
