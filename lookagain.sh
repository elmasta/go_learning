find . -name '*.sh' -type f -printf "%f\n"  | sed 's#./##' | cut -f 1 -d '.'