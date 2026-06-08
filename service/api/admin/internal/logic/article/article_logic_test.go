package article

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/infra/biz/bizheader"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"google.golang.org/grpc"
)

type stubArticleRPC struct {
	deleteReq        *articlerpc.DeletesArticleReq
	deleteResp       *articlerpc.DeletesArticleResp
	deleteErr        error
	getReq           *articlerpc.GetArticleReq
	getResp          *articlerpc.GetArticleResp
	getErr           error
	findReq          *articlerpc.FindArticleListReq
	findResp         *articlerpc.FindArticleListResp
	findErr          error
	updateReq        *articlerpc.UpdateArticleReq
	updateResp       *articlerpc.UpdateArticleResp
	updateErr        error
	updateDeleteReq  *articlerpc.UpdateArticleDeleteReq
	updateDeleteResp *articlerpc.UpdateArticleDeleteResp
	updateDeleteErr  error
	updateTopReq     *articlerpc.UpdateArticleTopReq
	updateTopResp    *articlerpc.UpdateArticleTopResp
	updateTopErr     error
}

func (s *stubArticleRPC) AnalysisArticle(context.Context, *articlerpc.AnalysisArticleReq, ...grpc.CallOption) (*articlerpc.AnalysisArticleResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) AddArticleVisits(context.Context, *articlerpc.AddArticleVisitsReq, ...grpc.CallOption) (*articlerpc.AddArticleVisitsResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) AddArticle(context.Context, *articlerpc.AddArticleReq, ...grpc.CallOption) (*articlerpc.AddArticleResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) UpdateArticle(_ context.Context, in *articlerpc.UpdateArticleReq, _ ...grpc.CallOption) (*articlerpc.UpdateArticleResp, error) {
	s.updateReq = in
	return s.updateResp, s.updateErr
}

func (s *stubArticleRPC) UpdateArticleDelete(_ context.Context, in *articlerpc.UpdateArticleDeleteReq, _ ...grpc.CallOption) (*articlerpc.UpdateArticleDeleteResp, error) {
	s.updateDeleteReq = in
	return s.updateDeleteResp, s.updateDeleteErr
}

func (s *stubArticleRPC) UpdateArticleTop(_ context.Context, in *articlerpc.UpdateArticleTopReq, _ ...grpc.CallOption) (*articlerpc.UpdateArticleTopResp, error) {
	s.updateTopReq = in
	return s.updateTopResp, s.updateTopErr
}

func (s *stubArticleRPC) DeletesArticle(_ context.Context, in *articlerpc.DeletesArticleReq, _ ...grpc.CallOption) (*articlerpc.DeletesArticleResp, error) {
	s.deleteReq = in
	return s.deleteResp, s.deleteErr
}

func (s *stubArticleRPC) GetArticle(_ context.Context, in *articlerpc.GetArticleReq, _ ...grpc.CallOption) (*articlerpc.GetArticleResp, error) {
	s.getReq = in
	return s.getResp, s.getErr
}

func (s *stubArticleRPC) GetArticleRelation(context.Context, *articlerpc.GetArticleRelationReq, ...grpc.CallOption) (*articlerpc.GetArticleRelationResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) FindArticleList(_ context.Context, in *articlerpc.FindArticleListReq, _ ...grpc.CallOption) (*articlerpc.FindArticleListResp, error) {
	s.findReq = in
	return s.findResp, s.findErr
}

