# Errors codes and messages

Example error response:
```javascript
{
  'ret_code': 10001,
  'ret_msg': 'params error!'
}
```

## Authentication
### ERROR IN PARAMS
* ```'ret_code': 10001,```
* ```'ret_msg': 'params error!'```
### REQUEST TIMEOUT
* ```'ret_code': 10002,```
* ```'ret_msg': 'request timeout'```
### INVALID API_KEY OR ILLEGAL IP
* ```'ret_code': 10003,```
* ```'ret_msg': 'error api_key or illegal ip!'```
### SIGN ERROR
* ```'ret_code': 10004,```
* ```'ret_msg': 'error sign!'```
### PERMISSION DENIED!
* ```'ret_code': 10005,```
* ```'ret_msg': 'Permission denied!'```
### REQUEST LIMIT EXCEEDED
* ```'ret_code': 10006,```
* ```'ret_msg': 'Too many visits!'```
### UNMATCHED IP
* ```'ret_code': 10010,```
* ```'ret_msg': 'unmatched IP, request_ip: xx.xxx.xxx.x'```

## Order
### ORDER DOES NOT EXIST
* ```'ret_code': 20001,```
* ```'ret_msg': 'Order not exists'```
### SIDE IS REQUIRED
* ```'ret_code': 20003,```
* ```'ret_msg': 'side is required'```
### SIDE IS NOT IN THE RANGE OF SELECTED
* ```'ret_code': 20004,```
* ```'ret_msg': 'side is not in the range of selected'```
### SYMBOL IS REQUIRED
* ```'ret_code': 20005,```
* ```'ret_msg': 'symbol is required'```
### SYMBOL IS NOT IN THE RANGE OF SELECTED
* ```'ret_code': 20006,```
* ```'ret_msg': 'symbol is not in the range of selected'```
### ORDER_TYPE IS REQUIRED
* ```'ret_code': 20007,```
* ```'ret_msg': 'order_type is required'```
### ORDER_TYPE IS NOT IN THE RANGE OF SELECTED
* ```'ret_code': 20008,```
* ```'ret_msg': 'order_type is not in the range of selected'```
### QTY IS REQUIRED
* ```'ret_code': 20009,```
* ```'ret_msg': 'qty is required'```
### QTY MUST BE GREATER THAN ZERO
* ```'ret_code': 20010,```
* ```'ret_msg': 'qty need greater than zero'```
### QTY MUST BE AN INTEGER
* ```'ret_code': 20011,```
* ```'ret_msg': 'qty is need integer'```
### QTY MUST BE BETWEEN 1 AND 1 000 000
* ```'ret_code': 20012,```
* ```'ret_msg': 'qty need between 1 - 1000000'```
### PRICE IS REQUIRED
* ```'ret_code': 20013,```
* ```'ret_msg': 'price is required'```
### PRICE MUST BE GREATER THAN ZERO
* ```'ret_code': 20014,```
* ```'ret_msg': 'price need greater than zero'```
### TIME_IN_FORCE IS REQUIRED
* ```'ret_code': 20015,```
* ```'ret_msg': 'time_in_force is required'```
### TIME_IN_FORCE IS NOT IN THE RANGE OF SELECTED
* ```'ret_code': 20016,```
* ```'ret_msg': 'time_in_force is not in the range of selected'```
### ORDER_ID IS REQUIRED
* ```'ret_code': 20017,```
* ```'ret_msg': 'order_id is required'```
### THE DATE FORMAT IS INVALID
* ```'ret_code': 20018,```
* ```'ret_msg': 'The date format is invalid'```
### STOP_PX IS REQUIRED
* ```'ret_code': 20019,```
* ```'ret_msg': 'stop_px is required'```
### BASE_PRICE IS REQUIRED
* ```'ret_code': 20020,```
* ```'ret_msg': 'base_price is required'```
### STOP_ORDER_ID IS REQUIRED
* ```'ret_code': 20021,```
* ```'ret_msg': 'stop_order_id is required'```
### ORDER_LINK_ID NOT UNIQUE
* ```'ret_code': 30001,```
* ```'ret_msg': 'order_link_id not unique'```
### ALREADY HAVE ACTIVE EXIT ORDER
* ```'ret_code': 30002,```
* ```'ret_msg': 'Already have active exit order'```
### THE NUMBER OF CONTRACTS EXCEEDS MAXIMUM LIMIT ALLOWED
* ```'ret_code': 30003,```
* ```'ret_msg': 'The number of contracts exceeds maximum limit allowed'```
### THE NUMBER OF CONTRACTS EXCEEDS MAXIMUM LIMIT ALLOWED: TOO LARGE
* ```'ret_code': 30004,```
* ```'ret_msg': 'The number of contracts exceeds maximum limit allowed: too large'```
### ORDER PRICE IS OUT OF PERMISSIBLE RANGE
* ```'ret_code': 30005,```
* ```'ret_msg': 'Order price is out of permissible range'```
### LAST_PRICE IS REQUIRED
* ```'ret_code': 30006,```
* ```'ret_msg': 'no last_price'```
### ORDER PRICE IS OUT OF PERMISSIBLE RANGE
* ```'ret_code': 30007,```
* ```'ret_msg': 'Order price is out of permissible range'```
### ORDER_TYPE INVALID
* ```'ret_code': 30008,```
* ```'ret_msg': 'order_type invalid'```
### NO POSITION FOR UID, SYMBOL
* ```'ret_code': 30009,```
* ```'ret_msg': 'no position for uid, symbol'```
### INSUFFICIENT WALLET BALANCE
* ```'ret_code': 30010,```
* ```'ret_msg': 'Insufficient wallet balance'```
### OPERATION NOT ALLOWED AS POSITION IS UNDERGOING LIQUIDATION
* ```'ret_code': 30011,```
* ```'ret_msg': 'Operation not allowed as position is undergoing liquidation'```
### OPERATION NOT ALLOWED AS POSITION IS UNDERGOING ADL
* ```'ret_code': 30012,```
* ```'ret_msg': 'Operation not allowed as position is undergoing ADL'```
### POSITION IS IN OTHER STATUS
* ```'ret_code': 30013,```
* ```'ret_msg': 'position is in other status'```
### INVALID EXIT ORDER: QTY SHOULD MATCH POSITION SIZE
* ```'ret_code': 30014,```
* ```'ret_msg': 'Invalid exit order: qty != size'```
### INVALID EXIT ORDER: SIDE SHOULD BE OPPOSITE OF POSITION
* ```'ret_code': 30015,```
* ```'ret_msg': 'Invalid exit order: side should be opposite'```
### PLEASE CANCEL TAKE PROFIT AND STOP LOSS FIRST
* ```'ret_code': 30016,```
* ```'ret_msg': 'Please cancel Take Profit and Stop Loss first'```
### ESTIMATED FILL PRICE:[%s] CANNOT BE LOWER THAN CURRENT BUY LIQ_PRICE:[%s]
* ```'ret_code': 30017,```
* ```'ret_msg': 'estimated fill price:[%s] cannot be lower than current Buy liq_price:[%s]'```
### ESTIMATED FILL PRICE:[%s] CANNOT BE HIGHER THAN CURRENT SELL LIQ_PRICE:{$POSITION->LIQ_PRICE}
* ```'ret_code': 30018,```
* ```'ret_msg': 'estimated fill price:[%s] cannot be higher than current Sell liq_price:{$position->liq_price}'```
### CANNOT ATTACH TP/SL PARAMS TO NON-ZERO POSITION
* ```'ret_code': 30019,```
* ```'ret_msg': 'cannot attach TP/SL params for none-zero position'```
### POSITION ALREADY HAS TP/SL PARAMS
* ```'ret_code': 30020,```
* ```'ret_msg': 'position already have TP/SL params'```
### CANNOT AFFORD ESTIMATED POSITION_MARGIN:[%s]
* ```'ret_code': 30021,```
* ```'ret_msg': 'cannot afford estimated position_margin:[%s]'```
### ESTIMATED BUY LIQ_PRICE:[%s] CANNOT BE HIGHER THAN CURRENT MARK_PRICE:[%s]
* ```'ret_code': 30022,```
* ```'ret_msg': 'estimated Buy liq_price:[%s] cannot be higher than current mark_price:[%s]'```
### ESTIMATED SELL LIQ_PRICE:[%s] CANNOT BE LOWER THAN CURRENT MARK_PRICE:[%s]
* ```'ret_code': 30023,```
* ```'ret_msg': 'estimated Sell liq_price:[%s] cannot be lower than current mark_price:[%s]'```
### CANNOT SET [%s] FOR ZERO POSITION
* ```'ret_code': 30024,```
* ```'ret_msg': 'cannot set [%s] for zero position'```
### SHOULD BE BIGGER THAN 10% OF MARK PRICE AND LESS THAN 1000000
* ```'ret_code': 30025,```
* ```'ret_msg': 'bigger  then 10% mark price and less then 1 million'```
### PRICE IS TOO HIGH
* ```'ret_code': 30026,```
* ```'ret_msg': 'price too high'```
### PRICE SET FOR TAKE PROFIT SHOULD BE HIGHER THAN LAST TRADED PRICE
* ```'ret_code': 30027,```
* ```'ret_msg': 'Price set for Take profit should be higher than Last Traded Price'```
### PRICE SET FOR STOP LOSS SHOULD BE BETWEEN LIQUIDATION PRICE AND LAST TRADED PRICE
* ```'ret_code': 30028,```
* ```'ret_msg': 'Price set for Stop loss should be between Liquidation price and Last Traded Price'```
### PRICE SET FOR STOP LOSS SHOULD BE BETWEEN LAST TRADED PRICE AND LIQUIDATION PRICE
* ```'ret_code': 30029,```
* ```'ret_msg': 'Price set for Stop loss should be between Last Traded Price and Liquidation price'```
### PRICE SET FOR TAKE PROFIT SHOULD BE LOWER THAN LAST TRADED PRICE
* ```'ret_code': 30030,```
* ```'ret_msg': 'Price set for Take profit should be lower than Last Traded Price'```
### INSUFFICIENT AVAILABLE BALANCE:[%s] FOR ORDER COST:[%s]
* ```'ret_code': 30031,```
* ```'ret_msg': 'Insufficient available balance:[%s] for order cost:[%s]'```
### ORDER HAS BEEN COMPLETED OR CANCELLED
* ```'ret_code': 30032,```
* ```'ret_msg': 'Order has been finished or canceled'```
### THE NUMBER OF STOP ORDERS (%s) EXCEEDS MAXIMUM LIMIT ALLOWED
* ```'ret_code': 30033,```
* ```'ret_msg': 'The number of stop orders (%s) exceeds maximum limit allowed'```
### NO STOP ORDER FOR [%s]
* ```'ret_code': 30034,```
* ```'ret_msg': 'no stop order for [%s]'```
### TOO FAST TO CANCEL, TRY IT LATER
* ```'ret_code': 30035,```
* ```'ret_msg': 'Too fast to cancel, Try it later'```
### THE EXPECTED POSITION VALUE AFTER ORDER EXECUTION EXCEEDS THE CURRENT RISK LIMIT
* ```'ret_code': 30036,```
* ```'ret_msg': 'The expected position value after order execution exceeds the current risk limit'```
### ORDER ALREADY CANCELLED
* ```'ret_code': 30037,```
* ```'ret_msg': 'Order already cancelled'```
### NO POSITION RECORD FOR USER_ID, SYMBOL
* ```'ret_code': 30041,```
* ```'ret_msg': 'No position record foruser_id, symbol'```
### INSUFFICIENT WALLET BALANCE
* ```'ret_code': 30042,```
* ```'ret_msg': 'Insufficient wallet balance'```
### OPERATION NOT ALLOWED AS POSITION IS UNDERGOING LIQUIDATION
* ```'ret_code': 30043,```
* ```'ret_msg': 'Operation not allowed as position is undergoing liquidation'```
### OPERATION NOT ALLOWED AS POSITION IS UNDERGOING ADL
* ```'ret_code': 30044,```
* ```'ret_msg': 'Operation not allowed as position is undergoing ADL'```
### POSITION IS IN OTHER STATUS
* ```'ret_code': 30045,```
* ```'ret_msg': 'position is in other status'```
### REQUESTED QUANTITY OF CONTRACTS EXCEEDS RISK LIMIT, PLEASE ADJUST YOUR RISK LIMIT LEVEL BEFORE TRYING AGAIN
* ```'ret_code': 30057,```
* ```'ret_msg': 'Requested quantity of contracts exceeds risk limit, please adjust your risk limit level before trying again'```
### REDUCE-ONLY RULE NOT SATISFIED
* ```'ret_code': 30063,```
* ```'ret_msg': 'reduce-only rule not satisfied'```

