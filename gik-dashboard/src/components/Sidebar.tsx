import styles from "../styles/Sidebar.module.scss";

import { Burger, createStyles, Navbar, Group, Code, getStylesRef, rem, ActionIcon, useMantineTheme, useMantineColorScheme } from "@mantine/core";

import { useNavigate } from "react-router-dom";

import { ReactElement, useEffect, useState } from "react";

import { FileInvoice, Qrcode, BuildingWarehouse, Settings, Users, ArrowsRightLeft, Logout, Shield, ReportAnalytics, Notes, Sun, MoonStars, BuildingBank} from "tabler-icons-react";
//import {Item} from "../types/item";

import logo from '../assets/Logo.png';

const useStyles = createStyles((theme) => ({
  header: {
    paddingBottom: theme.spacing.md,
    marginBottom: `calc(${theme.spacing.md} * 1.5)`,
    borderBottom: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]
    }`,
  },

  footer: {
    paddingTop: theme.spacing.md,
    marginTop: theme.spacing.md,
    borderTop: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2]
    }`,
  },

  link: {
    ...theme.fn.focusStyles(),
    display: 'flex',
    alignItems: 'center',
    textDecoration: 'none',
    fontSize: theme.fontSizes.sm,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[1] : theme.colors.gray[7],
    padding: `${theme.spacing.xs} ${theme.spacing.sm}`,
    borderRadius: theme.radius.sm,
    fontWeight: 500,

    '&:hover': {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
      color: theme.colorScheme === 'dark' ? theme.white : theme.black,

      [`& .${getStylesRef('icon')}`]: {
        color: theme.colorScheme === 'dark' ? theme.white : theme.black,
      },
    },
  },

  logo: {
    ...theme.fn.focusStyles(),
    display: 'flex',
    alignItems: 'center',
    textDecoration: 'none',
    fontSize: theme.fontSizes.sm,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[1] : theme.colors.gray[7],
    padding: `${theme.spacing.xs} ${theme.spacing.sm}`,
    borderRadius: theme.radius.sm,
    fontWeight: 500,
  },

  linkIcon: {
    ref: getStylesRef('icon'),
    color: theme.colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.gray[6],
    marginRight: theme.spacing.sm,
  },

  linkActive: {
    '&, &:hover': {
      backgroundColor: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).background,
      color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
      [`& .${getStylesRef('icon')}`]: {
        color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
      },
    },
  },
}));

  const data = [
    // { needsAdmin: false, label: 'Location', icon: BuildingBank, route: 'Location' },
    { needsAdmin: false, label: 'Analytics', icon: ReportAnalytics, route: 'Analytics' },
    //{ needsAdmin: false, label: 'Scanner', icon: Qrcode, route: 'Scanner' },
    { needsAdmin: false, label: 'Transaction', icon: ArrowsRightLeft, route: 'Transaction' },
    { needsAdmin: false, label: 'Inventory', icon: BuildingWarehouse, route: 'Inventory' },
    { needsAdmin: false, label: 'Clients & Donors', icon: Users, route: 'Clientsdonors' },
    { needsAdmin: false, label: 'Invoice', icon: FileInvoice, route: 'Invoice' },
    { needsAdmin: true, label: 'Audit Logs', icon: Notes, route: 'Audit' },
    { needsAdmin: true, label: 'Admin', icon: Shield, route: 'Admin' },
    { needsAdmin: false, label: 'Settings', icon: Settings, route: 'Settings' },
  ];

const Sidebar = () => {
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();
    const theme = useMantineTheme();

    const { classes, cx } = useStyles();
    const navigate = useNavigate();

    const [selected, setSelected] = useState<string>("Analytics");

    const [username, setUsername] = useState<string>("Analytics");

    const [isAdmin, setIsAdmin] = useState<boolean>(false);

    const links = data.map((item) => (
      <>
        {(!item.needsAdmin || isAdmin) && 
          <div
            className={cx(classes.link, { [classes.linkActive]: item.route === selected })}
            key={item.route}
            onClick={(event) => {
              event.preventDefault();
              setSelected(item.route);
            }}
          >
            <item.icon className={classes.linkIcon} strokeWidth={1.5} size={24} stroke="currentColor"/>
            <span>{item.label}</span>
          </div>
        }
      </>
    ));

    useEffect(() => {
        checkAdminStatus();
        getUsername()
    }, []);

    useEffect(() => {
        navigate(`/dashboard/${selected.toLowerCase()}`, { replace: true });
    }, [selected]);

    const checkAdminStatus = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/admin/status`,
            {
                credentials: "include",
            }
        );

        setIsAdmin(response.status === 200);

    };

    const getUsername = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/info/currentusername`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: string;
        } = await response.json();

        setUsername(data.data);
    };

    const DoLogout = async () => {
      console.log("LOGGING OUT")
      await fetch(`${process.env.REACT_APP_API_URL}/auth/logout`,
          {credentials: "include",});
      window.location.reload();
  }

    return (
      <Navbar width={{ sm: 275 }} p="md" zIndex={0.5}>
        <Navbar.Section grow>
          <Group className={classes.header} position="apart">
            <div className={classes.logo}>
              <img src={logo} style={{height:'34px'}} alt="Logo" />
            </div>

            <ActionIcon
                onClick={() => toggleColorScheme()}
                size="lg"
                sx={(theme) => ({
                backgroundColor:
                    theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
                    color: theme.colorScheme === 'dark' ? theme.colors.teal[4] : theme.colors.teal[6],
                })}
            >
                {colorScheme === 'dark' ? <Sun size="1.2rem" /> : <MoonStars size="1.2rem" />}
            </ActionIcon>
          </Group>
          {links}
        </Navbar.Section>
  
        <Navbar.Section className={classes.footer}>
          <div className={classes.link} onClick={DoLogout}>
            <Logout className={classes.linkIcon} strokeWidth={1.5} />
            <span>Logout</span>
          </div>
        </Navbar.Section>
      </Navbar>
    );
};

export default Sidebar;