func (s *stubArticleRPC) FindArticlePreviewList(context.Context, *articlerpc.FindArticleListReq, ...grpc.CallOption) (*articlerpc.FindArticlePreviewListResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) LikeArticle(context.Context, *articlerpc.LikeArticleReq, ...grpc.CallOption) (*articlerpc.LikeArticleResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) FindUserLikeArticle(context.Context, *articlerpc.FindUserLikeArticleReq, ...grpc.CallOption) (*articlerpc.FindLikeArticleResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) AddCategory(context.Context, *articlerpc.AddCategoryReq, ...grpc.CallOption) (*articlerpc.AddCategoryResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) UpdateCategory(context.Context, *articlerpc.UpdateCategoryReq, ...grpc.CallOption) (*articlerpc.UpdateCategoryResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) GetCategory(context.Context, *articlerpc.GetCategoryReq, ...grpc.CallOption) (*articlerpc.GetCategoryResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) DeletesCategory(context.Context, *articlerpc.DeletesCategoryReq, ...grpc.CallOption) (*articlerpc.DeletesCategoryResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) FindCategoryList(context.Context, *articlerpc.FindCategoryListReq, ...grpc.CallOption) (*articlerpc.FindCategoryListResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) AddTag(context.Context, *articlerpc.AddTagReq, ...grpc.CallOption) (*articlerpc.AddTagResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) UpdateTag(context.Context, *articlerpc.UpdateTagReq, ...grpc.CallOption) (*articlerpc.UpdateTagResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) GetTag(context.Context, *articlerpc.GetTagReq, ...grpc.CallOption) (*articlerpc.GetTagResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) DeletesTag(context.Context, *articlerpc.DeletesTagReq, ...grpc.CallOption) (*articlerpc.DeletesTagResp, error) {
	panic("unexpected call")
}

func (s *stubArticleRPC) FindTagList(context.Context, *articlerpc.FindTagListReq, ...grpc.CallOption) (*articlerpc.FindTagListResp, error) {
	panic("unexpected call")
}

func TestDeleteArticleBuildsDeleteRequest(t *testing.T) {
	articleRPC := &stubArticleRPC{
		deleteResp: &articlerpc.DeletesArticleResp{SuccessCount: 1},
	}
	logic := NewDeleteArticleLogic(context.Background(), &svc.ServiceContext{ArticleRpc: articleRPC})

	resp, err := logic.DeleteArticle(&types.IdReq{Id: 42})
	if err != nil {
		t.Fatalf("DeleteArticle returned error: %v", err)
	}

	if articleRPC.deleteReq == nil {
		t.Fatal("expected delete request to be sent")
	}

	if len(articleRPC.deleteReq.Ids) != 1 || articleRPC.deleteReq.Ids[0] != 42 {
		t.Fatalf("unexpected delete ids: %#v", articleRPC.deleteReq.Ids)
	}

	if resp.SuccessCount != 1 {
		t.Fatalf("unexpected success count: %d", resp.SuccessCount)
	}
}

