// Copyright (c) 2016 Uber Technologies, Inc.
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

// @generated Code generated by thrift-gen. Do not modify.

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
	Health(ctx thrift.Context) (*HealthResult_, error)
	Write(ctx thrift.Context, req *WriteRequest) error
}

// TChanNode is the interface that defines the server handler and client interface.
type TChanNode interface {
	Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error)
	FetchBlocks(ctx thrift.Context, req *FetchBlocksRequest) (*FetchBlocksResult_, error)
	FetchBlocksMetadata(ctx thrift.Context, req *FetchBlocksMetadataRequest) (*FetchBlocksMetadataResult_, error)
	FetchRawBatch(ctx thrift.Context, req *FetchRawBatchRequest) (*FetchRawBatchResult_, error)
	Health(ctx thrift.Context) (*HealthResult_, error)
	Write(ctx thrift.Context, req *WriteRequest) error
	WriteBatch(ctx thrift.Context, req *WriteBatchRequest) error
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
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanClusterClient) Health(ctx thrift.Context) (*HealthResult_, error) {
	var resp ClusterHealthResult
	args := ClusterHealthArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "health", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
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
		if e := resp.Err; e != nil {
			err = e
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
		"health",
		"write",
	}
}

func (s *tchanClusterServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "fetch":
		return s.handleFetch(ctx, protocol)
	case "health":
		return s.handleHealth(ctx, protocol)
	case "write":
		return s.handleWrite(ctx, protocol)

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

func (c *tchanNodeClient) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	var resp NodeFetchResult
	args := NodeFetchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetch", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBlocks(ctx thrift.Context, req *FetchBlocksRequest) (*FetchBlocksResult_, error) {
	var resp NodeFetchBlocksResult
	args := NodeFetchBlocksArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBlocks", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchBlocksMetadata(ctx thrift.Context, req *FetchBlocksMetadataRequest) (*FetchBlocksMetadataResult_, error) {
	var resp NodeFetchBlocksMetadataResult
	args := NodeFetchBlocksMetadataArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchBlocksMetadata", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) FetchRawBatch(ctx thrift.Context, req *FetchRawBatchRequest) (*FetchRawBatchResult_, error) {
	var resp NodeFetchRawBatchResult
	args := NodeFetchRawBatchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "fetchRawBatch", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanNodeClient) Health(ctx thrift.Context) (*HealthResult_, error) {
	var resp NodeHealthResult
	args := NodeHealthArgs{}
	success, err := c.client.Call(ctx, c.thriftService, "health", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
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
		if e := resp.Err; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanNodeClient) WriteBatch(ctx thrift.Context, req *WriteBatchRequest) error {
	var resp NodeWriteBatchResult
	args := NodeWriteBatchArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "writeBatch", &args, &resp)
	if err == nil && !success {
		if e := resp.Err; e != nil {
			err = e
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
		"fetch",
		"fetchBlocks",
		"fetchBlocksMetadata",
		"fetchRawBatch",
		"health",
		"write",
		"writeBatch",
	}
}

func (s *tchanNodeServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "fetch":
		return s.handleFetch(ctx, protocol)
	case "fetchBlocks":
		return s.handleFetchBlocks(ctx, protocol)
	case "fetchBlocksMetadata":
		return s.handleFetchBlocksMetadata(ctx, protocol)
	case "fetchRawBatch":
		return s.handleFetchRawBatch(ctx, protocol)
	case "health":
		return s.handleHealth(ctx, protocol)
	case "write":
		return s.handleWrite(ctx, protocol)
	case "writeBatch":
		return s.handleWriteBatch(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
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

func (s *tchanNodeServer) handleFetchBlocks(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBlocksArgs
	var res NodeFetchBlocksResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBlocks(ctx, req.Req)

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

func (s *tchanNodeServer) handleFetchBlocksMetadata(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchBlocksMetadataArgs
	var res NodeFetchBlocksMetadataResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchBlocksMetadata(ctx, req.Req)

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

func (s *tchanNodeServer) handleFetchRawBatch(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeFetchRawBatchArgs
	var res NodeFetchRawBatchResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.FetchRawBatch(ctx, req.Req)

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

func (s *tchanNodeServer) handleWriteBatch(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req NodeWriteBatchArgs
	var res NodeWriteBatchResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.WriteBatch(ctx, req.Req)

	if err != nil {
		switch v := err.(type) {
		case *WriteBatchErrors:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for err returned non-nil error type *WriteBatchErrors but nil value")
			}
			res.Err = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}
