package tests

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	gen "github.com/zippunov/protorpc/tests/generated"
)

/**
type ParkingServiceService interface {
	CurrentCapacity(ctx context.Context, in *empty.Empty) (*CapacityResponse, error)

	TakeSlot(ctx context.Context, in *PlateRequest) (*TakeSlotResponse, error)

	CurrentBilling(ctx context.Context, in *PlateRequest) (*CurrentBillingResponse, error)

	FreeSlot(ctx context.Context, in *PlateRequest) (*FreeSlotResponse, error)
}
*/

type ParkingImplementation struct {
}

func (i ParkingImplementation) CurrentCapacity(ctx context.Context, in *empty.Empty) (*gen.CapacityResponse, error) {
	return &gen.CapacityResponse{
		AvailableSlots: 1,
		ReservedSlots:  3,
		TakenSlots:     20,
		TotalSlots:     24,
	}, nil
}

func (i ParkingImplementation) TakeSlot(ctx context.Context, in *gen.PlateRequest) (*gen.TakeSlotResponse, error) {
	return &gen.TakeSlotResponse{
		Code:      true,
		Reason:    "",
		StartTime: time.Now().Unix(),
		Success:   true,
	}, nil
}

func (i ParkingImplementation) CurrentBilling(ctx context.Context, in *gen.PlateRequest) (*gen.CurrentBillingResponse, error) {
	return &gen.CurrentBillingResponse{
		Billing: ,
	}, nil
}
