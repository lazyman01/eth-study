## EIP-55: Mixed-case checksum address encoding提案

EIP-55规定：在将账户地址转换为十六进制字符时，当十六禁止地址的第i位时字母（即abcdef之一）时：

如果小写的十六禁止地址的哈希值的第i*4位比特为1，则大写形式打印，否则小写形式打印。

