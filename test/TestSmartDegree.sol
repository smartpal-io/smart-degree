pragma solidity ^0.4.11;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/SmartDegree.sol";

contract TestSmartDegree {


  function testAddDegree() public {
    SmartDegree smartDegree = new SmartDegree();

    uint32 studentId = 1234;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 expected = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;

	smartDegree.addDegreeHash(studentId, degreeHash);

    Assert.equal(smartDegree.getHash(studentId), expected, "degree hash should be equal");
  }
  
  function testGetHashUnknownStudent() public {
    SmartDegree smartDegree = new SmartDegree();

    uint32 studentId = 1234;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 expected = 0x00;

    Assert.equal(smartDegree.getHash(studentId), expected, "degree hash should be equal");
  }
  
  function testVerifyValidDegree() public {
    SmartDegree smartDegree = new SmartDegree();

    uint32 studentId = 1234;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bool expected = true;

	smartDegree.addDegreeHash(studentId, degreeHash);

    Assert.equal(smartDegree.verify(studentId, degreeHash), expected, "verify should return true");
  }
  
  function testVerifyInvalidDegree() public {
    SmartDegree smartDegree = new SmartDegree();

    uint32 studentId = 1234;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	bytes32 wrongDegreeHash = 0x99834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;

	bool expected = false;

	smartDegree.addDegreeHash(studentId, degreeHash);

    Assert.
	
	equal(smartDegree.verify(studentId, wrongDegreeHash), expected, "verify should return false");
  }
  
  function testVerifyUnknownStudent() public {
    SmartDegree smartDegree = new SmartDegree();

    uint32 studentId = 1234;
	bytes32 degreeHash = 0xa7834034bd059ecf00b0661f88f1e7242450bf1951c1e76803e80ce4182e2e9c;
	uint32 unkownStudentId = 9876;

	bool expected = false;

	smartDegree.addDegreeHash(studentId, degreeHash);

    Assert.
	
	equal(smartDegree.verify(unkownStudentId, degreeHash), expected, "verify should return false");
  }
}