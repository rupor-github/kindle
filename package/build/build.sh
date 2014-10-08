#!/bin/sh -e

HACKNAME="rupor"
HACKDIR="accal"
PKGNAME="${HACKNAME}"
PKGVER="0.2"

# We need kindletool (https://github.com/NiLuJe/KindleTool) in $PATH
if (( $(kindletool version | wc -l) == 1 )) ; then
	HAS_KINDLETOOL="true"
fi

if [[ "${HAS_KINDLETOOL}" != "true" ]] ; then
	echo "You need KindleTool (https://github.com/NiLuJe/KindleTool) to build this package."
	exit 1
fi

# We also need GNU tar
if [[ "$(uname -s)" == "Darwin" ]] ; then
	TAR_BIN="gtar"
else
	TAR_BIN="tar"
fi
if ! ${TAR_BIN} --version | grep "GNU tar" > /dev/null 2>&1 ; then
	echo "You need GNU tar to build this package."
	exit 1
fi

# check if we have necessary components
if [[ ! -d "../../goprojects/bin/linux_arm" ]] ; then
	echo "Nothing to build..."
	exit 1
fi

# Always take fresh build
[ ! -d ../src/${HACKDIR}/bin ] && mkdir -p ../src/${HACKDIR}/bin
cp ../../goprojects/bin/linux_arm/kal ../src/${HACKDIR}/bin/.
cp ../../goprojects/config ../src/${HACKDIR}/.

## Install

# Archive custom directory
${TAR_BIN} --owner root --group root --exclude-vcs -cvjf ${HACKNAME}.tar.bz2 -C ../src ${HACKDIR} extensions

# Copy the script to our working directory, to avoid storing crappy paths in the update package
cp ../src/install.sh ./

# Build the install package (PaperWhite 2)
kindletool create ota2 -d paperwhite2 libotautils5 install.sh ${HACKNAME}.tar.bz2 Update_${PKGNAME}_install_pw2.bin

## Uninstall
# Copy the script to our working directory, to avoid storing crappy paths in the update package
cp ../src/uninstall.sh ./

# Build the uninstall package
kindletool create ota2 -d paperwhite2 libotautils5 uninstall.sh Update_${PKGNAME}_uninstall_pw2.bin

## Cleanup
# Remove package specific temp stuff
rm -f ./install.sh ./uninstall.sh ./${HACKNAME}.tar.bz2

# Move our updates
git log --format=medium -n 10 >HISTORY
[ ! -f ../${PKGNAME}_${PKGVER}.7z ] && rm -f ../${PKGNAME}_${PKGVER}.7z
7z a ../${PKGNAME}_${PKGVER}.7z README HISTORY Update_*.bin
rm -f HISTORY *.bin
