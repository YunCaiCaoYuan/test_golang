package main
//package integration_test

import (
	"context"
	"git.2tianxin.com/platform/tars-go/tars/servant"
	"git.2tianxin.com/platform/utils/goim/libs/define"
	"time"
	"yesserver/util/errors"
	"yesserver/util/protocol/pb"
	"yesserver/yes/channel/cache"
	"yesserver/yes/channel/dao"
	"yesserver/yes/channel/handler"
	"yesserver/yes/channel/local"
	"yesserver/yes/channel/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	var (
		CObj               = &handler.ChannelExtObj{}
		ctxData            = map[string]string{}
		channelId    int64 = 20089635
		manPlayerId        = "1041514" // 管理员
		playerIdJoin       = "1001020" // 已加入
		playerId1    int64 = 1040635   // 未加入
		playerId2    int64 = 1001003   // 未加入
		applyId      int64 = 1         // 申请id，拒绝
		applyId2     int64 = 2         // 申请id，同意
		applyId3     int64 = 11        // 不存在的申请
	)

	BeforeEach(func() {
		local.InitLocal()

	})

	JustBeforeEach(func() {
		obj := &model.ChannelJoinApply{Id: applyId, ChannelId: channelId, PlayerId: playerId1, Reason: "test", Status: pb.ChannelApplyStatus_CAS_APPLY, CreatedAt: time.Now()}
		Expect(dao.ChannelJoinApplyDB.Create(obj)).To(BeNil())
		obj2 := &model.ChannelJoinApply{Id: applyId2, ChannelId: channelId, PlayerId: playerId2, Reason: "test", Status: pb.ChannelApplyStatus_CAS_APPLY, CreatedAt: time.Now()}
		Expect(dao.ChannelJoinApplyDB.Create(obj2)).To(BeNil())

		cache.DelPlayerDailyApply(channelId, playerId1)
		cache.DelPlayerDailyApply(channelId, playerId2)
	})

	Describe("ApplyJoin", func() {
		Context("ApplyJoin-申请理由为空", func() {
			It("ApplyJoin-申请理由为空", func() {
				req := &pb.ApplyJoinReq{ChannelId: channelId}
				ctxData[define.UID] = playerIdJoin
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				_, err := CObj.ApplyJoin(ctx, req)
				Expect(err).To(Equal(errors.ErrApplyReasonEmpty))
			})
		})

		Context("ApplyJoin-已经在频道中", func() {
			It("ApplyJoin-已经在频道中", func() {
				req := &pb.ApplyJoinReq{ChannelId: channelId, Reason: "test"}
				ctxData[define.UID] = playerIdJoin
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				_, err := CObj.ApplyJoin(ctx, req)
				Expect(err).To(Equal(errors.ErrPlayerInChannel))
			})
		})

	})

	Describe("处理申请", func() {
		Context("处理申请-通过", func() {
			It("处理申请-通过-记录已被处理", func() {
				ctxData[define.UID] = manPlayerId
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				req := &pb.HandleChannelApplyReq{Type: pb.ChannelHandleType_CHT_PASS, ApplyId: applyId3}
				_, err := CObj.HandleChannelApply(ctx, req)
				Expect(err).To(Equal(errors.ErrApplyNotFound))
			})
		})
		Context("处理申请-通过", func() {
			It("处理申请-通过", func() {
				ctxData[define.UID] = manPlayerId
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				req := &pb.HandleChannelApplyReq{Type: pb.ChannelHandleType_CHT_PASS, ApplyId: applyId2}
				_, err := CObj.HandleChannelApply(ctx, req)
				Expect(err).To(BeNil())
			})
		})

		Context("处理申请-拒绝", func() {
			It("处理申请-拒绝", func() {
				ctxData[define.UID] = manPlayerId
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				req := &pb.HandleChannelApplyReq{Type: pb.ChannelHandleType_CHT_REFUSE, ApplyId: applyId}
				_, err := CObj.HandleChannelApply(ctx, req)
				Expect(err).To(BeNil())
			})
		})

	})

	JustAfterEach(func() {
		Expect(dao.ChannelJoinApplyDB.Del(applyId)).To(BeNil())
		Expect(dao.ChannelJoinApplyDB.Del(applyId2)).To(BeNil())
	})

})

