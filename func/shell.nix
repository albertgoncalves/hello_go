{ pkgs ? import <nixpkgs> {} }:

with pkgs; mkShell {
    name = "go";

    buildInputs = [ go_1_11
                    tmux
                  ];

    shellHook = ''
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
