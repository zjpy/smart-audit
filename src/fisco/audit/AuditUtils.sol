pragma solidity ^0.4.24;


library AuditUtils {
    /// @dev 比较两个字符串是否相等.
    /// @param a 第一个字符串.
    /// @param b 第二个字符串.
    /// @return 相等则返回true，否则返回false.
    function compareStrings(string memory a, string memory b)
        public
        pure
        returns (bool)
    {
        return (keccak256(abi.encodePacked((a))) ==
            keccak256(abi.encodePacked((b))));
    }
}
