// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { IWETH } from "src/dispute/interfaces/IWETH.sol";

/// @title IDelayedWETH
/// @notice Interface for the DelayedWETH contract.
interface IDelayedWETH is IWETH {
    /// @notice Emitted when an unwrap is started.
    /// @param src The address that started the unwrap.
    /// @param wad The amount of WETH that was unwrapped.
    event Unwrap(address indexed src, uint256 wad);

    /// @notice Returns the withdrawal delay in seconds.
    /// @return The withdrawal delay in seconds.
    function delay() external view returns (uint256);

    /// @notice Returns the withdrawal request for the given address.
    /// @param owner The address to query the withdrawal request of.
    /// @return The withdrawal request for the given address.
    function withdrawals(address owner) external view returns (uint256);

    /// @notice Unlocks withdrawals for the sender's account, after a time delay.
    function unlock() external;

    /// @notice Allows the owner to recover from error cases by pulling ETH out of the contract.
    /// @param wad The amount of WETH to recover.
    function recover(uint256 wad) external;

    /// @notice Allows the owner to recover from error cases by pulling ETH from a specific owner.
    /// @param guy The address to recover the WETH from.
    /// @param wad The amount of WETH to recover.
    function hold(address guy, uint256 wad) external;
}
