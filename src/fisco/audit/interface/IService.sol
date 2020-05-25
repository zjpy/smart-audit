pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;


// 用于封装oracle服务的接口，所有相关合约都必须实现自该接口
interface IService {
    /// @dev 注册一个规则项.
    /// @param args 注册规则参数，以数组形式表示.
    /// @return ruleID 返回在预言机服务中注册后对应的规则ID
    /// @return errorCode 错误码，如果为0则表示没有错误，否则发生注册错误.
    /// @return message 返回结果信息.
    function register(string[] args)
        public
        returns (
            uint32 ruleID,
            uint32 errorCode,
            string memory message
        );

    /// @dev 验证一个规则项.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param args 验证规则所需的值，以数组形式表示.
    /// @return errorCode 错误码，如果为0则表示没有错误，否则发生注册错误.
    /// @return message 返回结果信息.
    function validate(uint32 ruleID, string[] args)
        public
        returns (uint32 errorCode, string memory message);
}
