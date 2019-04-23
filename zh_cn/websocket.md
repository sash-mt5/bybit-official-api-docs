### websocket推送
```
// websocket连接地址

// 测试网地址
wss://stream-testnet.bybit.com/realtime

// 主网地址
wss://stream.bybit.com/realtime


// 新版行情

// 测试网地址
wss://ws2-testnet.bybit.com/realtime

// 主网地址
wss://ws2.bybit.com/realtime
```

### 连接限制

一个api_key可同时建立20个连接,超过20个连接后再建立连接会被拒绝

### 身份验证

对于公共类topic不需要进行身份验证即可订阅,个人类topic则需要先进行身份验证

目前有两种方式进行身份验证

1. 在建立连接的请求上附加上身份验证信息进行认证
2. 建立连接后通过auth指令认证

```js
// 第一种认证方式
var api_key = "";
var secret = "";
// http请求的失效时间，防止重放攻击
// 单位:毫秒
var expires = time.now()+1000;

// 消息加签
var signature = hex(HMAC_SHA256(secret, 'GET/realtime' + expires));

// 参数列表
var param = "api_key={api_key}&expires={expires}&signature={signature}";

// 建立连接
var ws = new WebSocket("wsurl?param");

// --------------------------------------------------------------------------

// 第二种认证方式
var ws = new WebSocket("wsurl")
// signature加签方式与第一种方式一致
ws.send('{"op":"auth","args":["{api_key}",expires,"{signature}"]}');
```

### 如何发送业务心跳包
连接建立后，通过发送json格式的心跳包来进行心跳探测,具体格式如下
```js
ws.send('{"op":"ping"}');

// 返应格式如下
{
    "success":true,
    "ret_msg":"pong",
    "request":{
        "op":"ping",
        "args":null
    }
}
```

### 如何订阅topic

连接建立后，通过发送json格式的订阅指令来进行订阅topic,具体格式如下
```js
ws.send('{"op":"subscribe","args":["topic","topic.filter"]}');

// 同一个类型的filter有多个时，以'|'分割
// 如订阅BTCUSD一分钟和三分钟的kline
ws.send('{"op":"subscribe","args":["kline.BTCUSD.1m|3m"]}');

// 订阅同一个类型filter的所有数据时请使用'*'
// 如订阅所有产品的所有interval kline
ws.send('{"op":"subscribe","args":["kline.*.*"]}')


// 订阅topic结果
// 每一个订阅指令都会有一个响应,响应格式如下
{
    "success":true, // 订阅是否成功
    "ret_msg":"",   // 订阅成功时为空，失败时为具体错误信息
    "request":{     // 请求订阅的指令
        "op":"subscribe",
        "args":[
            "kline.BTCUSD.1m"
        ]
    }
}

```

## 目前支持的topic

