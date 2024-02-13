import {
    Box,
    Button,
    Center,
    Input,
    Table,
    SegmentedControl,
    Accordion,
    Pagination,
    Space,
    LoadingOverlay,
    TextInput,
} from "@mantine/core";
import { Tags } from "tabler-icons-react";

import { DatePickerInput } from "@mantine/dates";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { AdvancedLog } from "../../types/logs";

import styles from "../../styles/AuditLog.module.scss";
import { containerStyles } from "../../styles/container";
import { SimpleUser } from "../../types/user";

interface DisplayAdvanceLog {
    ID: number;
    control: string;
    ipAddress: string;
    userAgent: string;
    method: string;
    path: string;
    userId: string;
    timestamp: string;
}

const AdvancedLogs = ({
    actionFilter,
    dateFilter,
    userFilter,
    setVisible,
    usernames,
    getUsernames,
}: {
    actionFilter: string;
    dateFilter: [Date | null, Date | null] | undefined;
    setVisible: Dispatch<SetStateAction<boolean>>;
    userFilter: string;
    usernames: Map<number, string>;
    getUsernames: () => Promise<void>;
}) => {
    const [data, setData] = useState<AdvancedLog[]>([]);

    const [currentPage, setCurrentPage] = useState<number>(1);

    const [totalPages, setTotalPages] = useState<number>(1);


    const getData = async () => {
        setVisible(true);

        const response = await fetch(
            `${
                process.env.REACT_APP_API_URL
            }/logs/advanced?page=${currentPage}&action=${actionFilter}&date=${
                dateFilter && dateFilter[0] && dateFilter[1]
                    ? `${dateFilter[0].getTime() / 1000},${
                          dateFilter[1].getTime() / 1000
                      }`
                    : ""
            }&user=${userFilter}`,
            {
                credentials: "include",
            }
        );

        setVisible(false);

        const data: {
            success: boolean;
            data: {
                data: AdvancedLog[];
                total: number;
                currentPage: number;
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setData(data.data.data);
            setTotalPages(data.data.totalPages);
        }
    };

    useEffect(() => {
        getUsernames();
        getData();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        getUsernames();
        getData();
    }, [actionFilter, dateFilter, userFilter]);

    const displayAdvanceLogs : DisplayAdvanceLog[] = [];
    for(let idx = 0; idx<data.length; idx++){
        const strUserId = "User name: " + usernames.get(data[idx].userId); //.toString();
        const strIpAddress = "IpAddress: " + data[idx].ipAddress;
        const strUserAgent = "User Agent: " + data[idx].userAgent;
        const strMethod = "Method: " + data[idx].method;
        const strPath = "Path: " + data[idx].path;
        const strTimeStamp = "Time: " + new Date(data[idx].timestamp * 1000).toLocaleString();
        //
        let controlStr = "Log ID: ";
        controlStr += data[idx].ID.toString();
        controlStr = controlStr + "; " + strUserId;
        controlStr = controlStr + "; " + strTimeStamp;
        controlStr = controlStr + "; " + strPath;
        const newDisplayAdvanceLog : DisplayAdvanceLog = {
            ID : data[idx].ID,
            control : controlStr,
            ipAddress : strIpAddress,
            userAgent : strUserAgent,
            method : strMethod,
            path : strPath,
            userId : strUserId,
            timestamp : strTimeStamp,
        }
        displayAdvanceLogs.push(newDisplayAdvanceLog);
    }

    const advancedLogItems = displayAdvanceLogs.map((log) => (
        <Accordion.Item key={log.ID} value={log.ID.toString()}>
            <Accordion.Control 
            icon={
                <Tags
                  style={{ color: 'var(--mantine-color-red-filled'}}
                />
              }>
                {log.control}</Accordion.Control>
            <Accordion.Panel>{log.userId}</Accordion.Panel>
            <Accordion.Panel>{log.ipAddress}</Accordion.Panel>
            <Accordion.Panel>{log.userAgent}</Accordion.Panel>
            <Accordion.Panel>{log.method}</Accordion.Panel>
            <Accordion.Panel>{log.path}</Accordion.Panel>
            <Accordion.Panel>{log.timestamp}</Accordion.Panel>
        </Accordion.Item>
    ));


    return (
        <>
        <Accordion multiple>
            {advancedLogItems}
        </Accordion>
        
            <Center
                sx={{
                    marginTop: "1rem",
                }}
            >
            <Pagination
                value={currentPage}
                onChange={setCurrentPage}
                total={totalPages}
            />
            </Center>
        </>
    );
};

const SimpleLogRow = ({
    log,
    usernames,
}: { 
    log: AdvancedLog;
    usernames: Map<number, string>;
}) => {
    return (
        <>
            <tr>
                <td>{log.ID}</td>
                <td>{usernames.get(log.userId)}</td>
                <td>{log.ipAddress}</td>
                <td>{log.action}</td>
            </tr>
        </>
    );
};

const SimpleLogs = ({
    actionFilter,
    dateFilter,
    userFilter,
    setVisible,
    usernames,
    getUsernames,
}: {
    actionFilter: string;
    dateFilter: [Date | null, Date | null] | undefined;
    setVisible: Dispatch<SetStateAction<boolean>>;
    userFilter: string;
    usernames: Map<number, string>;
    getUsernames: () => Promise<void>;
}) => {
    const [data, setData] = useState<AdvancedLog[]>([]);

    const [currentPage, setCurrentPage] = useState<number>(1);

    const [totalPages, setTotalPages] = useState<number>(1);

    const getData = async () => {
        setVisible(true);

        const response = await fetch(
            `${
                process.env.REACT_APP_API_URL
            }/logs/simple?page=${currentPage}&action=${actionFilter}&date=${
                dateFilter && dateFilter[0] && dateFilter[1]
                    ? `${dateFilter[0].getTime() / 1000},${
                          dateFilter[1].getTime() / 1000
                      }`
                    : ""
            }&user=${userFilter}`,
            {
                credentials: "include",
            }
        );

        setVisible(false);

        const data: {
            success: boolean;
            data: {
                data: AdvancedLog[];
                total: number;
                currentPage: number;
                totalPages: number;
            };
        } = await response.json();

        if (data.success) {
            setData(data.data.data);
            setTotalPages(data.data.totalPages);
        }
    };

    useEffect(() => {
        getUsernames();
        getData();
    }, [currentPage]);

    useEffect(() => {
        setCurrentPage(1);
        getUsernames();
        getData();
    }, [actionFilter, dateFilter, userFilter]);

    return (
        <>
            <Table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Username</th>
                        <th>IP</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((log) => (
                        <SimpleLogRow key={log.ID} log={log} usernames={usernames} />
                    ))}
                </tbody>
            </Table>
            <Center
                sx={{
                    marginTop: "1rem",
                }}
            >
                <Pagination
                    value={currentPage}
                    onChange={setCurrentPage}
                    total={totalPages}
                />
            </Center>
        </>
    );
};

