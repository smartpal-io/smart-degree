var SmartDegree = artifacts.require("SmartDegree");
contract('SmartDegree', function() {
  it("should add hash to the contract", function() {
    var smartDegree;
    var result;
    return SmartDegree.deployed().then(function(instance) {
      smartDegree = instance;
      smartDegree.addDegreeHash('234564535',"0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function() {
      return smartDegree.getHash('234564535');
    }).then(function(hash) {
      result = hash;
    }).then(function() {
      assert.ok(result,"0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c", "Error valid hash not found");
    });
  });
  it("should found the good hash", function() {
    var smartDegree;
    var result;
    return SmartDegree.deployed().then(function(instance) {
      smartDegree = instance;
      smartDegree.addDegreeHash('234564536', "0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function() {
      smartDegree.addDegreeHash('234564537', "0xf7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9f");
    }).then(function() {
      return smartDegree.getHash('234564536');
    }).then(function(hash) {
      result = hash;
    }).then(function() {
      assert.ok(result,"0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c", "Error valid hash not found");
    });
  });
});
