package p_rpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"strconv"
)

func CommitContentCheck(playerId int64, cType int32, content *common2.ContentData, carrier int32, optType base_pb.AuditOp) (*base_pb.CommitContentCheckRes, error) {
	defer common.CheckGoPanic()
	res := new(base_pb.CommitContentCheckRes)
	bin, err := json.Marshal(content)
	if err != nil {
		appzaplog.Error("CommitAudit_json_error", zap.Error(err), zap.Any("playerId", playerId), zap.Any("cType", cType), zap.Any("content", content))
		return res, err
	}
	checkReq := &base_pb.CommitContentCheckReq{
		AppId:            env.AppId,
		PlayerId:         playerId,
		CType:            cType,
		Carrier:          carrier,
		Content:          bin,
		Op:               optType,
		IsHard:           false,
		AdvertAliPicture: true,
		OnlyMachine:      true,
	}
	tarsReq := &tarsServiceReq{
		Obj:  "mizhua.advert.AdvertExtObj",
		Func: "CommitContentCheck",
		Req:  checkReq,
	}
	pbresp, err := rpcTarsService(tarsReq)
	if pbresp == nil {
		return res, err
	}
	if err := proto.Unmarshal(pbresp.Rsp, res); err != nil {
		appzaplog.Error("rpc CommitAuditByOptType unmarshal response err", zap.Any("err", err))
	}

	appzaplog.Info("rpc CommitAuditByOptType", zap.Int64("playerId", playerId), zap.Int32("cType", cType), zap.Any("content", content), zap.Any("optType", optType), zap.Any("res", res), zap.Any("pbresp", pbresp))
	return res, err
}
