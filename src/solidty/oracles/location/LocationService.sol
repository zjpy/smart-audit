pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./interface/IService.sol";


// 该智能合约用于封装定位相关的预言机服务
contract LocationService is IService {
    constructor() public {
        // 初始化定位预言机服务相关信息……
    }

    /// @dev 注册一个规则项.
    /// @param args 注册规则参数，以数组形式表示.
    /// @return ruleID 返回在预言机服务中注册后对应的规则ID
    function register(string[] args) public returns (uint32 ruleID) {
        require(args.length > 0, "注册规则所需参数不足");

        // fixme 实际商用时实现定位服务预言机，然后在这里调用预言机服务
        return 0;
    }

    /// @dev 验证一个规则项.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param args 验证规则所需的值，以数组形式表示.
    function validate(uint32 ruleID, string[] args) public {
        require(args.length > 0, "注册规则所需参数不足");

        // fixme 实际商用时实现定位服务预言机，然后在这里调用预言机服务
    }
}
