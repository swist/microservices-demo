// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/home", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set log handler
	ctx := req.Context()
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{}
	ctx = context.WithValue(ctx, ctxKeyLog{}, log)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()
	svc := new(frontendServer)
	svc.homeHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v; want 200", res.Code)
	}
}
