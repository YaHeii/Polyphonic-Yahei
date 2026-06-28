package website

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/common/constant"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/articlerpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/configrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/newsrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/websiterpc"
	"google.golang.org/grpc"
)

type stubWebsiteConfigRPC struct {
	configrpc.ConfigRpc
	findReq  *configrpc.FindConfigReq
	findResp *configrpc.FindConfigResp
	saveReq  *configrpc.SaveConfigReq
}

func (s *stubWebsiteConfigRPC) FindConfig(_ context.Context, in *configrpc.FindConfigReq, _ ...grpc.CallOption) (*configrpc.FindConfigResp, error) {
	s.findReq = in
	return s.findResp, nil
}

func (s *stubWebsiteConfigRPC) SaveConfig(_ context.Context, in *configrpc.SaveConfigReq, _ ...grpc.CallOption) (*configrpc.SaveConfigResp, error) {
	s.saveReq = in
	return &configrpc.SaveConfigResp{}, nil
}

type stubWebsiteRPC struct {
	websiterpc.WebsiteRpc
	analysisReq  *websiterpc.AnalysisVisitReq
	analysisResp *websiterpc.AnalysisVisitResp
	trendReq     *websiterpc.FindVisitTrendReq
	trendResp    *websiterpc.FindVisitTrendResp
}

func (s *stubWebsiteRPC) AnalysisVisit(_ context.Context, in *websiterpc.AnalysisVisitReq, _ ...grpc.CallOption) (*websiterpc.AnalysisVisitResp, error) {
	s.analysisReq = in
	return s.analysisResp, nil
}

func (s *stubWebsiteRPC) FindVisitTrend(_ context.Context, in *websiterpc.FindVisitTrendReq, _ ...grpc.CallOption) (*websiterpc.FindVisitTrendResp, error) {
	s.trendReq = in
	return s.trendResp, nil
}

type stubWebsiteAccountRPC struct {
	accountrpc.AccountRpc
	areaReq  *accountrpc.AnalysisUserAreasReq
	areaResp *accountrpc.AnalysisUserAreasResp
	userReq  *accountrpc.AnalysisUserReq
	userResp *accountrpc.AnalysisUserResp
}

func (s *stubWebsiteAccountRPC) AnalysisUserAreas(_ context.Context, in *accountrpc.AnalysisUserAreasReq, _ ...grpc.CallOption) (*accountrpc.AnalysisUserAreasResp, error) {
	s.areaReq = in
	return s.areaResp, nil
}

func (s *stubWebsiteAccountRPC) AnalysisUser(_ context.Context, in *accountrpc.AnalysisUserReq, _ ...grpc.CallOption) (*accountrpc.AnalysisUserResp, error) {
	s.userReq = in
	return s.userResp, nil
}

type stubWebsiteArticleRPC struct {
	articlerpc.ArticleRpc
	analysisReq  *articlerpc.AnalysisArticleReq
	analysisResp *articlerpc.AnalysisArticleResp
	findReq      *articlerpc.FindArticleListReq
	findResp     *articlerpc.FindArticleListResp
}

func (s *stubWebsiteArticleRPC) AnalysisArticle(_ context.Context, in *articlerpc.AnalysisArticleReq, _ ...grpc.CallOption) (*articlerpc.AnalysisArticleResp, error) {
	s.analysisReq = in
	return s.analysisResp, nil
}

func (s *stubWebsiteArticleRPC) FindArticleList(_ context.Context, in *articlerpc.FindArticleListReq, _ ...grpc.CallOption) (*articlerpc.FindArticleListResp, error) {
	s.findReq = in
	return s.findResp, nil
}

type stubWebsiteNewsRPC struct {
	newsrpc.NewsRpc
	analysisReq  *newsrpc.AnalysisMessageReq
	analysisResp *newsrpc.AnalysisMessageResp
}

func (s *stubWebsiteNewsRPC) AnalysisMessage(_ context.Context, in *newsrpc.AnalysisMessageReq, _ ...grpc.CallOption) (*newsrpc.AnalysisMessageResp, error) {
	s.analysisReq = in
	return s.analysisResp, nil
}

