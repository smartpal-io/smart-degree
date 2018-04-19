var SmartDegree = artifacts.require("SmartDegree");

contract('SmartDegree', function() {
  it("should add hash to the contract", function() {
    var smartDegree;
    var result;

    return SmartDegree.deployed().then(function(instance) {
      smartDegree = instance;
      return smartDegree.addDegreeHash('234564535', "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function(result) {
      return smartDegree.verify('234564535',"0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function(isValid) {
      result = isValid;
    }).then(function() {
      assert.ok(result, "Invalid hash found");
    });
  });
});
