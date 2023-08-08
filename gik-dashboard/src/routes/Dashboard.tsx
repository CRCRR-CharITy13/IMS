import { useEffect, useState } from "react";
import Analytics from "../components/dashboard/Analytics";
import Sidebar from "../components/Sidebar";

import styles from "../styles/Dashboard.module.scss";

import { useNavigate, useParams } from "react-router-dom";
import AuditLog from "../components/dashboard/AuditLog";
import Scanner from "../components/dashboard/Scanner";
import Admin from "../components/dashboard/Admin";
import Settings from "../components/dashboard/Settings";
import Inventory from "../components/dashboard/Inventory";
import ClientsDonors from "../components/dashboard/ClientsDonors";
import Invoice from "../components/dashboard/Invoice";
import Transactions from "../components/dashboard/Transactions";
import {Button, Text} from "@mantine/core";
import {openConfirmModal} from "@mantine/modals";

const Dashboard = () => {
    const navigate = useNavigate();
    const [pane, setPane] = useState<JSX.Element>(<Analytics />);

    const [showTfaSetup, setShowTfaSetup] = useState<boolean>(false);

    const { handle } = useParams();

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
            case "scanner":
                requiredPane = <Scanner />;
                break;
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
        <>
            <div className={styles.wrapper}>
                <Sidebar />
                <div className={styles.pane}>{pane}</div>
            </div>
        </>
    );
};

export default Dashboard;
