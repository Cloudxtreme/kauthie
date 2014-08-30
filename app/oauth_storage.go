// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"errors"

	"github.com/RangelReale/osin"
	"gopkg.in/mgo.v2"
)

type Storage *osin.Storage
type Client *osin.Client

type OAuthStorage struct {
	collection *mgo.Collection
}

func NewOAuthStorage() *OAuthStorage {
	return &OAuthStorage{}
}

func (s *OAuthStorage) Clone() osin.Storage {
	return s
}

func (s *OAuthStorage) Close() {
}

func (s *OAuthStorage) GetClient(id string) (osin.Client, error) {
	var client osin.Client
	if client != nil {
		return client, nil
	}
	return nil, errors.New("Client not found")
}

func (s *OAuthStorage) SetClient(id string, client osin.Client) error {
	// Save client data from client
	return nil
}

func (s *OAuthStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	// Save authorization to mongo
	return nil
}

func (s *OAuthStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	var data *osin.AuthorizeData
	if data != nil {
		return data, nil
	}
	return nil, errors.New("Authorize not found")
}

func (s *OAuthStorage) RemoveAuthorize(code string) error {
	// Delete authorization
	return nil
}

func (s *OAuthStorage) SaveAccess(data *osin.AccessData) error {
	// Save AccessData to mongo
	// If refresh token -> inster refresh token too
	return nil
}

func (s *OAuthStorage) LoadAccess(code string) (*osin.AccessData, error) {
	var data *osin.AccessData
	if data != nil {
		return data, nil
	}
	return nil, errors.New("Access not found")
}

func (s *OAuthStorage) RemoveAccess(code string) error {
	// Delete access token from storage
	return nil
}

func (s *OAuthStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	// Load refresh token
	var accessToken string
	if accessToken != "" {
		return s.LoadAccess(accessToken)
	}
	return nil, errors.New("Refresh not found")
}

func (s *OAuthStorage) RemoveRefresh(code string) error {
	// Delete from mongo
	return nil
}
