# DISYS-AuctionHouse

This repository contains the source code for the DISYS Mini Project 3.

Made by:

* Danyal Yorulmaz (dayo)
* Deniz Isik (deni)
* Deniz Yildirim (deyi)
* Jakob Henriksen (jarh)

## How To Use

To start up the servers, simply go to AuctionHouse/server, open 3 terminals and use the following commands:

```powershell
go run .\server.go -port 7000
go run .\server.go -port 8000
go run .\server.go -port 9000
```

To start up the clients, simply go to Auctionhouse/client and open up as many terminals as you want. Run the following:

```powershell
go run .\client.go
```

Once the client has connected to the servers, you need to type `bid` followed by an amount. Example: `bid 1000`

Type `result` in the terminal to get the result of highest bidder.

Enjoy!