## Position
### ORDER PRICE IS OUT OF PERMISSIBLE RANGE
* ```'ret_code': 30005,```
* ```'ret_msg': 'Order price is out of permissible range'```
### LAST_PRICE IS REQUIRED
* ```'ret_code': 30006,```
* ```'ret_msg': 'no last_price'```
### LEVERAGE IS REQUIRED
* ```'ret_code': 20022,```
* ```'ret_msg': 'leverage is required'```
### LEVERAGE MUST BE NUMERIC
* ```'ret_code': 20023,```
* ```'ret_msg': 'leverage need numeric'```
### LEVERAGE MUST BE GREATER THAN ZERO
* ```'ret_code': 20031,```
* ```'ret_msg': 'leverage need greater than zero'```
### MARGIN IS REQUIRED
* ```'ret_code': 20070,```
* ```'ret_msg': 'margin is required'```
### MARGIN SHOULD BE GREATER THAN ZERO
* ```'ret_code': 20071,```
* ```'ret_msg': 'margin should greater than zero'```
### ANY ADJUSTMENTS MADE WILL TRIGGER IMMEDIATE LIQUIDATION
* ```'ret_code': 30040,```
* ```'ret_msg': 'Any adjustments made will trigger immediate liquidation'```
### INSUFFICIENT AVAILABLE BALANCE
* ```'ret_code': 30049,```
* ```'ret_msg': 'Available balance not enough!'```
### ANY ADJUSTMENTS MADE WILL TRIGGER IMMEDIATE LIQUIDATION
* ```'ret_code': 30050,```
* ```'ret_msg': 'Any adjustments made will trigger immediate liquidation'```
### CANNOT SET LEVERAGE TO [%s] AS IT WOULD EXCEED YOUR RISK_LIMIT
* ```'ret_code': 30051,```
* ```'ret_msg': 'due to risk limit, cannot set leverage to [%s]'```
### CANNOT SET LEVERAGE TO [value]
* ```'ret_code': 30052,```
* ```'ret_msg': 'cannot set leverage to [value]'```
### FIXED_NEW_POSITION_MARGIN[%s] INVALID
* ```'ret_code': 30054,```
* ```'ret_msg': 'fixed_new_position_margin[%s] invalid'```
### INSUFFICIENT AVAILABLE BALANCE
* ```'ret_code': 30067,```
* ```'ret_msg': 'Available balance not enough'```
### EXIT VALUE MUST BE POSITIVE
* ```'ret_code': 30068,```
* ```'ret_msg': 'exitValue must positive'```
