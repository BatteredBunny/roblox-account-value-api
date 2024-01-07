clean:
	go clean
	rm -rf result

docker:
	nix run .#docker.copyToDockerDaemon

docker-push: docker
	docker push ghcr.io/batteredbunny/roblox-account-value-api:latest