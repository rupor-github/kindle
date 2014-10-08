#!/bin/sh

HACKNAME="rupor"

# Pull libOTAUtils for logging & progress handling
[ -f ./libotautils5 ] && source ./libotautils5

# Directories

HACKDIR="accal"

# Result codes
OK=0
ERR=${OK}

otautils_update_progressbar

# Always do clean install (for now)
[ -d /mnt/us/${HACKDIR} ] && rm -rf /mnt/us/${HACKDIR}
[ -d /mnt/us/extensions/${HACKDIR} ] && rm -rf /mnt/us/extensions/${HACKDIR}

# Extract it
logmsg "I" "install" "" "installing our code"
tar -xvjf ${HACKNAME}.tar.bz2 -C /mnt/us
# Do check if that went well
_RET=$?
if [ ${_RET} -ne 0 ] ; then
    logmsg "C" "install" "code=${_RET}" "failed to extract our code to /mnt/us"
    return 1
fi
# Just in case
chmod +x /mnt/us/${HACKDIR}/bin/kal

otautils_update_progressbar

logmsg "I" "install" "" "Patching /etc/lipc-daemon-events.conf"

# Patch /etc/lipc-daemon-events.conf
sed -i "/fullScanFinish\s\+com\.lab126\.scanner.\+$/ d" /etc/lipc-daemon-events.conf
sed -i "/PowerButtonHeld\s\+com\.lab126\.powerd.\+$/ afullScanFinish    com.lab126.scanner    /mnt/us/${HACKDIR}/bin/kal -action=sync" /etc/lipc-daemon-events.conf

otautils_update_progressbar

logmsg "I" "install" "" "cleaning up"
rm -f ${HACKNAME}.tar.bz2

otautils_update_progressbar

logmsg "I" "install" "" "done"

otautils_update_progressbar

return ${OK}
