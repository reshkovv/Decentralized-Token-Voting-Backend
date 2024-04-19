#!/bin/sh
mkdir -p goAdmin/genFiles
abigen --bin solc_build_info/GovToken.bin --abi solc_build_info/GovToken.abi --type GovToken --pkg genFiles -out goAdmin/genFiles/GovToken.go
abigen --bin solc_build_info/MyGovernor.bin --abi solc_build_info/MyGovernor.abi --type MyGovernor --pkg genFiles -out goAdmin/genFiles/MyGovernor.go
abigen --bin solc_build_info/TimeLock.bin --abi solc_build_info/TimeLock.abi --type TimeLock --pkg genFiles -out goAdmin/genFiles/TimeLock.go