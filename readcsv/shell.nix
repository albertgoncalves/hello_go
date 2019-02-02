{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [ go_1_11
                    python36Packages.csvkit
                  ];
    shellHook = ''
        export GOPATH=`pwd`

        if [ $(uname -s) = "Darwin" ]; then
            alias ls='ls --color=auto'
            alias ll='ls -al'
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
        alias csvlook="csvlook --no-inference -d=';'"
    '';
}
