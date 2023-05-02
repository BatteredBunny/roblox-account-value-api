<h1 align="center">Roblox account value api</h1>

# API

## Account collectibles value API
```
GET https://roblox-account-value-api.sly.ee/api/collectibles-account-value?userid=XXX
{
	"total_robux": 0,
	"in_euro": 0,
	"collectibles": [
	    "name": "Collectibles",
	    "price": 0,
	    "id": 0,
	    "serialnumber": 0,
		"thumbnail": "https://tr.rbxcdn.com/"
	]
}
```

## Can view inventory API
```
GET https://roblox-account-value-api.sly.ee/api/can-view-inventory?userid=XXX
true
```

## Profile info API
```
GET https://roblox-account-value-api.sly.ee/api/profile-info?userid=XXX
{
    "username": "username",
    "displayname": "displayname",
    "avatar": "https://tr.rbxcdn.com/avatar"
}
```

## Exchange Rate API
```
GET https://roblox-account-value-api.sly.ee/api/exchange-rate
{
    "robux_per_euro": 60
}
```

## 
# Build program
```
nix build .
```

# Build docker container
```
nix run .#docker.copyToDockerDaemon
```