package integration_test

import (
	"context"
	"git.2tianxin.com/platform/tars-go/tars/servant"
	"git.2tianxin.com/platform/utils/goim/libs/define"
	"strconv"
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
		CObj                                = &handler.ChannelExtObj{}
		ctxData                             = map[string]string{}
		channelId                     int64 = 20089635
		manPlayerId                         = "1041514" // 管理员
		noPermPlayerId                int64 = 1001091
		playerIdJoin                        = "1001020"  // 已加入
		applyId3                      int64 = 11         // 不存在的申请
		playerIdPass, applyIdPass     int64 = 1001003, 2 // 申请id，同意
		playerIdRefuse, applyIdRefuse int64 = 1040635, 1 // 申请id，拒绝
		playerIdHasApply, applyIdHas  int64 = 2001186, 4 // 已经有申请
	)

	BeforeSuite(func() {
		local.InitLocal()

		obj := &model.ChannelJoinApply{Id: applyIdRefuse, ChannelId: channelId, PlayerId: playerIdRefuse, Reason: "test", Status: pb.ChannelApplyStatus_CAS_APPLY, CreatedAt: time.Now()}
		Expect(dao.ChannelJoinApplyDB.Create(obj)).To(BeNil())
		obj2 := &model.ChannelJoinApply{Id: applyIdPass, ChannelId: channelId, PlayerId: playerIdPass, Reason: "test", Status: pb.ChannelApplyStatus_CAS_APPLY, CreatedAt: time.Now()}
		Expect(dao.ChannelJoinApplyDB.Create(obj2)).To(BeNil())
		obj4 := &model.ChannelJoinApply{Id: applyIdHas, ChannelId: channelId, PlayerId: playerIdHasApply, Reason: "test", Status: pb.ChannelApplyStatus_CAS_APPLY, CreatedAt: time.Now()}
		Expect(dao.ChannelJoinApplyDB.Create(obj4)).To(BeNil())

		cache.DelPlayerDailyApply(channelId, playerIdRefuse)
		cache.DelPlayerDailyApply(channelId, playerIdPass)
	})

	Describe("ApplyJoin", func() {
		Context("ApplyJoin", func() {
			It("ApplyJoin-申请理由为空", func() {
				req := &pb.ApplyJoinReq{ChannelId: channelId}
				ctxData[define.UID] = playerIdJoin
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				_, err := CObj.ApplyJoin(ctx, req)
				Expect(err).To(Equal(errors.ErrApplyReasonEmpty))
			})
			It("ApplyJoin-已经在频道中", func() {
				req := &pb.ApplyJoinReq{ChannelId: channelId, Reason: "test"}
				ctxData[define.UID] = playerIdJoin
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				_, err := CObj.ApplyJoin(ctx, req)
				Expect(err).To(Equal(errors.ErrPlayerInChannel))
			})
			It("ApplyJoin-存在未处理的申请", func() {
				req := &pb.ApplyJoinReq{ChannelId: channelId, Reason: "test"}
				ctxData[define.UID] = strconv.Itoa(int(playerIdHasApply))
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				_, err := CObj.ApplyJoin(ctx, req)
				Expect(err).To(Equal(errors.ErrApplyNotPass))
			})
			Context("ApplyJoin-1天只能申请一次", func() {
				It("ApplyJoin-1天只能申请一次", func() {
					Expect(dao.ChannelJoinApplyDB.Del(playerIdHasApply, channelId)).To(BeNil())
					_ = cache.SetPlayerDailyApply(channelId, playerIdHasApply)
					req := &pb.ApplyJoinReq{ChannelId: channelId, Reason: "test"}
					ctxData[define.UID] = strconv.Itoa(int(playerIdHasApply))
					ctx := servant.NewOutgoingContext(context.Background(), ctxData)
					_, err := CObj.ApplyJoin(ctx, req)
					Expect(err).To(Equal(errors.ErrRepeatApply))
				})
			})
			Context("ApplyJoin-成功用例", func() {
				It("ApplyJoin-成功用例", func() {
					Expect(dao.ChannelJoinApplyDB.Del(playerIdHasApply, channelId)).To(BeNil())
					cache.DelPlayerDailyApply(channelId, playerIdHasApply)
					req := &pb.ApplyJoinReq{ChannelId: channelId, Reason: "test"}
					ctxData[define.UID] = strconv.Itoa(int(playerIdHasApply))
					ctx := servant.NewOutgoingContext(context.Background(), ctxData)
					_, err := CObj.ApplyJoin(ctx, req)
					Expect(err).To(BeNil())
				})
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
			It("处理申请-通过-权限检查", func() {
				ctxData[define.UID] = strconv.Itoa(int(noPermPlayerId))
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				req := &pb.HandleChannelApplyReq{Type: pb.ChannelHandleType_CHT_PASS, ApplyId: applyIdPass}
				_, err := CObj.HandleChannelApply(ctx, req)
				Expect(err).To(Equal(errors.ErrNoPermission))
			})
			It("处理申请-通过", func() {
				ctxData[define.UID] = manPlayerId
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				req := &pb.HandleChannelApplyReq{Type: pb.ChannelHandleType_CHT_PASS, ApplyId: applyIdPass}
				_, err := CObj.HandleChannelApply(ctx, req)
				Expect(err).To(BeNil())
			})
			It("处理申请-拒绝", func() {
				ctxData[define.UID] = manPlayerId
				ctx := servant.NewOutgoingContext(context.Background(), ctxData)
				req := &pb.HandleChannelApplyReq{Type: pb.ChannelHandleType_CHT_REFUSE, ApplyId: applyIdRefuse}
				_, err := CObj.HandleChannelApply(ctx, req)
				Expect(err).To(BeNil())
			})
		})

	})

	AfterSuite(func() {
		Expect(dao.ChannelJoinApplyDB.DelById(applyIdRefuse)).To(BeNil())
		Expect(dao.ChannelJoinApplyDB.DelById(applyIdPass)).To(BeNil())
		Expect(dao.ChannelJoinApplyDB.DelById(applyIdHas)).To(BeNil())

		_, err := handler.LeaveChannel(context.TODO(), playerIdPass, &pb.LeaveChannelReq{ChannelId: channelId})
		Expect(err).To(BeNil())
	})

})