func TestDeleteArticlePropagatesRPCError(t *testing.T) {
	wantErr := errors.New("delete failed")
	logic := NewDeleteArticleLogic(context.Background(), &svc.ServiceContext{
		ArticleRpc: &stubArticleRPC{deleteErr: wantErr},
	})

	_, err := logic.DeleteArticle(&types.IdReq{Id: 1})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestGetArticleConvertsRPCResponse(t *testing.T) {
	articleRPC := &stubArticleRPC{
		getResp: &articlerpc.GetArticleResp{
			Article: &articlerpc.ArticleDetails{
				Id:             9,
				ArticleCover:   "cover",
				ArticleTitle:   "title",
				ArticleContent: "content",
				ArticleType:    2,
				OriginalUrl:    "https://example.com",
				IsTop:          true,
				IsDelete:       false,
				Status:         3,
				CreatedAt:      100,
				UpdatedAt:      200,
				Category:       &articlerpc.ArticleCategory{CategoryName: "Go"},
				TagList: []*articlerpc.ArticleTag{
					{TagName: "tag-a"},
					{TagName: "tag-b"},
				},
				LikeCount: 7,
				ViewCount: 8,
			},
		},
	}
	logic := NewGetArticleLogic(context.Background(), &svc.ServiceContext{ArticleRpc: articleRPC})

	resp, err := logic.GetArticle(&types.IdReq{Id: 9})
	if err != nil {
		t.Fatalf("GetArticle returned error: %v", err)
	}

	if articleRPC.getReq == nil || articleRPC.getReq.Id != 9 {
		t.Fatalf("unexpected get request: %#v", articleRPC.getReq)
	}

	if resp.Id != 9 || resp.CategoryName != "Go" || len(resp.TagNameList) != 2 {
		t.Fatalf("unexpected article response: %#v", resp)
	}
}

func TestFindArticleListBuildsQueryAndMapsPage(t *testing.T) {
	articleRPC := &stubArticleRPC{
		findResp: &articlerpc.FindArticleListResp{
			Pagination: &articlerpc.PageResp{Page: 2, PageSize: 20, Total: 99},
			List: []*articlerpc.ArticleDetails{
				{
					Id:           1,
					ArticleTitle: "first",
					Category:     &articlerpc.ArticleCategory{CategoryName: "Cat"},
				},
			},
		},
	}
	logic := NewFindArticleListLogic(context.Background(), &svc.ServiceContext{ArticleRpc: articleRPC})
	req := &types.QueryArticleReq{
		PageQuery: types.PageQuery{
			Page:     2,
			PageSize: 20,
			Sorts:    []string{"created_at desc"},
		},
		ArticleTitle: "hello",
		ArticleType:  3,
		IsTop:        true,
		IsDelete:     true,
		Status:       4,
		CategoryName: "Cat",
		TagName:      "Tag",
	}

	resp, err := logic.FindArticleList(req)
	if err != nil {
		t.Fatalf("FindArticleList returned error: %v", err)
	}

	if articleRPC.findReq == nil {
		t.Fatal("expected find request to be sent")
	}

	if articleRPC.findReq.Paginate == nil || articleRPC.findReq.Paginate.Page != 2 || articleRPC.findReq.Paginate.PageSize != 20 {
		t.Fatalf("unexpected paginate request: %#v", articleRPC.findReq.Paginate)
	}

	if articleRPC.findReq.ArticleTitle != "hello" || articleRPC.findReq.CategoryName != "Cat" || articleRPC.findReq.TagName != "Tag" {
		t.Fatalf("unexpected query request: %#v", articleRPC.findReq)
	}

	if articleRPC.findReq.IsTop == nil || !*articleRPC.findReq.IsTop || articleRPC.findReq.IsDelete == nil || !*articleRPC.findReq.IsDelete {
		t.Fatalf("expected optional bool filters to be set: %#v", articleRPC.findReq)
	}

	if resp.Page != 2 || resp.PageSize != 20 || resp.Total != 99 {
		t.Fatalf("unexpected page response: %#v", resp)
	}

	list, ok := resp.List.([]*types.ArticleBackVO)
	if !ok || len(list) != 1 || list[0].ArticleTitle != "first" {
		t.Fatalf("unexpected page list: %#v", resp.List)
	}
}

func TestUpdateArticleUsesCurrentUserAndMapsResponse(t *testing.T) {
	articleRPC := &stubArticleRPC{
		updateResp: &articlerpc.UpdateArticleResp{
			Article: &articlerpc.ArticlePreview{Id: 77},
		},
	}
	ctx := context.WithValue(context.Background(), bizheader.HeaderUid, "u-1")
	logic := NewUpdateArticleLogic(ctx, &svc.ServiceContext{ArticleRpc: articleRPC})
	req := &types.NewArticleReq{
		Id:             77,
		ArticleCover:   "cover",
		ArticleTitle:   "title",
		ArticleContent: "content",
		ArticleType:    1,
		OriginalUrl:    "https://example.com",
		IsTop:          true,
		Status:         2,
		CategoryName:   "Cat",
		TagNameList:    []string{"t1", "t2"},
	}

	resp, err := logic.UpdateArticle(req)
	if err != nil {
		t.Fatalf("UpdateArticle returned error: %v", err)
	}

	if articleRPC.updateReq == nil {
		t.Fatal("expected update request to be sent")
	}

	if articleRPC.updateReq.UserId != "u-1" || articleRPC.updateReq.Id != 77 || articleRPC.updateReq.CategoryName != "Cat" {
		t.Fatalf("unexpected update request: %#v", articleRPC.updateReq)
	}

	if resp.Id != 77 {
		t.Fatalf("unexpected update response: %#v", resp)
	}
}

func TestUpdateArticleDeleteBuildsRequest(t *testing.T) {
	articleRPC := &stubArticleRPC{
		updateDeleteResp: &articlerpc.UpdateArticleDeleteResp{},
	}
	logic := NewUpdateArticleDeleteLogic(context.Background(), &svc.ServiceContext{ArticleRpc: articleRPC})

	resp, err := logic.UpdateArticleDelete(&types.UpdateArticleDeleteReq{Id: 5, IsDelete: true})
	if err != nil {
		t.Fatalf("UpdateArticleDelete returned error: %v", err)
	}

	if articleRPC.updateDeleteReq == nil || articleRPC.updateDeleteReq.ArticleId != 5 || !articleRPC.updateDeleteReq.IsDelete {
		t.Fatalf("unexpected update delete request: %#v", articleRPC.updateDeleteReq)
	}

	if resp == nil {
		t.Fatal("expected empty response")
	}
}

func TestUpdateArticleTopBuildsRequest(t *testing.T) {
	articleRPC := &stubArticleRPC{
		updateTopResp: &articlerpc.UpdateArticleTopResp{},
	}
	logic := NewUpdateArticleTopLogic(context.Background(), &svc.ServiceContext{ArticleRpc: articleRPC})

	resp, err := logic.UpdateArticleTop(&types.UpdateArticleTopReq{Id: 6, IsTop: true})
	if err != nil {
		t.Fatalf("UpdateArticleTop returned error: %v", err)
	}

	if articleRPC.updateTopReq == nil || articleRPC.updateTopReq.ArticleId != 6 || !articleRPC.updateTopReq.IsTop {
		t.Fatalf("unexpected update top request: %#v", articleRPC.updateTopReq)
	}

	if resp == nil {
		t.Fatal("expected empty response")
	}
}

func TestExportArticleListWritesMarkdownFiles(t *testing.T) {
	articleRPC := &stubArticleRPC{
		findResp: &articlerpc.FindArticleListResp{
			List: []*articlerpc.ArticleDetails{
				{
					Id:             1,
					ArticleTitle:   "export-me",
					ArticleCover:   "cover",
					ArticleType:    2,
					ArticleContent: "body",
					CreatedAt:      1,
					Category:       &articlerpc.ArticleCategory{CategoryName: "Cat"},
					TagList:        []*articlerpc.ArticleTag{{TagName: "Tag"}},
				},
			},
		},
	}
	logic := NewExportArticleListLogic(context.Background(), &svc.ServiceContext{ArticleRpc: articleRPC})

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd failed: %v", err)
	}

	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Chdir failed: %v", err)
	}
	defer func() { _ = os.Chdir(wd) }()

	_, err = logic.ExportArticleList(&types.IdsReq{Ids: []int64{1}})
	if err != nil {
		t.Fatalf("ExportArticleList returned error: %v", err)
	}

	if articleRPC.findReq == nil || len(articleRPC.findReq.Ids) != 1 || articleRPC.findReq.Ids[0] != 1 {
		t.Fatalf("unexpected export find request: %#v", articleRPC.findReq)
	}

	outputPath := filepath.Join(tempDir, "runtime", "article", "export-me.md")
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("expected exported markdown file: %v", err)
	}

	content := string(data)
	if content == "" || !containsAll(content, []string{"export-me", "cover", "Cat", "body"}) {
		t.Fatalf("unexpected exported markdown content: %s", content)
	}
}

func containsAll(s string, subs []string) bool {
	for _, sub := range subs {
		if !contains(s, sub) {
			return false
		}
	}
	return true
}

func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}