### 公共类topic
* [orderBook25](#orderBook25) `// 25档orderBook` -----这是过时的，推荐使用下面描述的V2版本
* [kline](#kline) `// K线`
* [trade](#trade) `// 实时交易`
* [insurance](#insurance) `// 每日保险基金更新`
* [instrument](#instrument) `// 产品最新信息`

### 新版行情topic
* [orderBook25](#orderBook25_v2) `// 25档orderBook`

### 个人类topic
* [position](#position) `// 仓位变化`
* [execution](#execution) `// 委托单成交信息`
* [order](#order) `// 委托单的更新`

<hr>

### <span id="orderBook25">订阅25档orderBook</span>
```js
// 发送订阅指令
ws.send('{"op": "subscribe", "args": ["orderBook25.BTCUSD"]}');

// 推送的消息格式
{
    "topic":"orderBook25",
    "action":"snapshot",
    "data":{
        "lastUpdateId":709,
        "symbol":"BTCUSD",
        "bids":[],
        "asks":[
            {"price":6400,"quantity":38},
            {"price":6300,"quantity":5},
        ]
    }
}

```

<hr>

### <span id="kline">k线</span>

* 目前支持的interval
* 1m 3m 5m 15m 30m
* 1h 2h 3h 4h 6h
* 1d 3d
* 1w 2w
* 1M
```js
ws.send('{"op":"subscribe","args":["kline.BTCUSD.1m"]}');

// 推送的消息格式
{
    "topic":"kline.BTCUSD.1m",
    "data":{
        "id":563,
        "symbol":"BTCUSD",
        "open_time":1539918000,
        "open":5900,
        "high":6501,
        "low":6501,
        "close":6501,
        "volume":9,
        "turnover":0.0013844,
        "interval":"1m"
    }
}
```

<hr>

### <span id="trade">实时交易信息</span>

```js
 ws.send('{"op":"subscribe","args":["trade"]}')

 // 推送的消息格式
{
    "topic":"trade.BTCUSD",
    "data":[
        {
            "timestamp":"2019-01-22T15:04:33.461Z",
            "symbol":"BTCUSD",
            "side":"Buy",
            "size":980,
            "price":3563.5,
            "tick_direction":"PlusTick",
            "trade_id":"9d229f26-09a8-42f8-aff3-0ba047b0449d",
            "cross_seq":163261271
        }
    ]
}
```

<hr>

### <span id="insurance">每日保险基金更新</span>

```js
ws.send('{"op":"subscribe","args":["insurance"]}')

// 推送的消息格式
 {
     "topic":"insurance.BTC",
     "action":"update",
     "data":{
        "currency":"BTC",
        "timestamp":"2018-10-24T12:00:00.000Z",
        "wallet_balance":140224705439 // 单位:聪
     }
 }
```

<hr>

### <span id="instrument">产品最新行情</span>

```js
ws.send('{"op":"subscribe","args":["instrument.BTCUSD"]}')

// 推送的消息格式
// NOTE: 这里data下的字段只在其变化时推送，无变化时推送的数据无该字段
// 比如index_price和mark_price变化了，而成交价没有变化，则推送的数据只有symbol、index_price、mark_price而没有last_price
 {
     "topic":"instrument.BTCUSD",
     "data":{
        "symbol": "BTCUSD",
        "mark_price": 5000.5, // 标记价格
        "index_price": 5000.5, // 指数价格
        "last_price": 5000.5 // 最新成交价
     }
 }
```
<hr>


### <span id="orderBook25_v2">订阅新版25档orderBook</span>
```js
// 发送订阅指令 以BTCUSD为例
ws.send('{"op": "subscribe", "args": ["order_book_25L1.BTCUSD"]}');

// 推送的消息格式 
// 当在一条websocket连接上订阅成功后返回第一条消息的type是snapshot类型的
// 后续的消息的type均是delta类型的，在接收到的snapshot消息上进行计算得到最新的orderbook

//snapshot类型消息格式
{
     "topic":"order_book_25L1.BTCUSD",
     "type":"snapshot",
     "data":[
        {
            "price":"2999.00",
            "symbol":"BTCUSD",
            "id":29990000,
            "side":"Buy",
            "size":9
        },
        {
            "price":"3001.00",
            "symbol":"BTCUSD",
            "id":30010000,
            "side":"Sell",
            "size":10
        }
     ],
     "cross_seq":11518,
     "timestamp_e6":1555647164875373
}

//delta类型的消息包含delete update insert三类数据，
//价格档根据price字段或id字段来作为唯一标示
//delta类型的消息包含delete update insert三类数据，
//delete表示某个或多个价格档挂单量变为0
//update表示某个或多个价格档的size更新
//insert表示新增某个或多个价格档的挂挡
 
//delta类型消息格式
{
     "topic":"order_book_25L1.BTCUSD",
     "type":"delta",
     "data":{
          "delete":[
			 {
                   "price":"3001.00",
                   "symbol":"BTCUSD",
                   "id":30010000,
                   "side":"Sell"
             }
          ],
          "update":[
             {
                   "price":"2999.00",
                   "symbol":"BTCUSD",
                   "id":29990000,
                   "side":"Buy",
                   "size":8
             }
          ],
          "insert":[
             {
                   "price":"2998.00",
                   "symbol":"BTCUSD",
                   "id":29980000,
                   "side":"Buy",
                   "size":8
             }
          ],
          "transactTimeE6":0
     },
     "cross_seq":11519,
     "timestamp_e6":1555647221331673
}

```

<hr>

### <span id="position">仓位变化消息</position>

```js
ws.send('{"op":"subscribe","args":["position"]}')

// 推送的消息格式
{
    "topic":"position",
    "action":"update",
    "data":[
        {
            "symbol":"BTCUSD",                  // 产品
            "side":"Sell",                      // 方向
            "size":11,                          // 数量
            "entry_price":6907.291588174717,    // 开仓价
            "liq_price":7100.234,               // 强平价
            "bust_price":7088.1234,             // 破产价
            "take_profit":0,                    // 止盈价格
            "stop_loss":0,                      // 止损价格
            "trailing_stop":0,                  // 追踪止损点数
            "position_value":0.00159252,        // 仓位名义价值
            "leverage":1,                       // 杠杆
            "position_status":"Normal",         // 仓位状态(Normal:正常 Liq:强平中 Adl:被减仓中)
            "auto_add_margin":0,                // 是否自动追加保证金(0:否 1:是)
            "position_seq":14                   // 仓位版本号
        }
    ]
}
```


<hr>

### <span id="execution">成交信息</span>
```js
 ws.send('{"op":"subscribe","args":["execution"]}')

 // 推送的消息格式
{
    "topic":"execution",
    "data":[
        {
            "symbol":"BTCUSD",
            "side":"Sell",
            "order_id":"095f99d2-844c-4358-9a8f-4a973eb5c418",
            "exec_id":"ce6e3ec9-0fb1-4972-8b66-c3d2fcd352f6",
            "order_link_id":"4119012214309",
            "price":3559,
            "exec_qty":1028,
            "exec_fee":-0.00007221,
            "leaves_qty":0,
            "is_maker":true,
            "trade_time":"2019-01-22T14:49:38.000Z"
        },
    ]
}

```

<hr>

### <span id="order">委托更新</span>

```js
 ws.send('{"op":"subscribe","args":["order"]}')

 // 推送的消息格式
{
    "topic":"order",
    "data":[
        {
            "order_id":"xxxxxxxx-xxxx-xxxx-832b-1eca710bf0a6",
            "order_link_id":"xxxxxxxx",
            "symbol":"BTCUSD",
            "side":"Sell",
            "order_type":"Limit",
            "price":3559.5,
            "qty":850,
            "time_in_force":"GoodTillCancel",
            "order_status":"Cancelled",
            "leaves_qty":0,
            "cum_exec_qty":0,
            "cum_exec_value":0,
            "cum_exec_fee":0,
            "timestamp":"2019-01-22T14:49:38.000Z"
        }
    ]
}
```
