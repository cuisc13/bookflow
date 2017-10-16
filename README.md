# Book Flow

![图书漂流](https://upload.wikimedia.org/wikipedia/commons/e/e4/BookcrossingLyon.jpg)图片来自[Wikipedia 图书漂流 词条](https://zh.wikipedia.org/wiki/%E5%9B%BE%E4%B9%A6%E6%BC%82%E6%B5%81)

# [文档](docs/SUMMARY.md)

图书漂流管理系统, 尝试性项目， 利用一些新学的技术搭建一个管理 陌生人之间图书借阅（漂流）的系统。

使用SSDB数据库保存数据。使用类似区块链的数据结构保存图书漂流交易历史数据。

每本书维护一个数据链，每发生一次交易，记录图书的去向，以及该本数上一本书的交易。保证交易能够链式保存。

现支持功能：

1. 添加、删除图书
1. 查看图书（列表、根据ID）
1. 发起、确定交易
1. 查看交易（列表、根据图书ID）
