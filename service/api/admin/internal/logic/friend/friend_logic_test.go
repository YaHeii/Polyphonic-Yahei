package friend

import (
	"context"
	"errors"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/socialrpc"
	"google.golang.org/grpc"
)

type stubSocialRPC struct {
	addFriendReq     *socialrpc.AddFriendReq
	addFriendResp    *socialrpc.AddFriendResp
	addFriendErr     error
	updateFriendReq  *socialrpc.UpdateFriendReq
	updateFriendResp *socialrpc.UpdateFriendResp
	updateFriendErr  error
	deleteFriendReq  *socialrpc.DeletesFriendReq
	deleteFriendResp *socialrpc.DeletesFriendResp
	deleteFriendErr  error
	findFriendReq    *socialrpc.FindFriendListReq
	findFriendResp   *socialrpc.FindFriendListResp
	findFriendErr    error
}

func (s *stubSocialRPC) AddFriend(_ context.Context, in *socialrpc.AddFriendReq, _ ...grpc.CallOption) (*socialrpc.AddFriendResp, error) {
	s.addFriendReq = in
	return s.addFriendResp, s.addFriendErr
}

func (s *stubSocialRPC) UpdateFriend(_ context.Context, in *socialrpc.UpdateFriendReq, _ ...grpc.CallOption) (*socialrpc.UpdateFriendResp, error) {
	s.updateFriendReq = in
	return s.updateFriendResp, s.updateFriendErr
}

func (s *stubSocialRPC) DeletesFriend(_ context.Context, in *socialrpc.DeletesFriendReq, _ ...grpc.CallOption) (*socialrpc.DeletesFriendResp, error) {
	s.deleteFriendReq = in
	return s.deleteFriendResp, s.deleteFriendErr
}

func (s *stubSocialRPC) FindFriendList(_ context.Context, in *socialrpc.FindFriendListReq, _ ...grpc.CallOption) (*socialrpc.FindFriendListResp, error) {
	s.findFriendReq = in
	return s.findFriendResp, s.findFriendErr
}

func (s *stubSocialRPC) AddTalk(context.Context, *socialrpc.AddTalkReq, ...grpc.CallOption) (*socialrpc.AddTalkResp, error) {
	panic("unexpected call")
}

func (s *stubSocialRPC) UpdateTalk(context.Context, *socialrpc.UpdateTalkReq, ...grpc.CallOption) (*socialrpc.UpdateTalkResp, error) {
	panic("unexpected call")
}

func (s *stubSocialRPC) DeletesTalk(context.Context, *socialrpc.DeletesTalkReq, ...grpc.CallOption) (*socialrpc.DeletesTalkResp, error) {
	panic("unexpected call")
}

func (s *stubSocialRPC) GetTalk(context.Context, *socialrpc.GetTalkReq, ...grpc.CallOption) (*socialrpc.GetTalkResp, error) {
	panic("unexpected call")
}

func (s *stubSocialRPC) FindTalkList(context.Context, *socialrpc.FindTalkListReq, ...grpc.CallOption) (*socialrpc.FindTalkListResp, error) {
	panic("unexpected call")
}

func (s *stubSocialRPC) LikeTalk(context.Context, *socialrpc.LikeTalkReq, ...grpc.CallOption) (*socialrpc.LikeTalkResp, error) {
	panic("unexpected call")
}

func (s *stubSocialRPC) FindUserLikeTalk(context.Context, *socialrpc.FindUserLikeTalkReq, ...grpc.CallOption) (*socialrpc.FindUserLikeTalkResp, error) {
	panic("unexpected call")
}

func TestAddFriendBuildsRequest(t *testing.T) {
	rpc := &stubSocialRPC{
		addFriendResp: &socialrpc.AddFriendResp{
			Friend: &socialrpc.Friend{Id: 1, LinkName: "link"},
		},
	}
	logic := NewAddFriendLogic(context.Background(), &svc.ServiceContext{SocialRpc: rpc})

	resp, err := logic.AddFriend(&types.NewFriendReq{
		Id:          1,
		LinkName:    "link",
		LinkAvatar:  "avatar",
		LinkAddress: "address",
		LinkIntro:   "intro",
	})
	if err != nil {
		t.Fatalf("AddFriend returned error: %v", err)
	}

	if rpc.addFriendReq == nil || rpc.addFriendReq.LinkName != "link" || rpc.addFriendReq.LinkAvatar != "avatar" {
		t.Fatalf("unexpected add request: %#v", rpc.addFriendReq)
	}

	if resp.Id != 1 || resp.LinkName != "link" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestDeleteFriendPropagatesRPCError(t *testing.T) {
	wantErr := errors.New("delete failed")
	logic := NewDeletesFriendLogic(context.Background(), &svc.ServiceContext{
		SocialRpc: &stubSocialRPC{deleteFriendErr: wantErr},
	})

	_, err := logic.DeletesFriend(&types.IdsReq{Ids: []int64{1}})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestFindFriendListBuildsQueryAndMapsPage(t *testing.T) {
	rpc := &stubSocialRPC{
		findFriendResp: &socialrpc.FindFriendListResp{
			Pagination: &socialrpc.PageResp{Page: 2, PageSize: 10, Total: 3},
			List: []*socialrpc.Friend{
				{Id: 7, LinkName: "friend"},
			},
		},
	}
	logic := NewFindFriendListLogic(context.Background(), &svc.ServiceContext{SocialRpc: rpc})

	resp, err := logic.FindFriendList(&types.QueryFriendReq{
		PageQuery: types.PageQuery{
			Page:     2,
			PageSize: 10,
			Sorts:    []string{"created_at desc"},
		},
		LinkName: "fri",
	})
	if err != nil {
		t.Fatalf("FindFriendList returned error: %v", err)
	}

	if rpc.findFriendReq == nil || rpc.findFriendReq.Paginate == nil {
		t.Fatal("expected find request to be sent")
	}

	if rpc.findFriendReq.Paginate.Page != 2 || rpc.findFriendReq.Paginate.PageSize != 10 || rpc.findFriendReq.LinkName != "fri" {
		t.Fatalf("unexpected find request: %#v", rpc.findFriendReq)
	}

	list, ok := resp.List.([]*types.FriendBackVO)
	if !ok || len(list) != 1 || list[0].Id != 7 {
		t.Fatalf("unexpected list payload: %#v", resp.List)
	}
}
