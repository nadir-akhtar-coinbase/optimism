// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "src/universal/ISemver.sol";

import { IDelayedWETH } from "src/dispute/interfaces/IDelayedWETH.sol";
import { IWETH } from "src/dispute/interfaces/IWETH.sol";
import { WETH98 } from "src/dispute/weth/WETH98.sol";

import { SuperchainConfig } from "src/L1/SuperchainConfig.sol";

/// @title DelayedWETH
/// @notice DelayedWETH is an extension to WETH9 that allows for delayed withdrawals. Accounts must
///         trigger an unlock function before they can withdraw their WETH. Accounts can trigger
///         the unlock function at any time, but must wait a delay period before they can withdraw
///         after the unlock function is triggered. DelayedWETH is designed to be used by the
///         DisputeGame contracts where unlock will only be triggered after a dispute is resolved.
///         DelayedWETH is designed to sit behind a proxy contract and has an owner address that
///         can pull WETH from any account and can recover ETH from the contract itself.
///         Variable and function naming vaguely follows the vibe of WETH9.
contract DelayedWETH is OwnableUpgradeable, WETH98, IDelayedWETH, ISemver {
    /// @notice Semantic version.
    /// @custom:semver 0.2.0
    string public constant version = "0.2.0";

    /// @inheritdoc IDelayedWETH
    mapping(address => uint256) public withdrawals;

    /// @notice Withdrawal delay in seconds.
    uint256 internal immutable DELAY_SECONDS;

    /// @notice Address of the SuperchainConfig contract.
    SuperchainConfig public config;

    /// @param wait The delay for withdrawals in seconds.
    constructor(uint256 wait) {
        DELAY_SECONDS = wait;
        initialize({ owner: address(0), cfg: SuperchainConfig(address(0)) });
    }

    /// @notice Initializes the contract.
    /// @param owner The address of the owner.
    /// @param cfg Address of the SuperchainConfig contract.
    function initialize(address owner, SuperchainConfig cfg) public initializer {
        __Ownable_init();
        _transferOwnership(owner);
        config = cfg;
    }

    /// @inheritdoc IDelayedWETH
    function delay() external view returns (uint256) {
        return DELAY_SECONDS;
    }

    /// @inheritdoc IDelayedWETH
    function unlock() external {
        require(!config.paused());
        withdrawals[msg.sender] = block.timestamp;
    }

    /// @inheritdoc IWETH
    function withdraw(uint256 wad) public override(WETH98, IWETH) {
        require(!config.paused());
        require(withdrawals[msg.sender] > 0);
        require(withdrawals[msg.sender] + DELAY_SECONDS <= block.timestamp);
        super.withdraw(wad);
    }

    /// @inheritdoc IDelayedWETH
    function recover(uint256 wad) external {
        require(msg.sender == owner());
        require(address(this).balance >= wad);
        payable(msg.sender).transfer(wad);
    }

    /// @inheritdoc IDelayedWETH
    function hold(address guy, uint256 wad) external {
        require(msg.sender == owner());
        allowance[msg.sender][guy] = wad;
        emit Approval(msg.sender, guy, wad);
    }
}
