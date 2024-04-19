#!/bin/sh
solc --bin --abi --include-path node_modules/ --base-path . contracts/GovToken.sol -o solc_build_info --overwrite --optimize
solc --bin --abi --include-path node_modules/ --base-path . contracts/MyGovernor.sol -o solc_build_info --overwrite --optimize
solc --bin --abi --include-path node_modules/ --base-path . contracts/TimeLock.sol -o solc_build_info --overwrite --optimize