const AuditLog = () => {
    const [viewMode, setViewMode] = useState<"smp" | "adv">("smp");
    const [visible, setVisible] = useState<boolean>(false);

    // Initialize filtering
    const [actionFilter, setActionFilter] = useState<string>("");
    const [actionFilterEditing, setActionFilterEditing] = useState<string>("");

    const [dateFilter, setDateFilter] = useState<
        [Date | null, Date | null] | undefined
    >();
    const [dateFilterEditing, setDateFilterEditing] = useState<
        [Date | null, Date | null] | undefined
    >();

    const [userFilter, setUserFilter] = useState<string>("");
    const [userFilterEditing, setUserFilterEditing] = useState<string>("");

    const doFilter = async () => { 
        setActionFilter(actionFilterEditing);
        setDateFilter(dateFilterEditing);
        setUserFilter(userFilterEditing);
    };

    // Get usernames
    const [usernames, setUsernames] = useState<Map<number, string>>(new Map());

    const getUsernames = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/admin/lists`,
            {
                credentials: "include",
            }
        );

        const data: {
            success: boolean;
            data: [SimpleUser[], SimpleUser[]];
        } = await response.json();

        const dict = new Map();

        [...data.data[0], ...data.data[1]].map((username) => {
            dict.set(Number(username.ID), username.username);
        })

        setUsernames(dict);
    };

    return (
        <>
            <div
                style={containerStyles}
            >
                <LoadingOverlay visible={visible} />
                <Box
                    sx={containerStyles}
                >
                    <SegmentedControl
                        data={[
                                { label: "Simple", value: "smp" },
                                { label: "Advanced", value: "adv" },
                        ]}
                        sx={{
                            marginBottom: "1rem",
                        }}
                        onChange={(value: "smp" | "adv") => setViewMode(value)}
                    />

                    <h1>
                        {viewMode === "smp" ? "Simple" : "Advanced"} Audit Logs
                    </h1>
                    <h2>View and filter actions and transactions.</h2>
                    <Center
                        sx={{
                            justifyContent: "space-between",
                            gap: ".5rem",
                            alignItems: "flex-end",
                            marginTop: "1rem",
                            marginBottom: "1rem",
                        }}
                    >
                        <DatePickerInput
                            type="range"
                            placeholder="Pick Date Range"
                            onChange={setDateFilterEditing}
                        />
                        <TextInput
                            sx={{
                                display: "flex",
                            }}
                            placeholder="User"
                            onChange={(e: any) =>
                                setUserFilterEditing(e.target.value)
                            }
                        />
                        <TextInput
                            placeholder="Action"
                            sx={{
                                display: "flex",
                                flexGrow: 1,
                            }}
                            onChange={(e: any) =>
                                setActionFilterEditing(e.target.value)
                            }
                        />
                        <Button color="teal" onClick={doFilter}>
                            Filter
                        </Button>
                    </Center>
                    {viewMode === "smp" ? (
                        <SimpleLogs
                            actionFilter={actionFilter}
                            dateFilter={dateFilter}
                            setVisible={setVisible}
                            userFilter={userFilter}
                            usernames={usernames}
                            getUsernames={getUsernames}
                        />
                    ) : (
                        <AdvancedLogs
                            actionFilter={actionFilter}
                            dateFilter={dateFilter}
                            setVisible={setVisible}
                            userFilter={userFilter}
                            usernames={usernames}
                            getUsernames={getUsernames}
                        />
                    )}
                </Box>
            </div>
        </>
    );
};

export default AuditLog;
