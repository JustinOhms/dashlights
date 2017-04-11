#!/bin/bash
export DASHLIGHT_FOO_2112_BGWHITE="Foo was here."
export DASHLIGHT_BAR_1F4A9="Poo was here."

echo "$ ./dashlights"
./dashlights

echo
echo "$ ./dashlights -diag"
./dashlights -diag

echo
echo "$ ./dashlights -clear"
./dashlights -clear

echo
echo "$ ./dashlights -listcolors"
./dashlights -listcolors

