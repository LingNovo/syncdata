<?xml version="1.0" encoding="utf-8" ?>
<SyncConfig RootOrgId="eb75d4e4-355a-434b-b7dd-fe7f98d30128" IsPaddingDept="false" IsSyncTable="false" IsPullData = "false"  IsSyncData="true">
  <Table Source="Ding_Department" Target="Sys_Organize" IsClear = "true">
      <Column Source="Id" Target="F_UnionId"></Column>
      <Column Source="Name" Target="F_FullName"></Column>
      <Column Source="ShortName" Target="F_ShortName" Unique="true"></Column>
      <Column Source="Order" Target="F_EnCode"></Column>
      <Column Target="F_ParentId" DefaultType="fn_dept_parentid" DefaultType="fn"></Column>
      <Column Target="F_CategoryId" DefaultType="fn_dept_categoryid" DefaultType="fn"></Column>
      <Column Target="F_EnabledMark" DefaultValue="true" DefaultType="bool"></Column>
  </Table>
  <Table Source="Ding_User" Target="Sys_User" IsClear = "true">
      <Column Source="Unionid" Target="F_UnionId"></Column>
      <Column Target="F_Account" DefaultValue="fn_account" DefaultType="fn"></Column>
      <Column Source="Name" Target="F_RealName" Unique="true"></Column>
      <Column Target="F_NickName" DefaultValue="fn_account" DefaultType="fn"></Column>
      <Column Source="HiredDate" Target="F_RuZhiDate"></Column>
      <Column Target="F_Gender" DefaultValue="true" DefaultType="bool"></Column>
      <Column Source="Mobile" Target="F_MobilePhone"></Column>
      <Column Source="Email" Target="F_Email"></Column>
      <Column Target="F_WeChat" DefaultValue="fn_account" DefaultType="fn"></Column>
      <Column Target="F_RoleId" DefaultValue="bf71db13-de68-45ca-adfb-140496ba2003"></Column>
      <Column Target="F_OrganizeId" DefaultValue="eb75d4e4-355a-434b-b7dd-fe7f98d30128"></Column>
      <Column Target="F_DepartmentId" DefaultValue="fn_user_deptid" DefaultType="fn"></Column>
      <Column Target="F_EnabledMark" DefaultValue="true" DefaultType="bool"></Column>
      <Column Target="IsPaiMing" DefaultValue="true" DefaultType="bool"></Column>
      <Column Source="Jobnumber" Target="F_JobNumber"></Column>
      <Column Source="DeptShortName" Target="F_DeptShortName"></Column>
  </Table>
  <Table Source="Sys_User" Target="Sys_UserLogOn">
      <Column Source="F_Id" Target="F_UserId" Unique="true"></Column>
      <Column Target="F_UserPassword" DefaultValue="44c35ab35cb0603e90d168642ca51fb6"></Column>
      <Column Target="F_UserSecretkey" DefaultValue="57d3031d6fc4a34d"></Column>
      <Column Source="Jobnumber" Target="F_JobNumber"></Column>
  </Table>
</SyncConfig>
