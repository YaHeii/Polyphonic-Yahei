package articlerpc

import "github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/articlerpc"

type (
	AddArticleReq              = articlerpc.AddArticleReq
	AddArticleResp             = articlerpc.AddArticleResp
	AddArticleVisitsReq        = articlerpc.AddArticleVisitsReq
	AddArticleVisitsResp       = articlerpc.AddArticleVisitsResp
	AddCategoryReq             = articlerpc.AddCategoryReq
	AddCategoryResp            = articlerpc.AddCategoryResp
	AddTagReq                  = articlerpc.AddTagReq
	AddTagResp                 = articlerpc.AddTagResp
	AnalysisArticleReq         = articlerpc.AnalysisArticleReq
	AnalysisArticleResp        = articlerpc.AnalysisArticleResp
	ArticleCategory            = articlerpc.ArticleCategory
	ArticleDetails             = articlerpc.ArticleDetails
	ArticlePreview             = articlerpc.ArticlePreview
	ArticleTag                 = articlerpc.ArticleTag
	Category                   = articlerpc.Category
	CategoryDetails            = articlerpc.CategoryDetails
	DeletesArticleReq          = articlerpc.DeletesArticleReq
	DeletesArticleResp         = articlerpc.DeletesArticleResp
	DeletesCategoryReq         = articlerpc.DeletesCategoryReq
	DeletesCategoryResp        = articlerpc.DeletesCategoryResp
	DeletesTagReq              = articlerpc.DeletesTagReq
	DeletesTagResp             = articlerpc.DeletesTagResp
	FindArticleListReq         = articlerpc.FindArticleListReq
	FindArticleListResp        = articlerpc.FindArticleListResp
	FindArticlePreviewListResp = articlerpc.FindArticlePreviewListResp
	FindCategoryListReq        = articlerpc.FindCategoryListReq
	FindCategoryListResp       = articlerpc.FindCategoryListResp
	FindLikeArticleResp        = articlerpc.FindLikeArticleResp
	FindTagListReq             = articlerpc.FindTagListReq
	FindTagListResp            = articlerpc.FindTagListResp
	FindUserLikeArticleReq     = articlerpc.FindUserLikeArticleReq
	GetArticleRelationReq      = articlerpc.GetArticleRelationReq
	GetArticleRelationResp     = articlerpc.GetArticleRelationResp
	GetArticleReq              = articlerpc.GetArticleReq
	GetArticleResp             = articlerpc.GetArticleResp
	GetCategoryReq             = articlerpc.GetCategoryReq
	GetCategoryResp            = articlerpc.GetCategoryResp
	GetTagReq                  = articlerpc.GetTagReq
	GetTagResp                 = articlerpc.GetTagResp
	LikeArticleReq             = articlerpc.LikeArticleReq
	LikeArticleResp            = articlerpc.LikeArticleResp
	PageReq                    = articlerpc.PageReq
	PageResp                   = articlerpc.PageResp
	Tag                        = articlerpc.Tag
	TagDetails                 = articlerpc.TagDetails
	UpdateArticleDeleteReq     = articlerpc.UpdateArticleDeleteReq
	UpdateArticleDeleteResp    = articlerpc.UpdateArticleDeleteResp
	UpdateArticleReq           = articlerpc.UpdateArticleReq
	UpdateArticleResp          = articlerpc.UpdateArticleResp
	UpdateArticleTopReq        = articlerpc.UpdateArticleTopReq
	UpdateArticleTopResp       = articlerpc.UpdateArticleTopResp
	UpdateCategoryReq          = articlerpc.UpdateCategoryReq
	UpdateCategoryResp         = articlerpc.UpdateCategoryResp
	UpdateTagReq               = articlerpc.UpdateTagReq
	UpdateTagResp              = articlerpc.UpdateTagResp
)
