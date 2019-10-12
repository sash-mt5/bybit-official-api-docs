## API Signature

### Getting Your API Key and Secret Key

// For Testnet

<a href="https://testnet.bybit.com/user/api-management">https://testnet.bybit.com/user/api-management</a>

 // For Mainnet

<a href="https://www.bybit.com/app/user/api-management">https://www.bybit.com/app/user/api-management</a>

### HTTP Requests
* `GET` requests should be in the `application/x-www-form-urlencoded` format
* `POST` requests should be in the `application/json` format

### Limits

* API Request Rate Limits
    * For order related API, such as *'place active order'*, *'get active order'*, the rate limit for each account is 80 requests per minute.
    * For position related API, such as *'leverage adjustment'*, *'get position'*, the rate limit for each account is 60 requests per minute.

* Order Count Limits
    * Each account can hold up to 200 active orders yet to be filled entirely simultaneously.
    * Each account can hold up to 5 conditional orders in the same side simultaneously.

* How To Raised API Limit Threshold
    * Please send application email to api@bybit.com, We will reply in 3-5 working days.

### Authentication

When calling API, you need to provide your API Key as an identification for every request. Meanwhile, a signature of the request you are making is asked. You need to sign the request data with the corresponding Secret Key.

#### Public Parameters
Name | Description | Type | Required | Default value| Comments
:- | :- | :- | :- | :- | :-
api_key | API_KEY requested from platform | string | yes | -- | For identification.
timestamp | timestamp when initiating request, unit: millisecond | int64 | yes | -- | UTC timestamp, server will verify this parameter when it receives the request. Rule of verification: timestamp < server_time + 1000.
recv_window| valid request timespan, unit: millisecond| int | no | 5000 | An http request will be invalid after this time: timestamp+recv_window. This is to prevent replay attacks.
sign | signature message |  string | yes | no | The signature message which is generated from a certain algorithm.

#### How to Sign
1. Concatenate all the public parameters except 'sign' in the format of http GET, by ascending order of parameters' name. To take *'leverage adjustment'* as example, it has two parameters('symbol' and *leverage*), and the result of concatenation is

``` js
var param_str = 'api_key=B2Rou0PLPpGqcU0Vu2&leverage=100&symbol=BTCUSD&timestamp=1542434791000';
```

2. Sign the parameters string.
```js
var secret = 't7T0YlFnYXk0Fx3JswQsDrViLg1Gh3DUU5Mr';
var sign = hex(HMAC_SHA256($secret, $param_str));
// sign = 670e3e4aa32b243f2dedf1dafcec2fd17a440e71b05681550416507de591d908
```

3. Append the signature at the end of parameters string, and send the http request. Currently, we support the following two types of requesting parameters:

```http
POST /user/leverage/save HTTP/1.1
Host: api-testnet.bybit.com
Content-Type: application/x-www-form-urlencoded

api_key=B2Rou0PLPpGqcU0Vu2&leverage=100&symbol=BTCUSD&timestamp=1542434791000&sign=670e3e4aa32b243f2dedf1dafcec2fd17a440e71b05681550416507de591d908

```

```http
POST /user/leverage/save HTTP/1.1
Host: api-testnet.bybit.com
content-type: application/json

{
    "api_key":"B2Rou0PLPpGqcU0Vu2",
    "leverage":100,
    "symbol":"BTCUSD",
    "timestamp":1542434791000,
    "sign":"670e3e4aa32b243f2dedf1dafcec2fd17a440e71b05681550416507de591d908"
}
```

### Response Messages

Name | Description | Sample value|
:-: | :-: | :-:
ret_code | return code (0：success, otherwise failure) | 0
ret_msg | return message | ok
ext_code | external code error| null
result | refer to each API|
rate_limit_status | Number of remaining calls in current period (1 minute)

 ### Errors

`10004:error sign`

This means the signature you signed is not the same as the server signed.
You may need to visit the `How to sign` above.

<hr>

`10002:invalid request`

```
if (timestamp < (server_time + 1000) && (server_time - timestamp) <= recv_window) {
  // process request
} else {
  // invalid request
  // You may need to check timestamp and recv_window
  // Server time can be get from the HTTP response.
}
```

<hr>

`10007:Login failed`

This means the server didn't find a parameter named api_key in your HTTP request.
