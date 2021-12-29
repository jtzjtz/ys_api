package rabbitmq

import (
	pool "github.com/jtzjtz/kit/conn/rabbitmq_pool"
)

var RabbitmqConn *pool.Service

func GetRabbitmqConn() {
	RabbitmqConn = new(pool.Service)
	pool.InitAmqp()
}
