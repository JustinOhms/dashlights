#!/bin/bash
export DASHLIGHT_FOO_2112_FGHIWHITE="Look, your name in lights!"
export DASHLIGHT_BAR_1F4A9="Something is rotten in the state of Denmark."

if [ ! -f "./dashlights" ]; then
  go build -o ./dashlights
fi

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

