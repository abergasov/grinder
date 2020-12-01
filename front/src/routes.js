import AppHome from './components/AppHome';
import AppProfile from "./components/AppProfile";

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
];

export default routes;
