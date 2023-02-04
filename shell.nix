{ pkgs ? import ( <nixpkgs> ) { } }:

pkgs.mkShell {
  nativeBuildInputs = [
    pkgs.go
    pkgs.tmux
    pkgs.burpsuite
  ];

  shellHook = ''
    eval `ssh-agent`
    ssh-add ~/.ssh/id_rsa
    git pull
  '';
}
