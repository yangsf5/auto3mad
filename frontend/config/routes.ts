export default [
  {
    path: '/daily',
    icon: 'aim',
    name: 'Daily',
    routes: [
      {
        name: '日拱记录',
        icon: '',
        path: '/daily/event',
        component: './daily/event',
      },
      {
        name: '日拱统计',
        icon: '',
        path: '/daily/stat',
        component: './daily/stat',
      },
    ],
  },
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
    path: '/day',
    icon: 'calendar',
    name: 'Day',
    routes: [
      {
        name: '时间戳',
        icon: '',
        path: '/day/time',
        component: './day/time',
      },
      {
        name: '纪念日',
        icon: '',
        path: '/day/memorial',
        component: './day/memorial',
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
