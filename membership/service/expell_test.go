package service

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	"github.com/maxpoletaev/kv/internal/grpcutil"
	"github.com/maxpoletaev/kv/membership"
	"github.com/maxpoletaev/kv/membership/proto"
)

func TestExpell(t *testing.T) {
	ctrl := gomock.NewController(t)
	memberRepo := NewMockMemberRegistry(ctrl)
	svc := NewMembershipService(memberRepo)
	ctx := context.Background()
	memberRepo.EXPECT().Expell(membership.NodeID(1)).Return(nil)
	_, err := svc.Expell(ctx, &proto.ExpellRequest{MemberId: 1})
	assert.NoError(t, err)
}

func TestExpellFails_MemberNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	memberRepo := NewMockMemberRegistry(ctrl)
	svc := NewMembershipService(memberRepo)
	ctx := context.Background()
	memberRepo.EXPECT().Expell(membership.NodeID(1)).Return(membership.ErrMemberNotFound)
	_, err := svc.Expell(ctx, &proto.ExpellRequest{MemberId: 1})
	assert.Equal(t, codes.NotFound, grpcutil.ErrorCode(err))
}