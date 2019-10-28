# old-sql 超导收银系统数据库对接
*** 不确定解释
## 数据库
### tXsSkArea   交换区
    - AreaCode 收款区id
    - AreaName 收款区路径
    - ExchgPath 交换区路径
### tXsPosState Pos状态
    - PosCode POSID
    - IsChgPlu (***POS商品数据有更新)
### tXsPos POS状态
    - PosCode POSID
    - AreaCode 收款区id
    - IsEnable 是否启用
### tBmPlu 商品
    - PluCode 自编码
    - PluName 商品名称
    - PluAbbrName 别名
    - BarCode 条码
    - Unit 单位
    - Spec 规格
    - Weigt 重量
    - DepCode 部门ID
    - Produce 国家税收分类编码(原产地)
    - Grade 国家税收分类名(原等级)
    - SupCode 供应商ID
    - ClsCode 品类ID
    - BrandCode 品牌ID
    - HJPrice 含税进价
    - WJPrice 未税进价
    - SPrice 零售价
    - HyPrice 会员价
    - PfPrice 批发价
    - JTaxRate 进项税率
    - XTaxRate 销项税率
    - CgyCode 采购员ID
    - CgyName 采购员名称
    - XgDate 修改时间
    - LrDate 录入时间
    - UserCode 用户ID
    - UserName 用户名称
    - PackPriceType 打包类型
    - PackPriceRate 打包价格率
    - PluStatus 商品状态(0,1,2,3,4,A,B)[0新品试销 1正常 2禁用 3删除 4新增商品 5删除 6禁止下单 7清仓 8季节性禁止下单 9新品评估 A暂停零售 B禁止零售]

    
### tBmPluPack 打包商品
    - PackCode 包自编码
    - PluCode 商品自编码
    - pluName 商品名称
    - Unit 单位
    - Spec 规格
    - SPrice 零售价
    - ColorCode
    - pluCount 数量
    - PkTotal 打包价
    - Percents
    - IsDecimal
    - MgType
### tBmMulBar 商品条码信息
    - BarCode 条码
    - PluCode 商品编码
    - DepCode 部门编码
    - PluName 商品名称
    - PluAbbrName 商品别名
    - ColorCode
    - SizeCode
    - Spec 规格
### tBmDept  部门
    - DepCode 部门编码
    - DepName 部门名称
    - DepType 部门类型(0-销售部门,1-其他部门)
    - IsLast (***层级)
    - DepArea 部门编辑
    - ReMark 部门说明
