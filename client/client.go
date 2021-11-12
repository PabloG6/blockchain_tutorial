package client

import "net/rpc"
type RpcClient struct {
	
}

func NewClient() error {
	return rpc.Register(&RpcClient{})
}


func(rpc *RpcClient) GetTransaction() {
	
}

/**
adds transaction to the mempool
**/
func(rpc *RpcClient) AddTransaction() {

}


/**
	allows validator nodes to make an initial request for funding for a set period of time.
	//todo load cut off date as time from config file
**/

func(rpc *RpcClient) RequestFunding() {
	
}

