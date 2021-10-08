package p_rpc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	deviceType = "9"
)

// 对方tars上的配置信息，在"公共后台——密钥管理——密钥"添加/查询对应项目，设备类型，名字的密钥
type tarsReq2HTTPOpt struct {
	DeviceType  string //设备类型，1.安卓, 2.ios ... 9.后端
	Name        string //业务方
	AppKey      string //对方tars上配置的密钥
	ProjectId   string //业务方projectId
	CommonProxy string //对方tars的公共网关
}

type tarsServiceReq struct {
	Obj  string        // 服务名称（全名，形如yes.channel.ChannelExtObj）
	Func string        // 服务接口名
	Req  proto.Message // 服务参数
}

func rpcTarsService(tarsReq *tarsServiceReq) (*proto2.RPCOutput, error) {
	tOpt := &tarsReq2HTTPOpt{
		ProjectId:  strconv.Itoa(env.ProjectId),
		Name:       env.AppName,
		DeviceType: deviceType,
	}
	switch tarsReq.Obj {
	case "mizhua.advert.AdvertExtObj", "mizacommon.dynconf.DynconfExtObj", "mizhuaTmp.advert.AdvertExtObj", "mizacommonTmp.dynconf.DynconfExtObj":
		tOpt.CommonProxy = config.AppCfg.Http2tars.ProxyCommon
		tOpt.AppKey = config.AppCfg.Http2tars.AppKey
	}
	return tarsReq2HTTP(tarsReq.Req, tarsReq.Func, tarsReq.Obj, tOpt)
}

// 审核接口，将tars请求转为http
func tarsReq2HTTP(m proto.Message, reqFuncName string, reqServant string, opt *tarsReq2HTTPOpt) (*proto2.RPCOutput, error) {
	pbresp := new(proto2.RPCOutput)
	pbbin, err := proto.Marshal(m)
	if err != nil {
		appzaplog.Error("rpc tarsReq2HTTP proto Marshal err", zap.Any("err", err))
		return pbresp, err
	}
	unix := strconv.Itoa(int(time.Now().Unix()))
	// 在"公共后台——密钥管理——密钥"添加/查询对应项目，设备类型，名字的密钥
	key := sha256.Sum256([]byte(reqServant + reqFuncName + opt.ProjectId + opt.DeviceType + opt.Name + string(pbbin) + unix + opt.AppKey))
	sign := fmt.Sprintf("%x", key)
	xyreqid := fmt.Sprintf("%d", time.Now().UnixNano())
	appzaplog.Info("HandleProtobufRPC xyreqid", zap.String("xyreqid", xyreqid))
	rpcOpt := map[string]string{
		"project_id":  opt.ProjectId,
		"device_type": opt.DeviceType,
		"name":        opt.Name,
		"unix":        unix,
		"key":         sign,
		"xyreqid":     xyreqid,
	}
	rpcInput := proto2.RPCInput{
		Obj:  reqServant,
		Func: reqFuncName,
		Opt:  rpcOpt,
		Req:  pbbin,
	}
	pbbin2, err := proto.Marshal(&rpcInput)
	if err != nil {
		appzaplog.Error("rpc tarsReq2HTTP proto Marshal err", zap.Any("err", err))
		return pbresp, err
	}
	client := http.Client{}
	bytesreader := bytes.NewReader(pbbin2)
	url := opt.CommonProxy
	httpreq, err := http.NewRequest("POST", url, bytesreader)
	if err != nil {
		appzaplog.Error("rpc tarsReq2HTTP request url err", zap.Any("err", err))
		return pbresp, err
	}
	resp, err := client.Do(httpreq)
	if err != nil {
		appzaplog.Error("rpc tarsReq2HTTP http client do err", zap.Any("err", err))
		return pbresp, err
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	bodybin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		appzaplog.Error("rpc tarsReq2HTTP http client read err", zap.Any("err", err))
		return pbresp, err
	}
	if err := proto.Unmarshal(bodybin, pbresp); err != nil {
		appzaplog.Error("rpc tarsReq2HTTP unmarshal bodybin err", zap.Any("err", err))
		return pbresp, err
	}
	return pbresp, nil
}
