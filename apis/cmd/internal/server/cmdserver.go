// Code generated by goctl. DO NOT EDIT!
// Source: cmd.proto

package server

import (
	"context"

	"ylink/apis/cmd/internal/logic"
	"ylink/apis/cmd/internal/svc"
	"ylink/apis/cmd/pb"
)

type CmdServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCmdServer
}

func NewCmdServer(svcCtx *svc.ServiceContext) *CmdServer {
	return &CmdServer{
		svcCtx: svcCtx,
	}
}

func (s *CmdServer) PlayerFetchCsInfo(ctx context.Context, in *pb.PlayerFetchCsInfoReq) (*pb.CmdResp, error) {
	l := logic.NewPlayerFetchCsInfoLogic(ctx, s.svcCtx)
	return l.PlayerFetchCsInfo(in)
}

func (s *CmdServer) PlayerFetchHistoryMsg(ctx context.Context, in *pb.PlayerFetchHistoryMsgReq) (*pb.CmdResp, error) {
	l := logic.NewPlayerFetchHistoryMsgLogic(ctx, s.svcCtx)
	return l.PlayerFetchHistoryMsg(in)
}

func (s *CmdServer) PlayerFetchMsg(ctx context.Context, in *pb.PlayerFetchMsgReq) (*pb.CmdResp, error) {
	l := logic.NewPlayerFetchMsgLogic(ctx, s.svcCtx)
	return l.PlayerFetchMsg(in)
}

func (s *CmdServer) PlayerSendMsg(ctx context.Context, in *pb.PlayerSendMsgReq) (*pb.CmdResp, error) {
	l := logic.NewPlayerSendMsgLogic(ctx, s.svcCtx)
	return l.PlayerSendMsg(in)
}

func (s *CmdServer) PlayerDisconnect(ctx context.Context, in *pb.PlayerDisconnectReq) (*pb.CmdResp, error) {
	l := logic.NewPlayerDisconnectLogic(ctx, s.svcCtx)
	return l.PlayerDisconnect(in)
}

func (s *CmdServer) CsFetchPlayerQueue(ctx context.Context, in *pb.CsFetchPlayerQueueReq) (*pb.CmdResp, error) {
	l := logic.NewCsFetchPlayerQueueLogic(ctx, s.svcCtx)
	return l.CsFetchPlayerQueue(in)
}

func (s *CmdServer) CsConnectPlayer(ctx context.Context, in *pb.CsConnectPlayerReq) (*pb.CmdResp, error) {
	l := logic.NewCsConnectPlayerLogic(ctx, s.svcCtx)
	return l.CsConnectPlayer(in)
}

func (s *CmdServer) CsFetchHistoryChat(ctx context.Context, in *pb.CsFetchHistoryChatReq) (*pb.CmdResp, error) {
	l := logic.NewCsFetchHistoryChatLogic(ctx, s.svcCtx)
	return l.CsFetchHistoryChat(in)
}

func (s *CmdServer) CsFetchHistoryMsg(ctx context.Context, in *pb.CsFetchHistoryMsgReq) (*pb.CmdResp, error) {
	l := logic.NewCsFetchHistoryMsgLogic(ctx, s.svcCtx)
	return l.CsFetchHistoryMsg(in)
}

func (s *CmdServer) CsFetchMsg(ctx context.Context, in *pb.CsFetchMsgReq) (*pb.CmdResp, error) {
	l := logic.NewCsFetchMsgLogic(ctx, s.svcCtx)
	return l.CsFetchMsg(in)
}

func (s *CmdServer) CsSendMsg(ctx context.Context, in *pb.CsSendMsgReq) (*pb.CmdResp, error) {
	l := logic.NewCsSendMsgLogic(ctx, s.svcCtx)
	return l.CsSendMsg(in)
}
