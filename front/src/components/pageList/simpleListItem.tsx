import React, {MouseEventHandler} from 'react';
import {ListItem, ListItemIcon, ListItemText} from '@material-ui/core';
import {Icons} from '../../const/const'

export function SimpleListItem(prop: { name: string, icon: string, onClick?: MouseEventHandler<HTMLDivElement> }) {
    return (
        <ListItem button onClick={prop.onClick}>
            <ListItemIcon>
                {React.createElement(Icons[prop.icon])}
            </ListItemIcon>
            <ListItemText primary={prop.name}/>
        </ListItem>
    )
}
