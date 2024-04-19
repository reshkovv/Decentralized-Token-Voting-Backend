#!/bin/sh
abigen --bin solc_build_info/GovToken.bin --abi solc_build_info/GovToken.abi --type GovToken --pkg genFiles -out deployer/genFiles/GovToken.go
abigen --bin solc_build_info/MyGovernor.bin --abi solc_build_info/MyGovernor.abi --type MyGovernor --pkg genFiles -out deployer/genFiles/MyGovernor.go
abigen --bin solc_build_info/TimeLock.bin --abi solc_build_info/TimeLock.abi --type TimeLock --pkg genFiles -out deployer/genFiles/TimeLock.go