//@File     esclient.go
//@Time     2022/06/16
//@Author   #Suyghur,

package es

import (
	"bytes"
	"context"
	"github.com/bytedance/sonic/encoder"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

type IEsClient interface {
	Insert(index string, data map[string]interface{})
}

type EsClient struct {
	client *elasticsearch.Client
}

func NewEsClient(conf EsConf) *EsClient {
	c := elasticsearch.Config{
		Addresses: conf.Addresses,
		Username:  conf.Username,
		Password:  conf.Password,
	}
	es, err := elasticsearch.NewClient(c)
	if err != nil {
		logx.WithContext(context.Background()).Error(err.Error())
		panic(err.Error())
	}
	return &EsClient{
		client: es,
	}
}

func (es *EsClient) Insert(index string, data interface{}) {
	var buf = bytes.NewBuffer(nil)
	if err := encoder.NewStreamEncoder(buf).Encode(data); err != nil {
		logx.WithContext(context.Background()).Error(err.Error())
	}

	req := esapi.IndexRequest{
		Index:   index,
		Body:    buf,
		Refresh: "true",
	}
	resp, err := req.Do(context.Background(), es.client)
	if err != nil {
		logx.WithContext(context.Background()).Errorf("error getting response: %s", err)
		return
	}
	logx.WithContext(context.Background()).Infof("%v", resp.String())
	defer resp.Body.Close()

	if resp.IsError() {
		logx.WithContext(context.Background()).Errorf("%s error indexing document data: %s", resp.Status(), data)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logx.WithContext(context.Background()).Error(err.Error())
		}
	}(resp.Body)
}
