const contract = require('truffle-contract');
const artifacts = require('./build/contracts/SmartDegree.json');
const SmartDegree = contract(artifacts);

var Web3 = require('web3');

if (typeof web3 !== 'undefined') {
  var web3 = new Web3(web3.currentProvider)
} else {
  var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'))
}
SmartDegree.setProvider(web3.currentProvider);

var smartDegree = SmartDegree.at("0x400770ad9af253d328e86931c941b80ef382b688");
smartDegree.deliverDegree("0x000000000000000000000000000000000000000000000000000000000000001", "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");

smartDegree.getDegreeHash("0x000000000000000000000000000000000000000000000000000000000000001")
.then((result) => {
	console.log(result);
})
.catch((err) => {
	console.log(err);
});
	

module.exports = function(callback) {
  // perform actions
}