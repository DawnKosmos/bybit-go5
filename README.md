# bybit-go5

# WIP

Implementation of the Bybit API v5 in golang

Version 5 unifies Spot/USDC/USDT/Inverse

[Docs](https://bybit-exchange.github.io/docs/v5/intro)

Names of Structs and Functions and their comments are generated from the above Documentation.
Due to limitations in the GoLang Type Systems some endpoints have 3 Functions. 
1 For each possible Response struct from the Api.
E.g. GetInstrumentsInfo depending on the **Category** linear/inverse, spot or options
You get another response struct



