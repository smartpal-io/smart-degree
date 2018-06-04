pragma solidity ^0.4.11;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/SmartDegree.sol";

contract TestSmartDegree {


  function testAddDegree() public {
    SmartDegree smartDegree = new SmartDegree();

    bytes32 studentId = 0x000000000000000000000000000000000000000000000000000000000000001;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 expected = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;

	smartDegree.deliverDegree(studentId, degreeHash);

    Assert.equal(smartDegree.getDegreeHash(studentId), expected, "degree hash should be equal");
  }
  
  function testgetDegreeHashUnknownStudent() public {
    SmartDegree smartDegree = new SmartDegree();

    bytes32 studentId = 0x000000000000000000000000000000000000000000000000000000000000001;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 expected = 0x00;

    Assert.equal(smartDegree.getDegreeHash(studentId), expected, "degree hash should be equal");
  }
  
  function testisValidValidDegree() public {
    SmartDegree smartDegree = new SmartDegree();

    bytes32 studentId = 0x000000000000000000000000000000000000000000000000000000000000001;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bool expected = true;

	smartDegree.deliverDegree(studentId, degreeHash);

    Assert.equal(smartDegree.isValid(studentId, degreeHash), expected, "isValid should return true");
  }
  
  function testisValidInvalidDegree() public {
    SmartDegree smartDegree = new SmartDegree();

    bytes32 studentId = 0x000000000000000000000000000000000000000000000000000000000000001;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 wrongDegreeHash = 0x99834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;

	bool expected = false;

	smartDegree.deliverDegree(studentId, degreeHash);

    Assert.
	
	equal(smartDegree.isValid(studentId, wrongDegreeHash), expected, "isValid should return false");
  }
  
  function testisValidUnknownStudent() public {
    SmartDegree smartDegree = new SmartDegree();

    bytes32 studentId = 0x000000000000000000000000000000000000000000000000000000000000001;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 unkownStudentId = 0x000000000000000000000000000000000000000000000000000000000000002;

	bool expected = false;

	smartDegree.deliverDegree(studentId, degreeHash);

    Assert.
	
	equal(smartDegree.isValid(unkownStudentId, degreeHash), expected, "isValid should return false");
  }
}