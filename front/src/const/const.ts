import {Page} from '../components/pageList/pageList';
import {Home, Info, SvgIconComponent} from '@material-ui/icons';
import MenuIcon from '@material-ui/icons/Menu';

export const PublicPages: Array<Page> = [
    {
        name: 'Home Page',
        root: '/',
        icon: 'Home',
    },
    {
        name: 'About',
        root: '/about',
        icon: 'Info'
    }
]

export const Icons = {
    'Home': Home,
    'Info': Info,
    'Menu': MenuIcon
} as { [index: string]: SvgIconComponent }