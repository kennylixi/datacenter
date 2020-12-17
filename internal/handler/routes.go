// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	common "datacenter/internal/handler/common"
	user "datacenter/internal/handler/user"
	votes "datacenter/internal/handler/votes"
	"datacenter/internal/svc"
	"github.com/tal-tech/go-zero/rest"
)

func dirhandler(patern, filedir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler := http.StripPrefix(patern, http.FileServer(http.Dir(filedir)))
		handler.ServeHTTP(w, req)

	}
}

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/MP_verify_NT04cqknJe0em3mT.txt",
				Handler: common.VotesVerificationHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/common/appinfo",
				Handler: common.AppInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/common/snsinfo",
				Handler: common.SnsInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/common/wx/ticket",
				Handler: common.WxTicketHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/common/qiuniu/token",
				Handler: common.QiuniuTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/ping",
				Handler: user.PingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/wx/login",
				Handler: user.WxloginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/wx/login",
				Handler: user.Code2SessionHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Usercheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/dc/info",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/votes/activity/info",
				Handler: votes.ActivityInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/votes/activity/view",
				Handler: votes.ActivityIcrViewHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/votes/enroll/info",
				Handler: votes.EnrollInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/votes/enroll/lists",
				Handler: votes.EnrollListsHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Usercheck},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/votes/vote",
					Handler: votes.VoteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/votes/enroll",
					Handler: votes.EnrollHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
