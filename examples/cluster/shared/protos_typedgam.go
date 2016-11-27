// Code generated by protoc-gen-gogo.
// source: protos.proto
// DO NOT EDIT!

/*
Package shared is a generated protocol buffer package.

It is generated from these files:
	protos.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	AddRequest
	AddResponse
*/
package shared

import log "log"
import time "time"
import errors "errors"
import github_com_AsynkronIT_gam_cluster "github.com/AsynkronIT/gam/cluster"
import github_com_AsynkronIT_gam_actor "github.com/AsynkronIT/gam/actor"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HelloRequestFuture struct {
	Value *HelloRequest
	Err   error
}

type HelloResponseFuture struct {
	Value *HelloResponse
	Err   error
}

type AddRequestFuture struct {
	Value *AddRequest
	Err   error
}

type AddResponseFuture struct {
	Value *AddResponse
	Err   error
}

var xHelloFactory func() Hello

func HelloFactory(factory func() Hello) {
	xHelloFactory = factory
}

func GetHelloGrain(id string) *HelloGrain {
	return &HelloGrain{Id: id}
}

type Hello interface {
	SayHello(*HelloRequest) (*HelloResponse, error)
	Add(*AddRequest) (*AddResponse, error)
}
type HelloGrain struct {
	Id string
}

func (g *HelloGrain) SayHello(r *HelloRequest, timeout time.Duration) (*HelloResponse, error) {
	pid := github_com_AsynkronIT_gam_cluster.Get(g.Id, "Hello")
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	gr := &github_com_AsynkronIT_gam_cluster.GrainRequest{Method: "SayHello", MessageData: bytes}
	r0 := pid.RequestFuture(gr, timeout)
	r1, err := r0.Result()
	if err != nil {
		return nil, err
	}
	switch r2 := r1.(type) {
	case *github_com_AsynkronIT_gam_cluster.GrainResponse:
		r3 := &HelloResponse{}
		err = proto.Unmarshal(r2.MessageData, r3)
		if err != nil {
			return nil, err
		}
		return r3, nil
	case *github_com_AsynkronIT_gam_cluster.GrainErrorResponse:
		return nil, errors.New(r2.Err)
	default:
		return nil, errors.New("Unknown response")
	}
}

func (g *HelloGrain) SayHelloChan(r *HelloRequest, timeout time.Duration) <-chan *HelloResponseFuture {
	c := make(chan *HelloResponseFuture, 1)
	go func() {
		defer close(c)
		res, err := g.SayHello(r, timeout)
		c <- &HelloResponseFuture{Value: res, Err: err}
	}()
	return c
}

func (g *HelloGrain) Add(r *AddRequest, timeout time.Duration) (*AddResponse, error) {
	pid := github_com_AsynkronIT_gam_cluster.Get(g.Id, "Hello")
	bytes, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	gr := &github_com_AsynkronIT_gam_cluster.GrainRequest{Method: "Add", MessageData: bytes}
	r0 := pid.RequestFuture(gr, timeout)
	r1, err := r0.Result()
	if err != nil {
		return nil, err
	}
	switch r2 := r1.(type) {
	case *github_com_AsynkronIT_gam_cluster.GrainResponse:
		r3 := &AddResponse{}
		err = proto.Unmarshal(r2.MessageData, r3)
		if err != nil {
			return nil, err
		}
		return r3, nil
	case *github_com_AsynkronIT_gam_cluster.GrainErrorResponse:
		return nil, errors.New(r2.Err)
	default:
		return nil, errors.New("Unknown response")
	}
}

func (g *HelloGrain) AddChan(r *AddRequest, timeout time.Duration) <-chan *AddResponseFuture {
	c := make(chan *AddResponseFuture, 1)
	go func() {
		defer close(c)
		res, err := g.Add(r, timeout)
		c <- &AddResponseFuture{Value: res, Err: err}
	}()
	return c
}

type HelloActor struct {
	inner Hello
}

func (a *HelloActor) Receive(ctx github_com_AsynkronIT_gam_actor.Context) {
	switch msg := ctx.Message().(type) {
	case *github_com_AsynkronIT_gam_cluster.GrainRequest:
		switch msg.Method {
		case "SayHello":
			req := &HelloRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.SayHello(req)
			if err == nil {
				bytes, err := proto.Marshal(r0)
				if err != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", err)
				}
				resp := &github_com_AsynkronIT_gam_cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &github_com_AsynkronIT_gam_cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
		case "Add":
			req := &AddRequest{}
			err := proto.Unmarshal(msg.MessageData, req)
			if err != nil {
				log.Fatalf("[GRAIN] proto.Unmarshal failed %v", err)
			}
			r0, err := a.inner.Add(req)
			if err == nil {
				bytes, err := proto.Marshal(r0)
				if err != nil {
					log.Fatalf("[GRAIN] proto.Marshal failed %v", err)
				}
				resp := &github_com_AsynkronIT_gam_cluster.GrainResponse{MessageData: bytes}
				ctx.Respond(resp)
			} else {
				resp := &github_com_AsynkronIT_gam_cluster.GrainErrorResponse{Err: err.Error()}
				ctx.Respond(resp)
			}
		}
	default:
		log.Printf("Unknown message %v", msg)
	}
}

func init() {
	github_com_AsynkronIT_gam_cluster.Register("Hello", github_com_AsynkronIT_gam_actor.FromProducer(func() github_com_AsynkronIT_gam_actor.Actor { return &HelloActor{inner: xHelloFactory()} }))
}