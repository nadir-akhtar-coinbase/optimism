// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Safe, Enum } from "safe-contracts/Safe.sol";

import { ISemver } from "src/universal/ISemver.sol";

/// @title ThresholdModule
/// @notice This module allows any owner of the Safe Wallet to execute a veto through this.
contract VetoModule is ISemver {
    /// @notice Semantic version.
    /// @custom:semver 1.0.0
    string public constant version = "1.0.0";

    /// @notice The Safe contract instance
    Safe internal immutable SAFE;

    /// @notice The delayed vetoable contract.
    address internal immutable DELAYED_VETOABLE;

    /// @notice The module constructor.
    /// @param safe The Safe Wallet addess.
    /// @param delayedVetoable The delay vetoable contract address.
    constructor(Safe safe, address delayedVetoable) {
        SAFE = safe;
        DELAYED_VETOABLE = delayedVetoable;
    }

    /// @notice Passthrough for any owner to execute a veto on the `DelayedVetoable` contract.
    /// @dev Revert if not called by a Safe Wallet owner address.
    function veto() external returns (bool) {
        // Ensure only a Safe Wallet owner can veto.
        require(SAFE.isOwner(msg.sender), "VetoModule: only owners can call veto");

        // Forward the call to the Safe Wallet, targeting the delayed vetoable contract.
        return SAFE.execTransactionFromModule(DELAYED_VETOABLE, 0, msg.data, Enum.Operation.Call);
    }
}