func TestGetAndUpdateWebsiteConfig(t *testing.T) {
	rpc := &stubWebsiteConfigRPC{
		findResp: &configrpc.FindConfigResp{
			ConfigValue: `{"admin_url":"https://admin","website_info":{"website_name":"Polyphonic"}}`,
		},
	}
	ctx := context.Background()

	getLogic := NewGetWebsiteConfigLogic(ctx, &svc.ServiceContext{ConfigRpc: rpc})
	getResp, err := getLogic.GetWebsiteConfig(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetWebsiteConfig returned error: %v", err)
	}
	if rpc.findReq == nil || rpc.findReq.ConfigKey != constant.ConfigKeyWebsite || getResp.AdminUrl != "https://admin" || getResp.WebsiteInfo.WebsiteName != "Polyphonic" {
		t.Fatalf("unexpected get website config flow: req=%#v resp=%#v", rpc.findReq, getResp)
	}

	updateLogic := NewUpdateWebsiteConfigLogic(ctx, &svc.ServiceContext{ConfigRpc: rpc})
	if _, err := updateLogic.UpdateWebsiteConfig(&types.WebsiteConfigVO{
		AdminUrl: "https://new-admin",
		WebsiteInfo: &types.WebsiteInfo{
			WebsiteName: "New Polyphonic",
		},
	}); err != nil {
		t.Fatalf("UpdateWebsiteConfig returned error: %v", err)
	}
	if rpc.saveReq == nil || rpc.saveReq.ConfigKey != constant.ConfigKeyWebsite || rpc.saveReq.ConfigValue == "" {
		t.Fatalf("unexpected save website config request: %#v", rpc.saveReq)
	}
}

func TestGetAndUpdateAboutMe(t *testing.T) {
	rpc := &stubWebsiteConfigRPC{
		findResp: &configrpc.FindConfigResp{
			ConfigValue: `{"content":"about me"}`,
		},
	}
	ctx := context.Background()

	getLogic := NewGetAboutMeLogic(ctx, &svc.ServiceContext{ConfigRpc: rpc})
	getResp, err := getLogic.GetAboutMe(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetAboutMe returned error: %v", err)
	}
	if rpc.findReq == nil || rpc.findReq.ConfigKey != constant.ConfigKeyAboutMe || getResp.Content != "about me" {
		t.Fatalf("unexpected get about me flow: req=%#v resp=%#v", rpc.findReq, getResp)
	}

	updateLogic := NewUpdateAboutMeLogic(ctx, &svc.ServiceContext{ConfigRpc: rpc})
	if _, err := updateLogic.UpdateAboutMe(&types.AboutMeVO{Content: "updated"}); err != nil {
		t.Fatalf("UpdateAboutMe returned error: %v", err)
	}
	if rpc.saveReq == nil || rpc.saveReq.ConfigKey != constant.ConfigKeyAboutMe || rpc.saveReq.ConfigValue == "" {
		t.Fatalf("unexpected save about me request: %#v", rpc.saveReq)
	}
}

func TestGetVisitStatsAndTrend(t *testing.T) {
	rpc := &stubWebsiteRPC{
		analysisResp: &websiterpc.AnalysisVisitResp{
			TodayUvCount: 1,
			TotalUvCount: 2,
			UvGrowthRate: 0.1,
			TodayPvCount: 3,
			TotalPvCount: 4,
			PvGrowthRate: 0.2,
		},
		trendResp: &websiterpc.FindVisitTrendResp{
			PvTrend: []*websiterpc.VisitDailyStatistics{
				{Date: "2026-06-01", Count: 10},
			},
			UvTrend: []*websiterpc.VisitDailyStatistics{
				{Date: "2026-06-01", Count: 5},
				{Date: "2026-06-02", Count: 6},
			},
		},
	}
	ctx := context.Background()

	statsLogic := NewGetVisitStatsLogic(ctx, &svc.ServiceContext{WebsiteRpc: rpc})
	statsResp, err := statsLogic.GetVisitStats(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetVisitStats returned error: %v", err)
	}
	if rpc.analysisReq == nil || statsResp.TotalPvCount != 4 || statsResp.TotalUvCount != 2 {
		t.Fatalf("unexpected visit stats flow: req=%#v resp=%#v", rpc.analysisReq, statsResp)
	}

	trendLogic := NewGetVisitTrendLogic(ctx, &svc.ServiceContext{WebsiteRpc: rpc})
	trendResp, err := trendLogic.GetVisitTrend(&types.GetVisitTrendReq{
		StartDate: "2026-06-01",
		EndDate:   "2026-06-02",
	})
	if err != nil {
		t.Fatalf("GetVisitTrend returned error: %v", err)
	}
	if rpc.trendReq == nil || rpc.trendReq.StartDate != "2026-06-01" || len(trendResp.VisitTrend) != 2 {
		t.Fatalf("unexpected visit trend flow: req=%#v resp=%#v", rpc.trendReq, trendResp)
	}
	if trendResp.VisitTrend[0].Date != "2026-06-01" || trendResp.VisitTrend[0].PvCount != 10 || trendResp.VisitTrend[0].UvCount != 5 {
		t.Fatalf("unexpected first trend item: %#v", trendResp.VisitTrend[0])
	}
}

