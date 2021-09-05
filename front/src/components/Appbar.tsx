import {
    AppBar,
    Avatar,
    createStyles,
    IconButton,
    makeStyles,
    Toolbar,
    MenuItem,
    Menu
} from '@material-ui/core';
import {Person} from '@material-ui/icons'
import {signIn, signOut, useSession} from 'next-auth/client';
import {Theme} from 'next-auth';
import React, {useState} from 'react';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        toolbar: {
            justifyContent: "flex-end"
        },
        button: {
        },
        empty: {
            flexGrow: 1,
        }
    }),
);


function UserButton(){
    const [session] = useSession()
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
    const isMenuOpen = Boolean(anchorEl);
    const handleProfileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };
    const handleMenuClose = () => {
        setAnchorEl(null);
    };
    let Icon: JSX.Element
    let menuItems: JSX.Element
    if (session && session.user) {
        let image = session.user.image;
        const username = session.user.name
        if (image === null){
            image = undefined
        }
        menuItems = (
            <>
                <MenuItem disabled>You are Loggined as<br/> {username}</MenuItem>
                <MenuItem onClick={()=>{signOut().then()}} >Logout</MenuItem>
            </>
        )
        Icon = (<Avatar src={image}/>)
    }
    else{
        menuItems = (
            <MenuItem onClick={()=>{signIn().then()}}>Login</MenuItem>
        )
        Icon = (<Person/>)
    }
    const menu = (
        <Menu open={isMenuOpen} onClose={handleMenuClose} anchorEl={anchorEl}>
            {menuItems}
        </Menu>
    )
    return (
        <>
            <IconButton onClick={handleProfileMenuOpen}>
                {Icon}
            </IconButton>
            {menu}
        </>
    )
}

export function MyAppBar(){
    const classes = useStyles();
    return (
        <AppBar position="static">
            <Toolbar className={classes.toolbar}>
                <UserButton/>
            </Toolbar>
        </AppBar>
    )
}