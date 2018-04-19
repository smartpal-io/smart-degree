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
     emit DegreeHashAdded(id,hash);
  }

  /**
   * Use these getter functions to access the degree hash
   * @param id of the degree
   * @param hash of the degree to verify
   */
  function verify(uint256 id, bytes32 hash) public view returns(bool) {
    return hashList_[id]==hash;
  }

}
