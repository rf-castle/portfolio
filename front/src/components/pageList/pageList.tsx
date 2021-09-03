import React, {createElement, MouseEventHandler} from 'react';
import Link from 'next/link'
import {
    createStyles,
    Divider,
    Drawer,
    List,
    ListSubheader,
    Theme,
    withStyles,
    WithStyles
} from '@material-ui/core';
import {SimpleListItem} from './simpleListItem';
import {PublicPages} from '../../const/const';
import clsx from 'clsx';


export interface Page {
    name: string
    root: string
    icon: string
}

function PageListItem(page: Page) {
    return (
        <Link href={page.root}>
            <a>
                <SimpleListItem name={page.name} icon={page.icon}/>
            </a>
        </Link>
    )
}


const styles = (theme: Theme) => {
    const drawerWidth = 240
    return createStyles({
        drawer: {
            width: drawerWidth,
            flexShrink: 0,
            whiteSpace: 'nowrap',
        },
        drawerOpen: {
            width: drawerWidth,
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
        },
        drawerClose: {
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen,
            }),
            overflowX: 'hidden',
            width: theme.spacing(7) + 1,
            [theme.breakpoints.up('sm')]: {
                width: theme.spacing(8) + 1,
            },
        },
        buttonLabel: {
            'justify-content': 'flex-start'
        }
    })
};

type Prop = {} & WithStyles<typeof styles>

type State = {
    secretPages: Array<Page>
    open: boolean
}

class PageList extends React.Component<Prop, State> {

    constructor(prop: Prop) {
        super(prop);
        this.state = {
            secretPages: [],
            open: false
        }
    }

    onLogin = () => {
        // Todo: Fetch Secret Page List
    }

    onLogout = () => {
        this.clearPages()
    }

    updatePages = (pages: Array<Page>) => {
        this.setState({
            secretPages: pages
        })
    }

    clearPages = () => {
        this.setState({
            secretPages: []
        })
    }

    setClose = () => {
        this.setState({
            open: false
        })
    }

    toggleOpen = () => {
        this.setState((state) => ({
            open: !state.open
        }))
    }

    render() {
        const classes = this.props.classes
        const secretPages = this.state.secretPages
        const callback = (page: Page) => (
            createElement(
                PageListItem,
                Object.assign({key: page.root}, page)
            )
        )
        let listItems: Array<JSX.Element> = [createElement(
            SimpleListItem,
            {name: "Page List", icon: "Menu", onClick: () => {this.toggleOpen()}, key: "list"}
        ), createElement(Divider, {key: "Divider1"})]
        listItems = listItems.concat(PublicPages.map(callback))
        if (secretPages.length > 0) {
            listItems.push(createElement(Divider, {key: "Divider2"}))
            listItems.push(createElement(ListSubheader, {key: 'secretPage'}, 'secretPage'))
            listItems = listItems.concat(secretPages.map(callback))
        }
        return (
            <div className="pageList">
                <Drawer
                    variant="permanent"
                    className={clsx(classes.drawer, {
                        [classes.drawerClose]: !this.state.open,
                        [classes.drawerOpen]: this.state.open,
                    })}
                    classes={{
                        paper: clsx({
                            [classes.drawerOpen]: this.state.open,
                            [classes.drawerClose]: !this.state.open,
                        }),
                    }}
                    onBlur={()=>{this.setClose()}}
                >
                    <Divider/>
                    <List >
                        {listItems}
                    </List>
                </Drawer>
            </div>
        );
    }
}


export default withStyles(styles, {withTheme: true})(PageList);