// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
)

var oauthServerConfig *osin.ServerConfig
var oauthServer *osin.Server
var oauthStorage *OAuthStorage

func init() {
	oauthServerConfig = &osin.ServerConfig{
		AuthorizationExpiration:   250,
		AccessExpiration:          3600,
		TokenType:                 "bearer",
		AllowedAuthorizeTypes:     osin.AllowedAuthorizeType{osin.CODE},
		AllowedAccessTypes:        osin.AllowedAccessType{osin.AUTHORIZATION_CODE},
		ErrorStatusCode:           200,
		AllowClientSecretInParams: false,
		AllowGetAccessRequest:     false,
	}

	oauthStorage = NewOAuthStorage()

	oauthServer = osin.NewServer(oauthServerConfig, oauthStorage)
}

func authorizeRoute(c *gin.Context) {
	r := c.Request
	w := c.Writer

	resp := oauthServer.NewResponse()
	defer resp.Close()

	if ar := oauthServer.HandleAuthorizeRequest(resp, r); ar != nil {

		// HANDLE LOGIN PAGE HERE

		ar.Authorized = true
		oauthServer.FinishAuthorizeRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}

func tokenRoute(c *gin.Context) {
	r := c.Request
	w := c.Writer

	resp := oauthServer.NewResponse()
	defer resp.Close()

	if ar := oauthServer.HandleAccessRequest(resp, r); ar != nil {
		ar.Authorized = true
		oauthServer.FinishAccessRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}