func TestGetUserAreaStats(t *testing.T) {
	accountRPC := &stubWebsiteAccountRPC{
		areaResp: &accountrpc.AnalysisUserAreasResp{
			List: []*accountrpc.UserArea{
				{Area: "Shanghai", Count: 3},
			},
		},
	}
	logic := NewGetUserAreaStatsLogic(context.Background(), &svc.ServiceContext{AccountRpc: accountRPC})

	resp, err := logic.GetUserAreaStats(&types.GetUserAreaStatsReq{UserType: 1})
	if err != nil {
		t.Fatalf("GetUserAreaStats returned error: %v", err)
	}
	if accountRPC.areaReq == nil || accountRPC.areaReq.UserType != 1 || len(resp.UserAreas) != 1 || resp.UserAreas[0].Name != "Shanghai" {
		t.Fatalf("unexpected user area flow: req=%#v resp=%#v", accountRPC.areaReq, resp)
	}
}

func TestGetAdminHomeInfo(t *testing.T) {
	accountRPC := &stubWebsiteAccountRPC{
		userResp: &accountrpc.AnalysisUserResp{UserCount: 10},
	}
	articleRPC := &stubWebsiteArticleRPC{
		analysisResp: &articlerpc.AnalysisArticleResp{
			ArticleCount: 20,
			CategoryList: []*articlerpc.CategoryDetails{
				{Id: 1, CategoryName: "Go", ArticleCount: 2},
			},
			TagList: []*articlerpc.TagDetails{
				{TagName: "grpc", ArticleCount: 3},
			},
			ArticleRankList: []*articlerpc.ArticlePreview{
				{Id: 3, ArticleTitle: "Hot", ViewCount: 100},
			},
		},
		findResp: &articlerpc.FindArticleListResp{
			List: []*articlerpc.ArticleDetails{
				{Id: 1, CreatedAt: 1717200000000},
				{Id: 2, CreatedAt: 1717200000000},
			},
		},
	}
	newsRPC := &stubWebsiteNewsRPC{
		analysisResp: &newsrpc.AnalysisMessageResp{MessageCount: 30},
	}
	logic := NewGetAdminHomeInfoLogic(context.Background(), &svc.ServiceContext{
		AccountRpc: accountRPC,
		ArticleRpc: articleRPC,
		NewsRpc:    newsRPC,
	})

	resp, err := logic.GetAdminHomeInfo(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetAdminHomeInfo returned error: %v", err)
	}
	if accountRPC.userReq == nil || articleRPC.analysisReq == nil || articleRPC.findReq == nil || newsRPC.analysisReq == nil {
		t.Fatalf("unexpected admin home requests: account=%#v articleAnalysis=%#v articleFind=%#v news=%#v", accountRPC.userReq, articleRPC.analysisReq, articleRPC.findReq, newsRPC.analysisReq)
	}
	if resp.UserCount != 10 || resp.ArticleCount != 20 || resp.MessageCount != 30 || len(resp.CategoryList) != 1 || len(resp.TagList) != 1 || len(resp.ArticleViewRanks) != 1 || len(resp.ArticleStatistics) != 1 {
		t.Fatalf("unexpected admin home response: %#v", resp)
	}
}

func TestGetSystemState(t *testing.T) {
	resp, err := NewGetSystemStateLogic(context.Background(), &svc.ServiceContext{}).GetSystemState(&types.EmptyReq{})
	if err != nil {
		t.Fatalf("GetSystemState returned error: %v", err)
	}
	if resp == nil {
		t.Fatalf("unexpected system state response: %#v", resp)
	}
	if resp.Os.GoVersion == "" || resp.Os.Goos == "" {
		t.Fatalf("unexpected os response: %#v", resp.Os)
	}
	if resp.Cpu.Cpus == nil || resp.Cpu.Cores < 0 {
		t.Fatalf("unexpected cpu response: %#v", resp.Cpu)
	}
	if resp.Ram.TotalMb < 0 || resp.Disk.TotalMb < 0 {
		t.Fatalf("unexpected memory or disk response: ram=%#v disk=%#v", resp.Ram, resp.Disk)
	}
}
