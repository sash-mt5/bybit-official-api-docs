### WebSocket Data
```
// websocket server address

//For Testnet
wss://stream-testnet.bybit.com/realtime

//For Mainnet
wss://stream.bybit.com/realtime

```
 
### Rate Limits
 
One single api_key can establish 20 connections simultaneously. Any additional connection after 20 connections will be rejected.
 
### Authentication
 
For public topics, no authentication is required. As for private topics, authentication is required.
 
Currently, there are two ways to authenticate your identity.
 
1. Apply for authentication when establishing a connection.
2. Apply for authentication after establishing a connection through auth request.
 
 
```js
// First way to authenticate
var api_key = "";
var secret = "";
// A UNIX timestamp after which the request become invalid. This is to prevent replay attacks.
// unit:millisecond
var expires = time.now()+1000;

// Signature
var signature = hex(HMAC_SHA256(secret, 'GET/realtime' + expires));
 
// Parameters string
var param = "api_key={api_key}&expires={expires}&signature={signature}";
 
// Establishing connection
var ws = new WebSocket("wsurl?param");
 
// --------------------------------------------------------------------------
 
// Second way to authenticate
var ws = new WebSocket("wsurl")
// Signature is the same as the first way's
ws.send('{"op":"auth","args":["{api_key}",expires,"{signature}"]}');
```

### How to Send The Heartbeat Packet
After establishing the connection,one can send a heartbeat packet to confirm the connection is normal by sending a json request.The specific formats are as follows: 
```js
ws.send('{"op":"ping"}');

// Every ping packet will have a response. The response is sent in the following format:
{
    "success":true,
    "ret_msg":"pong",
    "request":{
        "op":"ping",
        "args":null
    }
}
```
 
### How to Subscribe to a New Topic
 
After establishing the connection, one can subscribe to a new topic by sending a json request. The specific formats are as follows:
```js
ws.send('{"op":"subscribe","args":["topic","topic.filter"]}');
 
// Split the multiple filters by '|' if they belong to the same cluster of topics.
// For example, subscribing to BTCUSD one minute and three minutes kline.
ws.send('{"op":"subscribe","args":["kline.BTCUSD.1m|3m"]}');
 
// Use '*' when subscribing to all data of the same type filter.
// For exmaple, subscribing to all products' interval kline.
ws.send('{"op":"subscribe","args":["kline.*.*"]}')
 
 
// Result of subscribed topic
// Every subscription will have a response. The response is sent in the following format:
{
   "success":true, // Whether subscription is successful
   "ret_msg":"",   // For successful subscription it shows "", otherwise it shows error message
   "request":{     // Request to your subscription
       "op":"subscribe",
       "args":[
           "kline.BTCUSD.1m"
       ]
   }
}
 
```
 
## Currently Supported Topics
 
