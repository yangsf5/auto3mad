export default [
  {
    path: '/url',
    icon: 'link',
    name: 'URL',
    routes: [
      {
        name: 'OneTab',
        icon: '',
        path: '/url/one-tab',
        component: './url/one-tab',
      },
      {
        name: 'Backend API',
        icon: '',
        path: '/url/api',
        component: './url/api',
      },
    ],
  },
  {
    path: '/',
    redirect: '/url/one-tab',
  },
  {
    component: './404',
  },
];
