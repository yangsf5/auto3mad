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
        name: 'OneTab Edit',
        icon: '',
        path: '/url/one-tab-edit',
        component: './url/one-tab-edit',
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
