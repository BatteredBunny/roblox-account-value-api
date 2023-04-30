{ pkgs, buildGoModule, lib }: buildGoModule rec {
    src = ./.;

    name = "roblox-account-value-api";
    vendorSha256 = "sha256-pTw0KdOWjKDLb4rx87WGt9D6qok73J5uNRG+ecyAaaA=";

    ldflags = [
        "-s"
        "-w"
    ];
}