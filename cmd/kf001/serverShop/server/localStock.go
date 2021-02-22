package server

import (
	"base/pkg/redis"
)

const LuaScript = `
        local ticket_key = KEYS[1]
        local ticket_total_key = ARGV[1]
        local ticket_sold_key = ARGV[2]
        local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
        local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
		-- 查看是否还有余票,增加订单数量,返回结果值
        if(ticket_total_nums >= ticket_sold_nums) then
            return redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
        end
        return 0
`

//远程商品存储健值 Spike抢占
type RemoteSpikeKeys struct {
	SpikeGoodsIdKey string	//redis中商品hash结构key
	TotalInventoryKey string	//hash结构中商品总库存key
}

//redis扣库存
func (remoteSpikeKeys *RemoteSpikeKeys) RemoteDeductionStock(number int32) bool {

	// Eval(script string, keys []string, args ...interface{}) *Cmd
	//keys :=[]string{"dsafa","fdsa","fdsaaas"}
	//[]string{remoteSpikeKeys.SpikeOrderHashKey,remoteSpikeKeys.TotalInventoryKey,remoteSpikeKeys.TotalInventoryKey
	//redis.Client.Do()
	result, err :=redis.Client.Eval(LuaScript,[]string{remoteSpikeKeys.SpikeGoodsIdKey,remoteSpikeKeys.TotalInventoryKey},number).Result()


	//result, err := redis.Eval(LuaScript, []string{rl.lockKey}, rl.lockValue).Result()
	//result, err := newRedis.Int(lua.Eval( RemoteSpikeKeys.SpikeOrderHashKey, RemoteSpikeKeys.TotalInventoryKey, RemoteSpikeKeys.QuantityOfOrderKey))
	if err != nil {
		return false
	}
	return result != 0
}
