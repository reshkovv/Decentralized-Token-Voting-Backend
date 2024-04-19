Decentralized Token Voting Backend
==============

# Info
Contracts folder contains 3 smart contracts required for the DAO (GovToken.sol, MyGovernor.sol and TimeLock.sol). These contracts are deployed on sepolia network using the private key in .env file. This private key has eth funds on sepolia testnet for testing purposes.

GoAdmin folder contains golang code for fresh deployment of the contracts and go programs for interaction the smart contracts

Scripts folder contains shell scripts for compiling the contracts

index.js is the main file for the nodejs app and contains the API configuration and logic

# Instructions

## Smart contracts deployment (this step can be skipped because contracts are deployed, and go directly to `RUN APP` )

requirements (golang,solc and abigen installed globally) 


### Build

build contracts using the Solidity Compiler to generate the abi and bin files

run compile_contracts.sh script  ./scripts/compile_contracts.sh


### create go files

generate .go files which re required for the deployer

steps:
run generate_go_files.sh script

command:
./scripts/generate_go_files.sh


### install go dependencies

command:
cd goAdmin && go mod install


### Deploy contracts

deploy the contracts to sepolia testnet

command:
cd goAdmin && go run deployContracts.go


### update the config.yml file
update governorAddress, govTokenAddress and timeLockAddress inside goAdmin/config.yml with the new values from the previous step

### Delegate govToken
required to create snapshot for voting

command:
go run delegate.go


## RUN APP using nodejs
This command will start the API server, you can test the endpoints using the postman collection

command:
yarn install && yarn start

## RUN APP using golang
You can also test the app using go programs

commands:\
cd goAdmin (change working directory)\
go mod install (install dependencies)\
go run createProposal.go (create new proposal)\
go run getProposalInfo.go (get proposal info, proposalID in config.yml is used)\
go run proposeVote.go (vote for proposal, proposalID in config.yml is used)


