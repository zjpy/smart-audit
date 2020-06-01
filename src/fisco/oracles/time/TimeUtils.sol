pragma solidity ^0.4.24;


library TimeUtils {
    /// @dev 将一个字符串转换为uint256类型的整型数值.
    /// @param s 字符串.
    /// @return result 返回uint256类型的整型数值
    /// @return hasError 是否在转换中出现错误
    function stringToUint(string s)
        public
        constant
        returns (uint256 result, bool hasError)
    {
        bytes memory b = bytes(s);
        result = 0;
        hasError = false;

        uint256 oldResult = 0;
        for (uint256 i = 0; i < b.length; i++) {
            if (b[i] >= 48 && b[i] <= 57) {
                // 将旧的数值存储下来，然后检查是否溢出
                oldResult = result;
                result = result * 10 + (uint256(b[i]) - 48);
                if (oldResult > result) {
                    hasError = true;
                    break;
                }
            } else {
                hasError = true;
                break;
            }
        }
        return (result, hasError);
    }
}
