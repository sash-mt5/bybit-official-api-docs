# Errors codes and messages

Example error response:
```javascript
{
  'ret_code': 10001,
  'ret_msg': 'params error!'
}
```

## 认证
### 参数错误
* ```'ret_code': 10001,```
* ```'ret_msg': 'params error!'```
### 请求超时
* ```'ret_code': 10002,```
* ```'ret_msg': 'request timeout'```
### API_KEY错误 或 IP非法
* ```'ret_code': 10003,```
* ```'ret_msg': 'error api_key or illegal ip!'```
### signature错误
* ```'ret_code': 10004,```
* ```'ret_msg': 'error sign!'```
### api-key权限异常
* ```'ret_code': 10005,```
* ```'ret_msg': 'Permission denied!'```
### 请求次数超过限制
* ```'ret_code': 10006,```
* ```'ret_msg': 'Too many visits!'```
### IP不匹配
* ```'ret_code': 10010,```
* ```'ret_msg': 'unmatched IP, request_ip: xx.xxx.xxx.x'```

## 订单
### 订单不存在
* ```'ret_code': 20001,```
* ```'ret_msg': 'Order not exists'```
### 方向为必输字段
* ```'ret_code': 20003,```
* ```'ret_msg': 'side is required'```
### 买卖方向不合法
* ```'ret_code': 20004,```
* ```'ret_msg': 'side is not in the range of selected'```
### 合约为必输项
* ```'ret_code': 20005,```
* ```'ret_msg': 'symbol is required'```
### 合约字段不合法
* ```'ret_code': 20006,```
* ```'ret_msg': 'symbol is not in the range of selected'```
### 订单类型为必输项
* ```'ret_code': 20007,```
* ```'ret_msg': 'order_type is required'```
### 订单类型不合法
* ```'ret_code': 20008,```
* ```'ret_msg': 'order_type is not in the range of selected'```
### 下单数量为必输项
* ```'ret_code': 20009,```
* ```'ret_msg': 'qty is required'```
### 下单数量应该大于0
* ```'ret_code': 20010,```
* ```'ret_msg': 'qty need greater than zero'```
### 下单数量不合法
* ```'ret_code': 20011,```
* ```'ret_msg': 'qty is need integer'```
### 下单数量应是1-1000000之间的值
* ```'ret_code': 20012,```
* ```'ret_msg': 'qty need between 1 - 1000000'```
### 价格为必输项
* ```'ret_code': 20013,```
* ```'ret_msg': 'price is required'```
### 价格不合法
* ```'ret_code': 20014,```
* ```'ret_msg': 'price need greater than zero'```
### 成交类型是必输项
* ```'ret_code': 20015,```
* ```'ret_msg': 'time_in_force is required'```
### 成交类型输入不合法
* ```'ret_code': 20016,```
* ```'ret_msg': 'time_in_force is not in the range of selected'```
### ORDER_ID为必输项
* ```'ret_code': 20017,```
* ```'ret_msg': 'order_id is required'```
### 日期格式不合法
* ```'ret_code': 20018,```
* ```'ret_msg': 'The date format is invalid'```
### STOP_PX为必输项
* ```'ret_code': 20019,```
* ```'ret_msg': 'stop_px is required'```
### BASE_PRICE为必输项
* ```'ret_code': 20020,```
* ```'ret_msg': 'base_price is required'```
### 条件单ID为必输项
* ```'ret_code': 20021,```
* ```'ret_msg': 'stop_order_id is required'```
### ORDER_LINK_ID不唯一
* ```'ret_code': 30001,```
* ```'ret_msg': 'order_link_id not unique'```
### 已存在平仓订单
* ```'ret_code': 30002,```
* ```'ret_msg': 'Already hava active exit order'```
### 订单数量超过上限
* ```'ret_code': 30003,```
* ```'ret_msg': 'The number of contracts exceeds maximum limit allowed'```
### 订单数量超过上限
* ```'ret_code': 30004,```
* ```'ret_msg': 'The number of contracts exceeds maximum limit allowed: too large'```
### 订单价格范围不合法
* ```'ret_code': 30005,```
* ```'ret_msg': 'Order price is out of permissible range'```
### last_price为必输字段
* ```'ret_code': 30006,```
* ```'ret_msg': 'no last_price'```
### 订单价格范围不合法
* ```'ret_code': 30007,```
* ```'ret_msg': 'Order price is out of permissible range'```
### 订单类型不合法
* ```'ret_code': 30008,```
* ```'ret_msg': 'order_type invalid'```
### 用户没有对应的持仓
* ```'ret_code': 30009,```
* ```'ret_msg': 'no position for uid, symbol'```
### 钱包余额不足
* ```'ret_code': 30010,```
* ```'ret_msg': 'Insufficient wallet balance'```
### 由于仓位清算，不允许操作
* ```'ret_code': 30011,```
* ```'ret_msg': 'Operation not allowed as position is undergoing liquidation'```
### 由于ADL，不允许操作
* ```'ret_code': 30012,```
* ```'ret_msg': 'Operation not allowed as position is undergoing ADL'```
### 持仓处于其他状态
* ```'ret_code': 30013,```
* ```'ret_msg': 'position is in other status'```
### 下单数量应和持仓数量一致
* ```'ret_code': 30014,```
* ```'ret_msg': 'Invalid exit order: qty != size'```
### 平仓订单不合法，应和持仓方向相反
* ```'ret_code': 30015,```
* ```'ret_msg': 'Invalid exit order: side should be opposite'```
### 请先取消止盈止损单
* ```'ret_code': 30016,```
* ```'ret_msg': 'Please cancel Take Profit and Stop Loss first'```
### 不能低于强平价格
* ```'ret_code': 30017,```
* ```'ret_msg': 'estimated fill price:[%s] cannot be lower than current Buy liq_price:[%s]'```
### 不能高于卖方向强平价格
* ```'ret_code': 30018,```
* ```'ret_msg': 'estimated fill price:[%s] cannot be higher than current Sell liq_price:{$position->liq_price}'```
### 不能将TP/SL参数用于非0仓位
* ```'ret_code': 30019,```
* ```'ret_msg': 'cannot attach TP/SL params for none-zero position'```
### 持仓已有TP/SL设置
* ```'ret_code': 30020,```
* ```'ret_msg': 'position already have TP/SL params'```
### 预计保证金不足
* ```'ret_code': 30021,```
* ```'ret_msg': 'cannot afford estimated position_margin:[%s]'```
### 预计买方向强平价格不能高于标记价格
* ```'ret_code': 30022,```
* ```'ret_msg': 'estimated Buy liq_price:[%s] cannot be higher than current mark_price:[%s]'```
### 预计卖方向强平价不能低于标记价格
* ```'ret_code': 30023,```
* ```'ret_msg': 'estimated Sell liq_price:[%s] cannot be lower than current mark_price:[%s]'```
### 空仓无法设置
* ```'ret_code': 30024,```
* ```'ret_msg': 'cannot set [%s] for zero position'```
### 应该比标记价格高10%且不大于1000000
* ```'ret_code': 30025,```
* ```'ret_msg': 'bigger  then 10% mark price and less then 1 million'```
### 价格过高
* ```'ret_code': 30026,```
* ```'ret_msg': 'price too high'```
### 止盈价应高于上笔成交价
* ```'ret_code': 30027,```
* ```'ret_msg': 'Price set for Take profit should be higher than Last Traded Price'```
### 止损价应在强平价和上笔成交价之间
* ```'ret_code': 30028,```
* ```'ret_msg': 'Price set for Stop loss should be between Liquidation price and Last Traded Price'```
### 止损价格应在上笔成交价和清算价格之间
* ```'ret_code': 30029,```
* ```'ret_msg': 'Price set for Stop loss should be between Last Traded Price and Liquidation price'```
### 设定的止盈价格应低于上笔成交价
* ```'ret_code': 30030,```
* ```'ret_msg': 'Price set for Take profit should be lower than Last Traded Price'```
### 可用余额不足以支付订单费用
* ```'ret_code': 30031,```
* ```'ret_msg': 'Insufficient available balance:[%s] for order cost:[%s]'```
### 订单已成交或已取消
* ```'ret_code': 30032,```
* ```'ret_msg': 'Order has been finished or canceled'```
### 条件单的数目超过所容许的最高限额
* ```'ret_code': 30033,```
* ```'ret_msg': 'The number of stop orders (%s) exceeds maximum limit allowed'```
### 条件单不存在
* ```'ret_code': 30034,```
* ```'ret_msg': 'no stop order for [%s]'```
### 频率过快，请稍后尝试
* ```'ret_code': 30035,```
* ```'ret_msg': 'Too fast to cancel, Try it later'```
### 订单执行后的预期仓位值超过当前风险限额
* ```'ret_code': 30036,```
* ```'ret_msg': 'The expected position value after order execution exceeds the current risk limit'```
### 订单已取消
* ```'ret_code': 30037,```
* ```'ret_msg': 'Order already cancelled'```
### 无相应订单记录
* ```'ret_code': 30041,```
* ```'ret_msg': 'No position record foruser_id, symbol'```
### 钱包余额不足
* ```'ret_code': 30042,```
* ```'ret_msg': 'Insufficient wallet balance'```
### 由于仓位清算，不允许操作
* ```'ret_code': 30043,```
* ```'ret_msg': 'Operation not allowed as position is undergoing liquidation'```
### 不允许操作，因为位置正在进行ADL
* ```'ret_code': 30044,```
* ```'ret_msg': 'Operation not allowed as position is undergoing ADL'```
### 持仓处于其他状态
* ```'ret_code': 30045,```
* ```'ret_msg': 'position is in other status'```
### 平仓合约数量超过风险限额，请调整您的风险限额水平再试
* ```'ret_code': 30057,```
* ```'ret_msg': 'Requested quantity of contracts exceeds risk limit, please adjust your risk limit level before trying again'```
### 不满足只减仓的条件
* ```'ret_code': 30063,```
* ```'ret_msg': 'reduce-only rule not satisfied'```

