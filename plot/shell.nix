{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [
        go
        shellcheck
    ];
    shellHook = ''
        . .shellhook
    '';
}
