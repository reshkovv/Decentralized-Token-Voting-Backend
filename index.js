const express = require('express');
const { Web3 } = require('web3');
const governorContractABI = require('./governorAbi.json');
const app = express();
const port = process.env.PORT;
const RPC = process.env.RPC;
const governorAddress = process.env.GOVERNOR_CONTRACT_ADDRESS;

app.use(express.json());
const web3=new Web3(RPC);
const governorContract = new web3.eth.Contract(governorContractABI, governorAddress);

app.get('/proposalInfo/:proposalID', async(req, res) => {
  try {
    let proposalID = req.params.proposalID
    let proposer = await governorContract.methods.proposalProposer(proposalID).call();
    let votes = await governorContract.methods.proposalVotes(proposalID).call();
    let message = "proposal proposer:" + proposer + "\nproposal votes: votes_for:" + votes[1] +"  votes_against:" + votes[0] +"  votes_abstain:" + votes[2]
    res.send(message);  
  } catch (error) {
    res.send(error); 
  }
});

app.post('/propose', async(req, res) => {
  try {
    let postdata = req.body;
    const targets = postdata.targets; 
    const values = postdata.values; 
    const calldatas = postdata.calldatas;
   const description = postdata.description;
    
   const proposalData = governorContract.methods.propose(targets, values, calldatas, description).encodeABI();
  
   const account = process.env.ADMIN_ADDRESS;
   const privateKey = process.env.ADMIN_PK;
   const txCount = await web3.eth.getTransactionCount(account);
   const txData = {
     from: account,
     to: governorAddress,
     data: proposalData,
     nonce: web3.utils.toHex(txCount),
     gas: 2000000, // Gas limit
     gasPrice:await web3.eth.getGasPrice(), 
   };
   const signedTx = await web3.eth.accounts.signTransaction(txData, privateKey);
   const rawTransaction = signedTx.rawTransaction;
   web3.eth.sendSignedTransaction(rawTransaction)
    .on('receipt', (receipt) => {
      console.log('Transaction successful. Transaction hash:', receipt.transactionHash);
      var bn = BigInt(receipt.logs[0].data.substring(0,66));
      var proposalID = bn.toString(10);
      let message = 'Proposal created:'+ proposalID;
      res.send(message);
    })
    .on('error', (error) => {
      let message = 'Error sending transaction:'+ error;
      console.log(message);
      res.send(message);
    });  
  } catch (error) {
    res.send(error);
  }
});

app.post('/vote', async(req, res) => {
  try {
    let postdata = req.body;
    const proposalID = postdata.proposalID; 
    const support = postdata.support;
   const proposalData = governorContract.methods.castVote(proposalID,support).encodeABI();
  
   const account = process.env.ADMIN_ADDRESS;
   const privateKey = process.env.ADMIN_PK;
   const txCount = await web3.eth.getTransactionCount(account);
   const txData = {
     from: account,
     to: governorAddress,
     data: proposalData,
     nonce: web3.utils.toHex(txCount),
     gas: 2000000, // Gas limit
     gasPrice:await web3.eth.getGasPrice(), 
   };
   const signedTx = await web3.eth.accounts.signTransaction(txData, privateKey);
   const rawTransaction = signedTx.rawTransaction;
   await web3.eth.sendSignedTransaction(rawTransaction);
   res.send("voted");
  } catch (error) {
    res.send(error);
  }
});

app.listen(port, () => {
  console.log(`app listening at http://localhost:${port}`);
});