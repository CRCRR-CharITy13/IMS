import styles from "../styles/Sidebar.module.scss";

import { DiGoogleAnalytics } from "react-icons/di";
import { BsFillBinocularsFill } from "react-icons/bs";
import { FaWarehouse, FaCog, FaUsers } from "react-icons/fa";
import { AiOutlineQrcode } from "react-icons/ai";
import { Burger } from "@mantine/core";
import { MdAdminPanelSettings } from "react-icons/md";
import {GrTransaction} from "react-icons/gr"

import { useNavigate } from "react-router-dom";

import { ReactElement, useEffect, useState } from "react";

import { FileInvoice } from "tabler-icons-react";
//import {Item} from "../types/item";

import logo from '../assets/Logo.png';

const SidebarItem = ({
    label,
    icon,
    onClick,
    selected,
}: {
    label: string;
    icon: ReactElement;
    onClick?: () => void;
    selected?: boolean;
}) => {
    return (
        <>
            <div
                className={
                    styles.item + `${selected ? " " + styles.selected : ""}`
                }
                onClick={onClick}
            >
                <p
                    style={{
                        margin: 0,
                        display: "flex",
                        alignItems: "center",
                        gap: ".5rem",
                    }}
                >
                    {icon} {label}
                </p>
            </div>
        </>
    );
};

const Sidebar = () => {
    const navigate = useNavigate();

    const [selected, setSelected] = useState<string>("Analytics");

    const [username, setUsername] = useState<string>("Analytics");

    const [visible, setVisible] = useState<boolean>(false);

    const [innerWidth, setInnerWidth] = useState<number>(window.innerWidth);

    const [isAdmin, setIsAdmin] = useState<boolean>(false);

    window.onresize = () => {
        setInnerWidth(window.innerWidth);
    };

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

    return (
        <>
            <div className={styles.container}>
                <div
                    style={{
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "space-between",
                        paddingRight: "1rem",
                        boxSizing: "border-box",
                    }}
                >
                    <h1
                        style={{
                            margin: "1rem",
                            marginLeft: "2.4rem"
                        }}
                    >
                        Dashboard
                    </h1>
                    {innerWidth < 800 && (
                        <Burger
                            opened={visible}
                            onClick={() => {
                                setVisible(!visible);
                            }}
                        />
                    )}
                </div>
                {(visible || innerWidth >= 800) && (
                    <>
                        <SidebarItem
                            label="Analytics"
                            selected={selected === "Analytics"}
                            onClick={() => setSelected("Analytics")}
                            icon={<DiGoogleAnalytics size={25} />}
                        />
                        <SidebarItem
                            label="Scanner"
                            onClick={() => setSelected("Scanner")}
                            icon={<AiOutlineQrcode size={25} />}
                        />

                        <SidebarItem
                            label="Transaction"
                            selected={selected === "transaction"}
                            onClick={() => setSelected("transaction")}
                            icon={<GrTransaction size={25} />}
                        />
                        <SidebarItem
                            label="Inventory"
                            selected={selected === "Inventory"}
                            onClick={() => setSelected("Inventory")}
                            icon={<FaWarehouse size={25} />}
                        />
                        <SidebarItem
                            label="Clients & Donors"
                            selected={selected === "clientsdonors"}
                            onClick={() => setSelected("clientsdonors")}
                            icon={<FaUsers size={25} />}
                        />
                        <SidebarItem
                            label="Invoice"
                            selected={selected === "invoice"}
                            onClick={() => setSelected("invoice")}
                            icon={<FileInvoice size={25} />}
                        />


                        {isAdmin && (
                            <SidebarItem
                                label="Audit Logs"
                                selected={selected === "Audit"}
                                onClick={() => setSelected("Audit")}
                                icon={<BsFillBinocularsFill size={25} />}
                            />
                        )}
                        {isAdmin && (

                            <SidebarItem
                                label="Admin"
                                icon={<MdAdminPanelSettings size={25} />}
                                selected={selected === "Admin"}
                                onClick={() => setSelected("Admin")}
                            />
                        )}
                        <SidebarItem
                            label="Settings"
                            icon={<FaCog size={25} />}
                            selected={selected === "Settings"}
                            onClick={() => setSelected("Settings")}
                        />
                    </>
                )}
                {(innerWidth >= 800) && (
                    <>
                        <div style={{bottom: 0, position: 'relative', width:'100%'}}>
                            <img src={logo} style={{width: '50%', margin:'5%', marginLeft:'25%'}} alt="Logo" />
                            <h3 style={{marginTop:'0%',width:'100%', textAlign:'center'}}>
                                Welcome {username}
                            </h3>
                        </div>
                    </>
                )}
            </div>
        </>
    );
};

export default Sidebar;
