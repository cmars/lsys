{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {
  nativeBuildInputs = with pkgs.buildPackages; [ go_1_18 ];
  shellHook = ''
    export GOPATH="$HOME/.cache/gopaths/$(sha256sum <<<$(pwd) | awk '{print $1}')"
  '';
}

