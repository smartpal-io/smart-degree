pragma solidity ^0.4.11;

import "../node_modules/zeppelin-solidity/contracts/ownership/Ownable.sol";


/**
 * @title SmartDegree
 * @dev SmartDegree ....
 */
contract SmartDegree is Ownable {

  mapping (uint256 => bytes32) private hashList_;

  event DegreeHashAdded(uint256 indexed id, bytes32 indexed hash);


  /**
   * @notice Create a new SmartDegree Contract.
   */
  function SmartDegree() public {
  }

  /**
   * @notice Add a new DegreeHash to the contract.
   */
  function addDegreeHash(uint256 id, bytes32 hash) public onlyOwner {
     require(hashList_[id] == bytes32(0x0));
     hashList_[id]=hash;
     emit DegreeHashAdded(id,hash);
  }

  /**
   * Use these getter functions to access the degree hash
   * @param id of the degree
   * @return hash of the degree to verify
   */
  function getHash(uint256 id) public view returns(bytes32) {
    return hashList_[id];
  }
  
  function verify(uint256 id, bytes32 hash) public view returns(bool) {
    return hashList_[id]==hash;
  }

}
