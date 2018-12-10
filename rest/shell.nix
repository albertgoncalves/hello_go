{ pkgs ? import <nixpkgs> {} }:

with pkgs; mkShell {
    name = "go";

    buildInputs = [ go_1_11
                    tmux
                  ];

    shellHook = ''
        if [ ! -e ./src/ ]; then
            mkdir src
        fi

        if [ ! -e src/mux-1.6.2/ ]; then
            curl https://github.com/gorilla/mux/archive/v1.6.2.zip \
                -J \
                -L \
                -o src/mux.zip
            unzip src/mux.zip -d src/
        fi

        export GOPATH=`pwd`
    '';
}
