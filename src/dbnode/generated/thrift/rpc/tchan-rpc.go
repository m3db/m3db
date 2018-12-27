// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package rpc is generated code used to make or handle TChannel calls using Thrift.
package rpc

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanCluster is the interface that defines the server handler and client interface.
type TChanCluster interface {
	Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error)
	FetchTagged(ctx thrift.Context, req *FetchTaggedRequest) (*FetchTaggedResult_, error)
	Health(ctx thrift.Context) (*HealthResult_, error)
	Query(ctx thrift.Context, req *QueryRequest) (*QueryResult_, error)
	Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error)
	Write(ctx thrift.Context, req *WriteRequest) error
	WriteTagged(ctx thrift.Context, req *WriteTaggedRequest) error
}

// TChanNode is the interface that defines the server handler and client interface.
type TChanNode interface {
	Bootstrapped(ctx thrift.Context) (*NodeBootstrappedResult_, error)
	Debug(ctx thrift.Context) (*NodeDebugResult_, error)
	Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error)
	FetchBatchRaw(ctx thrift.Context, req *FetchBatchRawRequest) (*FetchBatchRawResult_, error)
	FetchBlocksMetadataRawV2(ctx thrift.Context, req *FetchBlocksMetadataRawV2Request) (*FetchBlocksMetadataRawV2Result_, error)
	FetchBlocksRaw(ctx thrift.Context, req *FetchBlocksRawRequest) (*FetchBlocksRawResult_, error)
	FetchTagged(ctx thrift.Context, req *FetchTaggedRequest) (*FetchTaggedResult_, error)
	GetPersistRateLimit(ctx thrift.Context) (*NodePersistRateLimitResult_, error)
	GetWriteNewSeriesAsync(ctx thrift.Context) (*NodeWriteNewSeriesAsyncResult_, error)
	GetWriteNewSeriesBackoffDuration(ctx thrift.Context) (*NodeWriteNewSeriesBackoffDurationResult_, error)
	GetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context) (*NodeWriteNewSeriesLimitPerShardPerSecondResult_, error)
	Health(ctx thrift.Context) (*NodeHealthResult_, error)
	Query(ctx thrift.Context, req *QueryRequest) (*QueryResult_, error)
	Repair(ctx thrift.Context) error
	SetPersistRateLimit(ctx thrift.Context, req *NodeSetPersistRateLimitRequest) (*NodePersistRateLimitResult_, error)
	SetWriteNewSeriesAsync(ctx thrift.Context, req *NodeSetWriteNewSeriesAsyncRequest) (*NodeWriteNewSeriesAsyncResult_, error)
	SetWriteNewSeriesBackoffDuration(ctx thrift.Context, req *NodeSetWriteNewSeriesBackoffDurationRequest) (*NodeWriteNewSeriesBackoffDurationResult_, error)
	SetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context, req *NodeSetWriteNewSeriesLimitPerShardPerSecondRequest) (*NodeWriteNewSeriesLimitPerShardPerSecondResult_, error)
	Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error)
	Write(ctx thrift.Context, req *WriteRequest) error
	WriteBatchRaw(ctx thrift.Context, req *WriteBatchRawRequest) error
	WriteTagged(ctx thrift.Context, req *WriteTaggedRequest) error
	WriteTaggedBatchRaw(ctx thrift.Context, req *WriteTaggedBatchRawRequest) error
}

// Implementation of a client and service handler.

type tchanClusterClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanClusterInheritedClient(thriftService string, client thrift.TChanClient) *tchanClusterClient {
	return &tchanClusterClient{
		thriftService,
		client,
	}
}

// NewTChanClusterClient creates a client that can be used to make remote calls.
func NewTChanClusterClient(client thrift.TChanClient) TChanCluster {
	return NewTChanClusterInheritedClient("Cluster", client)
}

