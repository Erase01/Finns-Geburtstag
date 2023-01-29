{ pkgs ? import ( <nixpkgs> ) { } }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.tmux
  ];

  shellHook = ''
    eval `ssh-agent`
    ssh-add ~/.ssh/id_rsa
    git pull
  '';
}
