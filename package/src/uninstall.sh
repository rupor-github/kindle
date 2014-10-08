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

logmsg "I" "install" "" "Patching /etc/lipc-daemon-events.conf"

# Patch /etc/lipc-daemon-events.conf
sed -i "/fullScanFinish\s\+com\.lab126\.scanner.\+$/ d" /etc/lipc-daemon-events.conf

otautils_update_progressbar

logmsg "I" "uninstall" "" "removing files & directories"

[ -d /mnt/us/${HACKDIR} ] && rm -rf /mnt/us/${HACKDIR}
[ -d /mnt/us/extensions/${HACKDIR} ] && rm -rf /mnt/us/extensions/${HACKDIR}

otautils_update_progressbar

logmsg "I" "uninstall" "" "done"

return ${OK}
