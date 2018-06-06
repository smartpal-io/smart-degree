pragma solidity ^0.4.24;

import "../node_modules/open-smartkit/contracts/integrity/Sealable.sol";

/**
 * @title SmartDegree
 * @dev SmartDegree ....
 */
contract SmartDegree is Sealable {

  event LogDegreeDelivered(bytes32 indexed id, bytes32 indexed hash);
  event LogAuthorityGranted(address grantee);

  /**
   * @notice Create a new SmartDegree Contract.
   */
  constructor() public {
    addAddressToWhitelist(msg.sender);
  }

  /**
   * @notice Register a new delegate authorized to add degree
   */
  function grantAuthority(address grantee) onlyOwner public returns(bool success) {  
	success = super.registerDelegate(grantee);
	emit LogAuthorityGranted(grantee);
	return success;
  }
  
   /**
   * Use this getter function to access the degree hash value
   * @param id of the seal
   * @return the seal
   */
  function getDegreeHash(bytes32 id) public view returns(bytes32) {
    return super.getSeal(id);
  }
  
  /**
   * @notice Deliver a degree
   */
  function deliverDegree(bytes32 id, bytes32 degreeHash) public onlyWhitelisted {
	 super.recordSeal(id, degreeHash);
     emit LogDegreeDelivered(id,degreeHash);
  }


  /**
   * Use these method functions to verify a degree hash
   * @param id of the degree
   * @param degreeHash of the degree to verify
   */
  function isValid(bytes32 id, bytes32 degreeHash) public view returns(bool) {
    return super.verifySeal(id, degreeHash);
  }

}
