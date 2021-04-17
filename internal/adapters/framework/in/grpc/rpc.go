package rpc

import (
	"hex-arch-go-grpc/internal/adapters/framework/in/grpc/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAddition returns the result of adding two numbers
func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required or one or more operation parameters set to 0")
	}

	answer, err := grpca.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected server error occurred")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

// GetSubtraction returns the result of subtracting one number from the other
func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required or one or more operation parameters set to 0")
	}

	answer, err := grpca.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected server error occurred")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

// GetMultiplication returns the result of multiplying one number by the other
func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required or one or more operation parameters set to 0")
	}

	answer, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected server error occurred")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

// GetDivision returns the result of dividing one number by the other
func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required or one or more operation parameters set to 0")
	}

	answer, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected server error occurred")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
