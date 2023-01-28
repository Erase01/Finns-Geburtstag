{ pkgs ? import ( <nixpkgs> ) { } }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.docker
  ];

  shellHook = ''
    eval `ssh-agent`
    ssh-add ~/.ssh/id_rsa
  '';
}
