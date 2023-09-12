import { useEffect, useState } from "react";
import Analytics from "../components/dashboard/Analytics";
import Sidebar from "../components/Sidebar";

import styles from "../styles/Dashboard.module.scss";

import { useNavigate, useParams } from "react-router-dom";
import AuditLog from "../components/dashboard/AuditLog";
// TEMPORARY COMMENTED BY TUAN IN 12 SEPT 2023
// import Scanner from "../components/dashboard/Scanner";
import Admin from "../components/dashboard/Admin";
import Settings from "../components/dashboard/Settings";
import Inventory from "../components/dashboard/Inventory";
import ClientsDonors from "../components/dashboard/ClientsDonors";
import Invoice from "../components/dashboard/Invoice";
import Transactions from "../components/dashboard/Transactions";
import {AppShell, Button, ScrollArea, Text, createStyles, rem, useMantineTheme} from "@mantine/core";
import {openConfirmModal} from "@mantine/modals";
import { Container } from "tabler-icons-react";
import { FALSE } from "sass";

const useStyles = createStyles((theme) => ({
    wrapper: {
        height: '100%',
        width: '100%',
        display: 'flex',
    },
    pane: {
        padding: rem(2),
        width: '100%',
        overflowY: 'auto'
    }
}));

const Dashboard = () => {
    const navigate = useNavigate();
    const [pane, setPane] = useState<JSX.Element>(<Analytics />);

    const [showTfaSetup, setShowTfaSetup] = useState<boolean>(false);

    const { handle } = useParams();

    const theme = useMantineTheme();

    const { classes, cx } = useStyles();

    const checkAuthStatus = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/status`,
            {
                credentials: "include",
            }
        );

        if (response.status !== 200) {
            navigate("/login", { replace: true });
            return;
        }

        //await checkTfaStatus();
    };

    useEffect(() => {
        checkAuthStatus();
    }, []);

    useEffect(() => {
        let requiredPane: JSX.Element = <></>;

        switch (handle) {
            case "analytics":
                requiredPane = <Analytics />;
                break;
            case "audit":
                requiredPane = <AuditLog />;
                break;
            // TEMPORARY COMMENTED BY TUAN IN 12 SEPT 2023
            // case "scanner":
            //     requiredPane = <Scanner />;
            //     break;
            case "admin":
                requiredPane = <Admin />;
                break;
            case "settings":
                requiredPane = <Settings />;
                break;
            case "inventory":
                requiredPane = <Inventory />;
                break;
            case "clientsdonors":
                requiredPane = <ClientsDonors />;
                break;
            case "invoice":
                requiredPane = <Invoice />;
                break;
            case "transaction":
                requiredPane = <Transactions />
                break;
        }

        setPane(requiredPane);
    }, [handle]);


    return (
        <div className={classes.wrapper}>
            <Sidebar />
            <div className={classes.pane}>{pane}</div>
        </div>

    );
};

export default Dashboard;
