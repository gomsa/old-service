# old-service 数据库动作解释
*** 不确定解释
## 下载商品信息到一楼收款区
### 查找一楼收款区配置
```
select AreaCode, AreaName, ExchgPath from tXsSkArea where AreaCode='0001' 
```
### 更改启动的pos(***可能锁死pos状态禁止在用管理的下载)
```
update tXsPosState set IsChgPlu='0' where PosCode in  (select PosCode from tXsPos where IsEnable="1" and AreaCode in   ("","0001" ) ) 
```
### 更新打包商品价格
```
update a set a.SPrice =ROUND(b.PluCount*c.SPrice*a.PackPriceRate/100,  1),               a.HyPrice=ROUND(b.PluCount*c.HyPrice*a.PackPriceRate/100, 1),               a.PfPrice=ROUND(b.PluCount*c.PfPrice*a.PackPriceRate/100, 1)   from tBmPlu a, tBmPluPack b, tBmPlu c  where a.PluCode=b.PackCode and a.PluType='1' and a.PluOrigin='3' and a.PackPriceType='1' and b.PluCode=c.PluCode 
```
### 查询不是销售部门的商品 DepType 1
```
select a.PluCode from tBmPlu a, tBmDept b  where a.DepCode=b.DepCode and b.DepType='1' 
### 
```
### 查询商品
```
select PluCode, PluName, PluAbbrName, BarCode, Unit, Spec,  DepCode, SupCode, ClsCode, BrandCode, SPrice,  HyPrice, PfPrice, XTaxRate, MgType,  PluType, IsWeight, IsDecimal, YhType, YhEndDate, YhPrice, PluStatus  from tBmPlu where PluStatus<>"2" and PluStatus<>"3" and PluStatus<>"4" and PluStatus<>"A" and PluStatus<>"B" 
```
### 查询一品多条码信息
```
select BarCode from  (select BarCode   from tBmPlu where (BarCode <> "") and (BarCode is not Null)   union all   select b.BarCode   from tBmPlu a, tBmMulBar b where a.PluCode=b.PluCode) a  group by BarCode Having count(BarCode)>1 
```
### 关联条形码查询商品
```
select BarCode, PluCode, PluName, PluAbbrName, Unit, Spec,  DepCode, SupCode, ClsCode, BrandCode, SPrice,  HyPrice, PfPrice, XTaxRate, MgType,  PluType, IsWeight, IsDecimal, YhType, YhEndDate, YhPrice, PluStatus  from tBmPlu where (BarCode <> "") and (BarCode is not Null) and (PluStatus<>"2" and PluStatus<>"3" and PluStatus<>"4" and PluStatus<>"A" and PluStatus<>"B") 
 union  select b.BarCode, b.PluCode, b.PluName, b.PluAbbrName, a.Unit, b.Spec,  a.DepCode, a.SupCode, a.ClsCode, a.BrandCode, a.SPrice,  a.HyPrice, a.PfPrice, a.XTaxRate, a.MgType,  a.PluType, a.IsWeight, a.IsDecimal, a.YhType, a.YhEndDate, a.YhPrice, a.PluStatus  from tBmPlu a, tBmMulBar b where a.PluCode=b.PluCode and (a.PluStatus<>"2" and a.PluStatus<>"3" and a.PluStatus<>"4" and a.PluStatus<>"A" and a.PluStatus<>"B") 
```
### 更改POS状态锁定
```
update tXsPosState set IsChgPlu='1' where PosCode in  (select PosCode from tXsPos where IsEnable="1" and AreaCode in   ("","0001" ) ) 
```
### 未知无数据的查询
```
select PluCode, BgnDate, EndDate, BgnTime, EndTime, LtNumber, LtSumNumber, XsNumber, YhPrice  from tBmPtLimitPlu 
select PluCode, SerialNo, PluName, YhType, MinNumber, MaxNumber, PfPrice  from tBmPosPfYhGrade 
select PluCode, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), SPrice, YhPrice  from tYhHalfPlu  where EndDate>='20191021 00:00:00.000' 
select PluCode, YhPluCode, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), SPrice, YhPrice  from tYhBuyAYhB  where EndDate>='20191021 00:00:00.000' 
select PluCode, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), SPrice, HyPrice, FlXsAmt, BuyType  from tYhVipBuyPlu  where EndDate>='20191021 00:00:00.000' 
select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), PluCode, BrandCode, XsNumber  from tYhBuyFreeMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode, ZsNumber  from tYhBuyFreeDetail  where Dno in (select Dno from tYhBuyFreeMaster where EndDate>='20191021 00:00:00.000') select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), ExType, XsNumber, XsAmt, PluCode, ZsNumber  from tYhMixFreeMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode  from tYhMixFreeDetail  where Dno in (select Dno from tYhMixFreeMaster where EndDate>='20191021 00:00:00.000') select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), ExType, XsNumber, XsAmt, SsAmt  from tYhMixBuyMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode  from tYhMixBuyDetail  where Dno in (select Dno from tYhMixBuyMaster where EndDate>='20191021 00:00:00.000') 
select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), PluCode, ZsNumber  from tYhCombinFreeMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode, XsNumber  from tYhCombinFreeDetail  where Dno in (select Dno from tYhCombinFreeMaster where EndDate>='20191021 00:00:00.000') 
select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), SsAmt, Remark  from tYhCombinBuyMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode, XsNumber, SPrice, YhPrice, SAmt, YhAmt, Percents  from tYhCombinBuyDetail  where Dno in (select Dno from tYhCombinBuyMaster where EndDate>='20191021 00:00:00.000') 
select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), XsNumber  from tYhMixCxMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode, SPrice, YhPrice  from tYhMixCxDetail  where Dno in (select Dno from tYhMixCxMaster where EndDate>='20191021 00:00:00.000') 
select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), CxType, XsAmt, YhAmt, Remark  from tYhExcessCxMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, GrpType, GrpCode  from tYhExcessCxDetail  where Dno in (select Dno from tYhExcessCxMaster where EndDate>='20191021 00:00:00.000') 
select Dno, BgnDate=convert(char(10),BgnDate,120), EndDate=convert(char(10),EndDate,120),  BgnTime=convert(char(8),BgnTime,108), EndTime=convert(char(8),EndTime,108), XsAmt, BuyAmt, MxType, Remark  from tYhExcessBuyMaster  where EndDate>='20191021 00:00:00.000' 
select Dno, PluCode, XsNumber, SPrice, YhPrice, SAmt, YhAmt, Percents  from tYhExcessBuyDetail  where Dno in (select Dno from tYhExcessBuyMaster where EndDate>='20191021 00:00:00.000') 
```