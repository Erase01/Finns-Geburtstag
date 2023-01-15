{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-22.11.tar.gz") {} }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.helix
    pkgs.gopls
  ];
}
