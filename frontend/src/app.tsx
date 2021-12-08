import { LinkOutlined } from '@ant-design/icons';
//import { PageLoading } from '@ant-design/pro-layout';
import type { RunTimeLayoutConfig } from 'umi';
import { history, Link, request as request_umi } from 'umi'

const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/';

async function queryGlobal() {
  return request_umi<API.CurrentUser>('/v2/day/countdowns', {
    method: 'GET'
  });
}

/** 获取用户信息比较慢的时候会展示一个 loading */
export const initialStateConfig = {
  //loading: <PageLoading />,
};

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
  //settings?: Partial<LayoutSettings>;
  currentUser?: API.CurrentUser;
}> {
  const fetchUserInfo = async () => {
    try {
      const currentUser = await queryGlobal();
      return currentUser;
    } catch (error) {
      history.push(loginPath);
    }
    return undefined;
  };

  const currentUser = await fetchUserInfo();
  return {
    currentUser: currentUser,
    //settings: {},
  };
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({ initialState }) => {
  return {
    title: 'Auto 3Mad',
    rightContentRender: () => (
      <span>日拱应该 {initialState.currentUser.data.ret}</span>
    ),
    disableContentMargin: false,
    links: isDev
      ? [
        <Link to="/umi/plugin/openapi" target="_blank">
          <LinkOutlined />
          <span>OpenAPI 文档</span>
        </Link>,
      ]
      : [],
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    // childrenRender: (children) => {
    //   if (initialState.loading) return <PageLoading />;
    //   return children;
    // },
    ...initialState?.settings,
  };
};
