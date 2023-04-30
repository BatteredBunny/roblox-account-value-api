<h1 align="center">Roblox account value api</h1>

# API 

## Account collectibles value API
```
GET https://roblox-account-value-api.sly.ee/api/limited-account-value?userid=XXX
{
	"total_robux": 0,
	"in_euro": 0
}
```

## Can view inventory API
```
GET https://roblox-account-value-api.sly.ee/api/can-view-inventory?userid=XXX
true
```

# Build program
```
nix build .
```

# Build docker container
```
nix run .#docker.copyToDockerDaemon
```