// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import {
    L1ArbitrumExtendedGateway,
    L1ArbitrumGateway,
    IL1ArbitrumGateway,
    ITokenGateway,
    TokenGateway,
    IERC20
} from "./L1ArbitrumExtendedGateway.sol";
import {L2USDCCustomGateway} from "../../arbitrum/gateway/L2USDCCustomGateway.sol";

/**
 * @title Custom gateway for USDC implementing Bridged USDC Standard.
 * @notice Reference to the Circle's Bridged USDC Standard:
 *         https://github.com/circlefin/stablecoin-evm/blob/master/doc/bridged_USDC_standard.md
 *
 * @dev    This contract can be used on new Orbit chains which want to provide USDC
 *         bridging solution and keep the possibility to upgrade to native USDC at
 *         some point later. This solution will NOT be used in existing Arbitrum chains.
 *
 *         Child chain custom gateway to be used along this parent chain custom gateway is L2USDCCustomGateway.
 *         This custom gateway differs from standard gateway in the following ways:
 *         - it supports a single parent chain - child chain USDC token pair
 *         - it is ownable
 *         - owner can one-time permanently pause deposits
 *         - owner can trigger burning all the USDC tokens locked in the gateway
 */
contract L1USDCCustomGateway is L1ArbitrumExtendedGateway {
    address public l1USDC;
    address public l2USDC;
    address public owner;
    bool public depositsPaused;

    event DepositsPaused();
    event GatewayUsdcBurned(uint256 amount);

    error L1USDCCustomGateway_DepositsAlreadyPaused();
    error L1USDCCustomGateway_DepositsPaused();
    error L1USDCCustomGateway_DepositsNotPaused();
    error L1USDCCustomGateway_InvalidL1USDC();
    error L1USDCCustomGateway_InvalidL2USDC();
    error L1USDCCustomGateway_NotOwner();
    error L1USDCCustomGateway_InvalidOwner();

    modifier onlyOwner() {
        if (msg.sender != owner) {
            revert L1USDCCustomGateway_NotOwner();
        }
        _;
    }

    function initialize(
        address _l2Counterpart,
        address _l1Router,
        address _inbox,
        address _l1USDC,
        address _l2USDC,
        address _owner
    ) public {
        if (_l1USDC == address(0)) {
            revert L1USDCCustomGateway_InvalidL1USDC();
        }
        if (_l2USDC == address(0)) {
            revert L1USDCCustomGateway_InvalidL2USDC();
        }
        if (_owner == address(0)) {
            revert L1USDCCustomGateway_InvalidOwner();
        }
        L1ArbitrumGateway._initialize(_l2Counterpart, _l1Router, _inbox);
        l1USDC = _l1USDC;
        l2USDC = _l2USDC;
        owner = _owner;
    }

    /**
     * @notice Pauses deposits. This can only be called by the owner.
     * @dev    Pausing is permanent and can't be undone. Pausing is prerequisite for burning escrowed USDC tokens.
     *         Incoming withdrawals are not affected. Pausing the withdrawals needs to be done separately on the child chain.
     */
    function pauseDeposits() external onlyOwner {
        if (depositsPaused) {
            revert L1USDCCustomGateway_DepositsAlreadyPaused();
        }
        depositsPaused = true;

        emit DepositsPaused();
    }

    /**
     * @notice Burns the USDC tokens escrowed in the gateway.
     * @dev    Can be called by owner after deposits are paused.
     *         Function signature complies by Bridged USDC Standard.
     */
    function burnLockedUSDC() external onlyOwner {
        if (!depositsPaused) {
            revert L1USDCCustomGateway_DepositsNotPaused();
        }
        uint256 gatewayBalance = IERC20(l1USDC).balanceOf(address(this));
        Burnable(l1USDC).burn(gatewayBalance);

        emit GatewayUsdcBurned(gatewayBalance);
    }

    /**
     * @notice Sets a new owner.
     */
    function setOwner(address newOwner) external onlyOwner {
        if (newOwner == address(0)) {
            revert L1USDCCustomGateway_InvalidOwner();
        }
        owner = newOwner;
    }

    /**
     * @notice entrypoint for depositing USDC, can be used only if deposits are not paused.
     */
    function outboundTransferCustomRefund(
        address _l1Token,
        address _refundTo,
        address _to,
        uint256 _amount,
        uint256 _maxGas,
        uint256 _gasPriceBid,
        bytes calldata _data
    ) public payable override returns (bytes memory res) {
        if (depositsPaused) {
            revert L1USDCCustomGateway_DepositsPaused();
        }
        return super.outboundTransferCustomRefund(
            _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data
        );
    }

    /**
     * @notice only parent chain - child chain USDC token pair is supported
     */
    function calculateL2TokenAddress(address l1ERC20)
        public
        view
        override(ITokenGateway, TokenGateway)
        returns (address)
    {
        if (l1ERC20 != l1USDC) {
            // invalid L1 USDC address
            return address(0);
        }
        return l2USDC;
    }
}

interface Burnable {
    /**
     * @notice Circle's referent USDC implementation exposes burn function of this signature.
     */
    function burn(uint256 _amount) external;
}
