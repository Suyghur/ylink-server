//@File     extdata.go
//@Time     2022/05/17
//@Author   #Suyghur,

package kafka

import "github.com/Shopify/sarama"

func GetTraceFromHeader(headers []*sarama.RecordHeader) string {
	for _, h := range headers {
		if string(h.Key) == "trace_id" {
			return string(h.Value)
		}
	}
	return ""
}