func (c *tchanClusterClient) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	var resp ClusterFetchResult
	args := ClusterFetchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetch", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetch")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) FetchTagged(ctx thrift.Context, req *FetchTaggedRequest) (*FetchTaggedResult_, error) {
	var resp ClusterFetchTaggedResult
	args := ClusterFetchTaggedArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchTagged", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetchTagged")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Health(ctx thrift.Context) (*HealthResult_, error) {
	var resp ClusterHealthResult
	args := ClusterHealthArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "health", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for health")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Query(ctx thrift.Context, req *QueryRequest) (*QueryResult_, error) {
	var resp ClusterQueryResult
	args := ClusterQueryArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "query", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for query")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	var resp ClusterTruncateResult
	args := ClusterTruncateArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "truncate", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for truncate")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Write(ctx thrift.Context, req *WriteRequest) error {
	var resp ClusterWriteResult
	args := ClusterWriteArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "write", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for write")
		}
	}

	return err
}

func (c *tchanClusterClient) WriteTagged(ctx thrift.Context, req *WriteTaggedRequest) error {
	var resp ClusterWriteTaggedResult
	args := ClusterWriteTaggedArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "writeTagged", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for writeTagged")
		}
	}

	return err
}

type tchanClusterServer struct {
	handler TChanCluster
}

// NewTChanClusterServer wraps a handler for TChanCluster so it can be
// registered with a thrift.Server.
func NewTChanClusterServer(handler TChanCluster) thrift.TChanServer {
	return &tchanClusterServer{
		handler,
	}
}

func (s *tchanClusterServer) Service() string {
	return "Cluster"
}

func (s *tchanClusterServer) Methods() []string {
	return []string{
		"fetch",
		"fetchTagged",
		"health",
		"query",
		"truncate",
		"write",
		"writeTagged",
	}
}

