<h1 align="center">Roblox account value api</h1>

## Simple companion API for the [web app](https://roblox-account-value.sly.ee/)

This is needed as browser itself can't do many of the requests to roblox so a proxy of sorts is needed.

# Recommended usage
The recommended way to use this is to use the docker image, check the [docker-compose.yml](https://github.com/ayes-web/roblox-account-value-api/blob/master/docker-compose.yml) in the repo :)

# Building docker image with nix
```
nix run github:ayes-web/roblox-account-value-api#docker.copyToDockerDaemon
```

# Building standalone program with nix
```
nix build github:ayes-web/roblox-account-value-api
```


# API
Handy info for utilizing the api yourself
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