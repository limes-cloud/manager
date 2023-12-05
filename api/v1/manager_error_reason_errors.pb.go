// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotFoundError.String() && e.Code == 200
}

func NotFoundErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotFoundError.String(), "数据不存在:"+fmt.Sprintf(format, args...))
}

func NotFoundError() *errors.Error {
	return errors.New(200, ErrorReason_NotFoundError.String(), "数据不存在")
}

func IsDatabaseError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DatabaseError.String() && e.Code == 200
}

func DatabaseErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DatabaseError.String(), "数据库错误:"+fmt.Sprintf(format, args...))
}

func DatabaseError() *errors.Error {
	return errors.New(200, ErrorReason_DatabaseError.String(), "数据库错误")
}

func IsMetadataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MetadataError.String() && e.Code == 200
}

func MetadataErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_MetadataError.String(), "元数据异常:"+fmt.Sprintf(format, args...))
}

func MetadataError() *errors.Error {
	return errors.New(200, ErrorReason_MetadataError.String(), "元数据异常")
}

func IsTransformError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TransformError.String() && e.Code == 200
}

func TransformErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_TransformError.String(), "数据转换失败:"+fmt.Sprintf(format, args...))
}

func TransformError() *errors.Error {
	return errors.New(200, ErrorReason_TransformError.String(), "数据转换失败")
}

func IsDepartmentPermissionsError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DepartmentPermissionsError.String() && e.Code == 200
}

func DepartmentPermissionsErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DepartmentPermissionsError.String(), "无此部门权限:"+fmt.Sprintf(format, args...))
}

func DepartmentPermissionsError() *errors.Error {
	return errors.New(200, ErrorReason_DepartmentPermissionsError.String(), "无此部门权限")
}

func IsRolePermissionsError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RolePermissionsError.String() && e.Code == 200
}

func RolePermissionsErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_RolePermissionsError.String(), "无此角色权限:"+fmt.Sprintf(format, args...))
}

func RolePermissionsError() *errors.Error {
	return errors.New(200, ErrorReason_RolePermissionsError.String(), "无此角色权限")
}

func IsUserPermissionsError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UserPermissionsError.String() && e.Code == 200
}

func UserPermissionsErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UserPermissionsError.String(), "无此用户权限:"+fmt.Sprintf(format, args...))
}

func UserPermissionsError() *errors.Error {
	return errors.New(200, ErrorReason_UserPermissionsError.String(), "无此用户权限")
}

func IsMenuPermissionsError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MenuPermissionsError.String() && e.Code == 200
}

func MenuPermissionsErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_MenuPermissionsError.String(), "无此菜单权限:"+fmt.Sprintf(format, args...))
}

func MenuPermissionsError() *errors.Error {
	return errors.New(200, ErrorReason_MenuPermissionsError.String(), "无此菜单权限")
}

func IsEditSystemDataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_EditSystemDataError.String() && e.Code == 200
}

func EditSystemDataErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_EditSystemDataError.String(), "系统数据，不允许修改:"+fmt.Sprintf(format, args...))
}

func EditSystemDataError() *errors.Error {
	return errors.New(200, ErrorReason_EditSystemDataError.String(), "系统数据，不允许修改")
}

func IsDeleteSystemDataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteSystemDataError.String() && e.Code == 200
}

func DeleteSystemDataErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DeleteSystemDataError.String(), "系统数据，不允许删除:"+fmt.Sprintf(format, args...))
}

func DeleteSystemDataError() *errors.Error {
	return errors.New(200, ErrorReason_DeleteSystemDataError.String(), "系统数据，不允许删除")
}

func IsNewCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NewCaptchaError.String() && e.Code == 200
}

func NewCaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NewCaptchaError.String(), "二维码生成失败:"+fmt.Sprintf(format, args...))
}

func NewCaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_NewCaptchaError.String(), "二维码生成失败")
}

func IsVerifyCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_VerifyCaptchaError.String() && e.Code == 200
}

func VerifyCaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_VerifyCaptchaError.String(), "验证码验证失败:"+fmt.Sprintf(format, args...))
}

func VerifyCaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_VerifyCaptchaError.String(), "验证码验证失败")
}

func IsRsaDecodeError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RsaDecodeError.String() && e.Code == 200
}

func RsaDecodeErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_RsaDecodeError.String(), "rsa解密失败:"+fmt.Sprintf(format, args...))
}

func RsaDecodeError() *errors.Error {
	return errors.New(200, ErrorReason_RsaDecodeError.String(), "rsa解密失败")
}

func IsPasswordFormatError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PasswordFormatError.String() && e.Code == 200
}

func PasswordFormatErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_PasswordFormatError.String(), "密码格式错误:"+fmt.Sprintf(format, args...))
}

func PasswordFormatError() *errors.Error {
	return errors.New(200, ErrorReason_PasswordFormatError.String(), "密码格式错误")
}

func IsPasswordExpireError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PasswordExpireError.String() && e.Code == 200
}

func PasswordExpireErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_PasswordExpireError.String(), "密码已过期:"+fmt.Sprintf(format, args...))
}

func PasswordExpireError() *errors.Error {
	return errors.New(200, ErrorReason_PasswordExpireError.String(), "密码已过期")
}

