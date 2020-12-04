import AppHome from './components/AppHome';
import AppProfile from "./components/AppProfile";
import AppUsers from "./components/AppUsers";

const routes = [
    {
        path: '/',
        name: 'home',
        component: AppHome
    },
    {
        path: '/profile',
        name: 'profile',
        component: AppProfile
    },
    {
        path: '/users',
        name: 'users',
        component: AppUsers
    },
];

export default routes;
