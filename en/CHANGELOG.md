<hr>

## 2019-11-04

### REST API
- [Cancel order](./rest_api.md##open-apiordercancelpost) [update]
    - Support cancel order by `order_link_id`
- [Cancel stop-order](./rest_api.md##open-apiordercancelpost) [update]
    - Support cancel stop-order by `order_link_id`
- [Get user API key info](./rest_api.md#open-apikeyget) [update]
    - Add extra info
    - Update ips field to return content
- [Update rest api rate limit](./rest_api_sign.md#rest-rate-limit)[update]
	- The rate limit is accurate to milliseconds
	- Refine the rate limit of the endpoints
	- add new response fields: rate_limit_reset_ms、rate_limit
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
