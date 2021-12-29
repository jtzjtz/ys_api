package es

import (
	"github.com/jtzjtz/kit/conn/es_pool"
	"github.com/olivere/elastic"
)

func GetEsConn() *elastic.Client {
	opts := &es_pool.Options{
		URL:         "",
		Index:       "-1",
		Username:    "",
		Password:    "",
		Shards:      0,
		Replicas:    0,
		Sniff:       nil,
		Healthcheck: nil,
		Infolog:     "",
		Errorlog:    "",
		Tracelog:    "",
	}

	return es_pool.NewPool(opts).GetClient()
}
