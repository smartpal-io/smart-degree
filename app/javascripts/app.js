// Import the page's CSS. Webpack will know what to do with it.
import "../stylesheets/app.css";

// Import libraries we need.
import { default as Web3} from 'web3';
import { default as contract } from 'truffle-contract';
import { default as ethUtil} from 'ethereumjs-util';
import { default as sigUtil} from 'eth-sig-util';


/*
 * When you compile and deploy your Voting contract,
 * truffle stores the abi and deployed address in a json
 * file in the build directory. We will use this information
 * to setup a Voting abstraction. We will use this abstraction
 * later to create an instance of the Voting contract.
 * Compare this against the index.js from our previous tutorial to see the difference
 * https://gist.github.com/maheshmurthy/f6e96d6b3fff4cd4fa7f892de8a1a1b4#file-index-js
 */

import smart_degree_artifacts from '../../build/contracts/SmartDegree.json'

var SmartDegree = contract(smart_degree_artifacts);


window.registerDegree = function(degree) {
  let degreeId = $("#register-degree-id").val();
  let degreeHash = $("#register-degree-hash").val();
  console.log("degree id : ", degreeId);
  console.log("degree hash : ", degreeHash);

  
  SmartDegree.deployed().then(function(contractInstance) {
	console.log("wallet used : ", web3.eth.accounts[0])
	contractInstance.addDegreeHash(degreeId,toBytes(degreeHash), {gas: 140000, from: web3.eth.accounts[0]});
  }).then(function() {
      $("#msg").html("Degree hash added")
  });
}

window.verifyDegree = function(degree) {
	let degreeId = $("#verify-degree-id").val();
    let degreeHash = $("#verify-degree-hash").val();
	
	SmartDegree.deployed().then(function(contractInstance) {
		return contractInstance.verify(degreeId,toBytes(degreeHash));
	}).then(function(result) {
      $("#msg").html("Verify hash result "+result)
    })
  
}

function toBytes(hash){
	return "0x".concat(hash);
}

$( document ).ready(function() {
  if (typeof web3 !== 'undefined') {
    console.warn("Using web3 detected from external source like Metamask")
    // Use Mist/MetaMask's provider
    window.web3 = new Web3(web3.currentProvider);
  } else {
    console.warn("No web3 detected. Falling back to http://localhost:8545. You should remove this fallback when you deploy live, as it's inherently insecure. Consider switching to Metamask for development. More info here: http://truffleframework.com/tutorials/truffle-and-metamask");
    // fallback - use your fallback strategy (local node / hosted node + in-dapp id mgmt / fail)
    window.web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
  }

  SmartDegree.setProvider(web3.currentProvider);
});
