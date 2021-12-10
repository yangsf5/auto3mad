// @ts-ignore
/* eslint-disable */

declare namespace API {
  type CurrentUser = {
    data?: { passed_day?: number; weekend?: number; holiday?: number; adapter?: number; ret?: number };
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };
}
