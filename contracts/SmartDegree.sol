pragma solidity ^0.4.24;

import "../node_modules/open-smartkit/contracts/integrity/Sealable.sol";
import "../node_modules/open-smartkit/contracts/lifecycle/Lockable.sol";

/**
 * @title SmartDegree
 * @dev SmartDegree ....
 */
contract SmartDegree is Sealable, Lockable {

    /**
     * @notice Create a new SmartDegree Contract.
     */
    constructor() Lockable(0) public {
        addAddressToWhitelist(msg.sender);
    }

    /**
     * @notice Register a new delegate authorized to add degree
     */
    function grantAuthority(address grantee) onlyOwner onlyUnlock public returns (bool success) {
        return super.registerDelegate(grantee);
    }

    /**
    * Use this getter function to access the degree hash value
    * @param id of the seal
    * @return the seal
    */
    function getDegreeHash(bytes32 id) public view returns (bytes32) {
        return super.getSeal(id);
    }

    /**
     * @notice Add a new DegreeHash to the contract.
     */
    function deliverDegree(bytes32 id, bytes32 degreeHash) public onlyWhitelisted onlyUnlock {
        return super.recordSeal(id, degreeHash);
    }

    /**
     * Use these method functions to verify a degree hash
     * @param id of the degree
     * @param degreeHash of the degree to verify
     */
    function isValid(bytes32 id, bytes32 degreeHash) public view returns (bool) {
        return super.verifySeal(id, degreeHash);
    }

    /**
     * @notice Define a date before blocking the contract
     * @param newDateLimit uint256 new date limit of the contract (timestamp as seconds since unix epoch)
     */
    function setDateLimit(uint256 newDateLimit) public onlyOwner {
        return super.setDateLimit(newDateLimit);
    }

    /**
   * @dev return contract status.
   * @return status (true if the contract is locked)
   */
    function isLocked() public view returns (bool){
        return super.isLocked();
    }
}
