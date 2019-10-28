<hr>

## 2019-10-26

### Websocket API
- [Kline](./websocket.md#kline) [Update]
	- delete `id` in response
	- delete two types `three days`-`3d` `two weeks`-`2w`
	- the Kline data will be more accurate

<hr>

## 2019-10-22

### REST API
- [Query active order (real-time)](./rest_api.md#v2-private-order) [new]

### Websocket API
- [Topic position](./websocket.md#position) [update]
    - Add extra info, eg. wallet_balance
- [Topic trade](./websocket.md#trade) [update]
    - Fix issue of sometimes push same trade multi times
    - Support pushing multi trades in single message
<hr>
