package ccanalyzer_client

import (
	"code-comment-analyzer/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CCAnalyzer interface {
	AnalyzeFileContent(language, fileContent string) (analyzedData map[string]interface{}, err error)
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

func (cca *cCAnalyzer) AnalyzeFileContent(language, fileContent string) (analyzedData map[string]interface{}, err error) {
	analyzeFileContentReq := &AnalyzeFileContentReq{
		Language:    language,
		FileContent: fileContent,
	}

	analyzeFileContentRes, err := cca.client.AnalyzeFileContent(context.Background(), analyzeFileContentReq)
	if err != nil {
		return nil, err
	}
	return analyzeFileContentRes.AnalyzedData.AsMap(), nil
}