func (s *tchanClusterServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "fetch":
		return s.handleFetch(ctx, protocol)
	case "fetchTagged":
		return s.handleFetchTagged(ctx, protocol)
	case "health":
		return s.handleHealth(ctx, protocol)
	case "query":
		return s.handleQuery(ctx, protocol)
	case "truncate":
		return s.handleTruncate(ctx, protocol)
	case "write":
		return s.handleWrite(ctx, protocol)
	case "writeTagged":
		return s.handleWriteTagged(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanClusterServer) handleFetch(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterFetchArgs
	var res ClusterFetchResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Fetch(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleFetchTagged(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterFetchTaggedArgs
	var res ClusterFetchTaggedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchTagged(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleHealth(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterHealthArgs
	var res ClusterHealthResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Health(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleQuery(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterQueryArgs
	var res ClusterQueryResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Query(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleTruncate(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterTruncateArgs
	var res ClusterTruncateResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Truncate(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleWrite(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterWriteArgs
	var res ClusterWriteResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Write(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanClusterServer) handleWriteTagged(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ClusterWriteTaggedArgs
	var res ClusterWriteTaggedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.WriteTagged(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

type tchanNodeClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanNodeInheritedClient(thriftService string, client thrift.TChanClient) *tchanNodeClient {
	return &tchanNodeClient{
		thriftService,
		client,
	}
}

// NewTChanNodeClient creates a client that can be used to make remote calls.
func NewTChanNodeClient(client thrift.TChanClient) TChanNode {
	return NewTChanNodeInheritedClient("Node", client)
}

func (c *tchanNodeClient) Bootstrapped(ctx thrift.Context) (*NodeBootstrappedResult_, error) {
	var resp NodeBootstrappedResult
	args := NodeBootstrappedArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "bootstrapped", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for bootstrapped")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Debug(ctx thrift.Context) (*NodeDebugResult_, error) {
	var resp NodeDebugResult
	args := NodeDebugArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "debug", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for debug")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	var resp NodeFetchResult
	args := NodeFetchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetch", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetch")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBatchRaw(ctx thrift.Context, req *FetchBatchRawRequest) (*FetchBatchRawResult_, error) {
	var resp NodeFetchBatchRawResult
	args := NodeFetchBatchRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBatchRaw", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetchBatchRaw")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBlocksMetadataRawV2(ctx thrift.Context, req *FetchBlocksMetadataRawV2Request) (*FetchBlocksMetadataRawV2Result_, error) {
	var resp NodeFetchBlocksMetadataRawV2Result
	args := NodeFetchBlocksMetadataRawV2Args{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBlocksMetadataRawV2", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetchBlocksMetadataRawV2")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBlocksRaw(ctx thrift.Context, req *FetchBlocksRawRequest) (*FetchBlocksRawResult_, error) {
	var resp NodeFetchBlocksRawResult
	args := NodeFetchBlocksRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBlocksRaw", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetchBlocksRaw")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchTagged(ctx thrift.Context, req *FetchTaggedRequest) (*FetchTaggedResult_, error) {
	var resp NodeFetchTaggedResult
	args := NodeFetchTaggedArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchTagged", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for fetchTagged")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) GetPersistRateLimit(ctx thrift.Context) (*NodePersistRateLimitResult_, error) {
	var resp NodeGetPersistRateLimitResult
	args := NodeGetPersistRateLimitArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "getPersistRateLimit", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for getPersistRateLimit")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) GetWriteNewSeriesAsync(ctx thrift.Context) (*NodeWriteNewSeriesAsyncResult_, error) {
	var resp NodeGetWriteNewSeriesAsyncResult
	args := NodeGetWriteNewSeriesAsyncArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "getWriteNewSeriesAsync", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for getWriteNewSeriesAsync")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) GetWriteNewSeriesBackoffDuration(ctx thrift.Context) (*NodeWriteNewSeriesBackoffDurationResult_, error) {
	var resp NodeGetWriteNewSeriesBackoffDurationResult
	args := NodeGetWriteNewSeriesBackoffDurationArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "getWriteNewSeriesBackoffDuration", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for getWriteNewSeriesBackoffDuration")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) GetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context) (*NodeWriteNewSeriesLimitPerShardPerSecondResult_, error) {
	var resp NodeGetWriteNewSeriesLimitPerShardPerSecondResult
	args := NodeGetWriteNewSeriesLimitPerShardPerSecondArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "getWriteNewSeriesLimitPerShardPerSecond", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for getWriteNewSeriesLimitPerShardPerSecond")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Health(ctx thrift.Context) (*NodeHealthResult_, error) {
	var resp NodeHealthResult
	args := NodeHealthArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "health", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for health")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Query(ctx thrift.Context, req *QueryRequest) (*QueryResult_, error) {
	var resp NodeQueryResult
	args := NodeQueryArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "query", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for query")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Repair(ctx thrift.Context) error {
	var resp NodeRepairResult
	args := NodeRepairArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "repair", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for repair")
		}
	}

	return err
}

func (c *tchanNodeClient) SetPersistRateLimit(ctx thrift.Context, req *NodeSetPersistRateLimitRequest) (*NodePersistRateLimitResult_, error) {
	var resp NodeSetPersistRateLimitResult
	args := NodeSetPersistRateLimitArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "setPersistRateLimit", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for setPersistRateLimit")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) SetWriteNewSeriesAsync(ctx thrift.Context, req *NodeSetWriteNewSeriesAsyncRequest) (*NodeWriteNewSeriesAsyncResult_, error) {
	var resp NodeSetWriteNewSeriesAsyncResult
	args := NodeSetWriteNewSeriesAsyncArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "setWriteNewSeriesAsync", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for setWriteNewSeriesAsync")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) SetWriteNewSeriesBackoffDuration(ctx thrift.Context, req *NodeSetWriteNewSeriesBackoffDurationRequest) (*NodeWriteNewSeriesBackoffDurationResult_, error) {
	var resp NodeSetWriteNewSeriesBackoffDurationResult
	args := NodeSetWriteNewSeriesBackoffDurationArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "setWriteNewSeriesBackoffDuration", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for setWriteNewSeriesBackoffDuration")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) SetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context, req *NodeSetWriteNewSeriesLimitPerShardPerSecondRequest) (*NodeWriteNewSeriesLimitPerShardPerSecondResult_, error) {
	var resp NodeSetWriteNewSeriesLimitPerShardPerSecondResult
	args := NodeSetWriteNewSeriesLimitPerShardPerSecondArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "setWriteNewSeriesLimitPerShardPerSecond", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for setWriteNewSeriesLimitPerShardPerSecond")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	var resp NodeTruncateResult
	args := NodeTruncateArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "truncate", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for truncate")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Write(ctx thrift.Context, req *WriteRequest) error {
	var resp NodeWriteResult
	args := NodeWriteArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "write", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for write")
		}
	}

	return err
}

func (c *tchanNodeClient) WriteBatchRaw(ctx thrift.Context, req *WriteBatchRawRequest) error {
	var resp NodeWriteBatchRawResult
	args := NodeWriteBatchRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "writeBatchRaw", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for writeBatchRaw")
		}
	}

	return err
}

func (c *tchanNodeClient) WriteTagged(ctx thrift.Context, req *WriteTaggedRequest) error {
	var resp NodeWriteTaggedResult
	args := NodeWriteTaggedArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "writeTagged", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for writeTagged")
		}
	}

	return err
}

func (c *tchanNodeClient) WriteTaggedBatchRaw(ctx thrift.Context, req *WriteTaggedBatchRawRequest) error {
	var resp NodeWriteTaggedBatchRawResult
	args := NodeWriteTaggedBatchRawArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "writeTaggedBatchRaw", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.Err != nil:
			err = resp.Err
		default:
			err = fmt.Errorf("received no result or unknown exception for writeTaggedBatchRaw")
		}
	}

	return err
}

type tchanNodeServer struct {
	handler TChanNode
}

// NewTChanNodeServer wraps a handler for TChanNode so it can be
// registered with a thrift.Server.
func NewTChanNodeServer(handler TChanNode) thrift.TChanServer {
	return &tchanNodeServer{
		handler,
	}
}

func (s *tchanNodeServer) Service() string {
	return "Node"
}

func (s *tchanNodeServer) Methods() []string {
	return []string{
		"bootstrapped",
		"debug",
		"fetch",
		"fetchBatchRaw",
		"fetchBlocksMetadataRawV2",
		"fetchBlocksRaw",
		"fetchTagged",
		"getPersistRateLimit",
		"getWriteNewSeriesAsync",
		"getWriteNewSeriesBackoffDuration",
		"getWriteNewSeriesLimitPerShardPerSecond",
		"health",
		"query",
		"repair",
		"setPersistRateLimit",
		"setWriteNewSeriesAsync",
		"setWriteNewSeriesBackoffDuration",
		"setWriteNewSeriesLimitPerShardPerSecond",
		"truncate",
		"write",
		"writeBatchRaw",
		"writeTagged",
		"writeTaggedBatchRaw",
	}
}

func (s *tchanNodeServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "bootstrapped":
		return s.handleBootstrapped(ctx, protocol)
	case "debug":
		return s.handleDebug(ctx, protocol)
	case "fetch":
		return s.handleFetch(ctx, protocol)
	case "fetchBatchRaw":
		return s.handleFetchBatchRaw(ctx, protocol)
	case "fetchBlocksMetadataRawV2":
		return s.handleFetchBlocksMetadataRawV2(ctx, protocol)
	case "fetchBlocksRaw":
		return s.handleFetchBlocksRaw(ctx, protocol)
	case "fetchTagged":
		return s.handleFetchTagged(ctx, protocol)
	case "getPersistRateLimit":
		return s.handleGetPersistRateLimit(ctx, protocol)
	case "getWriteNewSeriesAsync":
		return s.handleGetWriteNewSeriesAsync(ctx, protocol)
	case "getWriteNewSeriesBackoffDuration":
		return s.handleGetWriteNewSeriesBackoffDuration(ctx, protocol)
	case "getWriteNewSeriesLimitPerShardPerSecond":
		return s.handleGetWriteNewSeriesLimitPerShardPerSecond(ctx, protocol)
	case "health":
		return s.handleHealth(ctx, protocol)
	case "query":
		return s.handleQuery(ctx, protocol)
	case "repair":
		return s.handleRepair(ctx, protocol)
	case "setPersistRateLimit":
		return s.handleSetPersistRateLimit(ctx, protocol)
	case "setWriteNewSeriesAsync":
		return s.handleSetWriteNewSeriesAsync(ctx, protocol)
	case "setWriteNewSeriesBackoffDuration":
		return s.handleSetWriteNewSeriesBackoffDuration(ctx, protocol)
	case "setWriteNewSeriesLimitPerShardPerSecond":
		return s.handleSetWriteNewSeriesLimitPerShardPerSecond(ctx, protocol)
	case "truncate":
		return s.handleTruncate(ctx, protocol)
	case "write":
		return s.handleWrite(ctx, protocol)
	case "writeBatchRaw":
		return s.handleWriteBatchRaw(ctx, protocol)
	case "writeTagged":
		return s.handleWriteTagged(ctx, protocol)
	case "writeTaggedBatchRaw":
		return s.handleWriteTaggedBatchRaw(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanNodeServer) handleBootstrapped(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeBootstrappedArgs
	var res NodeBootstrappedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Bootstrapped(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleDebug(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeDebugArgs
	var res NodeDebugResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Debug(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetch(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchArgs
	var res NodeFetchResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Fetch(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchBatchRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBatchRawArgs
	var res NodeFetchBatchRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBatchRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchBlocksMetadataRawV2(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBlocksMetadataRawV2Args
	var res NodeFetchBlocksMetadataRawV2Result

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBlocksMetadataRawV2(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchBlocksRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBlocksRawArgs
	var res NodeFetchBlocksRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBlocksRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleFetchTagged(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchTaggedArgs
	var res NodeFetchTaggedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchTagged(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleGetPersistRateLimit(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeGetPersistRateLimitArgs
	var res NodeGetPersistRateLimitResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetPersistRateLimit(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleGetWriteNewSeriesAsync(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeGetWriteNewSeriesAsyncArgs
	var res NodeGetWriteNewSeriesAsyncResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetWriteNewSeriesAsync(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleGetWriteNewSeriesBackoffDuration(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeGetWriteNewSeriesBackoffDurationArgs
	var res NodeGetWriteNewSeriesBackoffDurationResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetWriteNewSeriesBackoffDuration(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleGetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeGetWriteNewSeriesLimitPerShardPerSecondArgs
	var res NodeGetWriteNewSeriesLimitPerShardPerSecondResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetWriteNewSeriesLimitPerShardPerSecond(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleHealth(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeHealthArgs
	var res NodeHealthResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Health(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleQuery(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeQueryArgs
	var res NodeQueryResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Query(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleRepair(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeRepairArgs
	var res NodeRepairResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Repair(ctx)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleSetPersistRateLimit(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeSetPersistRateLimitArgs
	var res NodeSetPersistRateLimitResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.SetPersistRateLimit(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleSetWriteNewSeriesAsync(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeSetWriteNewSeriesAsyncArgs
	var res NodeSetWriteNewSeriesAsyncResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.SetWriteNewSeriesAsync(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleSetWriteNewSeriesBackoffDuration(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeSetWriteNewSeriesBackoffDurationArgs
	var res NodeSetWriteNewSeriesBackoffDurationResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.SetWriteNewSeriesBackoffDuration(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleSetWriteNewSeriesLimitPerShardPerSecond(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeSetWriteNewSeriesLimitPerShardPerSecondArgs
	var res NodeSetWriteNewSeriesLimitPerShardPerSecondResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.SetWriteNewSeriesLimitPerShardPerSecond(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleTruncate(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeTruncateArgs
	var res NodeTruncateResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Truncate(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleWrite(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteArgs
	var res NodeWriteResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.Write(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleWriteBatchRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteBatchRawArgs
	var res NodeWriteBatchRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.WriteBatchRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *WriteBatchRawErrors:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *WriteBatchRawErrors but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleWriteTagged(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteTaggedArgs
	var res NodeWriteTaggedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.WriteTagged(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *Error:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *Error but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanNodeServer) handleWriteTaggedBatchRaw(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteTaggedBatchRawArgs
	var res NodeWriteTaggedBatchRawResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.WriteTaggedBatchRaw(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *WriteBatchRawErrors:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *WriteBatchRawErrors but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}
