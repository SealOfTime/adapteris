import { ClipboardListIcon, FireIcon, HomeIcon, UserIcon } from "@heroicons/react/outline";
import { BottomNavigation, BottomNavigationAction, Paper, useTheme } from "@mui/material";
import React, { CSSProperties, FC } from "react";
import { Link, matchPath, useLocation } from 'react-router-dom';

export const Navbar: FC = () => {
    const match = useRouteMatch(['/feed', '/school', '/results', '/profile']);
    const currentPage = match?.pattern?.path;
    const theme = useTheme();
    const iconStyle: CSSProperties = { color: theme.palette.primary.main, width: '32px', height: '32px' };
    const linkStyle: CSSProperties = { display: 'flex', alignItems: 'center', justifyContent: 'center', lineHeight: 'normal' };
    return (
        <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
            <BottomNavigation value={currentPage} showLabels>
                <BottomNavigationAction
                    label="Результаты"
                    value="/results"
                    icon={<ClipboardListIcon style={iconStyle} />}
                    component={Link} to="/results" style={linkStyle}
                />
                <BottomNavigationAction
                    label="Школа"
                    value="/school"
                    icon={<HomeIcon style={iconStyle} />}
                    component={Link} to="/school" style={linkStyle}
                />
                <BottomNavigationAction
                    label="Актуальное"
                    value="/feed"
                    icon={<FireIcon style={iconStyle} />}
                    component={Link} to="/feed" style={linkStyle}
                />
                <BottomNavigationAction
                    label="Профиль"
                    value="/profile"
                    icon={<UserIcon style={iconStyle} />}
                    component={Link} to="/profile" style={linkStyle}
                />
            </BottomNavigation>
        </Paper>
    )
};

function useRouteMatch(patterns) {
    const { pathname } = useLocation();

    for (let i = 0; i < patterns.length; i += 1) {
        const pattern = patterns[i];
        const possibleMatch = matchPath(pattern, pathname);
        if (possibleMatch !== null) {
            return possibleMatch;
        }
    }

    return null;
}
