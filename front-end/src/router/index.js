import Vue from 'vue';
import store from '../store';
import VueRouter from 'vue-router';
// plugins
import Swal from 'sweetalert2';
import Toast from '../plugins/toast';
// views
import Register from '@/views/Register';
import Login from '@/views/Login';
import ToDo from '@/views/ToDo';
import Diary from '@/views/Diary';
import Blog from '@/views/Blog';
import BlogDetail from '@/views/BlogDetail';
import Explore from '@/views/Explore';
import Settings from '@/views/Settings';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    redirect: '/explore'
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: {
      requireLogin: false,
      requireAdmin: false
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: {
      requireLogin: false,
      requireAdmin: false
    }
  },
  {
    path: '/todo',
    name: 'ToDo',
    component: ToDo,
    meta: {
      requireLogin: true,
      requireAdmin: false
    }
  },
  {
    path: '/diary',
    name: 'Diary',
    component: Diary,
    meta: {
      requireLogin: true,
      requireAdmin: false
    }
  },
  {
    path: '/blog',
    name: 'Blog',
    component: Blog,
    meta: {
      requireLogin: true,
      requireAdmin: false
    }
  },
  {
    path: '/blog/:id',
    name: 'BlogDetail',
    component: BlogDetail,
    meta: {
      requireLogin: false,
      requireAdmin: false,
      isPublic: true
    }
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings,
    meta: {
      requireLogin: true,
      requireAdmin: false
    }
  },
  {
    path: '/explore',
    name: 'Explore',
    component: Explore,
    meta: {
      requireLogin: false,
      requireAdmin: false,
      isPublic: true
    }
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('../views/Admin'),
    meta: {
      requireLogin: true,
      requireAdmin: true
    }
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/About.vue'),
    meta: {
      requireLogin: true,
      requireAdmin: false
    }
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  if (to.matched.length === 0) {
    // 前往的路由不存在时
    Swal.fire({
      position: 'top',
      icon: 'error',
      title: '404, NOT FOUND!',
      showConfirmButton: false,
      timer: 1500
    });
    if (from.name) {
      next({
        name: from.name
      });
    } else {
      next({
        path: '/'
      });
    }
  } else if (!store.getters.isLoggedIn && to.meta.requireLogin) {
    // 1.用户未登录，但想访问需要认证的相关路由时，跳转到登录页
    Toast.fire({
      icon: 'error',
      title: '请先登录！'
    });
    next({
      path: '/login'
    });
  } else if (
    store.getters.isLoggedIn &&
    !to.meta.requireLogin &&
    !to.meta.isPublic
  ) {
    // 2.用户已登录，但又去访问 登录/注册/请求重置密码/重置密码 页面时不让他过去
    Swal.fire({
      position: 'top',
      icon: 'error',
      title: '请勿重复登陆！',
      showConfirmButton: false,
      timer: 1500
    });
    next({
      path: '/'
    });
  } else if (!store.getters.isAdmin && to.meta.requireAdmin) {
    // 3.普通用户想在浏览器地址中直接访问 /admin ，提示他没有权限，并跳转到首页
    Swal.fire({
      position: 'top',
      icon: 'error',
      title: '没有管理员权限！',
      showConfirmButton: false,
      timer: 1500
    });
    next({
      path: '/'
    });
  } else {
    // 正常路由出口
    next();
  }
});

router.afterEach(() => {
  window.scroll(0, 0);
});

export default router;
