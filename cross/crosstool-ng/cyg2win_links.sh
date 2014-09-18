#!/usr/bin/bash

d="$1"

echo $cmd

recursedirs() {
        pushd "$1" > /dev/null
	for f in "$1"/*; do
		if [ -d "$f" ] && [ ! -L "$f" ]; then
            		recursedirs "$f"
		elif [ -d "$f" ] && [ -L "$f" ]; then
                        n=`readlink "$f"`
			if [ -L "$n" ]; then
				n=`readlink "$n"`
			fi
            		echo "Converting dir : ("$f") `realpath --relative-base="$1" --no-symlinks "$f"` to "$n" in `pwd`"
			rm "$f"
                        CYGWIN="winsymlinks:native" ln -s -v "$n" `realpath --relative-base="$1" --no-symlinks "$f"`
		elif [ -f "$f" ] && [ -L "$f" ]; then
                        n=`readlink "$f"`
            		echo "Converting file: ("$f") `realpath --relative-base="$1" --no-symlinks "$f"` to "$n" in `pwd`"
			rm "$f"
                        CYGWIN="winsymlinks:native" ln -s -v "$n" `realpath --relative-base="$1" --no-symlinks "$f"`
        	fi
	done
	popd > /dev/null
}

recursedirs $d
