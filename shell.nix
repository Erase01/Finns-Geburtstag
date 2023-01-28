{ pkgs ? import ( <nixpkgs> ) { } }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.docker
  ];

  shellHook = ''
    export GOPATH=$PWD
    export PATH=$PATH:$GOPATH/bin
    eval `ssh-agent`
    ssh-add ~/.ssh/id_rsa
  '';
}
