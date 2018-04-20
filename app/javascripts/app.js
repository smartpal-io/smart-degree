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


window.registerDegree = function(student) {
  let studentId = $("#register-student-id").val();
  let degreeHash = $("#register-degree-hash").val();
  console.log("student id : ", degreeHash);
  console.log("degree hash : ", studentId);

  $("#msg").html("Degree has been registered")
  
  SmartDegree.deployed().then(function(contractInstance) {
	console.log("wallet used : ", web3.eth.accounts[0])
	contractInstance.addDegreeHash(studentId,degreeHash, {gas: 140000, from: web3.eth.accounts[0]});
	console.log("addDegreeHash completed")
    /*contractInstance.voteForCandidate(candidateName, voterAddress, signature, {gas: 140000, from: web3.eth.accounts[0]}).then(function() {
      let div_id = candidates[candidateName];
      console.log(div_id);
      return contractInstance.totalVotesFor.call(candidateName).then(function(v) {
        console.log(v.toString());
        $("#" + div_id).html(v.toString());
        $("#msg").html("");
      });
    });*/
  });
}

window.verifyDegree = function(student) {
	console.log("verifyDegree invoked")
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
