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

        gofmts() {
            gofmt -w $1

            if (( $? == 0 )); then
                awk '{ gsub(/\t/, "    "); print }' < $1 > tmp
                cat tmp > $1
                rm tmp
            fi
        }

        export -f gofmts
        export GOPATH=`pwd`
    '';
}
