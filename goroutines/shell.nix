{ pkgs ? import <nixpkgs> {} }:

with pkgs; mkShell {
    name = "go";

    buildInputs = [ go_1_11
                    tmux
                  ];

    shellHook = ''
        export GOPATH=`pwd`
    '';
}
