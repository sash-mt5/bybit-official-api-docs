### Active Order

* [Place active order](#open-apiordercreatepost)

* [Get active order](#open-apiorderlistget)

* [Cancel active order](#open-apiordercancelpost)

### Conditional Order

* [Place conditional order](#open-apistop-ordercreatepost)

* [Get conditional order](#open-apistop-orderlistget)

* [Cancel conditional order](#open-apistop-ordercancelpost)

### Positions

* [User leverage](#userleverageget)

* [Change leverage](#userleveragesavepost)

* [My position](#positionlistget)

* [Change margin](#positionchange-position-marginpost)

### Funding

* [Funding rate](#lastfundingrate)

* [My funding fee](#userlastfundingfee)

* [Predicted funding](#open-apifundingpredicted-fundingget)

### Execution

* [Get the trade records of a order](#open-apiexecutionrecordslistget)

### Market data

* [Get the orderbook](#orderbook-L2)

* [Get latest information for symbol](#latest-information-for-symbol)

-----------
## <span id="open-apiordercreatepost"> Place Active Order </span>
#### API Function

>Parameters of 'side', 'symbol', 'order_type', 'qty', 'price', 'time_in_force' are required for all active orders. Other parameters are optional unless specified.

>Market price active order: A traditional market price order, will be filled at the best available price. 'price' and 'time_in_force' can set to be "" if and only if you are placing market price order.

>Limit price active order: You can set an execution price for your order. Only when last traded price reaches the order price, the system will fill your order.

>Take profit/Stop loss: You may only set a take-profit/stop-loss conditional order upon opening the position. Once you hold a position, the take profit and stop loss information u sent when placing an order will no longer be valid.

>Order quantity: This parameter indicates the quantity of perpetual contracts you want to buy or sell, currently Bybit only support order quantity in an integer.

>Order price: This parameter indicates the price of perpetual contracts you want to buy or sell, currently Bybit only support price increment of every 0.5.

>Customize conditional order ID: You may customize order IDs for active orders. We will link it to the system order ID , and return the unique system order ID to you after the active order is created successfully. You may use this order ID to cancel your active order. The customized order ID is asked to be unique, with a maximum length of 36 characters.

>Notes:
>* Each account can hold up to 200 active orders yet to be filled entirely simultaneously.
>* 'order_status' values explained:
>    * 'Created' indicates the order has been accepted by the system but not yet entered into the orderbook
>    * 'New' indicates the order has entered into the orderbook.

#### URL

> For Testnet:
> [https://api-testnet.bybit.com/open-api/order/create](https://api-testnet.bybit.com/open-api/order/create)

> For Mainnet:
> [https://api.bybit.com/open-api/order/create](https://api.bybit.com/open-api/order/create)

#### HTTP Request Method

> post

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|side |true |string |Side. Valid option: Buy, Sell    |
|symbol |true |string |Contract type. Valid option: BTCUSD, ETHUSD    |
|order_type |true |string |Active order type. Valid option: Limit, Market   |
|qty |true |integer |Order quantity. Maximum quantity of 1 million |
|price |true |integer |Order price. If you held no position, order price has to be more than 10% of the market price and less than 1 million. If you has held any position already, your order price has to be better than liquidation price. The minimum unit of order price's increment or decrement is 0.5.  |
|time_in_force |true |string |Time in force, Valid option: GoodTillCancel, ImmediateOrCancel, FillOrKill,PostOnly |
|take_profit |false |number |take profit price|
|stop_loss |false |number |stop loss price|
|reduce_only |false |bool |reduce only
|close_on_trigger |false |bool |close on trigger
|order_link_id |false |string |Customized order ID, maximum length at 36 characters, and order ID under the same agency has to be unique.|


#### Response example

```js

   {
       'ret_code':0   //Error code - True
       'ret_msg':'ok' //Error message
       'ext_code':''  ,
       'result':{
           'order_id': 'string',              //Unique order ID
           'user_id': 0,                      //User ID
           'symbol': 'string',                //Contract type
           'side': 'string',                  //Side
           'order_type': 'string',            //Order type
           'price': 0,                        //Order price
           'qty': 0,                          //Order quantity
           'time_in_force': 'string',         //Time in force
           'order_status': 'string',          //Order status: Created: order created; Rejected: order rejected; New: order pending; PartiallyFilled: order filled partially; Filled: order filled fully, Cancelled: order cancelled
           'last_exec_time': 0.000000,        //Last execution time
           'last_exec_price': 0,              //Last execution price
           'leaves_qty': 0,                   //Remaining order quantity
           'cum_exec_qty': 0,                 //Accumulated execution quantity
           'cum_exec_value': 0,               //Accumulated execution value
           'cum_exec_fee': 0,                 //Accumulated execution fee
           'reject_reason': 'string',         //Reason for rejection
           'order_link_id': 'string',         //Agency customized order ID
           'created_at':'2018-10-15T04:12:19.000Z',
           'updated_at':'2018-10-15T04:12:19.000Z',
       },
       'time_now':'1539778407.210858',    //UTC timestamp
   }

```

-----------
## <span id="open-apiorderlistget">Get Active Order </span>
#### API Function

> Get my active order list

#### URL

> For Testnet:
> [https://api-testnet.bybit.com/open-api/order/list](https://api-testnet.bybit.com/open-api/order/list)

> For Mainnet:
> [https://api.bybit.com/open-api/order/list](https://api.bybit.com/open-api/order/list)

#### HTTP Request Method

> get

#### Request parameters

|parameters|required|type|comments|
|:----- |:-------|:-----|----- |
|order_id |false |string |Order ID |
|order_link_id |false |string |Customized order ID |
|symbol |false |string | Contract type (BTCUSD, ETHUSD )    |
|sort |false |string |Sorted. Default sort by created time (created_at, cum_exec_qty, qty, last_exec_price, price, cum_exec_value, cum_exec_fee)    |
|order |false |string |Ascending or descending. Default as descending  (desc, asc )   |
|page |false |integer |Page. Default getting first page data |
|limit |false |integer |Limit for data size per page. Default as showing 20 pieces of data per page |
|order_status |false |string | Query your orders for all statuses if 'order_status' is empty. If you want to query orders with specific statuses , you can pass the order_status split by ','.  Available order_status: Created, New, PartiallyFilled, Filled, Cancelled, Rejected

#### Response example

```js

   {
       'ret_code':0,
       'ret_msg':'ok',
       'ext_code':'',
       'result':{
           'data':[
               {
                   'order_id': 'string',       //Unique order ID
                   'user_id': 0,               //User ID
                   'symbol': 'string',         //Contract type
                   'side': 'string',           //Side
                   'order_type': 'string',     //Order type
                   'price': 0,                 //Order price
                   'qty': 0,                   //Order quantity
                   'time_in_force': 'string',  //Time in force
                   'order_status': 'string',   //Order status: Created: order created; Rejected: order rejected; New: order pending; PartiallyFilled: order filled partially; Filled: order filled fully, Cancelled: order cancelled
                   'last_exec_time': 0.000000 , //Last execution time
                   'last_exec_price': 0,       //Last execution price
                   'leaves_qty': 0,            //Remaining order quantity
                   'cum_exec_qty': 0,          //Accumulated execution quantity
                   'cum_exec_value': 0,        //Accumulated execution value
                   'cum_exec_fee': 0,          //Accumulated execution fee
                   'reject_reason': 'string',  //Reason for rejection
                   'order_link_id': 'string',  //Agency customized order ID
                   'created_at':'2018-10-15T04:12:19.000Z',
                   'updated_at':'2018-10-15T04:12:19.000Z',
               }
           ],
           'current_page': 1,
           'total': 1
       },
       'time_now':'1539781050.462841'
   }

```

-----------
## <span id="open-apiordercancelpost">Cancel Active Order</span>
#### API Function

> 'order_id' is required for cancelling active order. The unique 36 characters order ID was returned to you when the active order was created successfully.
> 'symbol' is recommend filled, Otherwise, there will be a small probability of failure.

>You may cancel active order that are unfilled and partially filled. Fully filled order cannot be cancelled.

#### URL

> For Testnet
> [https://api-testnet.bybit.com/open-api/order/cancel](https://api-testnet.bybit.com/open-api/order/cancel)

> For Mainnet
> [https://api.bybit.com/open-api/order/cancel](https://api.bybit.com/open-api/order/cancel)


####  HTTP Request Method

> post

####  Request Parameters

|parameters|required|type|comments|
|:----- |:-------|:-----|----- |
|order_id |true |string |Your active order ID. The unique order ID returned to you when the corresponding active order was created |
|symbol |false |string |Contract type. Valid option: BTCUSD, ETHUSD,EOSUSD,XRPUSD |


#### Response example

```js

   {
       'ret_code':0   //Error code - True
       'ret_msg':'ok' //Error message
       'ext_code':''  ,
       'result':{
           'order_id': 'string',              //Unique order ID
           'user_id': 0,                      //User ID
           'symbol': 'string',                //Contract type
           'side': 'string',                  //Side
           'order_type': 'string',            //Order type
           'price': 0,                        //Order price
           'qty': 0,                          //Order quantity
           'time_in_force': 'string',         //Time in force
           'order_status': 'string',          //Order status: Created: order created; Rejected: order rejected; New: order pending; PartiallyFilled: order filled partially; Filled: order filled fully, Cancelled: order cancelled
           'last_exec_time': 0.000000,        //Last execution time
           'last_exec_price': 0,              //Last execution price
           'leaves_qty': 0,                   //Remaining order quantity
           'cum_exec_qty': 0,                 //Accumulated execution quantity
           'cum_exec_value': 0,               //Accumulated execution value
           'cum_exec_fee': 0,                 //Accumulated execution fee
           'reject_reason': 'string',         //Reason for rejection
           'order_link_id': 'string',         //Agency customized order ID
           'created_at':'2018-10-15T04:12:19.000Z',
           'updated_at':'2018-10-15T04:12:19.000Z',
       },
       'time_now':'1539778407.210858',    //UTC timestamp
   }

```

-----------
## <span id="open-apistop-ordercreatepost"> Place Conditional Order </span>
#### API Function

>Parameters of 'side', 'symbol', 'order_type', 'qty', 'price', 'base_price', 'stop_px', 'time_in_force' are required for all active orders. Other parameters are optional unless specified.

>Market price conditional order: A traditional market price order, will be filled at the best available price. 'price' and 'time_in_force' can set to be "" if and only if you are placing market price order.

>Limit price conditional order: You can set an execution price for your order. Only when last traded price reaches the order price, the system will fill your order.

>Take profit/Stop loss: You may only set a take-profit/stop-loss conditional order upon opening the position. Once you hold a position, the take profit and stop loss information u sent when placing an order will no longer be valid.

>Order quantity: This parameter indicates the quantity of perpetual contracts you want to buy or sell, currently Bybit only support order quantity in an integer.

>Order price: This parameter indicates the price of perpetual contracts you want to buy or sell, currently Bybit only support price increment of every 0.5.

>Conditional order trigger price: You may set a trigger price for your conditional order. conditional order will not enter the order book until the last price hits the trigger price. When last price hits trigger price: 1) your limit conditional order will enter order book, and wait to be executed; 2) your market conditional order will be executed immediately at the best available market price.

>Customize conditional order ID: You may customize order IDs for active orders. We will link it to the system order ID , and return the unique system order ID to you after the active order is created successfully. You may use this order ID to cancel your active order. The customized order ID is asked to be unique, with a maximum length of 36 characters.

>Note: Take profit/Stop loss is not supported in placing conditional orders. One can only use these 2 functions when placing active orders. Moreover, each account can hold up to 10 conditional orders yet to be filled entirely simultaneously.



#### URL
> For Testnet
> [https://api-testnet.bybit.com/open-api/stop-order/create](https://api-testnet.bybit.com/open-api/stop-order/create)

> For Mainnet
> [https://api.bybit.com/open-api/stop-order/create](https://api.bybit.com/open-api/stop-order/create)

#### HTTP Request Method

> post

#### Request Parameters


 |parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|side |true |string |Side. Valid option: Buy, Sell    |
|symbol |true |string |Contract type. Valid option: BTCUSD, ETHUSD    |
|order_type |true |string |Conditional order type. Valid option: Limit, Market |
|qty |true |integer |Order quantity. Maximum quantity of 1 million |
|price| true | integer | Execution price for conditional order|
|base_price |true |integer | Send current market price. It will be used to compare with the value of 'stop_px', to decide whether your conditional order will be triggered by crossing trigger price from upper side or lower side. Mainly used to identify the expected direction of the current conditional order. |
|stop_px | true | integer | Trigger price |
|time_in_force |true |string |Time in force, Valid option: GoodTillCancel, ImmediateOrCancel, FillOrKill,PostOnly |
|close_on_trigger |false |bool |close on trigger
|order_link_id |false |string |Customized order ID, maximum length at 36 characters, and order ID under the same agency has to be unique.|


#### Response example

```js

   {
       'ret_code':0   //Error code - True
       'ret_msg':'ok' //Error message
       'ext_code':''  ,
       'result':{
           'stop_order_id': 'string',         //Unique order ID
           'user_id': 0,                      //User ID
           'stop_order_status': 'string',     //Order status: Untriggered: order waits to be triggered; Triggered: order is triggered; Cancelled: order is cancelled, Active: order is triggered and placed successfully; Rejected: Order is triggered but fail to be placed
           'symbol': 'string',                //Contract type
           'side': 'string',                  //Side
           'order_type': 'string',            //Price type
           'price': 0,                        //Order price
           'qty': 0,                          //Order quantity
           'time_in_force': 'string',         //Time in force
           'stop_order_type': 'string',       //Order type
           'base_price': 0,                   //
           'stop_px': 0,                      //trigger price
           'order_link_id': 'string',         //Agency customized order ID
           'created_at':'2018-10-15T04:12:19.000Z',
           'updated_at':'2018-10-15T04:12:19.000Z',
       },
       'time_now':'1539778407.210858',    //UTC timestamp
   }

```

-----------
## <span id="open-apistop-orderlistget">Get Conditional Order </span>
#### API Function

> Get my conditional order list。

#### URL
> For Testnet
> [https://api-testnet.bybit.com/open-api/stop-order/list](https://api-testnet.bybit.com/open-api/stop-order/list)

> For Mainnet
> [https://api.bybit.com/open-api/stop-order/list](https://api.bybit.com/open-api/stop-order/list)

#### HTTP Request Method

> get

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|stop_order_id |false |string |Order ID of conditional order|
|order_link_id |false |string |Agency customized order ID|
|symbol |false |string |Contract type (BTCUSD ETHUSD )    |
|sort |false |string |Sorted. Default sort by created time (created_at)    |
|order |false |string |Ascending or descending. Default as descending  (desc, asc )   |
|page |false |integer |Page. Default getting first page data |
|limit |false |integer |Limit for data size per page. Default as showing 20 pieces of data per page |


#### Response example

```js

   {
       'ret_code':0,
       'ret_msg':'ok',
       'ext_code':'',
       'result':{
           'data':[
               {
                   'stop_order_id': 'string',         //Unique order ID
                   'user_id': 0,                      //User ID
                   'stop_order_status': 'string',     //Order status: Untriggered: Pending order waits to be triggered; Triggered: order is triggered; Cancelled: order is cancelled, Active: order is triggered and placed successfully; Rejected: Order is triggered but fail to be placed
                   'symbol': 'string',                //Contract type
                   'side': 'string',                  //Side
                   'order_type': 'string',            //Price type
                   'price': 0,                        //Order price
                   'qty': 0,                          //Order quantity
                   'time_in_force': 'string',         //Time in force
                   'stop_order_type': 'string',       //Order type
                   'base_price': 0,                   //
                   'stop_px': 0,                      //trigger price
                   'order_link_id': 'string',         //Agency customized order ID
                   'created_at':'2018-10-15T04:12:19.000Z',
                   'updated_at':'2018-10-15T04:12:19.000Z',
               }
           ],
           'current_page': 1,
           'total': 1
       },
       'time_now':'1539781050.462841'
   }
```

-----------
## <span id="open-apistop-ordercancelpost">Cancel Conditional Order </span>
#### API Function


>'stop_order_id' is required for cancelling conditional order. The unique 36 characters order ID was returned to you when the condional order was created successfully.

>You may cancel all untriggered conditional orders. Essentially, after a conditional order is triggered, it will become an active order. So, when a conditional order is triggered, cancellation has to be done through the active order port for all unfilled or partial filled active order. Similarly, order that has been fully filled cannot be cancelled.

#### URL

> For Testnet
> [https://api-testnet.bybit.com/open-api/stop-order/cancel](https://api-testnet.bybit.com/open-api/stop-order/cancel)

> For Mainnet
> [https://api.bybit.com/open-api/stop-order/cancel](https://api.bybit.com/open-api/stop-order/cancel)

#### HTTP Request Method

> post

#### Request Parameters

|parameter|required|type | comments|
|:----- |:-------|:-----|----- |
|stop_order_id |true |string | Order ID. The unique order ID returned to you when the corresponding order was created. |


#### Response example

```js

   {
       'ret_code':0   //Error code - True
       'ret_msg':'ok' //Error message
       'ext_code':''  ,
       'result':{
           'stop_order_id': 'string',         //Unique order ID
           'user_id': 0,                      //User ID
           'stop_order_status': 'string',     //Order status: Untriggered: order waits to be triggered; Triggered: order is triggered; Cancelled: order is cancelled, Active: order is triggered and placed successfully; Rejected: Order is triggered but fail to be placed
           'symbol': 'string',                //Contract type
           'side': 'string',                  //Side
           'order_type': 'string',            //Price type
           'price': 0,                        //Order price
           'qty': 0,                          //Order quantity
           'time_in_force': 'string',         //Time in force
           'stop_order_type': 'string',       //Order type
           'base_price': 0,                   //
           'stop_px': 0,                      //trigger price
           'order_link_id': 'string',         //Agency customized order ID
           'created_at':'2018-10-15T04:12:19.000Z',
           'updated_at':'2018-10-15T04:12:19.000Z',
       },
       'time_now':'1539778407.210858',    //UTC timestamp
   }

```


-----------
## <span id="userleverageget">User Leverage</span>
#### API Function

> Get user leverage

#### URL

 > For Testnet
> [https://api-testnet.bybit.com/user/leverage](https://api-testnet.bybit.com/user/leverage)

> For Mainnet
> [https://api.bybit.com/user/leverage](https://api.bybit.com/user/leverage)


#### HTTP Request Method

> get

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |


#### Response example

```js

{
    'ret_code': 0,    // return code (0: successful, -1: failed)
    'ret_msg': 'ok',  // error message
    'ext_code': '',   // error code
    'result': {
        'BTCUSD': {
            'leverage': 1
        },
        'EOSUSD': {
            'leverage': 1
        },
        'ETHUSD': {
            'leverage': 1
        },
        'XRPUSD': {
            'leverage': 1
        }
    },
    'ext_info': null,
    'time_now': '1567608910.732004',  // UTC timestamp
    'rate_limit_status': 74
}

```

-----------
## <span id="userleveragesavepost"> Change User Leverage</span>
#### API Function

> Change user leverage

#### URL

> For Testnet
> [https://api-testnet.bybit.com/user/leverage/save](https://api-testnet.bybit.com/user/leverage/save)

> For Mainnet
> [https://api.bybit.com/user/leverage/save](https://api.bybit.com/user/leverage/save)


#### HTTP Request Method

> post

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|symbol |true |string | Contract type (BTCUSD ETHUSD )    |
|leverage |true |string |leverage  |


####  Response example

```js

{
   'ret_code':0   //return code (0: successful, -1: failed)
   'ret_msg':'ok' //error message,
   'ext_code':''  //error code,
   'result': null, //One can decide whether a request is successful depending on ret_code. Will always return 'null'
   'time_now':'1539778407.210858',    //UTC timestamp
}

```

-----------
## <span id="positionlistget"> My Position</span>
#### API Function

> Get my position list

#### URL

> For Testnet
> [https://api-testnet.bybit.com/position/list](https://api-testnet.bybit.com/position/list)

> For Mainnet
> [https://api.bybit.com/position/list](https://api.bybit.com/position/list)

#### HTTP Request Method

> get

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |


#### Response example

```js

{
   'ret_code': 0,    //return code (0: successful, -1: failed)
   'ret_msg': 'ok',  //error message
   'ext_code': '',   //error code
   'result': [
       {
           'id': 1,                //position ID
           'user_id': 1,           //user ID
           'risk_id': 1,           //risk limit ID
           'symbol': 'BTCUSD',     //Contract type (BTCUSD,ETHUSD)
           'side': 'None',         //position Side  (None, buy, sell)
           'size': 0,              //position size
           'position_value': 0,    //position value
           'entry_price': 0,       //entry price
           'leverage': 1,          //user leverage
           'auto_add_margin': 0,   //auto margin replenishment switch
           'position_margin': 0,   //position margin
           'liq_price': 999999,    //liquidation price
           'bust_price': 999999,   //bankruptcy price
           'occ_closing_fee': 0,   //position closing
           'occ_funding_fee': 0,   //funding fee
           'take_profit': 0,       //take profit price
           'stop_loss': 0,         //stop loss price
           'trailing_stop': 0,     //trailing stop point
           'position_status': 'Normal',   //Status Normal(normal), Liq(Liquidation in process), ADL(ADL in process)
           'deleverage_indicator': 1,
           'oc_calc_data': '{\'blq\':\'0\',\'bmp\':\'0\',\'slq\':\'0\',\'smp\':\'0\'}',
           'order_margin': 0,      //Used margin by order
           'wallet_balance': 0,    //wallet balance
           'unrealised_pnl': 0,    //unrealised profit and loss
           'realised_pnl': 0,      //daily realized profit and loss
           'cum_realised_pnl': 0,  //Total realized profit and loss
           'cum_commission': 0,    //Total commissions
           'cross_seq': 0,         //
           'position_seq': 2,      //position sequence number
           'created_at': '2018-10-18T07:15:51.000Z',
           'updated_at': '2018-10-20T13:43:21.000Z'
       }
   ],
   'time_now': '1540043097.995523'      UTC timestamp
}
*    
```

-----------
## <span id="positionchange-position-marginpost"> Change Position Margin</span>
#### API Function

> Update margin

#### URL

> For Testnet
> [https://api-testnet.bybit.com/position/change-position-margin](https://api-testnet.bybit.com/position/change-position-margin)

> For Mainnet
> [https://api.bybit.com/position/change-position-margin](https://api.bybit.com/position/change-position-margin)



#### HTTP Request Method

> post

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|symbol |true |string |Contract type (BTCUSD ETHUSD )    |
|margin |true |string |margin  |


#### Response example

```js

{
   'ret_code':0   //return code(0: successful, -1: failed)
   'ret_msg':'ok' //error message
   'ext_code':'', //error code
  'result': null,  //One can decide whether a request is successful depending on ret_code. Will always return 'null'
   'time_now':'1539778407.210858',    //UTC timestamp
}

```

-----------
## <span id="lastfundingrate"> Get the Last Funding Rate</span>
#### API Function

> Funding rate is generated every 8 hours at 00:00 UTC, 08:00 UTC and 16:00 UTC.
> If it's 12:00 UTC now, what you will get is the funding rate generated at 08:00 UTC.

#### URL

> For Testnet
> [https://api-testnet.bybit.com/open-api/funding/prev-funding-rate](https://api-testnet.bybit.com/open-api/funding/prev-funding-rate)

> For Mainnet
> [https://api.bybit.com/open-api/funding/prev-funding-rate](https://api.bybit.com/open-api/funding/prev-funding-rate)



#### HTTP Request Method

> get

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|symbol |true |string |Contract type (BTCUSD ETHUSD )    |


#### Response example

```js

{
    'ret_code':0   // return code(0：successful、-101：Parameters verification failed)
    'ret_msg':'ok' // error message
    'ext_code':'', // additional error code
    'result': {
        'symbol':'BTCUSD',
        'funding_rate':'0.00375000', // When the Funding Rate is positive, longs pay shorts. When it is negative, shorts pay longs.
        'funding_rate_timestamp':1539950401 //The time of funding rate generated, UTC timestamp
    },
    'time_now':'1539778407.210858',    UTC timestamp
}

```


-----------
## <span id="userlastfundingfee"> Get My Last Funding Fee</span>
#### API Function

> Funding settlement occurs every 8 hours at 00:00 UTC, 08:00 UTC and 16:00 UTC.
> The current intervals's fund fee settlement is based on the previous intervals's fund rate.
> For example, at 16 o'clock, the settlement is based on the fund rate generated at 8 o'clock.
> The fund rate generated at 16 o'clock will be used at 0:00 on the next day.

#### URL

> For Testnet
> [https://api-testnet.bybit.com/open-api/funding/prev-funding](https://api-testnet.bybit.com/open-api/funding/prev-funding)

> For Mainnet
> [https://api.bybit.com/open-api/funding/prev-funding](https://api.bybit.com/open-api/funding/prev-funding)



#### HTTP Request Method

> get

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|symbol |true |string |Contract type (BTCUSD ETHUSD )    |


#### Response example

```js

{
    'ret_code':0   // return code(0：successful、-101：Parameters verification failed)
    'ret_msg':'ok' // error message
    'ext_code':'', // additional error code
    'result': {
        'symbol':'BTCUSD',
        'side': 'Buy', // Your position side at the time of settlement
        'size': 10,   // Your position size at the time of settlement
        'funding_rate':'0.00375000', // Funding rate for settlement. When the Funding rate is positive, longs pay shorts. When it is negative, shorts pay longs.
        'exec_fee': 0.00000116, // Funding fee.  
        'exec_timestamp': 1539950401, // The time of funding settlement occurred, UTC timestamp
    },
    'time_now':'1539778407.210858',    UTC timestamp
    'rate_limit_status':10, // The number of remaining calls for this type of api in the current period(1 min).
}

```

-----------
## <span id="open-apifundingpredicted-fundingget">Get Predicted Funding Rate and Funding Fee</span>
#### API Function

> Get predicted funding rate and funding fee

#### URL

```
> For Testnet
https://api-testnet.bybit.com/open-api/funding/predicted-funding

> For Mainnet
https://api.bybit.com/open-api/funding/predicted-funding
```

#### HTTP Request Method

> get

#### Request Parameters

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|symbol |true |string |Contract type (BTCUSD ETHUSD )    |


#### Response example

```js

{
    'ret_code':0   // return code(0：successful、-101：Parameters verification failed)
    'ret_msg':'ok' // error message
    'ext_code':'', // additional error code
    'result': {
        'predicted_funding_rate': 0.001, // predicted funding rate. When the Funding rate is positive, longs pay shorts. When it is negative, shorts pay longs.
        'predicted_funding_fee': 0.000234 // predicted funding fee
    },
    'time_now':'1539778407.210858',    UTC timestamp
}

```


 -----------
## <span id="open-apiexecutionrecordslistget">Get the trade records of  a  order </span>
#### API Function

> Get the trade records of a order

#### URL

```
> For Testnet
https://api-testnet.bybit.com/v2/private/execution/list

> For Mainnet
https://api.bybit.com/v2/private/execution/list
```

#### HTTP Request Method

> get

#### 请求参数

|parameter|required|type|comments|
|:----- |:-------|:-----|----- |
|order_id |true |string |orderID |

#### 返回示例

```js

{
    'ret_code': 0,                                        // return code
    'ret_msg': 'OK',                                      // error message
    'ext_code': '',                                       // additional error code
    'ext_info': '',                                       // additional error info
    'result': {
        'order_id': 'd854bb13-3fb9-4608-ade4-828f50210778',   // Unique order ID
        'trade_list': [{
            'closed_size': 0,                                // Closed size
            'cross_seq': 3154097,                            // CrossSeq
            'exec_fee': '-0.00000005',                       // Execution fee
            'exec_id': 'b3551383-19b1-4aa6-8ac2-f996bea6e07c', // Unique exec ID
            'exec_price': '4202',                              // Exec Price
            'exec_qty': 1,                                   // Exec Qty
            'exec_time': '1545203567',                       // Exec time
            'exec_type': 'Trade',                            // Exec type
            'exec_value': '0.00023798',                      // Exec value
            'fee_rate': '-0.00025',                          // Fee rate
            'last_liquidity_ind': 'AddedLiquidity',          // AddedLiquidity/RemovedLiquidity
            'leaves_qty': 0,                                 // Leave Qty
            'nth_fill': 7,                                   // Nth Fill
            'order_id': 'd854bb13-3fb9-4608-ade4-828f50210778', // Unique order ID
            'order_price': '4202',                           // Order's price
            'order_qty': 1,                                  // Order's qty
            'order_type': 'Limit',                           // Order's type
            'side': 'Sell',                                  // Side
            'symbol': 'BTCUSD',                              // Symbol
            'user_id': 155446                                // UserID
        }]
    },
    'time_now': '1551340186.761136'
}

```

-------
## <span id="orderbook-L2">Get Orderbook</span>
#### API Function

> Get the orderbook

> Response is in the snapshot format

#### URL

> For Testnet:
> [https://api-testnet.bybit.com/v2/public/orderBook/L2](https://api-testnet.bybit.com/v2/public/orderBook/L2)

> For Mainnet:
> [https://api.bybit.com/v2/public/orderBook/L2](https://api.bybit.com/v2/public/orderBook/L2)

#### HTTP Request Method

> get

#### Request parameters

|parameters|required|type|comments|
|:----- |:-------|:-----|----- |
|symbol |true |string |Valid options: BTCUSD, ETHUSD, EOSUSD, XRPUSD |

#### Response example

```js

{
    "ret_code": 0,                              // return code
    "ret_msg": "OK",                            // error message
    "ext_code": "",                             // additional error code
    "ext_info": "",                             // additional error info
    "result": [
        {
            "symbol": "BTCUSD",                 // symbol
            "price": "9487",                    // price
            "size": 336241,                     // size (in USD contracts)
            "side": "Buy"                       // side
        },
        {
            "symbol": "BTCUSD",                 // symbol
            "price": "9487.5",                  // price
            "size": 522147,                     // size (in USD contracts)
            "side": "Sell"                      // side
        }
    ],
    "time_now": "1567108756.834357"             // UTC timestamp
}

```

-------
## <span id="latest-information-for-symbol">Latest information for symbol</span>
#### API Function

> Get the latest information for symbol

#### URL

> For Testnet:
> [https://api-testnet.bybit.com/v2/public/tickers](https://api-testnet.bybit.com/v2/public/tickers)

> For Mainnet:
> [https://api.bybit.com/v2/public/tickers](https://api.bybit.com/v2/public/tickers)

#### HTTP Request Method

> get

#### Request parameters

|parameters|required|type|comments|
|:----- |:-------|:-----|----- |

#### Response example

```js

// other symbols omitted - the normal response will have additional dictionaries within the "result" class detailing the
// different symbols

{
    "ret_code": 0,                                   // return code
    "ret_msg": "OK",                                 // error message
    "ext_code": "",                                  // additional error code
    "ext_info": "",                                  // additional error info
    "result": [
        {
            "symbol": "BTCUSD",                             // instrument name
            "bid_price": "9499.5",                          // the bid price
            "ask_price": "9500",                            // the ask price
            "last_price": "9499.50",                        // the latest price
            "last_tick_direction": "ZeroMinusTick",         // the direction of last tick:PlusTick,ZeroPlusTick,MinusTick,ZeroMinusTick
            "prev_price_24h": "9659.00",                    // the price of prev 24h
            "price_24h_pcnt": "-0.02",                      // the current last price percentage change from prev 24h price
            "high_price_24h": "9737.00",                    // the highest price of prev 24h
            "low_price_24h": "9322.50",                     // the lowest price of prev 24h
            "prev_price_1h": "9469.00",                     // the price of prev 1h
            "price_1h_pcnt": "0.00",                        // the current lastprice percentage change from prev 1h price
            "mark_price": "9509.24",                        // mark price
            "index_price": "9509.00",                       // index price
            "open_interest": 78923133,                      // open interest quantity  Attention,the update is not timimmediately，the slowlest update is 1 minute
            "open_value": "8078.86",                        // open value quantity  Attention,the update is not immediately， the slowlest update is 1 minute
            "total_turnover": "6540981.84",                 // total turnover
            "turnover_24h": "59236.52",                     // 24h turnover
            "total_volume": 62971273761,                    // total volume
            "volume_24h": 564198332,                        // 24h volume
            "funding_rate": "0.00",                         // funding rate
            "predicted_funding_rate": "-0.00",              // predicted funding rate
            "next_funding_time": "2019-08-30T00:00:00Z",    // next funding time
            "countdown_hour": 4                             // the rest time to settle funding fee
        }
    ],
    "time_now": "1567109419.049271"
}

```
