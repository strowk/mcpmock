package fxm

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	foxyevent "github.com/strowk/foxy-contexts/pkg/foxy_event"
	"github.com/strowk/foxy-contexts/pkg/foxytest"
	"github.com/strowk/foxy-contexts/pkg/jsonrpc2"
	"github.com/strowk/foxy-contexts/pkg/mcp"
	"github.com/strowk/foxy-contexts/pkg/server"
	"github.com/strowk/foxy-contexts/pkg/stdio"
)

// This package helps to generate mock implementation for MCP server using foxy-contexts library and is based on MCP Story format.
// In order to use this package, you need to define a set of scenarios in one or more YAML files, then provide the path
// to the folder with these files and that then mock server would be generated.
//
// Aside from following MCP Story format, every case within the story must have at least one input that would be further
// referred as anchor input. When mock server starts it builds a map of anchor inputs and prepares list of
// expected inputs/outputs for each anchor input. When the server receives an input, it looks for the matching anchor in the map
// and uses all other inputs in same case as expectations and produces expected outputs.
// Anchor input is always the first input in the case.
// Server always would wait for any expected inputs to be received before producing expected outputs.

// MockServer is a mock implementation of MCP server that is generated from the scenarios in the provided folder.
type MockServer interface {
	Run() error
	Stop(ctx context.Context) error
}

type Expectation interface {
	check(*mockServer)
}

type expectedInput struct {
	value map[string]interface{}
}

func (e *expectedInput) check(srv *mockServer) {
	// TODO: wait till input is received
}

type expectedOutput struct {
	value map[string]interface{}
}

func (e *expectedOutput) check(srv *mockServer) {
	srv.outputsChan <- e.value
}

type mockCase struct {
	name         string
	expectations []Expectation
}

type mockServer struct {
	path      string
	anchorMap map[string]mockCase

	outputsChan chan map[string]interface{}

	stopChan chan struct{}

	responsesChan chan jsonrpc2.JsonRpcResponse
	tp            server.Transport
	logger        foxyevent.Logger
}

func (s *mockServer) GetResponses() chan jsonrpc2.JsonRpcResponse {
	return s.responsesChan
}

func (s *mockServer) Handle(b []byte) {
	var unmarshalled map[string]interface{}
	err := json.Unmarshal(b, &unmarshalled)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(unmarshalled)
	if err != nil {
		panic(err)
	}

	mc := s.anchorMap[string(data)]

	for _, exp := range mc.expectations {
		exp.check(s)
	}
}

func NewMockServer(path string) *mockServer {
	logger := foxyevent.NewSlogLogger(slog.Default())
	logger.UseLogLevel(slog.LevelDebug)

	mock := &mockServer{
		path:          path,
		outputsChan:   make(chan map[string]interface{}),
		responsesChan: make(chan jsonrpc2.JsonRpcResponse),
		stopChan:      make(chan struct{}),
		logger:        logger,
		anchorMap:     make(map[string]mockCase),
	}

	ts, err := foxytest.Read(path)
	if err != nil {
		panic(err)
	}

	for _, t := range ts.GetTests() {
		for _, c := range t.GetTestCases() {
			mc := &mockCase{
				name: c.GetName(),
			}
			anchorInput := c.GetInputs()[0]

			for i, input := range c.GetInputs() {
				if i == 0 {
					continue
				}
				mc.expectations = append(mc.expectations, &expectedInput{value: input})
			}

			for _, o := range c.GetOutputs() {
				mc.expectations = append(mc.expectations, &expectedOutput{value: o})
			}

			marshalledAnchor, err := json.Marshal(anchorInput)
			if err != nil {
				panic(err)
			}

			mock.anchorMap[string(marshalledAnchor)] = *mc
		}
	}

	tp := stdio.NewTransport(stdio.WithNewServerFunc(func(
		capabilities *mcp.ServerCapabilities,
		serverInfo *mcp.Implementation,
		options ...server.ServerOption,
	) server.Server {
		return mock
	},
	))
	mock.tp = tp

	return mock
}

func (s *mockServer) Run() error {
	go func() {
		for {
			select {
			case <-s.stopChan:
				return
			case out := <-s.outputsChan:
				outId := out["id"]
				var id jsonrpc2.RequestId
				if stringId, ok := outId.(string); ok {
					id = jsonrpc2.NewStringRequestId(stringId)
				} else if intId, ok := outId.(int); ok {
					id = jsonrpc2.NewIntRequestId(intId)
				} else {
					log.Printf("fxm: unexpected id type: %v", outId)
					continue
				}
				result := out["result"]
				if objRes, ok := result.(jsonrpc2.Result); ok {
					s.responsesChan <- jsonrpc2.JsonRpcResponse{
						Id:     id,
						Result: &objRes,
					}
				} else if err, ok := result.(*jsonrpc2.Error); ok {
					s.responsesChan <- jsonrpc2.JsonRpcResponse{
						Id:    id,
						Error: err,
					}
				} else {
					log.Printf("fxm: unexpected output registered: %v", out)
				}
			}
		}
	}()
	return s.tp.Run(
		&mcp.ServerCapabilities{},
		&mcp.Implementation{},
	)
}

func (s *mockServer) Stop(ctx context.Context) error {
	close(s.stopChan)
	return s.tp.Shutdown(ctx)
}

func (s *mockServer) GetLogger() foxyevent.Logger {
	return s.logger
}

func (s *mockServer) SetLogger(logger foxyevent.Logger) {
	s.logger = logger
}

// Neither of following methods need to be implemented in the mockServer struct
// , as the server is configured in a different way altogether.

func (s *mockServer) SetNotificationHandler(request jsonrpc2.Request, handler func(req jsonrpc2.Request)) {
	panic("unimplemented")
}

func (s *mockServer) SetRequestHandler(request jsonrpc2.Request, handler func(req jsonrpc2.Request) (jsonrpc2.Result, *jsonrpc2.Error)) {
	panic("unimplemented")
}
