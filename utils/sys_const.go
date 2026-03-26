package utils

const USER_STATUS_FORBIDDEN int8 = 0 //用户禁用状态
const USER_STATUS_NORMAL int8 = 1    //用户正常状态

const BIZ_APPLY_STATUS_ACCEPT int = 3  //企业认证通过
const BIZ_APPLY_STATUS_REJECT int = 2  //企业认证拒绝
const BIZ_APPLY_STATUS_DEFAULT int = 1 //企业认证待处理

const ORDER_NOT_PAY_STATUS int = 1    //订单未支付状体
const ORDER_HAS_PAY_STATUS int = 3    //订单已支付状体
const ORDER_NO_REFUND_STATUS int = 1  //订单无退款
const ORDER_HAS_REFUND_STATUS int = 3 //订单有退款

const ORDER_ITEM_NOT_PROCESS int = 1 //订单商品未处理
const ORDER_ITEM_HAS_PROCESS int = 3 //订单商品已处理

const PROD_TYPE_AD int = 1                 //信息栏类型
const PROD_TYPE_PROFILE_COUNT int = 2      //简历查看类型
const PROD_TYPE_REFRESH_COUNT int = 3      //职位刷新类型
const PROD_TYPE_LIVE_REFRESH_COUNT int = 4 //直播刷新类型
const PROD_TYPE_LIVE_COUNT int = 5         //招聘会类型
