package api

import (
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/permissionrpc"
	"github.com/go-openapi/spec"
)

func convertApiTypes(req *permissionrpc.Api) *types.ApiBackVO {

	children := make([]*types.ApiBackVO, 0)
	for _, v := range req.Children {
		m := convertApiTypes(v)
		children = append(children, m)
	}

	out := &types.ApiBackVO{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Name:      req.Name,
		Path:      req.Path,
		Method:    req.Method,
		Traceable: req.Traceable,
		Status:    req.Status,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		Children:  children,
	}

	return out
}

func getRoutes(sp *spec.Swagger) map[string]map[string]*spec.Operation {
	// map[path][method] -> operation
	routes := make(map[string]map[string]*spec.Operation)

	for k, v := range sp.Paths.Paths {
		if routes[k] == nil {
			routes[k] = make(map[string]*spec.Operation)
		}

		if v.Get != nil {
			routes[k][http.MethodGet] = v.Get
		}

		if v.Put != nil {
			routes[k][http.MethodPut] = v.Put
		}

		if v.Post != nil {
			routes[k][http.MethodPost] = v.Post
		}

		if v.Delete != nil {
			routes[k][http.MethodDelete] = v.Delete
		}

		if v.Options != nil {
			routes[k][http.MethodOptions] = v.Options
		}

		if v.Head != nil {
			routes[k][http.MethodHead] = v.Head
		}
	}

	return routes
}
