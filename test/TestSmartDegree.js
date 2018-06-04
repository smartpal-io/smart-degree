var SmartDegree = artifacts.require("SmartDegree");
contract('SmartDegree', function() {
  it("should add hash to the contract", function() {
    var smartDegree;
    var result;
    return SmartDegree.deployed().then(function(instance) {
      smartDegree = instance;
      smartDegree.deliverDegree('0x000000000000000000000000000000000000000000000000000000000000001',"0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function() {
      return smartDegree.getDegreeHash('0x000000000000000000000000000000000000000000000000000000000000001');
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
      smartDegree.deliverDegree('0x000000000000000000000000000000000000000000000000000000000000002', "0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function() {
      smartDegree.deliverDegree('0x000000000000000000000000000000000000000000000000000000000000003', "0xf7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9f");
    }).then(function() {
      return smartDegree.getDegreeHash('0x000000000000000000000000000000000000000000000000000000000000002');
    }).then(function(hash) {
      result = hash;
    }).then(function() {
      assert.ok(result,"0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c", "Error valid hash not found");
    });
  });
   it("should isValid the hash", function() {
    var smartDegree;
    var result;
    return SmartDegree.deployed().then(function(instance) {
      smartDegree = instance;
      return smartDegree.deliverDegree('0x000000000000000000000000000000000000000000000000000000000000004', "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function(result) {
      return smartDegree.isValid('0x000000000000000000000000000000000000000000000000000000000000004',"0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function(isValid) {
      result = isValid;
    }).then(function() {
      assert.ok(result, "Error valid hash not detected");
    }).then(function() {
      return smartDegree.deliverDegree('0x000000000000000000000000000000000000000000000000000000000000005', "0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c");
    }).then(function(result) {
      return smartDegree.isValid('0x000000000000000000000000000000000000000000000000000000000000005',"0xe7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9f");
    }).then(function(isValid) {
      result = isValid;
    }).then(function() {
      assert.ok(!result, "Error invalid hash not detected");
    });
  });
});
