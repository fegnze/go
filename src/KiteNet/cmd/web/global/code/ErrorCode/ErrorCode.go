package ErrorCode

//ErrorCode 该web统一错误码
const (
	/**
	 * 错误码:成功0
	 */
	SUCCESS int = iota
	/**
	 * 错误码:密码错误1
	 */
	PasswordError
	/**
	 * 错误码:账号存在2
	 */
	AccountExists
	/**
	 * 错误码:平台验证失败3
	 */
	PlatformAuthenticationFailed
	/**
	 * 错误码:参数空4
	 */
	ParameterNull
	/**
	 * 错误码:参数错误5
	 */
	ParameterError
	/**
	 * 错误码:系统异常6
	 */
	UnexpectedError
	/**
	 * 错误码:未从服务器获得数据7
	 */
	FromServerNull
	/**
	 * 错误码:Web错误8
	 */
	WebError
	/**
	 * 错误码：停服9
	 */
	StopServer
	/**
	 * 错误码：账号不存在10
	 */
	AccountNotExists

	/**
	 * 错误码：兑换码不正确11
	 */
	CDKeyError

	/**
	 * 错误码：兑换码已被使用12
	 */
	CDKeyUsed

	/**
	 * 错误码：兑换码已过期13
	 */
	CDKeyExpired

	/**
	 * 错误码：兑换码超过领取次数14
	 */
	CDKeyTimesLimit

	/**
	 * 错误码：兑换码领取失败15
	 */
	CDKeyFailed

	/**
	 * 错误码：服务器维护中16
	 */
	ServerMaintain

	/**
	 * 错误码：订单生成失败17
	 */
	BillnoCreateFailed

	/**
	 * 错误码：订单号已存在18
	 */
	BillnoExist

	/**
	 * 错误码：支付异常19
	 */
	PayError

	/**
	 * 错误码：邮箱错误20
	 */
	MailError

	/**
	 * 错误码：邮箱已存在21
	 */
	MailExist

	/**
	 * 错误码：五星评价已评价22
	 */
	ReviewExist

	/**
	 * 错误码：邀请码获取失败23
	 */
	InviteCodeGetFailed

	/**
	 * 错误码：邀请码不存在24
	 */
	InviteCodeNotExists

	/**
	 * 错误码：邀请码已绑定过25
	 */
	InviteCodeBinded

	/**
	 * 错误码：邀请码积分不足26
	 */
	InviteCodeScoreLess

	/**
	 * 错误码：邀请码积分奖励已领取过27
	 */
	InviteCodeScoreRewardExists

	/**
	 * 错误码：邀请码无法绑定自己28
	 */
	InviteCodeNotBindMyself

	/**
	* 错误码: 操作数据库错误29
	*/
	DBUnexpectedError
)
