#! /bin/sh

# brew install graphviz
# sudo cpan Graph:Easy

cat <<EOF | graph-easy
[ root ] -> [ height <= 100 ] -> [ procompsognathus ]
[ root ] -> [ height > 100 ]
[ height > 100 ] -> [ carnivores? ] -> [ t-rex ]
[ height > 100 ] -> [ herbivores? ] -> [ brontosaurus ]
EOF
