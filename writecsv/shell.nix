{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [ go_1_11
                  ];
    shellHook = ''
        export GOPATH=`pwd`

        if [ $(uname -s) = "Darwin" ]; then
            alias ls='ls --color=auto'
            alias ll='ls -al'
        else
            alias open="xdg-open"
        fi

        gofmts() {
            gofmt -w -s -e $1

            if (( $? == 0 )); then
                awk '{ gsub(/\t/, "    "); print }' < $1 > tmp
                cat tmp > $1
                rm tmp
            fi
        }

        export -f gofmts
    '';
}
