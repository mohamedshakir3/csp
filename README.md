# CSPFinder


CSP Finder is a tool that compares potential ROI for selling various cash secured puts. .


Configuration file to accept inputs for CSP Finder

- **symbolpricemap**: a map of ticker symbol and the strike price
to compare csp options. Ideally, this is the price you are comfortable getting assigned at, depending on your risk tolerance
- **totalcapital**: total buying power in USD
- **outputformat**:
    - **groupbydate**: groups options by expiry date
    - **groupbysymbol**: groups options by their symbol/ticker