### Public Topic
* ~~[orderBook25](#orderBook25) `// OrderBook of 25 depth per side`~~  -----It's deprecated.The following V2 version [orderBookL2_25](#orderBook25_v2) is recommended to use
* [kline](#kline) `// Candlestick chart`
* [trade](#trade) `// Real-time trading information`
* [insurance](#insurance) `// Daily insurance fund update`
* ~~[instrument](#instrument) `// Lastet information for symbol`~~  -----It's deprecated. The following v2 version [instrument_info](#instrument_info) is recommended to use
  
### V2 Version System topic
* [orderBookL2_25](#orderBook25_v2) `// OrderBook of 25 depth per side`
* [instrument_info](#instrument_info) `//instrument's infomation`

### Private Topic
* [position](#position) `// Positions of your account`
* [execution](#execution) `// Execution message`
* [order](#order) `// Update for your orders`
 
<hr>
 
### <span id="orderBook25">OrderBook of 25 depth per side</span>
```js
// Send subscription request
ws.send('{"op": "subscribe", "args": ["orderBook25.BTCUSD"]}');
 
// Response content format
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
 
### <span id="kline">Candlestick chart</span>
 
* Currently supported interval
* 1m 3m 5m 15m 30m
* 1h 2h 3h 4h 6h
* 1d 3d
* 1w 2w
* 1M
```js
ws.send('{"op":"subscribe","args":["kline.BTCUSD.1m"]}');
 
// Response content format
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
 
### <span id="trade"> Real-time trading information </span>
 
```js
ws.send('{"op":"subscribe","args":["trade"]}')
 
// Response content format
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
 
### <span id="insurance">Daily insurance fund update</span>
 
```js
ws.send('{"op":"subscribe","args":["insurance"]}')
 
// Response content format
{
    "topic":"insurance.BTC",
    "action":"update",
    "data":{
       "currency":"BTC",
       "timestamp":"2018-10-24T12:00:00.000Z",
       "walletBalance":140224705439 // unit: Satoshi
    }
}
```
 
 
 <hr>
 
### <span id="instrument"> Lastet information for symbol</span>
 
```js
ws.send('{"op":"subscribe","args":["instrument.BTCUSD"]}')

// Response content format
// NOTE: Message belong to "data" class will only be sent to you when it changes. 
// For example, if 'index_price' and 'mark_price' changed and the 'transaction price' didn't change, then only 'symbol', 'index_price' and 'mark_price' will be sent to you, without 'last_price'.
 {
     "topic":"instrument.BTCUSD",
     "data":{
        "symbol": "BTCUSD",
        "mark_price": 5000.5, // mark price
        "index_price": 5000.5, // index price
        "last_price": 5000.5 // latest price
     }
 }
```
 
 <hr>

### <span id="orderBook25_v2">OrderBook of 25 depth per side in V2 version</span>
```js
// orderBookL2_25.BTCUSD
ws.send('{"op": "subscribe", "args": ["orderBookL2_25.BTCUSD"]}');

// Response content format
// NOTE: After subscribe succeed,the first response's type is snapshot , the following responses's type are all delta 

//snapshot type format，the data is ordered by price,from buy to sell
{
     "topic":"orderBookL2_25.BTCUSD",
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

//the delta response includes three types(delete update insert)
//delete :  delete some slots in orderbook where identified by id or price
//update :  modify some slots's size in orderbook where identified by id or price
//insert :  insert new slots in orderbook where identified by id or price
 
//delta type format
{
     "topic":"orderBookL2_25.BTCUSD",
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
The orderbook consists of two lists of the buying direction and the selling direction. The key of the list is the price, and the value of the list is the quantity.

When received the snapshot type package, the orderbook maintained before is cleared, and the calculation is started based on the snapshot package, the new orderbook will be built. In the case that the connection is not disconnected, only the delta type data package will be received. The delta type package contains three types of data ('delete', 'update', 'insert'), each type of data contains a direction, and the modified key is specified according to the direction. ‘delete’ data should be handled firstly,'update' data and 'insert' data should be handled secondly. 

The list of values, 'delete' data indicates that the number of pending orders of special price in the corresponding direction list is 0, 'update' data indicates that the number of pending orders of special price in the corresponding direction list is changed to the latest size, 'insert' data indicates the list in the corresponding direction add the pending order of special price and the quantity is the value of size.

 
<hr>

### <span id="instrument_info">Lastet information for symbol</span>

```js
ws.send('{"op":"subscribe","args":["instrument_info.100ms.BTCUSD"]}')

// Response content format
// NOTE: After subscribe succeed,the first response's type is snapshot , the following responses's type are all delta 
// e4 stands for the result of the data multi 10^4，e6 stands for the result of the data multi 10^6
// snapshot format
{
	"topic": "instrument_info.100ms.BTCUSD",
	"type": "snapshot",
	"data": {
		"id": 1,
		"symbol": "BTCUSD",                     //instrument name
		"last_price_e4": 100000000,             //the latest price
		"last_tick_direction": "ZeroPlusTick",  //the direction of last tick:PlusTick,ZeroPlusTick,MinusTick,ZeroMinusTick
		"prev_price_24h_e4": 100000000,         //the price of prev 24h
		"price_24h_pcnt_e6": 0,                 //the current lastprice percentage change from prev 24h price
		"high_price_24h_e4": 100000000,         //the highest price of prev 24h
		"low_price_24h_e4": 58000000,           //the lowest price of prev 24h
		"prev_price_1h_e4": 71000000,           //the price of prev 1h
		"price_1h_pcnt_e6": 408450,             //the current lastprice percentage change from prev 1h price
		"mark_price_e4": 96758100,              //mark price
		"index_price_e4": 97000000,             //index price
		"open_interest": 158666,                //open interest quantity  Attention,the update is not timimmediately，the slowlest update is 1 minute 
		"open_value_e8": 2004325380,            //open value quantity  Attention,the update is not immediately， the slowlest update is 1 minute 
		"total_turnover_e8": 257108049130,      //total turnover
		"turnover_24h_e8": 8969373218,          //24h turnover
		"total_volume": 15462289,               //total volume
		"volume_24h": 541359,                   //24h volume
		"funding_rate_e6": -3750,               //funding rate
		"predicted_funding_rate_e6": -3750,     //predicted funding rate
		"cross_seq": 7980,                      //sequence
		"created_at": "2018-10-17T11:53:15Z",   
		"updated_at": "2019-07-30T03:12:42Z",
		"next_funding_time": "2019-07-30T08:00:00Z",//next funding time 
		"countdown_hour": 5                     //the rest time to settle funding fee
	},
	"cross_seq": 7980,
	"timestamp_e6": 1564456370126493            //the timestamp of pruduce the infomation of instrument
}
// delta format  
// Only the data-update field has the update data,the data-delete and data-insert is null.If a field not change,the field will not exist in data-update
{
	"topic": "instrument_info.100ms.BTCUSD",
	"type": "delta",
	"data": {
		"delete": [],
		"update": [{
			"id": 1,
			"symbol": "BTCUSD",
			"total_turnover_e8": 257108059130,
			"turnover_24h_e8": 8969383218,
			"total_volume": 15462290,
			"volume_24h": 541360,
			"cross_seq": 7981,
			"created_at": "2018-10-17T11:53:15Z",
			"updated_at": "2019-07-30T03:12:52Z"
		}],
		"insert": []
	},
	"cross_seq": 7981,
	"timestamp_e6": 1564456372227451
}
```
<hr>
 
### <span id="position">Positions of your account</position>
 
```js
ws.send('{"op":"subscribe","args":["position"]}')
 
// Response content format
{
   "topic":"position",
   "action":"update",
   "data":[
       {
           "symbol":"BTCUSD",                  // the contract for this position
           "side":"Sell",                      // side
           "size":11,                          // the current position amount
           "entry_price":6907.291588174717,    // entry price
           "liq_price":7100.234,               // liquidation price
           "bust_price":7088.1234,             // bankruptcy price
           "take_profit":0,                    // take profit price
           "stop_loss":0,                      // stop loss price
           "trailing_stop":0,                  // trailing stop points
           "position_value":0.00159252,        // positional value
           "leverage":1,                       // leverage
           "position_status":"Normal",         // status of position (Normal:normal Liq:in the process of liquidation Adl:in the process of Auto-Deleveraging)
           "auto_add_margin":0,                // Auto margin replenishment enabled (0:no 1:yes)
           "position_seq":14                   // position version number
       }
   ]
}
<hr>
 
### <span id="execution">Execution message</span>
```js
ws.send('{"op":"subscribe","args":["execution"]}')
 
// Response content format
{
    "topic":"execution",
    "data":[
        {
            "symbol":"BTCUSD",
            "side":"Sell",
            "order_id":"xxxxxxxx-xxxx-xxxx-9a8f-4a973eb5c418",
            "exec_id":"xxxxxxxx-xxxx-xxxx-8b66-c3d2fcd352f6",
            "order_link_id":"xxxxxxx",
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
 
### <span id="order">Update for your orders</span>
 
```js
ws.send('{"op":"subscribe","args":["order"]}')
 
// Response content format
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
 
