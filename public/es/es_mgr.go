//@Author   : KaiShin
//@Time     : 2021/10/28

package es_mgr

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/farmerx/elasticsql"
	"github.com/tal-tech/go-zero/core/logx"
	"io"
)

type EsMgr struct {
	Es *elasticsearch.Client
}

type EsMgrInterface interface {
	Insert(index string, data map[string]interface{})
	Query(sql string) []interface{}
}

func New(conf EsConfig) EsMgrInterface {
	var sel = new(EsMgr)

	config := elasticsearch.Config{
		Addresses: conf.Addresses,
		Username:  conf.UserName,
		Password:  conf.Password,
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		logx.Error("[EsMgr.Init] elasticsearch.NewClient failed, err:", err)
		return nil
	}

	sel.Es = es
	logx.Info("[EsMgr.Init], address: ", conf.Addresses)
	return sel
}

func (sel *EsMgr) Insert(index string, data map[string]interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		logx.Errorf("[EsMgr.Insert] err:", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   index,
		Body:    &buf,
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), sel.Es)
	if err != nil {
		logx.Errorf("[EsMgr.Insert] Error getting response: %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		logx.Errorf("[EsMgr.Insert] [%s] Error indexing document data=%s", res.Status(), data)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

}

func (sel *EsMgr) Query(sql string) []interface{} {
	// dsl, index, err := elasticsql.Convert(sql)
	index, dsl, err := elasticsql.NewElasticSQL().SQLConvert(sql)
	if err != nil {
		logx.Errorf("[EsMgr.Query] Convert, err: %s", err)
		return nil
	}
	// logx.Infof("[DEBUG][EsMgr.Query] dsl:", dsl)

	var query map[string]interface{}
	err = json.Unmarshal([]byte(dsl), &query)
	if err != nil {
		logx.Errorf("[EsMgr.Query] json.Unmarshal err: %s", err)
		return nil
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		logx.Errorf("[EsMgr.Query] json.NewEncoder err: %s", err)
		return nil
	}

	////// todo sql query demo
	//q := map[string] interface{}{
	//	"query": sql,
	//}
	//jsonBody, _ := json.Marshal(q)
	//req := esapi.SQLQueryRequest{Body: bytes.NewReader(jsonBody)}
	//res, _ := req.Do(context.Background(), sel.Es)
	//
	//defer res.Body.Close()
	//log.Println(res.String())

	res, err := sel.Es.Search(
		sel.Es.Search.WithContext(context.Background()),
		sel.Es.Search.WithIndex(index),
		sel.Es.Search.WithBody(&buf),
		sel.Es.Search.WithPretty(),
	)
	if err != nil {
		logx.Errorf("[EsMgr.Query] Es.Search err: %s", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logx.Errorf("[EsMgr.Query] err: %s", err)
		return nil
	}
	_, ok := r["error"]
	if ok == true {
		logx.Infof("[EsMgr.Query] es search err, err:%+v", r)
		return nil
	}
	var dataList []interface{}
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, data := range hits {
		dataList = append(dataList, data.(map[string]interface{})["_source"])
	}
	return dataList
}

func (sel *EsMgr) parseConf(confStr string) ([]string, string, string) {
	var mapConfStr map[string]interface{}
	err := json.Unmarshal([]byte(confStr), &mapConfStr)
	if err != nil {
		return nil, "", ""
	}

	address, ok := mapConfStr["Addresses"].([]interface{})
	if ok != true {
		return nil, "", ""
	}
	var addStrList []string
	for _, addr := range address {
		addStrList = append(addStrList, addr.(string))
	}

	username, ok := mapConfStr["Username"].(string)
	if ok != true {
		return nil, "", ""
	}
	password, ok := mapConfStr["Password"].(string)
	if ok != true {
		return nil, "", ""
	}

	return addStrList, username, password
}
