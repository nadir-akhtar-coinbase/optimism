// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";

import { Safe } from "safe-contracts/Safe.sol";
import { BaseGuard } from "safe-contracts/base/GuardManager.sol";
import { Enum } from "safe-contracts/common/Enum.sol";

import { ISemver } from "src/universal/ISemver.sol";

/// @title OwnerGuard
/// @notice This Guard contract is used to enforce a maximum number of owners and a required
///         threshold for the Safe Wallet.
contract OwnerGuard is ISemver, BaseGuard {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

        /// @notice The Safe Wallet for which this contract will be the guard.
    Safe internal immutable SAFE;

    /// @notice The maximum number of owners. Can be changed by supermajority.
    uint8 public maxCount;

    /// @notice Constructor.
    /// @param safe_ The Safe Wallet for which this contract will be the guard.
    constructor(Safe safe_) {
        SAFE = safe_;

        // Get the current owner count of the Smart Wallet.
        uint256 ownerCount = safe_.getOwners().length;
        require(ownerCount <= type(uint8).max, "OwnerGuard: owner count too high");

        // Set the initial `maxCount`, to the greater between 7 and the current owner count.
        maxCount = uint8(FixedPointMathLib.max(7, ownerCount));
    }

    /// @notice Inherited hook from `BaseGuard` that is run right before the transaction is executed
    ///         by the Safe Wallet when `execTransaction` is called.
    function checkTransaction(
        address,
        uint256,
        bytes memory,
        Enum.Operation,
        uint256,
        uint256,
        uint256,
        address,
        address payable,
        bytes memory,
        address
    )
        external
    {
    }

    /// @notice Inherited hook from `BaseGuard` that is run right after the transaction has been executed
    ///         by the Safe Wallet when `execTransaction` is called.
    function checkAfterExecution(bytes32, bool) external view {
        // Ensure the length of the new set of owners is not above `maxCount`, and get the corresponding threshold.
        uint256 threshold_ = checkNewOwnerCount(SAFE.getOwners().length);

        // Ensure the Safe Wallet threshold always stays in sync with the 66% one.
        require(
            SAFE.getThreshold() == threshold_,
            "OwnerGuard: Safe must have a threshold of at least 66% of the number of owners"
        );
    }

    /// @notice Checks if the given `newCount` of owners is allowed and returns the corresponding 66% threshold.
    /// @dev Reverts if `newCount` is above `maxCount`.
    /// @param newCount The owners count to check.
    /// @return threshold_ The corresponding 66% threshold for `newCount` owners.
    function checkNewOwnerCount(uint256 newCount) public view returns (uint256 threshold_) {
        // Ensure we don't exceed the maximum number of allowed owners.
        require(newCount <= maxCount, "OwnerGuard: too many owners");

        // Compute the corresponding ceil(66%) threshold of owners.
        threshold_ = (newCount * 66 + 99) / 100;
    }

    /// @notice Update the maximum number of owners.
    /// @dev Reverts if not called by the Safe Wallet.
    /// @param newMaxCount The new possible `maxCount` of owners.
    function updateMaxCount(uint8 newMaxCount) external {
        // Ensure only the Safe Wallet can call this function.
        require(msg.sender == address(SAFE), "OwnerGuard: only Safe can call this function");

        // Ensure the given `newMaxCount` is not bellow the current number of owners.
        require(newMaxCount >= SAFE.getOwners().length);

        // Update the new`maxCount`.
        maxCount = newMaxCount;
    }

    /// @notice Getter function for the Safe contract instance
    /// @return safe_ The Safe contract instance
    function safe() public view returns (Safe safe_) {
        safe_ = SAFE;
    }
}
