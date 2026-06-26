package permissionrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/query"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"
)

type permissionHelper struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func newPermissionHelper(ctx context.Context, svcCtx *svc.ServiceContext) *permissionHelper {
	return &permissionHelper{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func convertAddApiIn(in *permissionrpc.AddApiReq) *model.TApi {
	return &model.TApi{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
	}
}

func convertUpdateApiIn(in *permissionrpc.UpdateApiReq) *model.TApi {
	return &model.TApi{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
	}
}

func convertApiOut(record *model.TApi) *permissionrpc.Api {
	if record == nil {
		return nil
	}

	return &permissionrpc.Api{
		Id:        record.Id,
		ParentId:  record.ParentId,
		Path:      record.Path,
		Name:      record.Name,
		Method:    record.Method,
		Traceable: record.Traceable,
		Status:    record.Status,
		CreatedAt: record.CreatedAt.UnixMilli(),
		UpdatedAt: record.UpdatedAt.UnixMilli(),
		Children:  []*permissionrpc.Api{},
	}
}

func buildApiTree(records []*model.TApi) []*permissionrpc.Api {
	nodes := make(map[int64]*permissionrpc.Api, len(records))
	order := make([]int64, 0, len(records))
	parentIDs := make(map[int64]int64, len(records))

	for _, record := range records {
		node := convertApiOut(record)
		nodes[record.Id] = node
		order = append(order, record.Id)
		parentIDs[record.Id] = record.ParentId
	}

	roots := make([]*permissionrpc.Api, 0)
	for _, id := range order {
		node := nodes[id]
		parentID := parentIDs[id]
		if parentID != 0 {
			if parent, ok := nodes[parentID]; ok {
				parent.Children = append(parent.Children, node)
				continue
			}
		}
		roots = append(roots, node)
	}

	return roots
}

func convertAddMenuIn(in *permissionrpc.AddMenuReq) *model.TMenu {
	meta := in.Meta
	if meta == nil {
		meta = &permissionrpc.MenuMeta{}
	}

	return &model.TMenu{
		Id:         in.Id,
		ParentId:   in.ParentId,
		Path:       in.Path,
		Name:       in.Name,
		Component:  in.Component,
		Redirect:   in.Redirect,
		Type:       meta.Type,
		Title:      meta.Title,
		Icon:       meta.Icon,
		Rank:       meta.Rank,
		Perm:       meta.Perm,
		Params:     meta.Params,
		KeepAlive:  meta.KeepAlive,
		AlwaysShow: meta.AlwaysShow,
		Visible:    meta.Visible,
		Status:     meta.Status,
		Extra:      "{}",
	}
}

func convertUpdateMenuIn(in *permissionrpc.UpdateMenuReq) *model.TMenu {
	meta := in.Meta
	if meta == nil {
		meta = &permissionrpc.MenuMeta{}
	}

	return &model.TMenu{
		Id:         in.Id,
		ParentId:   in.ParentId,
		Path:       in.Path,
		Name:       in.Name,
		Component:  in.Component,
		Redirect:   in.Redirect,
		Type:       meta.Type,
		Title:      meta.Title,
		Icon:       meta.Icon,
		Rank:       meta.Rank,
		Perm:       meta.Perm,
		Params:     meta.Params,
		KeepAlive:  meta.KeepAlive,
		AlwaysShow: meta.AlwaysShow,
		Visible:    meta.Visible,
		Status:     meta.Status,
		Extra:      "{}",
	}
}

func convertMenuOut(record *model.TMenu) *permissionrpc.Menu {
	if record == nil {
		return nil
	}

	return &permissionrpc.Menu{
		Id:        record.Id,
		ParentId:  record.ParentId,
		Path:      record.Path,
		Name:      record.Name,
		Component: record.Component,
		Redirect:  record.Redirect,
		CreatedAt: record.CreatedAt.UnixMilli(),
		UpdatedAt: record.UpdatedAt.UnixMilli(),
		Children:  []*permissionrpc.Menu{},
		Meta: &permissionrpc.MenuMeta{
			Type:       record.Type,
			Title:      record.Title,
			Icon:       record.Icon,
			Rank:       record.Rank,
			Perm:       record.Perm,
			Params:     record.Params,
			KeepAlive:  record.KeepAlive,
			AlwaysShow: record.AlwaysShow,
			Visible:    record.Visible,
			Status:     record.Status,
		},
	}
}

func buildMenuTree(records []*model.TMenu) []*permissionrpc.Menu {
	nodes := make(map[int64]*permissionrpc.Menu, len(records))
	order := make([]int64, 0, len(records))
	parentIDs := make(map[int64]int64, len(records))

	for _, record := range records {
		node := convertMenuOut(record)
		nodes[record.Id] = node
		order = append(order, record.Id)
		parentIDs[record.Id] = record.ParentId
	}

	roots := make([]*permissionrpc.Menu, 0)
	for _, id := range order {
		node := nodes[id]
		parentID := parentIDs[id]
		if parentID != 0 {
			if parent, ok := nodes[parentID]; ok {
				parent.Children = append(parent.Children, node)
				continue
			}
		}
		roots = append(roots, node)
	}

	return roots
}

func convertRoleOut(record *model.TRole) *permissionrpc.Role {
	if record == nil {
		return nil
	}

	return &permissionrpc.Role{
		Id:          record.Id,
		RoleKey:     record.RoleKey,
		RoleComment: record.RoleComment,
		Status:      record.Status,
		CreatedAt:   record.CreatedAt.UnixMilli(),
		UpdatedAt:   record.UpdatedAt.UnixMilli(),
	}
}

func convertRoleListOut(records []*model.TRole) []*permissionrpc.Role {
	roles := make([]*permissionrpc.Role, 0, len(records))
	for _, record := range records {
		roles = append(roles, convertRoleOut(record))
	}

	return roles
}

func buildApiQuery(in *permissionrpc.FindApiListReq) (int, int, string, string, []any) {
	opts := []query.Option{
		query.WithPage(0),
		query.WithSize(0),
		query.WithSorts("id asc"),
	}
	if in.Name != "" {
		opts = append(opts, query.WithCondition("name like ?", "%"+in.Name+"%"))
	}
	if in.Path != "" {
		opts = append(opts, query.WithCondition("path like ?", "%"+in.Path+"%"))
	}
	if in.Method != "" {
		opts = append(opts, query.WithCondition("method = ?", in.Method))
	}

	return query.NewQueryBuilder(opts...).Build()
}

func buildMenuQuery(in *permissionrpc.FindMenuListReq) (int, int, string, string, []any) {
	opts := []query.Option{
		query.WithPage(0),
		query.WithSize(0),
		query.WithSorts("rank asc", "id asc"),
	}
	if in.Title != "" {
		opts = append(opts, query.WithCondition("title like ?", "%"+in.Title+"%"))
	}
	if in.Name != "" {
		opts = append(opts, query.WithCondition("name like ?", "%"+in.Name+"%"))
	}

	return query.NewQueryBuilder(opts...).Build()
}

func buildRoleQuery(in *permissionrpc.FindRoleListReq) (int, int, string, string, []any) {
	var opts []query.Option
	if in.Paginate != nil {
		opts = append(opts, query.WithPage(int(in.Paginate.Page)))
		opts = append(opts, query.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, query.WithSorts(in.Paginate.Sorts...))
	}
	if in.RoleKey != "" {
		opts = append(opts, query.WithCondition("role_key like ?", "%"+in.RoleKey+"%"))
	}
	if in.Status != 0 {
		opts = append(opts, query.WithCondition("status = ?", in.Status))
	}

	return query.NewQueryBuilder(opts...).Build()
}

func buildPageResp(page int, size int, total int64) *permissionrpc.PageResp {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = int(total)
	}

	return &permissionrpc.PageResp{
		Page:     int64(page),
		PageSize: int64(size),
		Total:    total,
	}
}

func (h *permissionHelper) saveAddApiTree(in *permissionrpc.AddApiReq, parentID int64) (*permissionrpc.Api, int64, error) {
	entity := convertAddApiIn(in)
	if entity.ParentId == 0 {
		entity.ParentId = parentID
	}
	if _, err := h.svcCtx.TApiModel.Save(h.ctx, entity); err != nil {
		return nil, 0, err
	}

	out := convertApiOut(entity)
	var count int64 = 1
	for _, child := range in.Children {
		childOut, childCount, err := h.saveAddApiTree(child, entity.Id)
		if err != nil {
			return nil, count, err
		}
		out.Children = append(out.Children, childOut)
		count += childCount
	}

	return out, count, nil
}

func (h *permissionHelper) saveUpdateApiTree(in *permissionrpc.UpdateApiReq, parentID int64) (*permissionrpc.Api, int64, error) {
	entity := convertUpdateApiIn(in)
	if entity.ParentId == 0 {
		entity.ParentId = parentID
	}
	if _, err := h.svcCtx.TApiModel.Save(h.ctx, entity); err != nil {
		return nil, 0, err
	}

	out := convertApiOut(entity)
	var count int64 = 1
	for _, child := range in.Children {
		childOut, childCount, err := h.saveUpdateApiTree(child, entity.Id)
		if err != nil {
			return nil, count, err
		}
		out.Children = append(out.Children, childOut)
		count += childCount
	}

	return out, count, nil
}

func (h *permissionHelper) saveAddMenuTree(in *permissionrpc.AddMenuReq, parentID int64) (*permissionrpc.Menu, int64, error) {
	entity := convertAddMenuIn(in)
	if entity.ParentId == 0 {
		entity.ParentId = parentID
	}
	if _, err := h.svcCtx.TMenuModel.Save(h.ctx, entity); err != nil {
		return nil, 0, err
	}

	out := convertMenuOut(entity)
	var count int64 = 1
	for _, child := range in.Children {
		childOut, childCount, err := h.saveAddMenuTree(child, entity.Id)
		if err != nil {
			return nil, count, err
		}
		out.Children = append(out.Children, childOut)
		count += childCount
	}

	return out, count, nil
}
