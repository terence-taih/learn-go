package main

import (
	xresilient "training/task10/resilient"
)

func main() {
	//xlog.XCommonLog()
	//xlog.XStructLog()

	// dal.TestDal()

	//xcache.TestLocalCache()
	//xcache.TestLRUCache()

	//xredis.TestRedis()

	//xresilient.TestCircuitBreaker()
	xresilient.TestRetry()

}
