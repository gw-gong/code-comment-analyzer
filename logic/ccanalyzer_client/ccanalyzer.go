package ccanalyzer_client

import (
	"code-comment-analyzer/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CCAnalyzer interface {
	AddUser(string, uint32) (*UserResponse, error)
	Close()
}

type cCAnalyzer struct {
	conn   *grpc.ClientConn
	client CcAnalyzerClient
}

func NewCCAnalyzer(config config.CcAnalyzerConfig) (CCAnalyzer, error) {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := NewCcAnalyzerClient(conn)
	return &cCAnalyzer{
		conn:   conn,
		client: client,
	}, nil
}

func (cca *cCAnalyzer) Close() {
	err := cca.conn.Close()
	if err != nil {
		panic(err)
	}
}

func (cca *cCAnalyzer) AddUser(name string, age uint32) (*UserResponse, error) {
	userRequest := &UserRequest{
		Name: name,
		Age:  age,
	}
	resp, err := cca.client.AddUser(context.Background(), userRequest)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