## Position
### 订单价格超出了允许的范围
* ```'ret_code': 30005,```
* ```'ret_msg': 'Order price is out of permissible range'```
### last_price为必输字段
* ```'ret_code': 30006,```
* ```'ret_msg': 'no last_price'```
### leverag为必输字段
* ```'ret_code': 20022,```
* ```'ret_msg': 'leverage is required'```
### leverage必须是一个数字
* ```'ret_code': 20023,```
* ```'ret_msg': 'leverage need numeric'```
### leverage必须大于0
* ```'ret_code': 20031,```
* ```'ret_msg': 'leverage need greater than zero'```
### margin为必输字段
* ```'ret_code': 20070,```
* ```'ret_msg': 'margin is required'```
### margin应大于0
* ```'ret_code': 20071,```
* ```'ret_msg': 'margin should greater than zero'```
### 任何调整都将立即引发清算
* ```'ret_code': 30040,```
* ```'ret_msg': 'Any adjustments made will trigger immediate liquidation'```
### 可用余额不足
* ```'ret_code': 30049,```
* ```'ret_msg': 'Available balance not enough!'```
### 任何调整都将立即引发清算
* ```'ret_code': 30050,```
* ```'ret_msg': 'Any adjustments made will trigger immediate liquidation'```
### 杠杆设置非法，因为它将超出您的风险限额
* ```'ret_code': 30051,```
* ```'ret_msg': 'due to risk limit, cannot set leverage to [%s]'```
### 杠杆设置非法
* ```'ret_code': 30052,```
* ```'ret_msg': 'cannot set leverage to [value]'```
### 保证金设置不合法
* ```'ret_code': 30054,```
* ```'ret_msg': 'fixed_new_position_margin[%s] invalid'```
### 可用余额不足
* ```'ret_code': 30067,```
* ```'ret_msg': 'Available balance not enough'```
### exitValue必须大于0
* ```'ret_code': 30068,```
* ```'ret_msg': 'exitValue must positive'```
