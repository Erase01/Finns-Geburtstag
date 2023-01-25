{ pkgs ? import ( <nixpkgs> ) { } }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.helix
    pkgs.gopls
  ];
}