func IsUserDisableError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UserDisableError.String() && e.Code == 200
}

func UserDisableErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UserDisableError.String(), "用户已被禁用:"+fmt.Sprintf(format, args...))
}

func UserDisableError() *errors.Error {
	return errors.New(200, ErrorReason_UserDisableError.String(), "用户已被禁用")
}

func IsRoleDisableError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RoleDisableError.String() && e.Code == 200
}

func RoleDisableErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_RoleDisableError.String(), "角色已被禁用:"+fmt.Sprintf(format, args...))
}

func RoleDisableError() *errors.Error {
	return errors.New(200, ErrorReason_RoleDisableError.String(), "角色已被禁用")
}

func IsUserPasswordError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UserPasswordError.String() && e.Code == 200
}

func UserPasswordErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UserPasswordError.String(), "用户密码错误:"+fmt.Sprintf(format, args...))
}

func UserPasswordError() *errors.Error {
	return errors.New(200, ErrorReason_UserPasswordError.String(), "用户密码错误")
}

func IsNewTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NewTokenError.String() && e.Code == 200
}

func NewTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NewTokenError.String(), "token生成失败:"+fmt.Sprintf(format, args...))
}

func NewTokenError() *errors.Error {
	return errors.New(200, ErrorReason_NewTokenError.String(), "token生成失败")
}

func IsParseTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ParseTokenError.String() && e.Code == 200
}

func ParseTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_ParseTokenError.String(), "token解析失败:"+fmt.Sprintf(format, args...))
}

func ParseTokenError() *errors.Error {
	return errors.New(200, ErrorReason_ParseTokenError.String(), "token解析失败")
}

func IsRefreshTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RefreshTokenError.String() && e.Code == 401
}

func RefreshTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败:"+fmt.Sprintf(format, args...))
}

func RefreshTokenError() *errors.Error {
	return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败")
}

func IsEmptyTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_EmptyTokenError.String() && e.Code == 200
}

func EmptyTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_EmptyTokenError.String(), "token不能为空:"+fmt.Sprintf(format, args...))
}

func EmptyTokenError() *errors.Error {
	return errors.New(200, ErrorReason_EmptyTokenError.String(), "token不能为空")
}

func IsDeleteSelfRoleError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteSelfRoleError.String() && e.Code == 200
}

func DeleteSelfRoleErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DeleteSelfRoleError.String(), "不能删除自己的当前所属角色:"+fmt.Sprintf(format, args...))
}

func DeleteSelfRoleError() *errors.Error {
	return errors.New(200, ErrorReason_DeleteSelfRoleError.String(), "不能删除自己的当前所属角色")
}

func IsDeleteSelfDepartmentError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteSelfDepartmentError.String() && e.Code == 200
}

func DeleteSelfDepartmentErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DeleteSelfDepartmentError.String(), "不能删除自己的当前所属部门:"+fmt.Sprintf(format, args...))
}

func DeleteSelfDepartmentError() *errors.Error {
	return errors.New(200, ErrorReason_DeleteSelfDepartmentError.String(), "不能删除自己的当前所属部门")
}

func IsParentMenuError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ParentMenuError.String() && e.Code == 200
}

func ParentMenuErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_ParentMenuError.String(), "当前菜单的父菜单不能为自己:"+fmt.Sprintf(format, args...))
}

func ParentMenuError() *errors.Error {
	return errors.New(200, ErrorReason_ParentMenuError.String(), "当前菜单的父菜单不能为自己")
}

func IsSendEmailCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SendEmailCaptchaError.String() && e.Code == 200
}

func SendEmailCaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_SendEmailCaptchaError.String(), "发送邮箱验证码失败:"+fmt.Sprintf(format, args...))
}

func SendEmailCaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_SendEmailCaptchaError.String(), "发送邮箱验证码失败")
}

func IsUsernameFormatError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UsernameFormatError.String() && e.Code == 200
}

func UsernameFormatErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UsernameFormatError.String(), "用户名格式错误:"+fmt.Sprintf(format, args...))
}

func UsernameFormatError() *errors.Error {
	return errors.New(200, ErrorReason_UsernameFormatError.String(), "用户名格式错误")
}

func IsUsernameNotExistError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UsernameNotExistError.String() && e.Code == 200
}

func UsernameNotExistErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UsernameNotExistError.String(), "用户不存在:"+fmt.Sprintf(format, args...))
}

func UsernameNotExistError() *errors.Error {
	return errors.New(200, ErrorReason_UsernameNotExistError.String(), "用户不存在")
}

func IsDisableSelfUserError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DisableSelfUserError.String() && e.Code == 200
}

func DisableSelfUserErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DisableSelfUserError.String(), "不能禁用当前用户:"+fmt.Sprintf(format, args...))
}

func DisableSelfUserError() *errors.Error {
	return errors.New(200, ErrorReason_DisableSelfUserError.String(), "不能禁用当前用户")
}

func IsForbiddenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ForbiddenError.String() && e.Code == 200
}

func ForbiddenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_ForbiddenError.String(), "无接口权限:"+fmt.Sprintf(format, args...))
}

func ForbiddenError() *errors.Error {
	return errors.New(200, ErrorReason_ForbiddenError.String(), "无接口权限")
}
