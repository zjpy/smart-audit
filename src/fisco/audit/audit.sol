pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./Rules.sol";

contract Audit is Rules {

    /// @dev 规则合约的构造方法，这里实现Rules.
    /// @param face 人脸识别服务合约地址.
    /// @param identify 物体识别服务合约地址.
    /// @param time 时间服务合约地址.
    /// @param location 定位服务合约地址.
    constructor(
        address face,
        address identify,
        address time,
        address location
    ) Rules (face, identify, time, location) {

    }
}