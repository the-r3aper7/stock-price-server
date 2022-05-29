# Stock Price Server 

## Request & Response Examples

### API Resources

  - GET /[ticker]
  - GET /m/[tickers]

### GET /[ticker]

Example: http://example.com/[ticker]

Response body:

    {
        "data": {
            "symbol": ticker,
            "currency": "USD",
            "price": 169.69,
            "change": 169.69,
            "perChange": 169.69
        },
        "errDescription": ""
    }

### GET /m/[tickers]

Example: http://example.com/m/[tickers]

Response body:

    [
        {
            "data": {
                "symbol": "ticker 1",
                "currency": "USD",
                "price": 169.69,
                "change": 169.69,
                "perChange": 169.69
            },
            "errDescription": ""
        },
        {
            "data": {
                "symbol": "ticker 2",
                "currency": "",
                "price": 0,
                "change": 0,
                "perChange": 0
            },
            "errDescription": "No data found, symbol may be delisted"
        }
    ]
