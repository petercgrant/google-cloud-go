// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package cx_test

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
)

func ExampleNewAgentsClient() {
	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleAgentsClient_ListAgents() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.ListAgentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListAgents(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleAgentsClient_GetAgent() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.GetAgentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetAgent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAgentsClient_CreateAgent() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.CreateAgentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateAgent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAgentsClient_UpdateAgent() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.UpdateAgentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateAgent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAgentsClient_DeleteAgent() {
	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.DeleteAgentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteAgent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAgentsClient_ExportAgent() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.ExportAgentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportAgent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAgentsClient_RestoreAgent() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewAgentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.RestoreAgentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RestoreAgent(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
