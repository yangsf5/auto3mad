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
