pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;


library Definition {
    // 通过经纬度表示的地理位置
    struct Position {
        // 纬度，这里以uint64表示浮点数类型
        uint64 Lat;
        // 经度，这里以uint64表示浮点数类型
        uint64 Lon;
    }

    // 定义用于验证是否在某个范围内的结构
    struct ValidateArea {
        // 验证范围的中心地理位置
        Position Center;
        // 范围半径，以米为单位
        uint64 Radius;
    }
}
