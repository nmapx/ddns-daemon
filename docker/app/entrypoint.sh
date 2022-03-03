#!/usr/bin/env sh

set -e

if [ "$1" = "daemon" ]
  then exec make -f Makefile.app daemon
elif [ "$1" = "test" ]
  then exec make -f Makefile.app test
elif [ "$1" = "loop" ]
  then exec sh -c "while true; do echo ping; sleep 1; done"
fi

exec $@
