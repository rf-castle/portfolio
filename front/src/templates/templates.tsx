import React from 'react';
import PageList from '../components/pageList/pageList';
import {Drawer} from '@material-ui/core';
import clsx from 'clsx';
import {MyAppBar} from '../components/Appbar';

type Prop = {
}

type State = {
}

class Layout extends React.Component<Prop, State>{

    render(){
        return (
            <div>
                {/*<PageList/>*/}
                <MyAppBar/>
                <main>
                    {this.props.children}
                </main>
            </div>
        )
    }
}
export default Layout;