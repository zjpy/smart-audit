pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./interface/IService.sol";
import "./Utils.sol";


// 该智能合约用于模拟时间服务的预言机服务调用
contract DummyTimeService is IService {
    uint32 nonce;
    uint32 dummyID;

    /// @dev 构造方法，目前仅用于初始化Nonce值.
    constructor() public {
        // 这里初始化nonce值
        nonce = 10;
        dummyID = 0;
    }

    /// @dev 注册一个规则项.
    /// @param args 注册规则参数，以数组形式表示.
    /// @return ruleID 返回在预言机服务中注册后对应的规则ID
    function register(string[] args) public returns (uint32 ruleID) {
        // 由于假设是固定的验证规则，所以这里不需要额外工作
        ruleID = dummyID;
        dummyID++;
        return ruleID;
    }

    /// @dev 验证一个规则项.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param args 验证规则所需的值，以数组形式表示.
    function validate(uint32 ruleID, string[] args) public {
        // 由于在solidity语言对时间支持的不是很好，这里不再模拟真实输入时间是否落入给定
        //    的时间范围，而是简单让90%的情况通过验证
        if (nonce % 10 >= 9) {
            require(false, "时间超出正常工作范围");
        }
    }
}
