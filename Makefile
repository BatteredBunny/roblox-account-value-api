clean:
	go clean
	rm -rf result

docker:
	nix run .#docker.copyToDockerDaemon

docker-push: docker
	docker push ghcr.io/ayes-web/roblox-account-value-api:latest